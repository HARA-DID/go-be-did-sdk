package accountabstractionhandler

type Response struct {
	Success  bool   `json:"success" example:"true"`
	Errors   string `json:"errors" example:"No Error Message"`
	Returned any    `json:"returned,omitempty"`
}
