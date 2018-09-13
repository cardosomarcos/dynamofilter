package filter

import "strings"

type Filter map[string]Item

func NewFilter() Filter {
	return Filter{}
}

func (f Filter) Equals(property string, value interface{}) Filter {
	f[property] = itemBuilder(property, value, ExpressionEquals)
	return f
}

func (f Filter) Contains(property string, value interface{}) Filter {
	f[property] = itemBuilder(property, value, ExpressionContains)
	return f
}

func (f Filter) In(property string, value interface{}) Filter {
	f[property] = itemBuilder(property, value, ExpressionIn)
	return f
}

func (f Filter) Builder() (string, []interface{}) {
	var query []string
	var args []interface{}
	
	for _, v := range f {
		query = append(query, v.Query)
		args = append(args, v.Value)
	}

	return strings.Join(query, " AND "), args
}

func (f Filter) Get(key string) Item {
	return f[key]
}
