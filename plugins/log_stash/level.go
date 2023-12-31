package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l Level) String() string {
	var str string
	switch l {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warn"
	case ErrorLevel:
		str = "error"
	default:
		str = "其他"
	}
	return str
}
