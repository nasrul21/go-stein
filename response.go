package stein

type InsertReponse struct {
	UpdatedRange string `json:"updatedRange"`
}

type UpdateReponse struct {
	UpdatedRange string `json:"updatedRange"`
}

type DeleteResponse struct {
	ClearedRowsCount int `json:"clearedRowsCount"`
}
