package pipes

// NewOrder contains data for a new order.
type NewOrder struct {
	ID    string
	Value float32
}

// Order contains data related to an order.
type Order struct {
	ID         string  `json:"id"`
	InternalID string  `json:"internal_id"`
	Value      float32 `json:"value"`
	Success    bool    `json:"success,omitempty"`
}
