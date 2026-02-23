package iomanager

type Manager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
