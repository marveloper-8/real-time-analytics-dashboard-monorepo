package graph

import (
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	AnalyticsService *service.AnalyticsService
}
