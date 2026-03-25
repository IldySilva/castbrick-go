package castbrick

import (
	"context"
	"fmt"
	"time"
)

// BroadcastsResource handles broadcast operations.
type BroadcastsResource struct{ c *Client }

type CreateBroadcastOptions struct {
	Name          string
	Message       string
	ContactListID string
	SenderID      string
}

type UpdateBroadcastOptions struct {
	Name          string
	Message       string
	ContactListID string
	SenderID      string
	ScheduleAt    *time.Time
}

func (r *BroadcastsResource) List(ctx context.Context, page, pageSize int) (*PagedResult[Broadcast], error) {
	var out PagedResult[Broadcast]
	return &out, r.c.get(ctx, fmt.Sprintf("/broadcasts?pageNumber=%d&pageSize=%d", page, pageSize), &out)
}

func (r *BroadcastsResource) Get(ctx context.Context, id string) (*Broadcast, error) {
	var out Broadcast
	return &out, r.c.get(ctx, "/broadcasts/"+id, &out)
}

func (r *BroadcastsResource) Create(ctx context.Context, opts CreateBroadcastOptions) (string, error) {
	body := map[string]any{
		"name":    opts.Name,
		"message": opts.Message,
	}
	if opts.ContactListID != "" {
		body["contactListId"] = opts.ContactListID
	}
	if opts.SenderID != "" {
		body["senderId"] = opts.SenderID
	}
	var out string
	return out, r.c.post(ctx, "/broadcasts", body, &out)
}

func (r *BroadcastsResource) Update(ctx context.Context, id string, opts UpdateBroadcastOptions) (string, error) {
	body := map[string]any{
		"name":    opts.Name,
		"message": opts.Message,
	}
	if opts.ContactListID != "" {
		body["contactListId"] = opts.ContactListID
	}
	if opts.SenderID != "" {
		body["senderId"] = opts.SenderID
	}
	if opts.ScheduleAt != nil {
		body["scheduleAt"] = opts.ScheduleAt.UTC().Format(time.RFC3339)
	}
	var out string
	return out, r.c.put(ctx, "/broadcasts/"+id, body, &out)
}

func (r *BroadcastsResource) Send(ctx context.Context, id string) error {
	return r.c.post(ctx, "/broadcasts/"+id+"/send", map[string]any{}, nil)
}

func (r *BroadcastsResource) Cancel(ctx context.Context, id string) error {
	return r.c.post(ctx, "/broadcasts/"+id+"/cancel", map[string]any{}, nil)
}

func (r *BroadcastsResource) Duplicate(ctx context.Context, id string) (string, error) {
	var out string
	return out, r.c.post(ctx, "/broadcasts/"+id+"/duplicate", map[string]any{}, &out)
}

func (r *BroadcastsResource) Delete(ctx context.Context, id string) error {
	return r.c.delete(ctx, "/broadcasts/"+id)
}
