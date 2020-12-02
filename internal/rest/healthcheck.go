package rest

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"skat-vending.com/selection-info/internal/coder"
)

// Helthckeck godoc
// @Summary Returns service info
// @Description Returns service info
// @ID healthckeck
// @Tags sections
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} api.Healthcheck
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router / [get]
func (s *Service) healthcheck(w http.ResponseWriter, r *http.Request) {
	result := s.Health.Check(context.Background())
	if err := coder.WriteData(w, r, result, http.StatusOK); err != nil {
		logrus.WithError(err).Error("healthcheck writing response")
		return
	}
}
