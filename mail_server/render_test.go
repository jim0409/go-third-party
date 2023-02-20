package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	tmp := map[string]interface{}{
		"Name": "demo",
		"URL":  "https://demo.testfire.net",
	}

	// note: 呼應到 template.html 內的取代字英文大小寫需一致
	s, err := ParseTemplate("tpl/template.html", tmp)
	assert.Nil(t, err)

	fmt.Println(s)
}
