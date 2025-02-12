package main

// type EchoHandler struct {
// 	log *zap.Logger
// }

// func NewEchoHandler(log *zap.Logger) *EchoHandler {
// 	return &EchoHandler{log: log}
// }

// func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	_, wErr := io.Copy(w, r.Body)
// 	if wErr != nil {
// 		h.log.Error("failed to write response", zap.Error(wErr))

// 		return
// 	}

// 	h.log.Info("handler called successfully")
// }
