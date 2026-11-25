package ship

import "net/http"

type Container interface {
	Name() string
	RegisterRoutes(mux *http.ServeMux)
}
