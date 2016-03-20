package test

type TypeWithFactoryScopeImpl struct {
	TypeWithNoRequires `provide:"resource" scope:"factory"`
	Value              string
}

func NewTypeWithFactoryScopeImpl() (*TypeWithFactoryScopeImpl, error) {

	newInstance := new(TypeWithFactoryScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

func (self *TypeWithFactoryScopeImpl) Inject() error {
	return nil
}
