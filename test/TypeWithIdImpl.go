package test

type typeWithIdImpl struct {
	TypeWithNoRequires `provide:"resource" id:"testId"`
}

func NewTypeWithIdImpl(injectedType TypeWithNoRequires) (*typeWithIdImpl, error) {

	newInstance := new(typeWithIdImpl)

	return newInstance, newInstance.Inject()
}

func (self *typeWithIdImpl) Inject() error {
	return nil
}
