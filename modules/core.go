package modules

import (
    "io/ioutil"
    "strings"
)

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

}
