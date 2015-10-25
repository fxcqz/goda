package modules

import (
    "io/ioutil"
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
    Modules []string
    Active []string
}

func NewModuleData(_active []string) *ModuleData {
    return &ModuleData {
        Modules: ListModules(),
        Active: _active,
    }
}

func ListModules() []string {
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
