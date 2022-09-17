// Package stdout
//
// flags := log.LstdFlags | log.Lshortfile
// log.New creates a new Logger. The out variable sets the destination to which log data will be written.
//
// A Logger represents an active logging object that generates lines of output to an io.Writer.
// Each logging operation makes a single call to the Writer's Write method.
// A Logger can be used simultaneously from multiple goroutines; it guarantees to serialize access to the Writer.
//
// Package `stdout` common custom loggers are:
//
// - OutputToFile: calls os.Create to create a new file to log in.
// - LogInfo: calls log.Logger to set flags while printing `info` to the logger.
// - LogWarn: calls log.Logger to set flags while printing `warn` to the logger.
// - LogError: calls log.Logger to set flags while printing `error` to the logger.
// - LogP: aggregate all three log levels into one with `Println`.
package stdout

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// OutputToFile calls os.Create to create a new file to log in.
// log.New is handled via os.Stdout with the flags passed through.
func OutputToFile(filename string) {
	if filename == "" {
		filename = "file.log"
	}

	file, error := os.Create(filename) //	func Create(name string) (*File, error) { return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) }
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if error != nil {
		log.Fatalf("file.log not created, %v", error)
	}

	log.SetOutput(file) // SetOutput sets the output destination for the standard logger.
	log.Println("helloweb stdout")
	file.Close()

	color.Set(color.FgGreen)
	log.SetOutput(os.Stdout) // Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
	log.Println("\nPrinting into standard out again")
	color.Unset()

	lp := LogP()
	lp.Info("Info log")
}

// Struct `LogPack` aggregates all loggers into one
//
// A Logger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type LogPack struct {
	LogInfo  *log.Logger
	LogWarn  *log.Logger
	LogError *log.Logger
}

// LogP aggregate all three log levels into one with `Println`.
//
// It uses these flags: `flags := log.LstdFlags | log.Lshortfile`.
//
// Usage:
//
//	lp := LogP()                     // Instantiate.
//	lp.Info("This is an info log")   // info.
//	lp.Warn("This is an warn log")   // warn.
//	lp.Error("This is an error log") // error.
//
// Aggregates these three from `LogPack` struct:
//
// LogInfo  *log.Logger
// LogWarn  *log.Logger
// LogError *log.Logger
func LogP() LogPack {
	flags := log.LstdFlags | log.Lshortfile
	logInfo := log.New(os.Stdout, "INFO: ", flags)
	logWarn := log.New(os.Stdout, "WARN: ", flags)
	logError := log.New(os.Stdout, "ERROR: ", flags)

	// Instantiate aggregatedLogger and assign
	lp := LogPack{
		LogInfo:  logInfo,
		LogWarn:  logWarn,
		LogError: logError,
	}

	return lp
}

// The `Info` method
// `v ...interface{}` passes all args into l.LogInfo.Println
func (l *LogPack) Info(v ...interface{}) {
	l.LogInfo.Println(v...)
}

// The `Warn` method
// `v ...interface{}` passes all args into l.LogInfo.Println
func (l *LogPack) Warn(v ...interface{}) {
	l.LogInfo.Println(v...)
}

// The `Error` method
// `v ...interface{}` passes all args into l.LogInfo.Println
func (l *LogPack) Error(v ...interface{}) {
	l.LogInfo.Println(v...)
}

// LogInfo calls log.Logger to set flags while printing to the logger.
// log.New is handled via os.Stdout with the flags passed through
func LogInfo() *log.Logger {
	flags := log.LstdFlags | log.Lshortfile
	logInfo := log.New(os.Stdout, "INFO: ", flags)
	return logInfo
}

// LogWarn calls log.Logger to set flags while printing to the logger.
// log.New is handled via os.Stdout with the flags passed through
func LogWarn() *log.Logger {
	flags := log.LstdFlags | log.Lshortfile
	logWarn := log.New(os.Stdout, "WARN: ", flags)
	return logWarn
}

// LogError calls log.Logger to set flags while printing to the logger.
// `log.New()` is handled via os.Stdout with the flags passed through
func LogError() *log.Logger {
	flags := log.LstdFlags | log.Lshortfile
	logError := log.New(os.Stdout, "ERROR: ", flags)
	return logError
}

// OutputToFile calls os.Create to create a new file to log in.
// log.New is handled via os.Stdout with the flags passed through
func playground(filename string) {
	if filename == "" {
		filename = "file.log"
	}
	file, error := os.Create(filename) //	func Create(name string) (*File, error) { return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) }
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if error != nil {
		log.Fatalf("file.log not created, %v", error)
	}
	log.SetOutput(file) // SetOutput sets the output destination for the standard logger.
	log.Println("helloweb stdout")
	file.Close()

	color.Set(color.FgGreen)
	log.SetOutput(os.Stdout) // Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
	log.Println("\nPrinting into standard out again")
	color.Unset()

	LogInfo().Println("This is an info log")
	LogWarn().Println("This is an warning log")
	LogError().Println("This is an error log")

	flags := log.LstdFlags | log.Lshortfile
	logInfo := log.New(os.Stdout, "INFO: ", flags)
	logWarn := log.New(os.Stdout, "WARN: ", flags)
	logError := log.New(os.Stdout, "ERROR: ", flags)

	// Aggregate all three into one
	// Instantiate aggregatedLogger and assign
	lp := LogPack{
		LogInfo:  logInfo,
		LogWarn:  logWarn,
		LogError: logError,
	}

	lp.Info("This is an info log")
	lp.Warn("This is an warn log")
	lp.Error("This is an error log")

}
