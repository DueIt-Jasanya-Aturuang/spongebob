package converter

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
)

func NotifModelToResponse(n *domain.Notification) *domain.ResponseNotification {
	return &domain.ResponseNotification{
		ID:           n.ID,
		ProfileID:    n.ProfileID,
		UserConfigID: n.UserConfigID,
		Message:      n.Message,
		Title:        n.Title,
		Icon:         n.Icon,
		Status:       n.Status,
		CreatedAt:    n.CreatedAt,
	}
}
