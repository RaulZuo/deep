package v1

import (
	"context"
	"github.com/RaulZuo/deep/internal/apiserver/store"
)

// InsightDataSrv defines functions used to handle insight_data request.
type InsightDataSrv interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	Get(ctx context.Context) error
}

type insightDataService struct {
	store store.Factory
}

