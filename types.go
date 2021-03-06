package fu

import (
	"reflect"
	"time"
)

var Int = reflect.TypeOf(int(0))
var Int8 = reflect.TypeOf(int8(0))
var Int16 = reflect.TypeOf(int16(0))
var Int32 = reflect.TypeOf(int32(0))
var Int64 = reflect.TypeOf(int64(0))
var Uint = reflect.TypeOf(uint(0))
var Uint8 = reflect.TypeOf(uint8(0))
var Uint16 = reflect.TypeOf(uint16(0))
var Uint32 = reflect.TypeOf(uint32(0))
var Uint64 = reflect.TypeOf(uint64(0))
var Float32 = reflect.TypeOf(float32(0))
var Float64 = reflect.TypeOf(float64(0))
var Byte = reflect.TypeOf(byte(0))
var String = reflect.TypeOf("")
var Bool = reflect.TypeOf(true)
var Ts = reflect.TypeOf(time.Time{})

//var StructType = reflect.TypeOf(Struct{})
var Fixed8Type = reflect.TypeOf(Fixed8{0})
var Fixed8Zero = reflect.ValueOf(Fixed8{0})
var InterfaceType = reflect.TypeOf(interface{}(nil))

var IntZero = reflect.ValueOf(int(0))
var EmptyString = reflect.ValueOf("")
var Float32Zero = reflect.ValueOf(float32(0))
var Float64Zero = reflect.ValueOf(float64(0))
var TsZero = reflect.ValueOf(time.Time{})
var True = reflect.ValueOf(true)
var False = reflect.ValueOf(false)


func TypeOf(x interface{}) reflect.Type {
	v := reflect.ValueOf(x)
	for v.Kind() == reflect.Interface && v.IsValid() {
		v = v.Elem()
	}
	return v.Type()
}
