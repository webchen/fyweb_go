package core

// jsonMessage , 所有的信息输出，都是基于该结构体
type httpMessage struct {
	// http返回状态码
	statusCode int
	// 输出的数据
	jsonData *jsonMessage
}

// jsonMessage , 对应输出的数据
// 里面必须包含code参数，表明业务数据的状态
// 必须包含message参数，表明对该返回数据的消息说明
type jsonMessage struct {
	code    int
	message string
	data    interface{}
}

func (j *jsonMessage) ReSet() *jsonMessage {
	j.code = -1
	j.message = "no data"
	j.data = nil
	return j
}

func (j *jsonMessage) ToMap() map[string]interface{} {
	output := make(map[string]interface{}, 3)
	output["code"] = j.code
	output["message"] = j.message
	output["data"] = j.data
	return output
}
