package filter

import "strings"

type filter map[string]Item

type Filter interface {
	Equals(property string, value interface{}) filter
	Contains(property string, value interface{}) filter
	In(property string, value interface{}) filter
	Builder() (string, []interface{})
	Get(key string) Item
}

func NewFilter() Filter {
	return filter{}
}

func (f filter) Equals(property string, value interface{}) filter {
	f[property] = itemBuilder(property, value, ExpressionEquals)
	return f
}

func (f filter) Contains(property string, value interface{}) filter {
	f[property] = itemBuilder(property, value, ExpressionContains)
	return f
}

func (f filter) In(property string, value interface{}) filter {
	f[property] = itemBuilder(property, value, ExpressionIn)
	return f
}

func (f filter) Builder() (string, []interface{}) {
	var query []string
	var args []interface{}

	for _, v := range f {
		query = append(query, v.Query)
		args = append(args, v.Value)
	}

	return strings.Join(query, " AND "), args
}

func (f filter) Get(key string) Item {
	return f[key]
}
