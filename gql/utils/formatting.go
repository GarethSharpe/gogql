package utils

import (
	"fmt"
	"bytes"
	"reflect"
)

func GetFields(c interface{}) *string {
	t := reflect.TypeOf(c)
	buff := bytes.Buffer{}
	for i := 0; i < t.NumField(); i++ {
		buff.WriteString(fmt.Sprintf("%s ", t.Field(i).Name))
	}
	fields := buff.String()
	return &fields
}