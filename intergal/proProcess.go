package intergal

import "text/template"

const templateText = `<x ID="{{ID}}" t1="{{trans .key}} {{setOpe operation}}">
<List SizeOfView="9" ViewStartAt="0" CursorPos="0" Cycles="1" IsAutoInterval="0" Delay="30" PlayMode="0"/>
</x>`

var id = 0

func trans(byte2 byte) string {

}

func setOpe() string {

}
func work(content []byte) string {
	var Ops []string
	funcMap := template.FuncMap{
		"trans":  trans,
		"setOpe": setOpe,
	}
	for i := 0; i < len(content); i++ {
		Ops = append(Ops)
	}
}
