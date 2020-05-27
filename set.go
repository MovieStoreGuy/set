package set

// Set stores all the item as a distinct values
// to allow for easy comparison
type Set interface {

	// Insert will store item internally
	// panics when an item is an array, slice or map
	Insert(item interface{})

	// HasItem will check to see item exists within the set
	HasItem(item interface{}) bool

	// Items returns all the stored inside the set
	Items() []interface{}

	// Size returns the number of items stored within the set
	Size() int
}
