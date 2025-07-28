package response

type ServiceSectionResponse struct {
	ID       int64  `json:"id"`
	PathIcon string `json:"path_icon"`
	Name     string `json:"name"`
	Tagline  string `json:"Tagline"`
}
