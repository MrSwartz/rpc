package encoder

import (
	"encoding/json"
	"fmt"
	"os"
	"rpc/sharedLib/models"
)

func MarshallCompany(payload models.Company, method string, code int32) []byte {
	var resp models.ApiResponseCompany

	resp.Type = method
	resp.Code = code
	resp.Message = payload

	if obj, err := json.Marshal(resp); err == nil {
		return obj
	} else {
		fmt.Fprintf(os.Stderr, "Error in json mashalling : %v\n", err)
	}
	return nil
}

func UnMarshallCompany(payload []byte) (parse models.ApiResponseCompany) {
	if err := json.Unmarshal(payload, &parse); err != nil {
		fmt.Fprintf(os.Stderr, "Error in json unmarshalling : %v\n", err)
	}
	return parse
}

func MarshallEmployee(payload models.Employee, method string, code int32) []byte {
	var resp models.ApiResponseEmployee

	resp.Type = method
	resp.Code = code
	resp.Message = payload

	if obj, err := json.Marshal(resp); err == nil {
		return obj
	} else {
		fmt.Fprintf(os.Stderr, "Error in json mashalling : %v\n", err)
	}
	return nil
}

func UnMarshallEmployee(payload []byte) (parse models.ApiResponseEmployee) {
	if err := json.Unmarshal(payload, &parse); err != nil {
		fmt.Fprintf(os.Stderr, "Error in json unmarshalling : %v\n", err)
	}
	return parse
}
