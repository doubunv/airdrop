package errs

type UserError int32

// user error 1000000~2000000
const (
	// user
	UserNotExists UserError = 1000001

	// invite 1010000 ~ 1020000

	UserInviteErr           UserError = 1010000
	UserInviteeExists       UserError = 1010001
	UserInviteCodeNotExists UserError = 1010002
	UserInviteCodeNoChange  UserError = 1010003
	InviteCodeExists        UserError = 1010004
	InviteCodeNowAllow      UserError = 1010005
	UserInviteeSelf         UserError = 1010006
)

func (e UserError) Code() int32 {
	return int32(e)
}

func (e UserError) Error() string {
	if s, ok := UserErrMap[e]; ok {
		return s
	}

	return "Something wrong with user"
}

var UserErrMap = map[UserError]string{
	UserNotExists: "User not exists.",

	UserInviteeExists:       "Invitation code is exists.",
	UserInviteCodeNotExists: "Invalid invitation code.",
	UserInviteCodeNoChange:  "Invitation code can not change.",
	InviteCodeExists:        "The invitation code already exists, please try another code.",
	InviteCodeNowAllow:      "Invitation code not allow.",
	UserInviteeSelf:         "Invitation code is yourself please change.",
}
