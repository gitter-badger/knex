package test

type typeWithIDImpl struct {
	typeWithNoRequires `provide:"resource" id:"testId"`
}

func newTypeWithIDImpl(injectedType typeWithNoRequires) (*typeWithIDImpl, error) {

	newInstance := new(typeWithIDImpl)

	return newInstance, newInstance.Inject()
}

// Inject required dependencies
func (t *typeWithIDImpl) Inject() error {
	return nil
}
