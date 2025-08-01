package request

type ServiceDetailRequest struct {
	ServiceID   int64   `json:"service_id" validate:"required"`
	PathImage   string  `json:"path_image" validate:"required`
	Title       string  `json:"title" validate:"required`
	Description string  `json:"description" validate:"required`
	PathPdf     *string `json:"path_pdf" validate:"required`
	PathDocx    *string `json:"path_docx" validate:"required`
	ServiceName string  `json:"service_name" validate:"required`
}
