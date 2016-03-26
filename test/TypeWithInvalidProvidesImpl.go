package test

type TypeWithInvalidProvidesImpl struct {
	typeWithRequires `provide:"BadValue"`
}

func NewTypeWithInvalidProvidesImpl() (*TypeWithInvalidProvidesImpl, error) {

	newInstance := new(TypeWithInvalidProvidesImpl)

	return newInstance, newInstance.Inject()
}

func (self *TypeWithInvalidProvidesImpl) Inject() error {
	return nil
}
