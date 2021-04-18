package stein

type UpdateRequest struct {
	Condition map[string]interface{} `json:"condition"`
	Set       map[string]interface{} `json:"set"`
}

type DeleteRequest struct {
	Condition map[string]interface{} `json:"condition"`
}
