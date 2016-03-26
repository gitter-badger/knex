package test

type typeWithCircularDependencyImpl struct {
	typeWithCircularDependency `provide:"resource"`
	injectedType               typeWithCircularDependency `require:"true"`
}

func newTypeWithCircularDependencyImpl(injectedType typeWithNoRequires) (*typeWithCircularDependencyImpl, error) {

	newInstance := new(typeWithCircularDependencyImpl)

	return newInstance, newInstance.Inject(injectedType)
}

// Inject required dependencies
func (t *typeWithCircularDependencyImpl) Inject(injectedType typeWithCircularDependency) error {
	t.injectedType = injectedType
	return nil
}
