package response

type OurTeamResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	PathPhoto string `json:"path_photo"`
	Tagline   string `json:"tagline"`
}
