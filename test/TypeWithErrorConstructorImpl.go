package test

import "errors"

type TypeWithErrorInjectorImpl struct {
	TypeWithNoRequires `provide:"resource"`
}

func NewTypeWithErrorInjectorImpl() (*TypeWithErrorInjectorImpl, error) {

	newInstance := new(TypeWithErrorInjectorImpl)

	return nil, newInstance.Inject()
}

func (self *TypeWithErrorInjectorImpl) Inject() error {
	return errors.New("Test error")
}
