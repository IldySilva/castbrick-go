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
	ID            string     `json:"id"`
	ContactName   *string    `json:"contactName,omitempty"`
	RecipientPhone string    `json:"recipientPhone"`
	Message       string     `json:"message"`
	CampaignName  *string    `json:"campaignName,omitempty"`
	CampaignID    *string    `json:"campaignId,omitempty"`
	SenderID      *string    `json:"senderId,omitempty"`
	Status        string     `json:"status"`
	ErrorMessage  *string    `json:"errorMessage,omitempty"`
	RetryCount    int        `json:"retryCount"`
	ScheduledAt   *time.Time `json:"scheduledAt,omitempty"`
	SentAt        *time.Time `json:"sentAt,omitempty"`
	DeliveredAt   *time.Time `json:"deliveredAt,omitempty"`
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
