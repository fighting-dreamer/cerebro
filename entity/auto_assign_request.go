package entity

type AutoAssignRequest struct {
	ApplicationID string            `json:"id"`
	Type          string            `json:"type"`
	Params        map[string]string `json:"params"`
}
