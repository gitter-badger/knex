package test

type typeWithNoRequiresOneImpl struct {
	typeWithNoRequires `provide:"resource"`
}

func newTypeWithNoRequiresOneImpl() (*typeWithNoRequiresOneImpl, error) {

	newInstance := new(typeWithNoRequiresOneImpl)

	return newInstance, newInstance.Inject()
}

// Inject injects required dependencies
func (t *typeWithNoRequiresOneImpl) Inject() error {
	return nil
}
