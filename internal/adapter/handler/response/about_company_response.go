package response

type AboutCompanyResponse struct {
	ID              int64                         `json:"id"`
	Description     string                        `json:"description"`
	CompanyKeyNotes []AboutCompanyKeynoteResponse `json:"company_keynotes"`
}
