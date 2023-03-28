// kv package specifies the contract type used in ORM arguments
// and provides implementations of the contract for some popular types.
package kv

// String specifies the contract that key and value types should comply with.
//
// String allows clients to operate with keys and values of their own types.
// For example, some may find reasonable to store complex structs as JSONs.
type String interface {
	// Serialize converts the value into string representation.
	Serialize() string

	// Scan takes raw representation and should create a new value from it.
	Deserialize(raw string) String
}
