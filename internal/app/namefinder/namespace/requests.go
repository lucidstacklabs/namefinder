package namespace

type CreationRequest struct {
	Name string `json:"name" bson:"name" binding:"required"`
}

type UpdateRequest struct {
	Name string `json:"name" bson:"name"`
}
