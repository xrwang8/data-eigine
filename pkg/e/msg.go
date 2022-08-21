package e

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "fail",
	INVALID_PARAMS:            "请求参数错误",
	ERROR_ADD_DATASET_FAIL:    "新增数据集失败",
	ERROR_DELETE_DATASET_FAIL: "删除数据集失败",
	ERROR_SELECT_DATASET_FAIL: "查找数据集失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
