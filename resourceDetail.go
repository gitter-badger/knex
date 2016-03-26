package knex

import (
	"fmt"
	"reflect"
	"strings"
)

type resourceDetail struct {
	interfaceType reflect.Type
	provider      Provider
}

func newResourceDetail(implementationType interface{}) (*resourceDetail, error) {

	// Search each field for 'provide' tag.
	resourceDetail := &resourceDetail{provider: Provider{}}
	reflectType := reflect.TypeOf(implementationType).Elem()
	fieldCount := reflectType.NumField()
	for i := 0; i < fieldCount; i++ {

		// Check if field has 'provide' tag.
		field := reflectType.Field(i)
		provideTagValue := strings.ToUpper(strings.Trim(field.Tag.Get(provideTagName), " "))
		if provideTagValue != emptyString {

			// Check for valid 'provide' tag values.
			if provideTagValue != resourceValue {
				return nil, fmt.Errorf("Invalid provide value '%s'", provideTagValue)
			}
			resourceDetail.interfaceType = field.Type

			// Check for 'id' tag value.
			idTagValue := strings.Trim(field.Tag.Get(idTagName), " ")
			if idTagValue != emptyString {
				resourceDetail.provider.ID = idTagValue
			}

			// Check if 'scope' field value is valid.
			scopeValue := strings.ToUpper(strings.Trim(field.Tag.Get(scopeTagName), " "))
			if validateScopeValue(scopeValue) {
				resourceDetail.provider.Scope = scopeValue
			} else {
				return nil, fmt.Errorf("Invalid scope value '%s'", scopeValue)
			}
		}
	}
	return resourceDetail, nil
}

func newResourceDetailByProvider(provider *Provider) (*resourceDetail, error) {

	// Check if provider interface type does not exist.
	if provider.Type == nil {
		return nil, fmt.Errorf("Provider must have an interface type")
	}

	// Check if scope field is valid.
	provider.Scope = strings.ToUpper(strings.Trim(provider.Scope, " "))
	if !validateScopeValue(provider.Scope) {
		return nil, fmt.Errorf("Invalid scope value '%s'", provider.Scope)
	}

	// Build resourceDetail struct.
	returnValue := &resourceDetail{
		interfaceType: reflect.TypeOf(provider.Type).Elem(),
		provider:      *provider,
	}

	return returnValue, nil
}

func validateScopeValue(value string) bool {
	if value == emptyString || value == factoryValue || value == graphValue {
		return true
	}
	return false
}
