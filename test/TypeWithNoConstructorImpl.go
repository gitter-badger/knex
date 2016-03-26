package test

type typeWithNoInjectorImpl struct {
	typeWithNoRequires `provide:"resource"`
}
