package test

type TypeWithNoRequiresTwoImpl struct {
	TypeWithNoRequires `provide:"resource"`
}

func NewTypeWithNoRequiresTwoImpl() (*TypeWithNoRequiresTwoImpl, error) {

	newInstance := new(TypeWithNoRequiresTwoImpl)

	return newInstance, newInstance.Inject()
}

func (self *TypeWithNoRequiresTwoImpl) Inject() error {
	return nil
}
