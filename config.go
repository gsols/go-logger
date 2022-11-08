package logger

// WriterConfig is a struct with 6 fields.
// @property {string} Writer - The writer to use. This can be either "file" or "console".
// @property {string} Directory - The directory where the log files will be written.
// @property {string} FileName - The name of the file to write logs to.
// @property {int} MaxSize - The maximum size in megabytes of the log file before it gets rotated. It defaults to 100
// megabytes.
// @property {int} MaxBackups - The maximum number of old log files to retain.
// @property {int} MaxAge - The maximum age of a log file in days
type WriterConfig struct {
	Writer     string
	Directory  string
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// Config is a struct that contains a boolean, a `zerolog.Level` and an `io.Writer`.
// @property {bool} Debug - This is a boolean value that determines whether the application is running in debug mode or
// not.
// @property LogLevel - The level of logging you want to use.
// @property Writer - The writer to which the logs will be written.
type Config struct {
	Debug        bool
	Verbosity    int
	WriterConfig WriterConfig
}
