package status

type AppStatusCode = int

const (
	UNKNOW      AppStatusCode = 100
	SUCCESS     AppStatusCode = 200
	BAD_REQUEST AppStatusCode = 400
	NOT_FOUND   AppStatusCode = 404
	INTERNAL    AppStatusCode = 500

	// User stauts
	USER_INVALID_PARAMS     AppStatusCode = 1001
	USER_INVALID_ID         AppStatusCode = 1002
	USER_NOT_FOUND          AppStatusCode = 1003
	USER_MISSING_FIRST_NAME AppStatusCode = 1004
	USER_MISSING_LAST_NAME  AppStatusCode = 1005
	USER_MISSING_EMAIL      AppStatusCode = 1006
	USER_INVALID_EMAIL      AppStatusCode = 1007
	USER_MISSING_PHONE      AppStatusCode = 1008
	USER_INVALID_ROLE       AppStatusCode = 1019
	USER_INVALID_STATUS     AppStatusCode = 1010
)
