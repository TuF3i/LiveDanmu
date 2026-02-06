package response

type FinalResponse struct {
	Status uint        `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

type Response struct {
	Status uint   `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string {
	return r.Info
}

func InternalError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}
