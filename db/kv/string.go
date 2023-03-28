package kv

type Str string

// Validate contract in compile-time.
var _ String = Str("")

func (s Str) Serialize() string {
	return string(s)
}

func (Str) Deserialize(raw string) String {
	return Str(raw)
}
