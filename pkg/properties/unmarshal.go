package properties

import (
	"errors"
	"reflect"
)

func Unmarshall(data []byte, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	kv, err := readFromBytes(data)
	if err != nil {
		return err
	}

	if v.Kind() != reflect.Ptr || v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("not a pointer to struct")
	}

	for i := 0; i < v.Elem().NumField(); i++ {
		vf := v.Elem().Field(i)
		tf := v.Elem().Type().Field(i)
		tag := tf.Tag.Get("properties")
		if tag == "" {
			continue
		}

		vf.SetString(kv[tag])
	}

	return nil
}
