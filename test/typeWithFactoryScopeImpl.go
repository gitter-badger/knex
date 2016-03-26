package test

type typeWithFactoryScopeImpl struct {
	typeWithNoRequires `provide:"resource" scope:"factory"`
	Value              string
}

func newTypeWithFactoryScopeImpl() (*typeWithFactoryScopeImpl, error) {

	newInstance := new(typeWithFactoryScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

// Inject required dependencies
func (t *typeWithFactoryScopeImpl) Inject() error {
	return nil
}
