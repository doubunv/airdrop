package errs

type SystemError int32

func (e SystemError) Code() int32 {
	return int32(e)
}

// system err 10000

// db err  40000 ~ 40400
const (
	DbSelectErr SystemError = 40000
	DbUpdateErr SystemError = 40001
	DbDeleteErr SystemError = 40002
	DbInsertErr SystemError = 40003
)

func (e SystemError) Error() string {
	if s, ok := SystemErrMap[e]; ok {
		return s
	}

	return "Something wrong with system."
}

var SystemErrMap = map[SystemError]string{
	DbSelectErr: "Nothing select.",
	DbUpdateErr: "Nothing update.",
	DbDeleteErr: "Nothing delete.",
	DbInsertErr: "Nothing create.",
}
