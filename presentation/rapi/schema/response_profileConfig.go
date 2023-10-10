package schema

type ResponseProfileConfig struct {
	ID          string   `json:"profile_config_id"`
	ProfileID   string   `json:"profile_id"`
	ConfigName  string   `json:"config_name"`
	ConfigValue string   `json:"config_value"`
	Status      string   `json:"status"`
	Days        []string `json:"days,omitempty"`
	Token       string   `json:"token"`
}
