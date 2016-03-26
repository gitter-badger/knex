package test

type typeWithNoRequiresOneImpl struct {
	TypeWithNoRequires `provide:"resource"`
}

func newTypeWithNoRequiresOneImpl() (*typeWithNoRequiresOneImpl, error) {

	newInstance := new(typeWithNoRequiresOneImpl)

	return newInstance, newInstance.Inject()
}

// Inject injects required dependencies
func (self *typeWithNoRequiresOneImpl) Inject() error {
	return nil
}
