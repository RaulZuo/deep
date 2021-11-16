package insight_field

import "github.com/RaulZuo/deep/internal/apiserver/store"

// InsightFieldHandler creates an insight_field handler used to handle request for insight field.
type InsightFieldHandler struct {
	store store.Factory
}

// NewInsightFieldHandler creates an insight_field handler.
func NewInsightFieldHandler(store store.Factory) *InsightFieldHandler {
	return &InsightFieldHandler{
		store: store,
	}
}