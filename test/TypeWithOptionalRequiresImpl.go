package test

type TypeWithOptionalRequiresImpl struct {
	TypeWithRequires `provide:"resource"`
	InjectedType     TypeWithNoRequires `require:"false"`
}

func NewTypeWithOptionalRequiresImpl(injectedType TypeWithNoRequires) (*TypeWithOptionalRequiresImpl, error) {

	newInstance := new(TypeWithOptionalRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithOptionalRequiresImpl) Inject(injectedType TypeWithNoRequires) error {
	self.InjectedType = injectedType
	return nil
}
