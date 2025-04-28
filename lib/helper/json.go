package helper

import (
	"encoding/json"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"log"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}, statusCode int) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	return err
}

func WriteErrorResponse(w http.ResponseWriter, err error, detail interface{}) {
	exceptionModel, ok := err.(exception.CustomError)
	if !ok {
		exceptionModel = exception.InternalServerError
		if detail == nil {
			detail = err
		}
	}

	output := out.WebResponse{
		Status: out.WebStatus{
			Message: exceptionModel.Message,
			Code:    exceptionModel.Code,
			Detail:  detail,
		},
	}

	err = WriteToResponseBody(w, output, exceptionModel.HttpCode)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}, message string, code int, meta *out.WebMetaData) {
	output := out.WebResponse{
		Data: data,
		Status: out.WebStatus{
			Message: message,
		},
		MetaData: meta,
	}

	err := WriteToResponseBody(w, output, code)
	if err != nil {
		log.Fatal(err)
	}
}
