package test

type TypeWithSliceRequiresImpl struct {
	TypeWithRequires `provide:"resource"`
	InjectedType     []TypeWithNoRequires `require:"true"`
}

func NewTypeWithSliceRequiresImpl(injectedType []TypeWithNoRequires) (*TypeWithSliceRequiresImpl, error) {

	newInstance := new(TypeWithSliceRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithSliceRequiresImpl) Inject(injectedType []TypeWithNoRequires) error {
	self.InjectedType = injectedType
	return nil
}
