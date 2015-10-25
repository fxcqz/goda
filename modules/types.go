package modules

import (
    "errors"
    "reflect"
    "runtime"
)

type _typeRegistry map[string]reflect.Type

func (t _typeRegistry) Set(i interface{}) {
    //name string, typ reflect.Type
    t[reflect.TypeOf(i).String()] = reflect.TypeOf(i)
}

func (t _typeRegistry) Get(name string) (interface{}, error) {
    if _type, ok := t["*modules." + name]; ok {
        return reflect.New(_type).Elem().Interface(), nil
    }
    return nil, errors.New("no one")
}

var TypeRegistry = make(_typeRegistry)

func init() {
    TypeRegistry.Set(new(Crud))
    runtime.GC()
}
