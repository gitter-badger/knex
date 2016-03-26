package test

type typeWithRequiresWithIDImpl struct {
	typeWithRequiresWithID `provide:"resource"`
	InjectedType           typeWithNoRequires `require:"true" id:"testId"`
}

func newTypeWithRequiresWithIDImpl(injectedType typeWithNoRequires) (*typeWithRequiresWithIDImpl, error) {

	newInstance := new(typeWithRequiresWithIDImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject injects required dependencies
func (t *typeWithRequiresWithIDImpl) Inject(injectedType typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
