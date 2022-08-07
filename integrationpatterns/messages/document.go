package messages

// Document contains data and lets the receiver decide what, if anything, to do with the data
type Document struct {
	OrderID string
	Value   float64
}
