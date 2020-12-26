package intergal

import (
	"bytes"
	"strings"
	"text/template"
)

const templateText = `<x ID="{{.ID}}" t1="{{trans .key}} {{setOpe .operation}}">
<List SizeOfView="9" ViewStartAt="0" CursorPos="0" Cycles="1" IsAutoInterval="0" Delay="30" PlayMode="0"/>
</x>`

var b bytes.Buffer
var F map[int]string

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func trans(byte2 byte) interface{} {
	//if byte2 < 0{
	//	return "Left SHIFT"
	//}
	//return F[int(byte2)]
	return If(byte2 < 0, "Left SHIFT", F[int(byte2)])
}

func setOpe(id int) interface{} {
	return If(id == 1, "Down", "UP")
}

func getF() {
	// TODO 2020/12/27 @Jason.Z tobe improved.
	// unchanged
	var index []int
	// '
	index = append(index, 39)
	// , - . / 0-9
	for i := 44; i <= 57; i++ {
		index = append(index, i)
	}
	// ;
	index = append(index, 59)
	// =
	index = append(index, 61)
	// [\]
	for i := 91; i <= 93; i++ {
		index = append(index, i)
	}
	// ` + a-z
	for i := 96; i <= 122; i++ {
		index = append(index, i)
	}
	// 刻录
	for i, val := range index {
		if i <= 122 && i >= 97 {
			F[val] = string(rune(val - ('a' - 'A')))
		} else {
			F[val] = string(rune(val))
		}
	}

	index = []int{}
	// need shift
	// ! " # $ % &
	for i := 33; i <= 38; i++ {
		index = append(index, i)
	}
	// ()*+
	for i := 40; i <= 43; i++ {
		index = append(index, i)
	}
	// <
	index = append(index, 60)
	// > ? @ A-Z
	for i := 62; i <= 90; i++ {
		index = append(index, i)
	}
	// ^ _
	for i := 94; i <= 95; i++ {
		index = append(index, i)
	}
	// {|}~
	for i := 123; i <= 126; i++ {
		index = append(index, i)
	}
	// 刻录
	for _, val := range index {
		F[val] = "damnShift " + string(rune(val))
	}
	// special
	F[int(' ')] = "SPACEBAR"
	F[10] = "ENTER"
}

var funcMap = template.FuncMap{
	"trans":  trans,
	"setOpe": setOpe,
}
var t = template.Must(template.New("Combo").Funcs(funcMap).Parse(templateText))
var id = 0

func shiftDown() {
	data := map[string]interface{}{
		"ID":        id,
		"key":       -1,
		"operation": 1,
	}
	_ = t.Execute(&b, data)
	id++
}

func shiftUp() {
	data := map[string]interface{}{
		"ID":        id,
		"key":       -1,
		"operation": 0,
	}
	_ = t.Execute(&b, data)
	id++
}

func Work(content []byte) string {

	getF()
	// TODO 2020/12/27 @Jason.Z tobe improved.
	onShift := false
	for _, c := range content {
		data := map[string]interface{}{
			"ID":        id,
			"key":       c,
			"operation": 1,
		}
		res := F[int(c)]
		if strings.HasPrefix(res, "damnShift") {
			if onShift == false {
				shiftDown()
				onShift = true
			}
		} else {
			if onShift {
				shiftUp()
				onShift = false
			}
		}
		_ = t.Execute(&b, data)
		data["operation"] = 0
		id++
		_ = t.Execute(&b, data)
		id++

	}
	if onShift {
		shiftUp()
		onShift = false
	}
	return b.String()

}
