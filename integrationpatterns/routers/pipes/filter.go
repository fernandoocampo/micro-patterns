package pipes

import (
	"context"
	"encoding/json"
	"log"
)

type Storage struct {
	order         Order
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Logger struct {
	log           string
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Enrich struct {
	newID         string
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Validator struct {
	result        bool
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Collector struct {
	newOrder      NewOrder
	RequestStream chan interface{}
	NextStream    chan interface{}
}

func NewStorage(requestStream chan interface{}) *Storage {
	newStorage := Storage{
		RequestStream: requestStream,
	}

	return &newStorage
}

func (s *Storage) WithNextFilter(nextStream chan interface{}) *Storage {
	s.NextStream = nextStream

	return s
}

func (s *Storage) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-s.RequestStream:
				newOrder, ok := event.(Order)
				if !ok {
					log.Printf("unexpected order %T, with value: %+v", event, event)

					continue
				}

				s.order = newOrder
				s.order.Success = true
				s.NextStream <- true
			}
		}
	}()
}

func (s *Storage) GetOrder() Order {
	return s.order
}

func NewLogger(requestStream chan interface{}) *Logger {
	newLogger := Logger{
		RequestStream: requestStream,
	}

	return &newLogger
}

func (l *Logger) WithNextFilter(nextStream chan interface{}) *Logger {
	l.NextStream = nextStream

	return l
}

func (l *Logger) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-l.RequestStream:
				newOrder, ok := event.(Order)
				if !ok {
					log.Printf("unexpected order %T, with value: %+v", event, event)

					continue
				}

				stringOrder, err := json.Marshal(&newOrder)
				if err != nil {
					log.Printf("unexpected error marshalling order: %s", err)

					continue
				}

				l.log = string(stringOrder)
				l.NextStream <- newOrder
			}
		}
	}()
}

func (l *Logger) Logs() string {
	return l.log
}

func NewValidator(requestStream chan interface{}) *Validator {
	newValidator := Validator{
		RequestStream: requestStream,
	}

	return &newValidator
}

func (v *Validator) WithNextFilter(nextStream chan interface{}) *Validator {
	v.NextStream = nextStream

	return v
}

func (v *Validator) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-v.RequestStream:
				newOrder, ok := event.(NewOrder)
				if !ok {
					log.Printf("unexpected order %T, with value: %+v", event, event)

					continue
				}

				if newOrder.ID == "" || newOrder.Value < 1.0 {
					v.result = false

					log.Println("new order is invalid", newOrder)

					continue
				}

				v.result = true
				v.NextStream <- newOrder
			}
		}
	}()
}

func (v *Validator) Result() bool {
	return v.result
}

func NewEnrich(requestStream chan interface{}) *Enrich {
	newEnrich := Enrich{
		RequestStream: requestStream,
	}

	return &newEnrich
}

func (e *Enrich) WithNextFilter(nextStream chan interface{}) *Enrich {
	e.NextStream = nextStream

	return e
}

func (e *Enrich) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-e.RequestStream:
				newOrder, ok := event.(NewOrder)
				if !ok {
					log.Printf("unexpected order %T, with value: %+v", event, event)

					continue
				}

				e.newID = newOrder.ID + "-" + "Z1"

				order := Order{
					ID:         newOrder.ID,
					Value:      newOrder.Value,
					InternalID: e.newID,
				}

				e.NextStream <- order
			}
		}
	}()
}

func (e *Enrich) GetGeneratedID() string {
	return e.newID
}

func NewCollector(requestStream chan interface{}) *Collector {
	newCollector := Collector{
		RequestStream: requestStream,
	}

	return &newCollector
}

func (c *Collector) WithNextFilter(nextStream chan interface{}) *Collector {
	c.NextStream = nextStream

	return c
}

func (c *Collector) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-c.RequestStream:
				newOrder, ok := event.(NewOrder)
				if !ok {
					log.Printf("unexpected order %T, with value: %+v", event, event)

					continue
				}

				c.newOrder = newOrder
				c.NextStream <- newOrder
			}
		}
	}()
}

func (c *Collector) Value() NewOrder {
	return c.newOrder
}
