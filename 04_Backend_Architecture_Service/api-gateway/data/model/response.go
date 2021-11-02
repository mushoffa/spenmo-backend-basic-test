package model

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type BaseApiResponse struct {
	Data interface{} `json:"data,omitempty"`
	Meta `json:"meta"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type Meta struct {
	Code 		int 	`json:"code,omitempty"`
	Message 	string 	`json:"message,omitempty"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (m Meta) GetDefaultError() Meta {
	return Meta{
		Code:    400,
		Message: m.Message,
	}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (m Meta) GetDefaultSuccess() Meta {
	return Meta{
		Code:    200,
		Message: m.Message,
	}
}