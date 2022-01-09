package LOG

// -----------------------------------------------------------------------------
// module imports
// -----------------------------------------------------------------------------
import (
	"github.com/apsdehal/go-logger"
)

// -----------------------------------------------------------------------------
// type definitions
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// module global variables (non-exported)
// -----------------------------------------------------------------------------
var appLogger, appLoggerError = logger.New()
var appLoggerIsInitialized bool

// -----------------------------------------------------------------------------
// module global variables (exported)
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// non-exported methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// exported methods
// -----------------------------------------------------------------------------

// initialize logger
func Initialize(level logger.LogLevel) bool {
	if appLoggerError != nil {
		appLoggerIsInitialized = false
		return appLoggerIsInitialized
	}

	appLogger.SetLogLevel(level)
	appLogger.SetFormat("%{lvl} %{time} %{message}")

	appLoggerIsInitialized = true

	Dbg("logger was initialized")
	return appLoggerIsInitialized
}

// Logging priority: debug -> info -> notice -> warning -> error -> critical -> panic

// return default logging level
func GetDefaultLevel() logger.LogLevel {
	return logger.DebugLevel
}

// set logging level
func SetLevel(newLevel logger.LogLevel) {
	Dbg("setting log level '%d'", newLevel)
	appLogger.SetLogLevel(newLevel)
}

// log debug message
func Dbg(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.DebugF(format, a...)
	}
}

// log informnation message
func Inf(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.InfoF(format, a...)
	}
}

// log notice message
func Ntc(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.NoticeF(format, a...)
	}
}

// log warning message
func Wrn(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.WarningF(format, a...)
	}
}

// log error message
func Err(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.ErrorF(format, a...)
	}
}

// log critical message
func Crt(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.CriticalF(format, a...)
	}
}

// log stack as error
func ErrStack(message string) {
	if appLoggerIsInitialized {
		appLogger.StackAsError(message)
	}
}

// log stack as critical
func CrtStack(message string) {
	if appLoggerIsInitialized {
		appLogger.StackAsCritical(message)
	}
}

// panic
func Panic(format string, a ...interface{}) {
	if appLoggerIsInitialized {
		appLogger.PanicF(format, a...)
	}
}
