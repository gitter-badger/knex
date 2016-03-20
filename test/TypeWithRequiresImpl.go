package test

type TypeWithRequiresImpl struct {
	TypeWithRequires `provide:"resource"`
	InjectedType     TypeWithNoRequires `require:"true"`
}

func NewTypeWithRequiresImpl(injectedType TypeWithNoRequires) (*TypeWithRequiresImpl, error) {

	newInstance := new(TypeWithRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithRequiresImpl) Inject(injectedType TypeWithNoRequires) error {
	self.InjectedType = injectedType
	return nil
}
