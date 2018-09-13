package filter

import (
	"reflect"
	"strings"
)

func repeat(value interface{}) string {
	size := 1
	if reflect.ValueOf(value).Kind() == reflect.Slice {
		size = reflect.ValueOf(value).Len()
	}
	query := strings.Repeat("?,", size)
	return query[:len(query)-1]
}
