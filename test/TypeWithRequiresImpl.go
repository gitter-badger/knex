package test

type typeWithRequiresImpl struct {
	typeWithRequires `provide:"resource"`
	InjectedType     typeWithNoRequires `require:"true"`
}

func newTypeWithRequiresImpl(injectedType typeWithNoRequires) (*typeWithRequiresImpl, error) {

	newInstance := new(typeWithRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject injects required dependencies
func (t *typeWithRequiresImpl) Inject(injectedType typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
