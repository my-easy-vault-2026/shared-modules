package common

type HTTPMethod int // http方法
const (
	HTTP_METHOD_POST   = 1
	HTTP_METHOD_GEt    = 2
	HTTP_METHOD_PUT    = 3
	HTTP_METHOD_DELETE = 4
)

type APIAuthorityStatus int // api權限狀態
const (
	API_AUTHORITY_STATUS_ENABLED  APIAuthorityStatus = 1
	API_AUTHORITY_STATUS_DISABLED APIAuthorityStatus = 2
)
