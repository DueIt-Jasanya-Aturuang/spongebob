package response

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helper/message"
)

func DecodeReq(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err == io.EOF {
		return Err400(map[string][]string{
			"bad_request": {
				"empty body request",
			},
		}, err)
	}

	if err != nil {
		return err
	}

	return nil
}

func ParserMultipartForm(r *http.Request, data any) error {
	if err := r.ParseMultipartForm(3 << 20); err != nil {
		log.Err(err).Msg(message.ErrParseForm)
		return Err400(map[string][]string{
			"unexpected": {
				"unexpected end of multipart/form-data input",
			},
		}, err)
	}

	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("form")
		if tag != "" {
			switch val.Field(i).Kind() {
			case reflect.String:
				formField := r.FormValue(tag)
				val.Field(i).SetString(formField)
			case reflect.Ptr:
				if val.Field(i).Type() == reflect.TypeOf(&multipart.FileHeader{}) {
					_, fileHeader, err := r.FormFile(tag)
					if err != nil {
						log.Warn().Msgf("error form file : %v", err)
					}
					val.Field(i).Set(reflect.ValueOf(fileHeader))
				}
			}
		}
	}
	return nil
}
