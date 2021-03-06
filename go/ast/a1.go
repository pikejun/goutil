package main

import (
	"bytes"
	"errors"
	"go/format"
	"html/template"
	"os"
)

const suffix = "_msg_gen.go"

// tpl 生成代码需要用到模板
const tpl = `
// Code generated by github.com/mohuishou/gen-const-msg DO NOT EDIT

// {{.pkg}} const code comment msg
package {{.pkg}}

// noErrorMsg if code is not found, GetMsg will return this
const noErrorMsg = "unknown error"

// messages get msg from const comment
var messages = map[int]string{
	{{range $key, $value := .comments}}
	{{$key}}: "{{$value}}",{{end}}
}

// GetMsg get error msg
func GetMsg(code int) string {
	var (
		msg string
		ok  bool
	)
	if msg, ok = messages[code]; !ok {
		msg = noErrorMsg
	}
	return msg
}
`

// noErrorMsg if code is not found, GetMsg will return this
const noErrorMsg = "unknown error"
// messages get msg from const comment
var messages = map[int]string{}

// GetMsg get error msg
func GetMsg(code int) string {
	var (
		msg string
		ok  bool
	)
	if msg, ok = messages[code]; !ok {
		msg = noErrorMsg
	}
	return msg
}

//go:generate oooooooo
// gen 生成代码
func gen(comments map[string]string) ([]byte, error) {
var buf = bytes.NewBufferString("")

data := map[string]interface{}{
"pkg":      os.Getenv("GOPACKAGE"),
"comments": comments,
}

t, err := template.New("").Parse(tpl)
if err != nil {
return nil, errors.New( "template init err")
}

err = t.Execute(buf, data)
if err != nil {
return nil, errors.New("template data err")
}

return format.Source(buf.Bytes())
}