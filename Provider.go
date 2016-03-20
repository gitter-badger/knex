package knex

type Provider struct {
	Type     interface{}
	Id       string
	Scope    string
	Instance func() (interface{}, error)
}
