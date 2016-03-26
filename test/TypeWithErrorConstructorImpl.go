package test

import "errors"

type typeWithErrorInjectorImpl struct {
	typeWithNoRequires `provide:"resource"`
}

func newTypeWithErrorInjectorImpl() (*typeWithErrorInjectorImpl, error) {

	newInstance := new(typeWithErrorInjectorImpl)

	return nil, newInstance.Inject()
}

// Inject required dependencies
func (t *typeWithErrorInjectorImpl) Inject() error {
	return errors.New("Test error")
}
