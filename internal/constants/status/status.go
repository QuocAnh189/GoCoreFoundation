package status

type AppStatusCode = int

const (
	UNKNOW      AppStatusCode = 100
	SUCCESS     AppStatusCode = 200
	BAD_REQUEST AppStatusCode = 400
	NOT_FOUND   AppStatusCode = 404
	ERROR       AppStatusCode = 500
)
