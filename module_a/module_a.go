package module_a

import (
	"fmt"
	"package_module_test/module_register"
)

var name string = "module_a"

type ModuleA struct{}

func init() {
	module_register.Register(name, &ModuleA{})
}

func (m ModuleA) Init(info string) error {
	fmt.Println("this is a module A Init", info)
	return nil
}

func (m ModuleA) Prepare() error {
	fmt.Println("this is a module A Prepare")
	return nil
}

func (m ModuleA) Run() error {
	fmt.Println("this is a module A Run")
	return nil
}

func (m ModuleA) Clean() error {
	fmt.Println("this is a module A Clean")
	return nil
}
