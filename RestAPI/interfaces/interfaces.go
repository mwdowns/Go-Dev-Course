package interfaces

type Updater interface {
	Update() (string, error)
}

type Saver interface {
	Save() (string, error)
}

type Deleter interface {
	Delete() error
}

type DbSaver interface {
	SaveObject(query string) (string, error)
}
