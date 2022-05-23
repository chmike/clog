package clog

import (
	"fmt"
	"os"
)

// Fatal logs a fatal error and exit the program with a status 1.
func Fatal(args ...interface{}) {
	mainLog.log(FatalLevel, 0, fmt.Sprint(args...))
	activeHandler.close()
	os.Exit(1)
}

// Fatalln logs a fatal error and exit the program with a status 1.
// A white space is inserted between arguments.
func Fatalln(args ...interface{}) {
	mainLog.log(FatalLevel, 0, fmt.Sprintln(args...))
	activeHandler.close()
	os.Exit(1)
}

// Fatalf logs a formatted error and exits the program with a status 1.
func Fatalf(format string, args ...interface{}) {
	mainLog.log(FatalLevel, 0, fmt.Sprintf(format, args...))
	activeHandler.close()
	os.Exit(1)
}

// Error logs an error requiring an intervention.
func Error(args ...interface{}) {
	if mainLog.level >= ErrorLevel {
		mainLog.log(ErrorLevel, 0, fmt.Sprint(args...))
	}
}

// Errorln logs an error requiring an intervention.
// A white space is inserted between arguments.
func Errorln(args ...interface{}) {
	if mainLog.level >= ErrorLevel {
		mainLog.log(ErrorLevel, 0, fmt.Sprintln(args...))
	}
}

// Errorf logs a formatted error requiring an intervention.
func Errorf(format string, args ...interface{}) {
	if mainLog.level >= ErrorLevel {
		mainLog.log(ErrorLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Warning logs a self corrected error.
func Warning(args ...interface{}) {
	if mainLog.level >= WarningLevel {
		mainLog.log(WarningLevel, 0, fmt.Sprint(args...))
	}
}

// Warningln logs a self corrected error.
// A white space is inserted between arguments.
func Warningln(args ...interface{}) {
	if mainLog.level >= WarningLevel {
		mainLog.log(WarningLevel, 0, fmt.Sprintln(args...))
	}
}

// Warningf logs a formatted self corrected error.
func Warningf(format string, args ...interface{}) {
	if mainLog.level >= WarningLevel {
		mainLog.log(WarningLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Info logs a major progress notification.
func Info(args ...interface{}) {
	if mainLog.level >= InfoLevel {
		mainLog.log(InfoLevel, 0, fmt.Sprint(args...))
	}
}

// Infoln logs a major progress notification.
// A white space is inserted between arguments.
func Infoln(args ...interface{}) {
	if mainLog.level >= InfoLevel {
		mainLog.log(InfoLevel, 0, fmt.Sprintln(args...))
	}
}

// Infof logs a formatted major progress notification.
func Infof(format string, args ...interface{}) {
	if mainLog.level >= InfoLevel {
		mainLog.log(InfoLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Print logs a normal progress notification.
func Print(args ...interface{}) {
	if mainLog.level >= PrintLevel {
		mainLog.log(PrintLevel, 0, fmt.Sprint(args...))
	}
}

// Println logs a normal progress notification.
// A white space is inserted between arguments.
func Println(args ...interface{}) {
	if mainLog.level >= PrintLevel {
		mainLog.log(PrintLevel, 0, fmt.Sprintln(args...))
	}
}

// Printf logs a formatted normal progress notification.
func Printf(format string, args ...interface{}) {
	if mainLog.level >= PrintLevel {
		mainLog.log(PrintLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Debug logs a debug message for minimum verbosity.
func Debug(args ...interface{}) {
	if mainLog.level >= DebugLevel {
		mainLog.log(DebugLevel, 0, fmt.Sprint(args...))
	}
}

// Debugln logs a debug message for minimum verbosity.
// A white space is inserted between arguments.
func Debugln(args ...interface{}) {
	if mainLog.level >= DebugLevel {
		mainLog.log(DebugLevel, 0, fmt.Sprintln(args...))
	}
}

// Debugf logs a formatted debug message for minimum verbosity.
func Debugf(format string, args ...interface{}) {
	if mainLog.level >= DebugLevel {
		mainLog.log(DebugLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Debug1 logs a debug message for intermediate verbosity.
func Debug1(args ...interface{}) {
	if mainLog.level >= Debug1Level {
		mainLog.log(Debug1Level, 0, fmt.Sprint(args...))
	}
}

// Debug1ln logs a debug message for intermediate verbosity.
// A white space is inserted between arguments.
func Debug1ln(args ...interface{}) {
	if mainLog.level >= Debug1Level {
		mainLog.log(Debug1Level, 0, fmt.Sprintln(args...))
	}
}

// Debug1f logs a debug message for intermediate verbosity.
func Debug1f(format string, args ...interface{}) {
	if mainLog.level >= Debug1Level {
		mainLog.log(Debug1Level, 0, fmt.Sprintf(format, args...))
	}
}

// Debug2 logs a debug message fon maximum verbosity.
func Debug2(args ...interface{}) {
	if mainLog.level >= Debug2Level {
		mainLog.log(Debug2Level, 0, fmt.Sprint(args...))
	}
}

// Debug2ln logs a debug message fon maximum verbosity.
// A white space is inserted between arguments.
func Debug2ln(args ...interface{}) {
	if mainLog.level >= Debug2Level {
		mainLog.log(Debug2Level, 0, fmt.Sprintln(args...))
	}
}

// Debug2f logs a debug message fon maximum verbosity.
func Debug2f(format string, args ...interface{}) {
	if mainLog.level >= Debug2Level {
		mainLog.log(Debug2Level, 0, fmt.Sprintf(format, args...))
	}
}

// Fatal logs a fatal error and exit the program with a status 1.
func (c Clog) Fatal(args ...interface{}) {
	c.log(FatalLevel, 0, fmt.Sprint(args...))
	activeHandler.close()
	os.Exit(1)
}

// Fatalln logs a fatal error and exit the program with a status 1.
// A white space is inserted between arguments.
func (c Clog) Fatalln(args ...interface{}) {
	c.log(FatalLevel, 0, fmt.Sprintln(args...))
	activeHandler.close()
	os.Exit(1)
}

// Fatalf logs a formatted error and exits the program with a status 1.
func (c Clog) Fatalf(format string, args ...interface{}) {
	c.log(FatalLevel, 0, fmt.Sprintf(format, args...))
	activeHandler.close()
	os.Exit(1)
}

// Error logs an error requiring an intervention.
func (c Clog) Error(args ...interface{}) {
	if c.level >= ErrorLevel {
		c.log(ErrorLevel, 0, fmt.Sprint(args...))
	}
}

// Errorln logs an error requiring an intervention.
// A white space is inserted between arguments.
func (c Clog) Errorln(args ...interface{}) {
	if c.level >= ErrorLevel {
		c.log(ErrorLevel, 0, fmt.Sprintln(args...))
	}
}

// Errorf logs a formatted error requiring an intervention.
func (c Clog) Errorf(format string, args ...interface{}) {
	if c.level >= ErrorLevel {
		c.log(ErrorLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Warning logs a self corrected error.
func (c Clog) Warning(args ...interface{}) {
	if c.level >= WarningLevel {
		c.log(WarningLevel, 0, fmt.Sprint(args...))
	}
}

// Warningln logs a self corrected error.
// A white space is inserted between arguments.
func (c Clog) Warningln(args ...interface{}) {
	if c.level >= WarningLevel {
		c.log(WarningLevel, 0, fmt.Sprint(args...))
	}
}

// Warningf logs a formatted self corrected error.
func (c Clog) Warningf(format string, args ...interface{}) {
	if c.level >= WarningLevel {
		c.log(WarningLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Info logs a major progress notification.
func (c Clog) Info(args ...interface{}) {
	if c.level >= InfoLevel {
		c.log(InfoLevel, 0, fmt.Sprint(args...))
	}
}

// Infoln logs a major progress notification.
// A white space is inserted between arguments.
func (c Clog) Infoln(args ...interface{}) {
	if c.level >= InfoLevel {
		c.log(InfoLevel, 0, fmt.Sprintln(args...))
	}
}

// Infof logs a formatted major progress notification.
func (c Clog) Infof(format string, args ...interface{}) {
	if c.level >= InfoLevel {
		c.log(InfoLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Print logs a normal progress notification.
func (c Clog) Print(args ...interface{}) {
	if c.level >= PrintLevel {
		c.log(PrintLevel, 0, fmt.Sprint(args...))
	}
}

// Println logs a normal progress notification.
// A white space is inserted between arguments.
func (c Clog) Println(args ...interface{}) {
	if c.level >= PrintLevel {
		c.log(PrintLevel, 0, fmt.Sprintln(args...))
	}
}

// Printf logs a formatted normal progress notification.
func (c Clog) Printf(format string, args ...interface{}) {
	if c.level >= PrintLevel {
		c.log(PrintLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Debug logs a debug message for minimum verbosity.
func (c Clog) Debug(args ...interface{}) {
	if c.level >= DebugLevel {
		c.log(DebugLevel, 0, fmt.Sprint(args...))
	}
}

// Debugln logs a debug message for minimum verbosity.
// A white space is inserted between arguments.
func (c Clog) Debugln(args ...interface{}) {
	if c.level >= DebugLevel {
		c.log(DebugLevel, 0, fmt.Sprintln(args...))
	}
}

// Debugf logs a formatted debug message for minimum verbosity.
func (c Clog) Debugf(format string, args ...interface{}) {
	if c.level >= DebugLevel {
		c.log(DebugLevel, 0, fmt.Sprintf(format, args...))
	}
}

// Debug1 logs a debug message for intermediate verbosity.
func (c Clog) Debug1(args ...interface{}) {
	if c.level >= Debug1Level {
		c.log(Debug1Level, 0, fmt.Sprint(args...))
	}
}

// Debug1ln logs a debug message for intermediate verbosity.
// A white space is inserted between arguments.
func (c Clog) Debug1ln(args ...interface{}) {
	if c.level >= Debug1Level {
		c.log(Debug1Level, 0, fmt.Sprintln(args...))
	}
}

// Debug1f logs a debug message for intermediate verbosity.
func (c Clog) Debug1f(format string, args ...interface{}) {
	if c.level >= Debug1Level {
		c.log(Debug1Level, 0, fmt.Sprintf(format, args...))
	}
}

// Debug2 logs a debug message fon maximum verbosity.
func (c Clog) Debug2(args ...interface{}) {
	if c.level >= Debug2Level {
		c.log(Debug2Level, 0, fmt.Sprint(args...))
	}
}

// Debug2ln logs a debug message fon maximum verbosity.
// A white space is inserted between arguments.
func (c Clog) Debug2ln(args ...interface{}) {
	if c.level >= Debug2Level {
		c.log(Debug2Level, 0, fmt.Sprintln(args...))
	}
}

// Debug2f logs a debug message fon maximum verbosity.
func (c Clog) Debug2f(format string, args ...interface{}) {
	if c.level >= Debug2Level {
		c.log(Debug2Level, 0, fmt.Sprintf(format, args...))
	}
}
