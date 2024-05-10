package controllers

type Resp struct {
	Status int
	Body []byte
}

func newResponse(status int, body []byte) *Resp {
	return &Resp{ status, body }
}