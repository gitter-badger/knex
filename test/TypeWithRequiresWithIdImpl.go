package test

type TypeWithRequiresWithIdImpl struct {
	typeWithRequiresWithId `provide:"resource"`
	InjectedType           TypeWithNoRequires `require:"true" id:"testId"`
}

func NewTypeWithRequiresWithIdImpl(injectedType TypeWithNoRequires) (*TypeWithRequiresWithIdImpl, error) {

	newInstance := new(TypeWithRequiresWithIdImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithRequiresWithIdImpl) Inject(injectedType TypeWithNoRequires) error {
	self.InjectedType = injectedType
	return nil
}
