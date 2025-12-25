package accountabstractionhandler

type Response struct {
	Success  bool   `json:"success"`
	Errors   string `json:"errors"`
	Returned any    `json:"returned"`
}
