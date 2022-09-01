package dto

type CustomerResponse struct {
	Id      string `json:"id" xml:"id"`
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip"`
	Birth   string `json:"birth" xml:"birth"`
	Status  string `json:"status" xml:"status"`
}
