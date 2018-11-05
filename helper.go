package dynamofilter

import (
	"reflect"
	"strings"
)

func repeat(value interface{}) string {
	if reflect.ValueOf(value).Kind() != reflect.Slice || reflect.ValueOf(value).Len() < 1 {
		return "?"
	}
	query := strings.Repeat("?,", reflect.ValueOf(value).Len())
	return query[:len(query)-1]
}
