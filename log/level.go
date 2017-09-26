package log

type LogLevel uint // log 级别

const (
	OFF   LogLevel = iota + 1 // 1
	FATAL                     // 2
	ERROR                     // 3
	WARN                      // 4
	INFO                      // 5
	DEBUG                     // 6
	ALL                       // 7
)
