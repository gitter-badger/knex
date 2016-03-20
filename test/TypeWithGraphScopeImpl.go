package test

type TypeWithGraphScopeImpl struct {
	TypeWithNoRequires `provide:"resource" scope:"graph"`
	Value              string
}

func NewTypeWithGraphScopeImpl() (*TypeWithGraphScopeImpl, error) {

	newInstance := new(TypeWithGraphScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

func (self *TypeWithGraphScopeImpl) Inject() error {
	return nil
}
