package bencoding

type Dict map[string]interface{}

func NewDict() Dict {
	return make(Dict)
}

type List []interface{}

func NewList() List {
	return make(List, 0)
}
