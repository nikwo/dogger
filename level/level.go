package level

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

var levelsString = map[Level]string{
	TRACE: "trace",
	DEBUG: "debug",
	INFO:  "info",
	WARN:  "warn",
	ERROR: "error",
}

var stringLevels = map[string]Level{
	"trace": TRACE,
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
}

func (l Level) String() string {
	level, exists := levelsString[l]
	if !exists {
		return levelsString[INFO]
	}

	return level
}

func LogLevelFromString(level string) Level {
	str, exists := stringLevels[level]
	if !exists {
		return stringLevels["info"]
	}

	return str
}
