package dns

type RecordAdditionRequest struct {
	Name  string      `json:"name" binding:"required"`
	Type  RecordType  `json:"type" binding:"required"`
	Value string      `json:"value" binding:"required"`
	TTL   uint32      `json:"ttl" binding:"required"`
	Class RecordClass `json:"class" binding:"required"`
}

type RecordUpdateRequest struct {
	Name  string      `json:"name"`
	Type  RecordType  `json:"type"`
	Value string      `json:"value"`
	TTL   uint32      `json:"ttl"`
	Class RecordClass `json:"class"`
}
