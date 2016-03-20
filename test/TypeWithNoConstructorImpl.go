package test

type TypeWithNoInjectorImpl struct {
	TypeWithNoRequires `provide:"resource"`
}
