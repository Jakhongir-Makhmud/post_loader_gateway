package structs

var ErrInternalResponse = struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}{
	Status: "Failed",
	Msg:    "Internal server error occured",
}
