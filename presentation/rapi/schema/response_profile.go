package schema

type ResponseProfile struct {
	ProfileID string  `json:"profile_id"`
	Quote     *string `json:"quote"`
	Profesi   *string `json:"profesi"`
}
