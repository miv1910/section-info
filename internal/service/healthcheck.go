package service

import (
	"context"
	"database/sql"
	"strings"

	"skat-vending.com/selection-info/internal/dal"
	"skat-vending.com/selection-info/pkg/api"
)

// Health represents service for health check
type Health struct {
	dal  *dal.Sections
	conn string
}

// NewHealth returns new instance of Health service
func NewHealth(db *sql.DB, conn string) *Health {
	return &Health{
		dal:  dal.NewSection(db),
		conn: conn,
	}
}

// GetM0 for get
func (s *Health) Check(ctx context.Context) api.Healthcheck {
	var (
		DBName *string
		DBAddr *string
		DBTime string
	)
	dt, err := s.dal.GetDate(ctx)
	if err != nil {
		DBTime = err.Error()
	} else {
		DBTime = *dt
	}

	mcs := strings.Split(s.conn, ";")
	for _, item := range mcs {
		params := strings.Split(item, "=")
		if len(params) == 2 {
			switch {
			case params[0] == `database`:
				DBName = &params[1]
			case params[0] == `server`:
				DBAddr = &params[1]
			}
		}
	}

	response := api.Healthcheck{
		Version:  api.Version,
		Name:     api.Name,
		RootPath: api.Path,
		DBAddr:   DBAddr,
		DBName:   DBName,
		DBTime:   &DBTime,
	}

	return response
}
