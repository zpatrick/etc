package main

import (
	"encoding/json"
	"net/http"
)

type CertificateHandler struct {
	logger  logutil.Logger
	creator certificate.Creator
}

func (c *CertificateHandler) Create(r *http.Request) http.Handler {
	logger := logutil.WithMiddleware(c.logger, func(e *logutil.Event) *logutil.Event {
		e.Tags[LogTagHandler] = "Create"
		return e
	})

	var req CreateCertificateRequest
	if err := json.Unmarshal(r.Body, &req); err != nil {
		return handleError(logger, w, err)
	}

	cert, err := c.creator.Create(logger, req.OrgGUID, req.Domain)
	if err != nil {
		return handleError(logger, w, err)
	}

	return handle.JSON(w, http.StatusCreated, cert)
}
