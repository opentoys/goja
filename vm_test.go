package goja

import (
	"testing"

	"github.com/dop251/goja/file"
	"github.com/dop251/goja/unistring"
)

func TestTaggedTemplateArgExport(t *testing.T) {
	vm := New()
	vm.Set("f", func(v Value) {
		v.Export()
	})
	vm.RunString("f`test`")
}

func TestVM1(t *testing.T) {
	r := &Runtime{}
	r.init()

	vm := r.vm

	vm.prg = &Program{
		src: file.NewFile("dummy", "", 1),
		code: []instruction{
			&bindGlobal{vars: []unistring.String{"v"}},
			newObject,
			setGlobal("v"),
			loadVal{asciiString("test")},
			loadVal{valueInt(3)},
			loadVal{valueInt(2)},
			add,
			setElem,
			pop,
			loadDynamic("v"),
		},
	}

	vm.run()

	rv := vm.pop()

	if obj, ok := rv.(*Object); ok {
		if v := obj.self.getStr("test", nil).ToInteger(); v != 5 {
			t.Fatalf("Unexpected property value: %v", v)
		}
	} else {
		t.Fatalf("Unexpected result: %v", rv)
	}

}
