package helper

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func DecodeJson(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err == io.EOF {
		log.Warn().Msgf(util.LogErrDecode, err)
		return _error.HttpErrMapOfSlices(map[string][]string{
			"bad_request": {
				"tidak ada request body",
			},
		}, response.CM06)
	}

	if err != nil {
		log.Warn().Msgf(util.LogErrDecode, err)
		return err
	}

	return nil
}

func ParserMultipartForm(r *http.Request, data any) error {
	if err := r.ParseMultipartForm(3 << 20); err != nil {
		log.Warn().Msgf("failed parse multipart form data | err : %v", err)
		return _error.HttpErrString("unexpected end of multipart/form-data input", response.CM11)
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
