package modules

type Crud struct {}

func (crud *Crud) Commands() map[string]string {
    return map[string]string{
        "greet": "Greet",
    }
}

// end boilerplate

func (crud *Crud) Greet() string {
    return "Called Greet!\n"
}
