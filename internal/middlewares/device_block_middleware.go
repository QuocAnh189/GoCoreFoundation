package middlewares

import (
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/block"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

func DeviceBlockMiddleware(blockSvc *block.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			deviceUUID := r.Header.Get("Device-UUID")
			println(deviceUUID)

			ctx := r.Context()
			statusCode, _, err := blockSvc.CheckBlockByValue(ctx, deviceUUID)
			if err != nil {
				response.WriteJson(w, ctx, nil, err, statusCode)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
