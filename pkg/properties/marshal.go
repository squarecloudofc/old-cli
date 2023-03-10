package properties

import (
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var result = []byte{}
	rv := reflect.Indirect(reflect.ValueOf(v))
	
	for i := 0; i < rv.NumField(); i++ {
		to_add := ""
		ty := rv.Type().Field(i)
		field := rv.Field(i)
		value := field.String()
		if value == "" {
			continue
		}

		key := ty.Tag.Get("properties")

		to_add = fmt.Sprintf("%s=%s", key, value)

		if i+1 < rv.NumField() {
			to_add = to_add + "\n"
		}

		result = append(result, to_add...)
	}

	return result, nil
}
