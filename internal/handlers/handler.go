package handlers

// HTTPHandler satisfies "Handler" interface and used to provide all handler methods via basic http connection.
type HTTPHandler struct {
	controller Controller
}

// NewHTTPHandler returns new HTTPHandler, that will use provided controller.
func NewHTTPHandler(controller Controller) *HTTPHandler {
	return &HTTPHandler{controller: controller}
}
