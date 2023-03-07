package response

type Responses interface {
	SetCode(int32)
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	SetDebug(string)
	Clone() Responses
}