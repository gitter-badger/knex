package test

type typeWithNoScopeImpl struct {
	typeWithNoRequires `provide:"resource"`
	Value              string
}

func newTypeWithNoScopeImpl() (*typeWithNoScopeImpl, error) {

	newInstance := new(typeWithNoScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

// Inject injects required dependencies
func (t *typeWithNoScopeImpl) Inject() error {
	return nil
}
