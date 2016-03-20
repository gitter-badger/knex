package test

type TypeWithIdImpl struct {
	TypeWithNoRequires `provide:"resource" id:"testId"`
}

func NewTypeWithIdImpl(injectedType TypeWithNoRequires) (*TypeWithIdImpl, error) {

	newInstance := new(TypeWithIdImpl)

	return newInstance, newInstance.Inject()
}

func (self *TypeWithIdImpl) Inject() error {
	return nil
}
