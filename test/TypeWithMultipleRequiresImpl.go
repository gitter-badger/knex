package test

type TypeWithMultipleRequiresImpl struct {
	TypeWithRequires `provide:"resource"`
	InjectedTypeOne  TypeWithNoRequires `require:"true"`
	InjectedTypeTwo  TypeWithNoRequires `require:"true"`
}

func NewTypeWithMultipleRequiresImpl(injectedTypeOne TypeWithNoRequires, injectedTypeTwo TypeWithNoRequires) (*TypeWithMultipleRequiresImpl, error) {

	newInstance := new(TypeWithMultipleRequiresImpl)

	return newInstance, newInstance.Inject(injectedTypeOne, injectedTypeTwo)
}

func (self *TypeWithMultipleRequiresImpl) Inject(injectedTypeOne TypeWithNoRequires, injectedTypeTwo TypeWithNoRequires) error {
	self.InjectedTypeOne = injectedTypeOne
	self.InjectedTypeTwo = injectedTypeTwo
	return nil
}
