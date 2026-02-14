package service

type SkipReason string

const (
	SkipDuplicate     SkipReason = "duplicate"
	SkipAlreadyQueued SkipReason = "already_queued"
	SkipRateLimited   SkipReason = "rate_limited"
	SkipSendError     SkipReason = "send_error"
)
