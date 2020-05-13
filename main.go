package main

import (
	_ "package_module_test/module_a"
	_ "package_module_test/module_b"
	"package_module_test/module_register"
)

func main() {

	m_a, ok_a := module_register.Open("module_a", "info a")
	if ok_a == nil {
		m_a.Prepare()
		m_a.Run()
		m_a.Clean()
	}

	m_b, ok_b := module_register.Open("module_b", "info b")
	if ok_b == nil {
		m_b.Prepare()
		m_b.Run()
		m_b.Clean()
	}
}
