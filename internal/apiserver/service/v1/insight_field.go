package v1

import (
	"context"
	"github.com/RaulZuo/deep/internal/apiserver/store"
)

// InsightFieldSrv defines functions used to handle insight_field request.
type InsightFieldSrv interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	Get(ctx context.Context) error
}

type insightFieldService struct {
	store store.Factory
}

