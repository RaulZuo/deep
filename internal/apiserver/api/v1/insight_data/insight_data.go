package insight_data

import "github.com/RaulZuo/deep/internal/apiserver/store"

// InsightDataHandler create an insight_data handler used to handle request for insight data.
type InsightDataHandler struct {
	store store.Factory
}

// NewInsightDataHandler creates an insight_data handler
func NewInsightDataHandler(store store.Factory) *InsightDataHandler {
	return &InsightDataHandler{
		store: store,
	}
}


