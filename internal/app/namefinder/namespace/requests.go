package namespace

type CreationRequest struct {
	Name string `json:"name" bson:"name" binding:"required"`
}

type UpdateRequest struct {
	Name string `json:"name" bson:"name"`
}

type ApiKeyAccessRequest struct {
	ApiKeyID string   `json:"api_key_id" binding:"required"`
	Actions  []Action `json:"actions" binding:"required"`
}

type ApiKeyAccessDestroyRequest struct {
	ApiKeyID string `json:"api_key_id" binding:"required"`
}
