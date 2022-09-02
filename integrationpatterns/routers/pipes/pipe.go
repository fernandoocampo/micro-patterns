package pipes

// NewPipe creates a new event channel pipe.
func NewPipe() chan interface{} {
	return make(chan interface{})
}
