package test

type TypeWithCircularDependencyImpl struct {
	TypeWithCircularDependency `provide:"resource"`
	injectedType               TypeWithCircularDependency `require:"true"`
}

func NewTypeWithCircularDependencyImpl(injectedType TypeWithNoRequires) (*TypeWithCircularDependencyImpl, error) {

	newInstance := new(TypeWithCircularDependencyImpl)

	return newInstance, newInstance.Inject(injectedType)
}

func (self *TypeWithCircularDependencyImpl) Inject(injectedType TypeWithCircularDependency) error {
	self.injectedType = injectedType
	return nil
}
