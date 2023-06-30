package e

const (
	SCHOOL_EXIST     = 10001
	SCHOOL_NOT_EXIST = 10002
)

var SchoolErrMsg = map[int]string{
	SCHOOL_EXIST:     "School already exists",
	SCHOOL_NOT_EXIST: "School not found",
}
