package request

type ServiceSectionRequest struct {
	PathIcon string `json:"path_icon" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Tagline  string `json:"tag_line" validate:"required"`
}
