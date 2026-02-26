package interfaces

type DbBridge interface {
	Update() (string, error)
	Save() (string, error)
}
