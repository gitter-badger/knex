package knex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type implementationDetail struct {
	source         int
	resourceDetail resourceDetail
	implType       reflect.Type
	injector       reflect.Method
	fieldSlice     []reflect.StructField
	getResource    func(reflect.StructField, *typeSet, map[interface{}]reflect.Value) []reflect.Value
}

func newImplementationDetail(implementationType interface{}, getResourceFunc func(reflect.StructField, *typeSet, map[interface{}]reflect.Value) []reflect.Value) (*implementationDetail, error) {

	implDetail := &implementationDetail{
		source:      implementationSource,
		implType:    reflect.TypeOf(implementationType),
		getResource: getResourceFunc,
	}

	// Add required fields to ImplementationDetail.
	injectableFields, err := implDetail.getInjectableFields(implementationType)
	if err != nil {
		return nil, err
	}
	implDetail.fieldSlice = injectableFields

	// Add Resource details to ImplementationDetail.
	resourceDetail, err := newResourceDetail(implementationType)
	if err != nil {
		return nil, err
	}
	implDetail.resourceDetail = *resourceDetail

	// Add Injector function to ImplementationDetail.
	implDetail.injector = implDetail.getInjector(implementationType)

	return implDetail, nil
}

func newImplementationDetailByProvider(provider Provider) (*implementationDetail, error) {

	implDetail := &implementationDetail{
		source: providerSource,
	}

	// Add Resource details to ImplementationDetail.
	resourceDetail, err := newResourceDetailByProvider(&provider)
	if err != nil {
		return nil, err
	}
	implDetail.resourceDetail = *resourceDetail

	return implDetail, nil
}

func (self *implementationDetail) callInjector(typeSet *typeSet, graphScopeMap map[interface{}]reflect.Value) []reflect.Value {

	// Create new instance of implementation.
	newInstance := reflect.New(self.implType.Elem())

	// Get list of arguments to pass into injector method.
	arguments := []reflect.Value{newInstance}
	for _, field := range self.fieldSlice {
		resourceResult := self.getResource(field, typeSet, graphScopeMap)
		if !resourceResult[1].IsNil() {
			return resourceResult
		} else {
			arguments = append(arguments, resourceResult[0])
		}
	}

	// Call injector method.
	injectResult := self.injector.Func.Call(arguments)
	if !injectResult[0].IsNil() {
		return []reflect.Value{reflect.Zero(self.implType), injectResult[0]}
	}

	// Return implementation.
	return []reflect.Value{newInstance, reflect.Zero(reflect.TypeOf(errors.New("")))}
}

func (self *implementationDetail) GetImplType() reflect.Type {
	return self.implType
}

func (self *implementationDetail) HasInjector() bool {
	return self.injector == (reflect.Method{})
}

func (self *implementationDetail) getInjectableFields(implementationType interface{}) ([]reflect.StructField, error) {

	// Find all injectable fields.
	var returnValue []reflect.StructField
	reflectType := reflect.TypeOf(implementationType).Elem()
	fieldCount := reflectType.NumField()
	for i := 0; i < fieldCount; i++ {

		// Check if field has a valid 'require' tag.
		field := reflectType.Field(i)
		requireTagValue := strings.ToUpper(strings.Trim(field.Tag.Get(requireTagName), " "))
		if requireTagValue != emptyString {
			if requireTagValue == trueValue || requireTagValue == falseValue {
				returnValue = append(returnValue, field)
			} else {
				return nil, fmt.Errorf("Invalid require value '%s'", requireTagValue)
			}
		}
	}
	return returnValue, nil
}

func (self *implementationDetail) getInjector(implementationType interface{}) reflect.Method {

	// Check if implementation has an injector method.
	reflectType := reflect.TypeOf(implementationType)
	implementationInjector, hasInjector := reflectType.MethodByName(injectorName)
	if hasInjector {
		return implementationInjector
	} else {
		return reflect.Method{}
	}
}
