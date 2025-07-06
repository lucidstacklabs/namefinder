package apikey

type CreationRequest struct {
	Name string `json:"name" binding:"required"`
}
