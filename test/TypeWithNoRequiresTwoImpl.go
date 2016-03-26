package test

type typeWithNoRequiresTwoImpl struct {
	typeWithNoRequires `provide:"resource"`
}

func newTypeWithNoRequiresTwoImpl() (*typeWithNoRequiresTwoImpl, error) {

	newInstance := new(typeWithNoRequiresTwoImpl)

	return newInstance, newInstance.Inject()
}

// Inject injects required dependencies
func (t *typeWithNoRequiresTwoImpl) Inject() error {
	return nil
}
