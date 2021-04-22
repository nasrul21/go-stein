package stein

type InsertReponse struct {
	UpdatedRange string `json:"updatedRange"`
}

type UpdateReponse struct {
	TotalUpdatedRows int `json:"totalUpdatedRows"`
}

type DeleteResponse struct {
	ClearedRowsCount int `json:"clearedRowsCount"`
}
