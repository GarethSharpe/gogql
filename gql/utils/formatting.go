package utils

import (
	"fmt"
	"bytes"
	"reflect"
)

func GetFields(c interface{}) *string {
	t := reflect.TypeOf(c)
	buff := bytes.Buffer{}
	n := t.NumField()
	for i := 0; i < n-1; i++ {
		buff.WriteString(fmt.Sprintf("%s, ", t.Field(i).Name))
	}
	buff.WriteString(fmt.Sprintf("%s", t.Field(n-1).Name))
	fields := buff.String()
	return &fields
}