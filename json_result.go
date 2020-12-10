package simple

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
	Time      int64       `json:"time"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
		Time:      NowTimestamp(),
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
		Time:      NowTimestamp(),
	}
}

func JsonPageData(results interface{}, page *Paging) *JsonResult {
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonCursorData(results interface{}, cursor string) *JsonResult {
	return JsonData(&CursorResult{
		Results: results,
		Cursor:  cursor,
	})
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      true,
		Success:   true,
		Time:      NowTimestamp(),
	}
}

func JsonError(err *CodeError) *JsonResult {
	return &JsonResult{
		ErrorCode: err.Code,
		Message:   err.Message,
		Data:      err.Data,
		Success:   false,
		Time:      NowTimestamp(),
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
		Time:      NowTimestamp(),
	}
}

func JsonErrorCode(code int, message string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
		Time:      NowTimestamp(),
	}
}

func JsonErrorData(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   false,
		Time:      NowTimestamp(),
	}
}

type RspBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func NewRspBuilder(obj interface{}) *RspBuilder {
	return NewRspBuilderExcludes(obj)
}

func NewRspBuilderExcludes(obj interface{}, excludes ...string) *RspBuilder {
	return &RspBuilder{Data: StructToMap(obj, excludes...)}
}

func (builder *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) Build() map[string]interface{} {
	return builder.Data
}

func (builder *RspBuilder) JsonResult() *JsonResult {
	return JsonData(builder.Data)
}
