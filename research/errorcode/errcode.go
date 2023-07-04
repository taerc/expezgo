package errorcode


type ErrorCode int64

//go:generate stringer -type=ErrorCode
const (
	OK ErrorCode = iota + 200
	Failed
	UNK
)