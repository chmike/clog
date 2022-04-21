package clog

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Level uint32

const (
	FatalLevel   Level = iota // Terminate program.
	ErrorLevel                // Error requiring intervention.
	WarningLevel              // Self corrected error.
	InfoLevel                 // Major progress notification.
	PrintLevel                // Progress notification.
	DebugLevel                // Debug message with minimum verbosity.
	Debug1Level               // Debug message with intermediate verbosity.
	Debug2Level               // Debug message with maximum verbosity.
)

// Clog manages logging with a specific level and tag name.
type Clog struct {
	level Level  // Filtering level: smaller or equal levels are published.
	tag   string // Tag identifying component.
}

// mainLog is default main logging.
var mainLog = Clog{
	level: PrintLevel,
	tag:   "",
}

// SetLevel sets the main level.
func SetLevel(l Level) {
	mainLog.SetLevel(l)
}

// New return a logger with a distinct tag and level.
func New(tag string, level Level) Clog {
	if tag == "" {
		tag = "log"
	}
	return Clog{tag: tag, level: level}
}

// SetLevel sets the filtering level
func (c *Clog) SetLevel(l Level) {
	atomic.StoreUint32((*uint32)(&c.level), uint32(l))
}

// New return a new clog instance, inheriting the level and appending the
// tag name into "a.b".
func (c Clog) New(tag string) Clog {
	if tag == "" {
		tag = "log"
	}
	return Clog{tag: c.tag + "." + tag, level: c.level}
}

const levelChar = "FEWI .-="
const digits = "0123456789"

var bufPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

// log generates and outputs the message.
func (c *Clog) log(l Level, depth int, msg string) {
	if l > c.level {
		return
	}
	buf := bufPool.Get().(*strings.Builder)
	defer bufPool.Put(buf)
	buf.Reset()
	now := time.Now()
	_, month, day := now.Date()
	mon := int(month)
	hour, minute, second := now.Clock()
	micro := now.Nanosecond() / 1000
	var hdr [25]byte
	hdr[0] = levelChar[l]
	hdr[1] = digits[(mon/10)%10]
	hdr[2] = digits[mon%10]
	hdr[3] = digits[(day/10)%10]
	hdr[4] = digits[day%10]
	hdr[5] = ' '
	hdr[6] = digits[(hour/10)%10]
	hdr[7] = digits[hour%10]
	hdr[8] = ':'
	hdr[9] = digits[(minute/10)%10]
	hdr[10] = digits[minute%10]
	hdr[11] = ':'
	hdr[12] = digits[(second/10)%10]
	hdr[13] = digits[second%10]
	hdr[14] = '.'
	hdr[15] = digits[(micro/100000)%10]
	hdr[16] = digits[(micro/10000)%10]
	hdr[17] = digits[(micro/1000)%10]
	hdr[18] = digits[(micro/100)%10]
	hdr[19] = digits[(micro/10)%10]
	hdr[20] = digits[micro%10]
	hdr[21] = ' '
	buf.WriteString(string(hdr[:22]))
	_, file, line, ok := runtime.Caller(2 + depth)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteByte(' ')
	if c.tag != "" {
		buf.WriteString(c.tag)
		buf.WriteByte(':')
		buf.WriteByte(' ')
	}
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	if strings.Contains(msg, "\n") {
		msg = strings.ReplaceAll(msg, "\n", "\n"+strings.Repeat(" ", buf.Len()))
	}
	buf.WriteString(msg)
	buf.WriteByte('\n')
	activeHandler.output(buf.String())
}

func (l Level) String() string {
	switch l {
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarningLevel:
		return "warning"
	case InfoLevel:
		return "info"
	case PrintLevel:
		return "print"
	case DebugLevel:
		return "debug"
	case Debug1Level:
		return "debug1"
	case Debug2Level:
		return "debug2"
	default:
		return "unknown"
	}
}

// LevelFromString returns the level corresponding to s, or PrintLevel by default.
// It is case insensitive.
func LevelFromString(s string) Level {
	switch strings.ToLower(s) {
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warning":
		return WarningLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "debug1":
		return Debug1Level
	case "debug2":
		return Debug2Level
	default:
		return PrintLevel
	}
}

func (l *Level) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err == nil {
		*l = LevelFromString(s)
	}
	return
}

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
