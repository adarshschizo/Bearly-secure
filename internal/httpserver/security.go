package httpserver

import (
	"net/http"
	"time"
)

func securityTxt(responseWriter http.ResponseWriter, _ *http.Request) {
	expires := time.Now().UTC().Add(180 * 24 * time.Hour).Format(time.RFC3339)

	responseWriter.Header().Set("Content-Type", "text/plain")
	responseWriter.WriteHeader(http.StatusOK)

	_, _ = responseWriter.Write([]byte(
		"Contact: mailto:security@bearlysecure.example\n" +
			"Policy: https://bearlysecure.example/security-policy\n" +
			"Expires: " + expires + "\n\n",
	))
}
