package service

import (
	"context"
	"time"

	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/database"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/graph/model"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/pkg/utils"
)

type AnalyticsService struct {
	DB			*database.InfluxDB
	Logger	*utils.Logger
}

func NewAnalyticsService(db *database.InfluxDB, logger *utils.Logger) *AnalyticsService {
	return &AnalyticsService{
		DB:			db,
		Logger:	logger,
	}
}

func (s *AnalyticsService) GetAnalytics(ctx context.Context, startTime, endTime string) (*model.Analytics, error)  {
	return &model.Analytics{
		TotalVisits: 				1000,
		UniqueVisitors:			500,
		AvgSessionDuration:	180.5,
		BounceRate:					0.25,
		TopPages: []*model.PageVisit{
			{URL: "/home", Visits: 300},
			{URL: "/products", Visits: 200},
			{URL: "/about", Visits: 100},
		},
	}, nil
}

func (s *AnalyticsService) SubscribeToAnalytics(ctx context.Context) (<-chan *model.Analytics, error)  {
	ch := make(chan *model.Analytics)

	go func ()  {
		defer close(ch)

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				analytics, err := s.GetAnalytics(ctx, "", "")
				if err != nil {
					s.Logger.Error("Error fetching analytics")
					continue
				}
				ch <- analytics
			}
		}
	}()

	return ch, nil
}