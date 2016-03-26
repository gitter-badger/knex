package test

type typeWithSliceRequiresImpl struct {
	typeWithRequires `provide:"resource"`
	InjectedType     []typeWithNoRequires `require:"true"`
}

func newTypeWithSliceRequiresImpl(injectedType []typeWithNoRequires) (*typeWithSliceRequiresImpl, error) {

	newInstance := new(typeWithSliceRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject injects required dependencies
func (t *typeWithSliceRequiresImpl) Inject(injectedType []typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
