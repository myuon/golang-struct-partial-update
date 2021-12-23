package partial_update

import (
	"errors"
	"reflect"
)

var (
	ErrNotStruct = errors.New("updater is not a struct")
)

func PartialUpdate(value interface{}, updater interface{}) error {
	updaterType := reflect.TypeOf(updater)

	if updaterType.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	valueOf := reflect.ValueOf(&value)

	tmp := reflect.New(valueOf.Elem().Elem().Type()).Elem()
	tmp.Set(valueOf.Elem().Elem())

	for i := 0; i < updaterType.NumField(); i++ {
		field := updaterType.Field(i)
		fieldValue := reflect.ValueOf(updater).Field(i)

		if field.Type.Kind() == reflect.Ptr {
			if !fieldValue.IsNil() {
				tmp.Elem().FieldByName(field.Name).Set(fieldValue.Elem())
			}
		} else {
			tmp.Elem().FieldByName(field.Name).Set(fieldValue)
		}
	}

	return nil
}
