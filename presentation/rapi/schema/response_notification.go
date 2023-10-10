package schema

type ResponseNotification struct {
	ID           string `json:"id"`
	ProfileID    string `json:"profile_id"`
	UserConfigID string `json:"user_config_id"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Icon         string `json:"icon"`
	Status       string `json:"status"`
	CreatedAt    int64  `json:"created_at"`
}
