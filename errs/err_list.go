package errs

type ErrType int

const ( // 通用错误类型
	ERRTYPE_INTERNAL_SERVER ErrType = -1

	ERRTYPE_DATA_NOT_FOUND ErrType = 100 + iota
	// WRONG_VALUE     = 10000
	// 	NIL_VALUE       = 10002
	// 	NO_PERMISSION   = 10003
	// 	UPLOAD_DIR_ERR  = 10004
	// 	UPLOAD_FILE_ERR = 10005
	// 	MULTIPLE_ERR    =
)

const ( // 用户、角色类 10000-10999
	// 10000-10200 危险型错误

	ERRTYPE_USER_NOT_EXIST ErrType = 10201 + iota
	ERRTYPE_USER_IS_EXIST
	ERRTYPE_USER_EMAIL_NOT_EXIST
	ERRTYPE_USER_EMAIL_IS_EXIST
	ERRTYPE_USER_PHONE_NOT_EXIST
	ERRTYPE_USER_PHONE_IS_EXIST
	ERRTYPE_USER_WRONG_PASS
	ERRTYPE_USER_LOGIN_TIME_OUT
	ERRTYPE_USER_LOGOUT
	ERRTYPE_USER_LOGIN_ELSEWHERE
	ERRTYPE_USER_WRONG_STATUS
)
const ( // 权限类 11000-11499
	ERRTYPE_NO_PERMISSION ErrType = 11000 + iota
)
const ( // product  11500-11999
	ERRTYPE_PRODUCT_NUM_IS_EXIST ErrType = 11500 + iota
	ERRTYPE_PRODUCT_UNAPPLIED
	ERRTYPE_PRODUCT_VERIFY_FAILED_IP
	ERRTYPE_PRODUCT_VERIFY_FAILED_MP
)

var msgList = make(map[ErrType]string)

func init() {
	msgList[ERRTYPE_INTERNAL_SERVER] = "内部发生错误，请稍后重试"
	msgList[ERRTYPE_DATA_NOT_FOUND] = "查询的数据不存在"

	msgList[ERRTYPE_USER_NOT_EXIST] = "用户不存在"
	msgList[ERRTYPE_USER_IS_EXIST] = "用户已存在"
	msgList[ERRTYPE_USER_EMAIL_NOT_EXIST] = "用户邮箱不存在"
	msgList[ERRTYPE_USER_EMAIL_IS_EXIST] = "用户邮箱已存在"
	msgList[ERRTYPE_USER_PHONE_NOT_EXIST] = "用户手机号不存在"
	msgList[ERRTYPE_USER_PHONE_IS_EXIST] = "用户手机号已存在"
	msgList[ERRTYPE_USER_WRONG_PASS] = "用户密码错误"
	msgList[ERRTYPE_USER_LOGIN_TIME_OUT] = "登录超时，请重新登录"
	msgList[ERRTYPE_USER_LOGOUT] = "已退出登录，请重新登录"
	msgList[ERRTYPE_USER_LOGIN_ELSEWHERE] = "已在其他地方登录，若不是本人操作，请立即检查账户"
	msgList[ERRTYPE_USER_WRONG_STATUS] = "用户状态不正常"

	// 权限类 11000-11499
	msgList[ERRTYPE_NO_PERMISSION] = "无权限进行当前操作"

	// product  11500-11999
	msgList[ERRTYPE_PRODUCT_NUM_IS_EXIST] = "产品编号已存在"
	msgList[ERRTYPE_PRODUCT_UNAPPLIED] = "未开通当前产品"
	msgList[ERRTYPE_PRODUCT_VERIFY_FAILED_IP] = "产品IP校验不通过"
	msgList[ERRTYPE_PRODUCT_VERIFY_FAILED_MP] = "产品公众号AppID校验不通过"
}

func NewWithErrType(errType ErrType) error {
	e := &YsErr{Code: -1, Msg: msgList[ERRTYPE_INTERNAL_SERVER]}
	if msg, ok := msgList[errType]; ok {
		e.Msg = msg
	}
	e.caller()
	return e
}

func GetErrMessage(errType ErrType) string {
	return msgList[errType]
}
