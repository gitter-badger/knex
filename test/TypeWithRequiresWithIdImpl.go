package test

type typeWithRequiresWithIdImpl struct {
	typeWithRequiresWithId `provide:"resource"`
	InjectedType           typeWithNoRequires `require:"true" id:"testId"`
}

func newTypeWithRequiresWithIdImpl(injectedType typeWithNoRequires) (*typeWithRequiresWithIdImpl, error) {

	newInstance := new(typeWithRequiresWithIdImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject injects required dependencies
func (t *typeWithRequiresWithIdImpl) Inject(injectedType typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
