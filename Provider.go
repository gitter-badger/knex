package knex

// Provider defines a custom constructor method for type implementations.
type Provider struct {
	Type     interface{}
	ID       string
	Scope    string
	Instance func() (interface{}, error)
}
