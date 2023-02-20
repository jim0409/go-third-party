package main

import (
	"bytes"
	"html/template"
)

var MsgTemplate = make(map[int]string)

func init() {
	MsgTemplate[0] = "tpl/demo.html"
	MsgTemplate[1] = "tpl/template.html"
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
