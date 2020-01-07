package order_status

// OrderStatus is a type with underlying type of int.
type OrderStatus int

const (
	New OrderStatus = iota
	Preparing
	Ready
	Completed
	Cancelled
)
