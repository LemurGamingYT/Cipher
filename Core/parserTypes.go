package core

import "reflect"

var IntObjectType = reflect.TypeOf(&IntObject{})
var FloatObjectType = reflect.TypeOf(&FloatObject{})
var StringObjectType = reflect.TypeOf(&StringObject{})
var BoolObjectType = reflect.TypeOf(&BoolObject{})
var ArrayObjectType = reflect.TypeOf(&ArrayObject{})
