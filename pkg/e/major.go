package e

const (
	MAJOR_EXIST     = 20001
	MAJOR_NOT_EXIST = 20002
)

var MajorErrMsg = map[int]string{
	MAJOR_EXIST:     "Major already exists",
	MAJOR_NOT_EXIST: "Major not found",
}
