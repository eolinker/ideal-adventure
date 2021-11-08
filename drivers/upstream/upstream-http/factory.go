package upstream_http

import (
	"reflect"

	round_robin "github.com/eolinker/goku/upstream/round-robin"

	"github.com/eolinker/eosc"
)

//Register 注册http_proxy驱动工厂
func Register(register eosc.IExtenderDriverRegister) {
	register.RegisterExtenderDriver("upstream_http_proxy", NewFactory())
}

type factory struct {
	profession string
	name       string
	label      string
	desc       string
	params     map[string]string
}

//NewFactory 创建http_proxy驱动工厂
func NewFactory() eosc.IExtenderDriverFactory {
	round_robin.Register()
	return &factory{}
}

//Create 创建http_proxy驱动
func (f *factory) Create(profession string, name string, label string, desc string, params map[string]string) (eosc.IExtenderDriver, error) {
	return &driver{
		profession: profession,
		name:       name,
		label:      label,
		desc:       desc,
		driver:     driverName,
		configType: reflect.TypeOf((*Config)(nil)),
		params:     params,
	}, nil
}
