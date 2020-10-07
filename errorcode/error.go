package errorcode

const (
	SUCCESS = 200
	ERROR   = 500
)

var Msg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",
}

func ErrMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return Msg[ERROR]
}
