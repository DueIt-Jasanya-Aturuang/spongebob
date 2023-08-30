package helpers

import (
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"strings"

	"github.com/rs/zerolog/log"
)

func EmailFormat(email string) string {
	emailArr := strings.Split(email, "@")
	if len(emailArr) != 2 {
		log.Err(fmt.Errorf("%s", message.ErrInvalidEmail)).Msgf("INVALID EMAIL : %s", email)
		return email
	}
	return fmt.Sprintf("%c••••%c@%s", emailArr[0][0], emailArr[0][len(emailArr[0])-1], emailArr[1])
}
