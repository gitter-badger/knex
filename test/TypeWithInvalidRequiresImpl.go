package test

type TypeWithInvalidRequiresImpl struct {
	TypeWithRequires `provide:"resource"`
	InjectedType     TypeWithNoRequires `require:"BadValue"`
}

func NewTypeWithInvalidRequiresImpl(injectedType TypeWithNoRequires) (*TypeWithInvalidRequiresImpl, error) {

	newInstance := new(TypeWithInvalidRequiresImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithInvalidRequiresImpl) Inject(injectedType TypeWithNoRequires) error {
	self.InjectedType = injectedType
	return nil
}
