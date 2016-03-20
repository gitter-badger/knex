package test

type TypeWithNoRequiresOneImpl struct {
	TypeWithNoRequires `provide:"resource"`
}

func NewTypeWithNoRequiresOneImpl() (*TypeWithNoRequiresOneImpl, error) {

	newInstance := new(TypeWithNoRequiresOneImpl)

	return newInstance, newInstance.Inject()
}

func (self *TypeWithNoRequiresOneImpl) Inject() error {
	return nil
}
