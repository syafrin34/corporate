package response

type ContactUsResponse struct {
	ID           int64  `json:"id"`
	CompanyName  string `json:"company_name"`
	LocationName string `json:"location_name"`
	Address      string `json:"addres"`
	PhoneNumber  string `json:"phone_number"`
}
