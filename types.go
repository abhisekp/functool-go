package functool

type List []interface{}

type Predicate func(interface{}) bool

type Transformer func(interface{}) interface{}

type OverTransformer func(*List) interface{}
