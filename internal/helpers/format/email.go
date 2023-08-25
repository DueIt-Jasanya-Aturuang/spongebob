package format

import (
	"fmt"
	"strings"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/rs/zerolog/log"
)

func EmailFormat(email string) (string, error) {
	emailArr := strings.Split(email, "@")
	if len(emailArr) != 2 {
		log.Err(exception.Err500InvalidFormatEmail).Msgf("INVALID EMAIL : %s", email)
		return "", exception.Err500InvalidFormatEmail
	}
	return fmt.Sprintf("%c••••%c@%s", emailArr[0][0], emailArr[0][len(emailArr[0])-1], emailArr[1]), nil
}
