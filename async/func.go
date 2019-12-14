package async

type MapFunction func(interface{}) interface{}
type FilterFunction func(interface{}) bool
type SelectFunction MapFunction
type ScanFunction func(interface{}, interface{}) interface{}

