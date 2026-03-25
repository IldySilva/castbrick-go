package castbrick

import (
	"context"
	"fmt"
	"time"
)

// SmsResource handles SMS operations.
type SmsResource struct{ c *Client }

type SendSmsOptions struct {
	To            []string
	Content       string
	SenderID      string
	ScheduledAt   *time.Time
	ContactListID string
}

func (r *SmsResource) Send(ctx context.Context, opts SendSmsOptions) (*SendSmsResponse, error) {
	body := map[string]any{
		"recipients": opts.To,
		"content":    opts.Content,
	}
	if opts.SenderID != "" {
		body["senderId"] = opts.SenderID
	}
	if opts.ScheduledAt != nil {
		body["scheduledAt"] = opts.ScheduledAt.UTC().Format(time.RFC3339)
	}
	if opts.ContactListID != "" {
		body["contactListId"] = opts.ContactListID
	}

	var out SendSmsResponse
	return &out, r.c.post(ctx, "/sms/send", body, &out)
}

func (r *SmsResource) Get(ctx context.Context, id string) (*SmsMessage, error) {
	var out SmsMessage
	return &out, r.c.get(ctx, "/sms/"+id, &out)
}

func (r *SmsResource) List(ctx context.Context, page, pageSize int) (*PagedResult[SmsMessage], error) {
	var out PagedResult[SmsMessage]
	return &out, r.c.get(ctx, fmt.Sprintf("/sms?pageNumber=%d&pageSize=%d", page, pageSize), &out)
}

func (r *SmsResource) CancelScheduled(ctx context.Context, messageID string) error {
	return r.c.post(ctx, "/sms/cancel-scheduled", map[string]any{"messageId": messageID}, nil)
}
