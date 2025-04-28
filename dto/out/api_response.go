package out

type WebResponse struct {
	Data     interface{}  `json:"data,omitempty"`
	Status   WebStatus    `json:"status"`
	MetaData *WebMetaData `json:"meta_data,omitempty"`
}

type WebStatus struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Detail  interface{} `json:"detail,omitempty"`
}

type WebMetaData struct {
	TotalData int `json:"total_data"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
}
