package model

type errorCode struct {
	SUCCESS     int
	ERROR       int
	NotFound    int
	LoginError  int
}

// ErrorCode 错误码
var ErrorCode = errorCode{
	SUCCESS     : 200,
	ERROR       : 500,
	NotFound    : 404,
	LoginError  : 410, //用户名或密码错误
}









