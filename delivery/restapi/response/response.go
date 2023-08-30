package response

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewError(w http.ResponseWriter, _ *http.Request, err error) {
	var (
		errHTTP            *model.ErrHTTP
		unmarshalTypeError *json.UnmarshalTypeError
		syntaxError        *json.SyntaxError
		pqError            *pq.Error
	)
	switch {
	case errors.As(err, &unmarshalTypeError):
		if err.Error() != "failed to decode: schema: converter not found for multipart.FileHeader" {
			log.Info().Msgf("error pq : %v", err)
			unprocessableEntity := map[string]string{
				err.(*json.UnmarshalTypeError).Field: fmt.Sprintf(
					"invalid type input, type must be %s", err.(*json.UnmarshalTypeError).Type.String(),
				),
			}
			err = Err422(unprocessableEntity, err)
			break
		}
		log.Info().Msg(err.Error())
		unprocessableEntity := map[string]string{
			"image": "invalid file",
		}
		err = Err422(unprocessableEntity, err)

	case errors.As(err, &syntaxError):
		log.Info().Msgf("error pq : %v", err)
		err = Err400(map[string][]string{
			"unexpected": {
				"unexpected end of json input",
			},
		}, err)

	case errors.As(err, &pqError):
		err = Err500("INTERNAL SERVER ERROR", err)

	case errors.As(err, &errHTTP):
		break

	case errors.Is(err, context.DeadlineExceeded):
		err = Err408("REQUEST TIME OUT", err)

	case errors.Is(err, sql.ErrNoRows):
		err = Err404("DATA NOT FOUND", err)

	//	type error checking
	case errors.Is(err, model.ErrUnauthorization):
		err = Err401(err.Error(), err)
	case errors.Is(err, model.ErrUnauthorization):
		err = Err401(err.Error(), err)

	case errors.Is(err, model.ErrForbidden):
		err = Err403(err.Error(), err)

	default:
		log.Info().Msgf("error default : %v", err)
		err = Err500("INTERNAL SERVER ERROR", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.(*model.ErrHTTP).Code)
	resp := model.ResponseError{
		Errors:     &err.(*model.ErrHTTP).Message,
		ErrorsReal: err.(*model.ErrHTTP).Err,
	}
	if errEncode := json.NewEncoder(w).Encode(resp); errEncode != nil {
		log.Err(errEncode).Msg(message.ErrEncodeJson)
	}

}

func NewSucc(w http.ResponseWriter, _ *http.Request, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Err(err).Msg(message.ErrEncodeJson)
	}
}
