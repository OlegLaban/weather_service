package output

type PingDTO struct {
	Answer string `json:"result"`
}

func NewPing() *PingDTO {
	return &PingDTO{
		Answer: "pong",
	}
}
