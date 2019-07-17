package controllers

import (
	k "blueDigitalKwegonian/kwengo"
	"encoding/json"
	"fmt"
	"net/http"
)

// ConvertKwengonianToDecimal return a decimal conversion of a kwengonian number
func ConvertKwengonianToDecimal(rw http.ResponseWriter, req *http.Request) {
	var kenRaw k.Kwengo
	var response struct {
		Success  bool   `json:"success"`
		Message  string `json:"decimal"`
		ErrorMsg string `json:"error_msg"`
	}
	responseCode := 400
	if err := json.NewDecoder(req.Body).Decode(&kenRaw); err != nil {
		response.ErrorMsg = err.Error()
	} else {
		if list, res, err := kenRaw.CheckKwego(); res {
			if num, res, err := k.ConvertDecimal(list); res {
				responseCode = 200
				response.Success = true
				response.Message = fmt.Sprintf("%v", num)
			} else {
				response.ErrorMsg = err.Error()
			}
		} else {
			response.ErrorMsg = err.Error()
		}
	}
	rw.WriteHeader(responseCode)
	json.NewEncoder(rw).Encode(&response)
}
