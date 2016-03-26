package test

type typeWithNoScopeImpl struct {
	TypeWithNoRequires `provide:"resource"`
	Value              string
}

func newTypeWithNoScopeImpl() (*typeWithNoScopeImpl, error) {

	newInstance := new(typeWithNoScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

func (t *typeWithNoScopeImpl) Inject() error {
	return nil
}
