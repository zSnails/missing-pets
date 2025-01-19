package response

type Response[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
