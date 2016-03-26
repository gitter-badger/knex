package test

type typeWithInvalidProvidesImpl struct {
	typeWithRequires `provide:"BadValue"`
}

func newTypeWithInvalidProvidesImpl() (*typeWithInvalidProvidesImpl, error) {

	newInstance := new(typeWithInvalidProvidesImpl)

	return newInstance, newInstance.Inject()
}

// Inject required dependencies
func (t *typeWithInvalidProvidesImpl) Inject() error {
	return nil
}
