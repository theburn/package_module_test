package module_register

import (
	"errors"
	"sync"
)

type ScanPlugin interface {
	Init(info string) error
	Prepare() error
	Run() error
	Clean() error
}

var ScanPluginMap = make(map[string]ScanPlugin)

var pluginMutex sync.Mutex

func Register(name string, plugin ScanPlugin) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	if plugin == nil {
		panic("sql: Register plugin is nil")
	}
	if _, dup := ScanPluginMap[name]; dup {
		panic("sql: Register called twice for plugin " + name)
	}

	ScanPluginMap[name] = plugin
}

func Open(name, info string) (ScanPlugin, error) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	if _, ok := ScanPluginMap[name]; ok {
		if _, ok := ScanPluginMap[name].(ScanPlugin); ok {
			if err := ScanPluginMap[name].Init(info); err != nil {
				return nil, errors.New("plugin init error")
			}
			return ScanPluginMap[name], nil
		}
		return nil, errors.New("plugin interface error")
	}

	return nil, errors.New("plugin not found")
}
