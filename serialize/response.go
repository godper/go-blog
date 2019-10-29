package serialize

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TokenResponse 更新token
type TokenResponse struct {
	*Response
	NewToken string `json:"newtoken"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	TrackID string `json:"track_id"`
}

//BuildTokenRespon 绑定newtoken
func BuildTokenRespon(response *Response, token string) *TokenResponse {
	return &TokenResponse{
		response,
		token,
	}
}
