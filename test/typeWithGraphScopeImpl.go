package test

type typeWithGraphScopeImpl struct {
	typeWithNoRequires `provide:"resource" scope:"graph"`
	Value              string
}

func newTypeWithGraphScopeImpl() (*typeWithGraphScopeImpl, error) {

	newInstance := new(typeWithGraphScopeImpl)
	newInstance.Value = "Initial value"

	return newInstance, newInstance.Inject()
}

// Inject required dependencies
func (t *typeWithGraphScopeImpl) Inject() error {
	return nil
}
