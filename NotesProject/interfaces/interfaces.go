package interfaces

type SaveAndDisplay interface {
	Save() error
	SuccessMessage()
}
