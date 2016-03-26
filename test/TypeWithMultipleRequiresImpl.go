package test

type typeWithMultipleRequiresImpl struct {
	typeWithRequires `provide:"resource"`
	InjectedTypeOne  typeWithNoRequires `require:"true"`
	InjectedTypeTwo  typeWithNoRequires `require:"true"`
}

func newTypeWithMultipleRequiresImpl(injectedTypeOne typeWithNoRequires, injectedTypeTwo typeWithNoRequires) (*typeWithMultipleRequiresImpl, error) {

	newInstance := new(typeWithMultipleRequiresImpl)

	return newInstance, newInstance.Inject(injectedTypeOne, injectedTypeTwo)
}

// Inject required dependencies
func (t *typeWithMultipleRequiresImpl) Inject(injectedTypeOne typeWithNoRequires, injectedTypeTwo typeWithNoRequires) error {
	t.InjectedTypeOne = injectedTypeOne
	t.InjectedTypeTwo = injectedTypeTwo
	return nil
}
