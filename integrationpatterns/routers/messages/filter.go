package messages

import (
	"context"
	"encoding/json"
	"log"
)

type Storage struct {
	order         Order
	log           string
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Logger struct {
	logs          []string
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Enrich struct {
	newID         string
	RequestStream chan interface{}
	NextStream    chan interface{}
}

type Validator struct {
	result        []bool
	RequestStream chan interface{}
	EnrichStream  chan interface{}
	LoggerStream  chan interface{}
}

type Collector struct {
	newOrders     []NewOrder
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
				switch newEvent := event.(type) {
				case string:
					log.Println("log:", newEvent)
					s.log = newEvent
				case Order:
					log.Println("order:", newEvent)
					s.order = newEvent
					s.order.Success = true
				default:
					log.Printf("unexpected order %T, with value: %+v", event, event)
				}
			}

			s.NextStream <- true
		}
	}()
}

func (s *Storage) GetOrder() Order {
	return s.order
}

func (s *Storage) GetLog() string {
	return s.log
}

func NewLogger(requestStream chan interface{}) *Logger {
	newLogger := Logger{
		RequestStream: requestStream,
		logs:          make([]string, 0),
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
				switch v := event.(type) {
				case string:
					l.logs = append(l.logs, v)
				case Order:
					stringOrder, err := json.Marshal(&v)
					if err != nil {
						log.Printf("unexpected error marshalling order: %s", err)

						continue
					}

					l.logs = append(l.logs, string(stringOrder))
				default:
					log.Printf("unexpected order %T, with value: %+v", event, event)
				}

				l.NextStream <- event
			}
		}
	}()
}

func (l *Logger) Logs() []string {
	return l.logs
}

func NewValidator(requestStream chan interface{}) *Validator {
	newValidator := Validator{
		RequestStream: requestStream,
	}

	return &newValidator
}

func (v *Validator) WithEnrichFilter(nextStream chan interface{}) *Validator {
	v.EnrichStream = nextStream

	return v
}

func (v *Validator) WithLoggerFilter(nextStream chan interface{}) *Validator {
	v.LoggerStream = nextStream

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
					v.result = append(v.result, false)

					log.Println("new order is invalid", newOrder)

					v.LoggerStream <- "order id or order value is invalid"

					continue
				}

				log.Println("new order is valid", newOrder)

				v.result = append(v.result, true)
				v.EnrichStream <- newOrder
			}
		}
	}()
}

func (v *Validator) Result() []bool {
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

				c.newOrders = append(c.newOrders, newOrder)
				c.NextStream <- newOrder
			}
		}
	}()
}

func (c *Collector) Values() []NewOrder {
	return c.newOrders
}
