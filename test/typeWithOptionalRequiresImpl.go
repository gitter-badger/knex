package test

type typeWithOptionalRequiresImpl struct {
	typeWithRequires `provide:"resource"`
	InjectedType     typeWithNoRequires `require:"false"`
}

func newTypeWithOptionalRequiresImpl(injectedType typeWithNoRequires) (*typeWithOptionalRequiresImpl, error) {

	newInstance := new(typeWithOptionalRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject injects required dependencies
func (t *typeWithOptionalRequiresImpl) Inject(injectedType typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
