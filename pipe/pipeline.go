package pipe

type Pipeline interface {
	Append(name string, fn Process) Pipeline
	AppendFilter(f Filter) Pipeline
}
