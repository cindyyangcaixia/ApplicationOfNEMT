package e

const (
	ERROR_AUTH_CHECK_TOKEN_FAIL = 30001
	TOKEN_EXPIRED               = 30002
)

var TokenMsgFlags = map[int]string{
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token authentication failed",
	TOKEN_EXPIRED:               "Token is expired",
}
