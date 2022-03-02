package formatter

import (
	"encoding/json"
	"fmt"
	"julo-case-study/http/enum"
	"julo-case-study/models/responses"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader((statusCode))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SUCCESS(w http.ResponseWriter, statusCode int, payload interface{}) {
	data := make(map[string]interface{})
	marshalled, _ := json.Marshal(payload)
	json.Unmarshal(marshalled, &data)
	response := responses.Response{
		Data:   data,
		Status: enum.ResponseSuccess,
	}
	JSON(w, statusCode, response)
}

func FAIL(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		errorMessage := make(map[string]interface{})
		errorMessage["error"] = err.Error()
		response := responses.Response{
			Data:   errorMessage,
			Status: enum.ResponseFail,
		}
		JSON(w, statusCode, response)
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
