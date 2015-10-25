package modules

import (
    "io/ioutil"
    "reflect"
    "strings"
)

type Module struct {
    Name string
    Type interface{}
    // maps command to function
    Commands map[string]string
}

func NewModule(_name string, _type interface{}, _commands map[string]string) *Module {
    return &Module {
        Name: _name,
        Type: _type,
        Commands: _commands,
    }
}

type ModuleData struct {
    modules []string
    active []string
}

func NewModuleData(_active []string) *ModuleData {
    return &ModuleData {
        modules: listModules(),
        active: _active,
    }
}

func listModules() []string {
    dirInfo, _ := ioutil.ReadDir("modules/")
    var modules []string
    for _, files := range dirInfo {
        name := files.Name()
        if strings.HasSuffix(name, ".go") {
            modules = append(modules, name[:len(name)-3])
        }
    }
    return modules
}

func Hook(nick string, message string) {
    if message == "hello" {
        var c Crud
        reflect.ValueOf(&c).MethodByName("Greet").Call([]reflect.Value{})
    }
}
