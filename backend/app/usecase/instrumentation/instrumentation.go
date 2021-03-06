package instrumentation

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

// Instrumentation measures the internal operation of the system.
type Instrumentation struct {
	logger    fw.Logger
	tracer    fw.Tracer
	timer     fw.Timer
	metrics   fw.Metrics
	analytics fw.Analytics
	ctx       fw.ExecutionContext
}

// RedirectingAliasToLongLink tracks RedirectingAliasToLongLink event.
func (i Instrumentation) RedirectingAliasToLongLink(user *entity.User) {
	go func() {
		userID := i.getUserID(user)
		i.analytics.Track("RedirectingAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// RedirectedAliasToLongLink tracks RedirectedAliasToLongLink event.
func (i Instrumentation) RedirectedAliasToLongLink(user *entity.User) {
	go func() {
		userID := i.getUserID(user)
		i.analytics.Track("RedirectedAliasToLongLink", map[string]string{}, userID, i.ctx)
	}()
}

// LongLinkRetrievalSucceed tracks the successes when retrieving long links.
func (i Instrumentation) LongLinkRetrievalSucceed() {
	go func() {
		i.metrics.Count("long-link-retrieval-succeed", 1, 1, i.ctx)
	}()
}

// LongLinkRetrievalFailed tracks the failures when retrieving long links.
func (i Instrumentation) LongLinkRetrievalFailed(err error) {
	go func() {
		i.logger.Error(err)
		i.metrics.Count("long-link-retrieval-failed", 1, 1, i.ctx)
	}()
}

func (i Instrumentation) getUserID(user *entity.User) string {
	if user == nil {
		return i.ctx.RequestID
	}
	return user.Email
}

// NewInstrumentation initializes instrumentation code.
func NewInstrumentation(logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	metrics fw.Metrics,
	analytics fw.Analytics,
	ctx fw.ExecutionContext,
) Instrumentation {
	return Instrumentation{
		logger:    logger,
		tracer:    tracer,
		timer:     timer,
		metrics:   metrics,
		analytics: analytics,
		ctx:       ctx,
	}
}
