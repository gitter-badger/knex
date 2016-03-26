package test

type typeWithInvalidRequiresImpl struct {
	typeWithRequires `provide:"resource"`
	InjectedType     typeWithNoRequires `require:"BadValue"`
}

func newTypeWithInvalidRequiresImpl(injectedType typeWithNoRequires) (*typeWithInvalidRequiresImpl, error) {

	newInstance := new(typeWithInvalidRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject required dependencies
func (t *typeWithInvalidRequiresImpl) Inject(injectedType typeWithNoRequires) error {
	t.InjectedType = injectedType
	return nil
}
