package api

// Section represents info about section
type SectionRequest struct {
	IdOtdel    int `json:"id_otdel" example:"123"`
	IdRazdel   int `json:"id_razdel" example:"123123"`
	IdOperator int `json:"id_operator" example:"2"`
}
