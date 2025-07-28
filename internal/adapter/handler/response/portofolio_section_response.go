package response

type PortofolioSectionResponse struct {
	ID        int64  `json:"id"`
	Thumbnail string `json:"path_icon"`
	Name      string `json:"name"`
	Tagline   string `json:"tagline"`
}
