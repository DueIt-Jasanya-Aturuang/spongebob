package dto

type ProfileResp struct {
	ProfileID string `json:"profile_id"`
	Quote     string `json:"quote"`
}

type GetProfileReq struct {
	UserID string // request header 'id'
}
type StoreProfileReq struct {
	UserID string `json:"user_id"`
}
