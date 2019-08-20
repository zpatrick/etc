package health

type Reader interface {
	Read() (Health, error}
}
