package test

type typeWithIDImpl struct {
	TypeWithNoRequires `provide:"resource" id:"testId"`
}

func newTypeWithIDImpl(injectedType TypeWithNoRequires) (*typeWithIDImpl, error) {

	newInstance := new(typeWithIDImpl)

	return newInstance, newInstance.Inject()
}

func (self *typeWithIDImpl) Inject() error {
	return nil
}
