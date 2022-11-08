package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "waysbeans/dto/result"
)

func ResponseHelper(w http.ResponseWriter, err error, data interface{}) {
	fmt.Println("halo")
	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: "success", Data: data}
		json.NewEncoder(w).Encode(response)
	}
}
