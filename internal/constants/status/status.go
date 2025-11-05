package status

type Code = int

const (
	UNKNOW      Code = 100
	SUCCESS     Code = 200
	BAD_REQUEST Code = 400
	NOT_FOUND   Code = 404
	INTERNAL    Code = 500

	// User stauts
	USER_INVALID_PARAMS     Code = 1001
	USER_INVALID_ID         Code = 1002
	USER_NOT_FOUND          Code = 1003
	USER_MISSING_FIRST_NAME Code = 1004
	USER_MISSING_LAST_NAME  Code = 1005
	USER_MISSING_EMAIL      Code = 1006
	USER_INVALID_EMAIL      Code = 1007
	USER_MISSING_PHONE      Code = 1008
	USER_INVALID_ROLE       Code = 1019
	USER_INVALID_STATUS     Code = 1010
)
