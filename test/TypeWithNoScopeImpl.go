package test

type TypeWithNoScopeImpl struct {
	TypeWithNoRequires `provide:"resource"`
	Value              string
}

func NewTypeWithNoScopeImpl() (*TypeWithNoScopeImpl, error) {

	newInstance := new(TypeWithNoScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

func (self *TypeWithNoScopeImpl) Inject() error {
	return nil
}
