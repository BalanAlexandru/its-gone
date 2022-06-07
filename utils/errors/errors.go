package errors

type Level int

const (
	INFO Level = iota
	WARNING
	DANGER
)

type Error struct {
	ErrorLevel Level
	Message    string
}

type IndexOutOfBounds struct {
	Error
}
