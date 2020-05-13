package module_a

import (
	"fmt"
	"package_module_test/module_register"
)

var name string = "module_b"

type ModuleA struct{}

func init() {
	module_register.Register(name, &ModuleA{})
}

func (m ModuleA) Init(info string) error {
	fmt.Println("this is a module B Init", info)
	return nil
}

func (m ModuleA) Prepare() error {
	fmt.Println("this is a module B Prepare")
	return nil
}

func (m ModuleA) Run() error {
	fmt.Println("this is a module B Run")
	return nil
}

func (m ModuleA) Clean() error {
	fmt.Println("this is a module B Clean")
	return nil
}
