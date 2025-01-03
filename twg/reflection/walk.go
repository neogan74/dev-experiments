package reflection

import "reflect"

func walk(x interface{}, fn func(intput string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
