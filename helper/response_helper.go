package helper

import (
	"encoding/json"
	"net/http"
	dto "waysbeans/dto/result"
)

func ResponseHelper(w http.ResponseWriter, err error, data interface{}, httpStatusError int) {
	if err != nil {
		w.WriteHeader(httpStatusError)
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)

}
