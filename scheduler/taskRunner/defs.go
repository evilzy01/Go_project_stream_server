package taskRunner

const (
	READY_TO_DISPATCH = "d"
	RRADY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

type controlChnn chan string

type dataChnn chan interface{}

type fn func(dc dataChnn) error
