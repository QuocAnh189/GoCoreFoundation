package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants"
)

func WriteJson(w http.ResponseWriter, data any, err error) {
	payload := make(map[string]any)

	// If there's data, try to unmarshal data into being the payload
	if data != nil {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Printf("WriteJson: failed to marshal data: %v\n", err)
			return
		}
		var tmp map[string]any
		err = json.Unmarshal(dataBytes, &tmp)
		if err != nil || tmp == nil {
			// If this fails, just add the data to an empty payload as "result"
			payload["result"] = data
		} else {
			payload = tmp
		}
	}

	if err != nil {
		var appErr *AppError
		if errors.As(err, &appErr) {
			payload["mmessage"] = appErr.Message
			payload["status"] = appErr.Status
		}
	} else {
		payload["status"] = http.StatusOK
	}

	// Default to not set if not set
	if payload["status"] == 0 {
		payload["status"] = constants.UNKNOW
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
