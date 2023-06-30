package e

import "fmt"

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

var errMsg = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "invalid params",
}

func mergeMsgMaps(maps ...map[int]string) map[int]string {
	result := make(map[int]string)
	for _, m := range maps {
		for k, v := range m {
			_, ok := result[k]
			if ok {
				panic(fmt.Sprintf("duplicate error code: %d", k))
			}
			result[k] = v
		}
	}
	return result
}

var MsgFlags = mergeMsgMaps(SchoolErrMsg, MajorErrMsg, errMsg)

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
