package castbrick

import "time"

// PagedResult is a generic paginated response.
type PagedResult[T any] struct {
	Items          []T  `json:"items"`
	TotalCount     int  `json:"totalCount"`
	PageNumber     int  `json:"pageNumber"`
	TotalPages     int  `json:"totalPages"`
	HasNextPage    bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}

// ── SMS ──────────────────────────────────────────────────────────────────────

type SendSmsResponse struct {
	MessageID      string `json:"messageId"`
	Status         string `json:"status"`
	RecipientCount int    `json:"recipientCount"`
	Error          string `json:"error,omitempty"`
	Timestamp      string `json:"timestamp"`
}

type SmsMessage struct {
	ID          string     `json:"id"`
	PhoneNumber string     `json:"phoneNumber"`
	Message     string     `json:"message"`
	Status      string     `json:"status"`
	TenantID    string     `json:"tenantId"`
	SentAt      *time.Time `json:"sentAt,omitempty"`
	DeliveredAt *time.Time `json:"deliveredAt,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
}

// ── Contacts ─────────────────────────────────────────────────────────────────

type Contact struct {
	ID          string    `json:"id"`
	Name        string    `json:"name,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Email       string    `json:"email,omitempty"`
	TenantID    string    `json:"tenantId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ContactList struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TenantID     string    `json:"tenantId"`
	ContactCount int       `json:"contactCount"`
	CreatedAt    time.Time `json:"createdAt"`
}

// ── Broadcasts ───────────────────────────────────────────────────────────────

type Broadcast struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Status        string     `json:"status"`
	Message       string     `json:"message"`
	SenderID      string     `json:"senderId,omitempty"`
	ContactListID string     `json:"contactListId,omitempty"`
	ScheduledAt   *time.Time `json:"scheduledAt,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
}
