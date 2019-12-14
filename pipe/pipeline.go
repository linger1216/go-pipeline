package pipe

type Pipeline interface {
	Map(name string, fn Process) Pipeline
	MapF(f Filter) Pipeline
}
