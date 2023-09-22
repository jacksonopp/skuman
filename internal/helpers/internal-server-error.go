package helpers

import (
	"net/http"

	"github.com/jacksonopp/skuman/internal/logger"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorf("error: %v. path: %s", err, r.URL.Path)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
