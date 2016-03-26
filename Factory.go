package knex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var DefaultFactory = NewFactory()

type Factory struct {
	factoryScopeMap map[interface{}]reflect.Value
	idMap           map[string]*implementationDetail
	multipleTypeMap map[reflect.Type][]*implementationDetail
	parentSlice     []*Factory
	typeMap         map[reflect.Type]*implementationDetail
}

// NewFactory creates a new Factory struct.
func NewFactory() *Factory {
	return &Factory{
		factoryScopeMap: make(map[interface{}]reflect.Value),
		idMap:           make(map[string]*implementationDetail),
		multipleTypeMap: make(map[reflect.Type][]*implementationDetail),
		parentSlice:     make([]*Factory, 0),
		typeMap:         make(map[reflect.Type]*implementationDetail),
	}
}

// AddParent adds a parent factory.  If there is a circular dependency return
// a circular dependency error.
func (self *Factory) AddParent(parent *Factory) error {

	// If there is a circular dependency return an error.
	if parent.containsParent(self) {
		parentType := reflect.TypeOf(parent)
		return fmt.Errorf("Circular dependency detected with '%s/%s'", parentType.Elem().PkgPath(), parentType.Elem().Name())
	}

	self.parentSlice = append(self.parentSlice, parent)

	return nil
}

// GetAllOfType gets all implementations for the provided 'interfaceType'.  If
// there are no implementations it returns an empty slice.  If there is only
// one implementation it returns a slice with the one value.  Otherwise it
// returns a slice with all registered implementations.
func (self *Factory) GetAllOfType(interfaceType interface{}) (interface{}, error) {

	// Get the reflect.Type of the given type.
	reflectType := self.getReflectType(interfaceType)

	// If there are multiple implementations return a slice that contains each
	// implementation.
	implDetailSlice, exists := self.multipleTypeMap[reflectType]
	if exists {
		result := self.getAllByReflectTypeAndImplSlice(reflectType, implDetailSlice)
		err := self.valueToError(result[1])
		if err != nil {
			return nil, err
		}
		return self.valueToInterface(result[0]), nil
	}

	// If there is only one implementation return a slice that contains the one
	// implementation.
	implDetail, exists := self.typeMap[reflectType]
	if exists {
		result := self.getAllByReflectTypeAndImplDetail(reflectType, implDetail)
		err := self.valueToError(result[1])
		if err != nil {
			return nil, err
		}
		return self.valueToInterface(result[0]), nil
	}

	// Check if any of this factories' parents has the type.
	for _, parent := range self.parentSlice {

		// Check if parent has implementation(s) or propagate any error.
		result, err := parent.GetAllOfType(interfaceType)
		reflectSlice := reflect.ValueOf(result)
		if err != nil || reflectSlice.Len() > 0 {
			return result, err
		}
	}

	// If there are no implementations return an empty slice.
	return self.valueToInterface(reflect.MakeSlice(reflect.SliceOf(reflectType), 0, 0)), nil
}

// GetById gets an implementation based on the provided 'id'. If an
// implementation has not been registerd for the 'id' it returns an error.
// Otherwise it returns the implementation registered to 'id'.
func (self *Factory) GetById(id string) (interface{}, error) {

	// Instanciate implementation.
	result := self.getReflectValueById(id)
	err := self.valueToError(result[1])
	if err != nil {
		return nil, err
	}

	return self.valueToInterface(result[0]), nil
}

// GetByType gets an implementation based on the provided 'interfaceType'. If
// an implementation has not been registerd for the 'interfaceType then it
// returns an error.  If multiple implementations have been registerd for the
// 'interfaceType' then it returns an error. Otherwise it returns the one
// implementaion.
func (self *Factory) GetByType(interfaceType interface{}) (interface{}, error) {

	// Get the reflect.Type of the given type.
	reflectType := self.getReflectType(interfaceType)

	// There can only be one implementation for the given type,  if there is more
	// return an error.
	_, exists := self.multipleTypeMap[reflectType]
	if exists {
		return nil, fmt.Errorf("Multiple implementations for type '%s/%s' declared", reflectType.PkgPath(), reflectType.Name())
	}

	return self.getByReflectType(reflectType)
}

// Register adds an implementation to the factory.  If the implementation is
// improperly tagged it will return an error.
func (self *Factory) Register(implementationType interface{}) error {

	// Get implementation meta data
	implDetail, err := newImplementationDetail(implementationType, self.getByField)
	if err != nil {
		return err
	}

	// Register implementation based on its type.
	self.registerImplWithType(implDetail)

	// Register implementation based on its id.
	self.registerImplWithId(implDetail)

	return nil
}

// Register adds a provider to the factory.  If the provider is
// improperly defined it will return an error.
func (self *Factory) RegisterProvider(provider Provider) error {

	// Get implementation meta data
	implDetail, err := newImplementationDetailByProvider(provider)
	if err != nil {
		return err
	}

	// Register implementation based on its type.
	self.registerImplWithType(implDetail)

	// Register implementation based on its id.
	self.registerImplWithId(implDetail)

	return nil
}

func (self *Factory) containsParent(factory *Factory) bool {

	// Recursively checks if factroy is related.
	for _, parent := range self.parentSlice {
		if parent == factory {
			return true
		}
		if parent.containsParent(factory) {
			return true
		}
	}

	return false
}

func (self *Factory) errorValue(err error) []reflect.Value {

	// Convert error into reflect.Value slice.
	return []reflect.Value{
		reflect.ValueOf(nil),
		reflect.ValueOf(err),
	}
}

func (self *Factory) getAllByReflectTypeAndImplDetail(reflectType reflect.Type, implDetail *implementationDetail) []reflect.Value {

	// Get implementation.
	result := self.getByImplDetail(implDetail, newTypeSet(), make(map[interface{}]reflect.Value))
	err := self.valueToError(result[1])
	if err != nil {
		return result
	}

	// Create a one element slice containting the implementation above.
	reflectSlice := reflect.MakeSlice(reflect.SliceOf(reflectType), 0, 0)
	reflectSlice = reflect.Append(reflectSlice, result[0])
	return []reflect.Value{
		reflectSlice,
		self.nilErrorValue(),
	}
}

func (self *Factory) getAllByReflectTypeAndImplSlice(reflectType reflect.Type, implDetailSlice []*implementationDetail) []reflect.Value {

	// Build a slice with all registered implementations.
	reflectSlice := reflect.MakeSlice(reflect.SliceOf(reflectType), 0, len(implDetailSlice))
	for _, currImplDetail := range implDetailSlice {

		// Get implementation and add to slice, if unable to create instance return
		// and error.
		result := self.getByImplDetail(currImplDetail, newTypeSet(), make(map[interface{}]reflect.Value))
		err := self.valueToError(result[1])
		if err != nil {
			return result
		}
		reflectSlice = reflect.Append(reflectSlice, result[0])
	}

	// Return all implementations.
	return []reflect.Value{
		reflectSlice,
		self.nilErrorValue(),
	}
}

func (self *Factory) getByField(field reflect.StructField, typeSet *typeSet, graphScopeMap map[interface{}]reflect.Value) []reflect.Value {

	// Get implementation based on id tag.
	id := field.Tag.Get("id")
	if strings.Trim(id, " ") != "" {
		return self.getReflectValueById(field.Tag.Get("id"))
	}

	// Get reflect.Type regardless regardless if the field is a slice or not.
	reflectType := self.getFieldReflectType(field)

	// Check if there are multiple implementations.
	implDetailSlice, exists := self.multipleTypeMap[reflectType]
	if exists {

		// If the field is a slice then set the field, otherwise return an error
		if field.Type.Kind() == reflect.Slice {
			return self.getAllByReflectTypeAndImplSlice(reflectType, implDetailSlice)
		} else {
			return self.errorValue(fmt.Errorf("Multiple implementations for type '%s/%s' declared", reflectType.PkgPath(), reflectType.Name()))
		}
	}

	// Check if there is one implementation.
	implDetail, exists := self.typeMap[reflectType]
	if exists {

		// If the field is a slice then set the field as a slice otherwise set the field.
		if reflectType.Kind() == reflect.Slice {
			return self.getAllByReflectTypeAndImplDetail(reflectType, implDetail)
		} else {
			return self.getByImplDetail(implDetail, typeSet, graphScopeMap)
		}
	}

	// Check if any of this factories' parents has the type.
	for _, parent := range self.parentSlice {

		// Check if parent has implementation(s) or propagate any error.
		reflectResult := parent.getByField(field, typeSet, graphScopeMap)
		err := self.valueToError(reflectResult[1])
		if err == nil || !strings.HasPrefix(err.Error(), "Undeclared resource") {
			return reflectResult
		}
	}

	// An implementation has not been declared for this field type.
	return self.getUndeclaredField(field)
}

func (self *Factory) getByImplDetail(implDetail *implementationDetail, typeSet *typeSet, graphScopeMap map[interface{}]reflect.Value) []reflect.Value {

	// If there is an implementation available within scope return it.
	reuseValue, exists := self.getScopeImpl(implDetail, graphScopeMap)
	if exists {
		return []reflect.Value{
			reuseValue,
			self.nilErrorValue(),
		}
	}

	// Get the reflect.Type of the given implementation.
	reflectType := implDetail.resourceDetail.interfaceType

	// Get implementation based on struct with taged fields.
	if implDetail.source == implementationSource {

		// If implementation does not have a injector return an error.
		if implDetail.HasInjector() {
			return self.errorValue(fmt.Errorf("Resource '%s/%s' missing injector", reflectType.PkgPath(), reflectType.Name()))
		}

		// If there is a circular dependency return an error.
		implType := implDetail.GetImplType()
		if typeSet.get(implType) {
			return self.errorValue(fmt.Errorf("Circular dependency detected with '%s/%s'", implType.Elem().PkgPath(), implType.Elem().Name()))
		}

		// Call injector method.
		typeSet.add(implType)
		injectorResult := implDetail.callInjector(typeSet, graphScopeMap)
		err := self.valueToError(injectorResult[1])
		if err == nil {

			// Add resource to Factory or Graph scope if necessary.
			if implDetail.resourceDetail.provider.Scope == "FACTORY" {
				self.factoryScopeMap[implType] = injectorResult[0]
			} else if implDetail.resourceDetail.provider.Scope == "GRAPH" {
				graphScopeMap[implType] = injectorResult[0]
			}
		}
		typeSet.remove(implType)

		return injectorResult
	}

	// Get implementation based on Provider function.
	if implDetail.source == providerSource {

		// Call custom provider instance method.
		newInstance, err := implDetail.resourceDetail.provider.Instance()
		if err != nil {
			return self.errorValue(err)
		}
		reflectValue := reflect.ValueOf(newInstance)

		// Add resource to Factory or Graph scope if necessary.
		if implDetail.resourceDetail.provider.Scope == "FACTORY" {
			self.factoryScopeMap[&implDetail.resourceDetail.provider.Instance] = reflectValue
		} else if implDetail.resourceDetail.provider.Scope == "GRAPH" {
			graphScopeMap[&implDetail.resourceDetail.provider.Instance] = reflectValue
		}

		return []reflect.Value{
			reflectValue,
			self.nilErrorValue(),
		}
	}

	return self.errorValue(fmt.Errorf("Resource '%s/%s' has unknown source", reflectType.PkgPath(), reflectType.Name()))
}

func (self *Factory) getByReflectType(reflectType reflect.Type) (interface{}, error) {

	// Check if type has an implementation registered for it.
	implDetail, exists := self.typeMap[reflectType]
	if !exists {

		// Check if any of this factories' parents has the type.
		for _, parent := range self.parentSlice {

			// Check if parent has an implementation, propagate any error except for
			// "Undeclared resource...", otherwise move on to next parent.
			impl, err := parent.getByReflectType(reflectType)
			if err == nil {
				return impl, nil
			} else if !strings.HasPrefix(err.Error(), "Undeclared resource") {
				return nil, err
			}
		}

		return nil, fmt.Errorf("Undeclared resource '%s/%s'", reflectType.PkgPath(), reflectType.Name())
	}

	// Get implementation.
	result := self.getByImplDetail(implDetail, newTypeSet(), make(map[interface{}]reflect.Value))
	err := self.valueToError(result[1])
	if err != nil {
		return nil, err
	}

	// Convert reflect.Value to interface{}
	return self.valueToInterface(result[0]), nil
}

func (self *Factory) getFieldReflectType(field reflect.StructField) reflect.Type {

	// If the field type is a slice get the slices' element type.
	reflectType := field.Type
	if reflectType.Kind() == reflect.Slice {
		reflectType = field.Type.Elem()
	}

	return reflectType
}

func (self *Factory) getReflectType(interfaceType interface{}) reflect.Type {

	// Get reflect.Type of the given interfaceType.
	return reflect.TypeOf(interfaceType).Elem()
}

func (self *Factory) getReflectValueById(id string) []reflect.Value {

	// Check if there is an impementation for the given id.
	implDetail, exists := self.idMap[id]
	if !exists {

		// Check if any of this factories' parents has the type.
		for _, parent := range self.parentSlice {

			// Check if parent has an implementation, propagate any error except for
			// "Undeclared resource...", otherwise move on to next parent.
			reflectResult := parent.getReflectValueById(id)
			err := self.valueToError(reflectResult[1])
			if err == nil || !strings.HasPrefix(err.Error(), "Undeclared resource") {
				return reflectResult
			}
		}

		return self.errorValue(fmt.Errorf("Undeclared resource with id '%s'", id))
	}

	// Instanciate implementation.
	return self.getByImplDetail(implDetail, newTypeSet(), make(map[interface{}]reflect.Value))
}

func (self *Factory) getScopeImpl(implDetail *implementationDetail, graphScopeMap map[interface{}]reflect.Value) (reflect.Value, bool) {

	// Check factory scope map and then graph scope map for existing
	// implementations. If one exists return it.
	var scopeKey = self.getScopeKey(implDetail)
	var reuseValue reflect.Value
	var exists = false
	if implDetail.resourceDetail.provider.Scope == "FACTORY" {
		reuseValue, exists = self.factoryScopeMap[scopeKey]
	} else if implDetail.resourceDetail.provider.Scope == "GRAPH" {
		reuseValue, exists = graphScopeMap[scopeKey]
	}
	return reuseValue, exists
}

func (self *Factory) getScopeKey(implDetail *implementationDetail) interface{} {

	// Get scope key based on source of implementation detail.  If source is based
	// on tag value then use the implementations' reflect.Tag value, otherwise use
	// a pointer to the Instace function from the Provider. If source is not a
	// valid value return nil.
	if implDetail.source == implementationSource {
		return implDetail.GetImplType()
	} else if implDetail.source == providerSource {
		return &implDetail.resourceDetail.provider.Instance
	} else {
		return nil
	}
}

func (self *Factory) getUndeclaredField(field reflect.StructField) []reflect.Value {

	// Get reflect.Type regardless regardless if the field is a slice or not.
	reflectType := self.getFieldReflectType(field)

	// If field is a slice return empty slice.
	if field.Type.Kind() == reflect.Slice {
		return []reflect.Value{
			reflect.MakeSlice(reflect.SliceOf(reflectType), 0, 0),
			self.nilErrorValue(),
		}
	}

	// If field is required return error.
	requireTagValue := strings.ToUpper(strings.Trim(field.Tag.Get(requireTagName), " "))
	if requireTagValue == "TRUE" {
		return self.errorValue(fmt.Errorf("Undeclared resource '%s/%s'", reflectType.PkgPath(), reflectType.Name()))
	}

	// If field is  not required return zero value.
	return []reflect.Value{
		reflect.Zero(reflectType),
		self.nilErrorValue(),
	}
}

func (self *Factory) nilErrorValue() reflect.Value {

	// Return a reflect.Value that represents a nil error.
	return reflect.Zero(reflect.TypeOf(errors.New("")))
}

func (self *Factory) registerImplWithId(implDetail *implementationDetail) {

	// Add implemetaion based on id tag value.
	id := implDetail.resourceDetail.provider.Id
	if id != "" {
		self.idMap[id] = implDetail
	}
}

func (self *Factory) registerImplWithType(implDetail *implementationDetail) {

	// Get the reflect.Type of the given type.
	reflectType := implDetail.resourceDetail.interfaceType

	// If there is multiple implementations of the type add implementation to slice.
	impDetailSlice, exists := self.multipleTypeMap[reflectType]
	if exists {
		self.multipleTypeMap[reflectType] = append(impDetailSlice, implDetail)
		return
	}

	// If there is only one implementation create the slice and add implementation.
	existingImplDetail, exists := self.typeMap[reflectType]
	if exists {
		self.multipleTypeMap[reflectType] = []*implementationDetail{existingImplDetail, implDetail}
		delete(self.typeMap, reflectType)
		return
	}

	// There are no implementations so just add it.
	self.typeMap[reflectType] = implDetail
	return
}

func (self *Factory) valueToError(value reflect.Value) error {

	// Convert a reflect.Value to an error.
	if value.IsValid() && !value.IsNil() {
		return value.Interface().(error)
	} else {
		return nil
	}
}

func (self *Factory) valueToInterface(value reflect.Value) interface{} {

	// Convert a reflect.Value to an interface{}.
	if value.IsValid() && (value.Kind() == reflect.Struct || !value.IsNil()) {
		return value.Interface()
	} else {
		return nil
	}
}
