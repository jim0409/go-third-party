package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

var jsonStr = `{"one":"foobar","two":"foobar","three":"foobar","four":"foobar","five":"foobar","six":"foobar","seven":"foobar","eight":"foobar","nine":"foobar","ten":"foobar"}`

// ioutil
func IOUtilReadAll(reader io.Reader) (map[string]interface{}, error) {
	var (
		m    map[string]interface{}
		b, _ = ioutil.ReadAll(reader)
	)

	return m, json.Unmarshal(b, &m)
}

// io.Copy
func IOCopy(reader io.Reader) (map[string]interface{}, error) {
	var (
		m    map[string]interface{}
		buf  bytes.Buffer
		_, _ = io.Copy(&buf, reader)
	)

	return m, json.Unmarshal(buf.Bytes(), &m)
}

// json decode
func JsonDecoder(reader io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}

	return m, json.NewDecoder(reader).Decode(&m)
}

func BenchmarkJsonDecoder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsonDecoder(strings.NewReader(jsonStr))
	}
}

func BenchmarkIOUtilReadAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IOUtilReadAll(strings.NewReader(jsonStr))
	}
}

func BenchmarkIOCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IOCopy(strings.NewReader(jsonStr))
	}
}
