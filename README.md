# castbrick-go

Official Go SDK for the [CastBrick](https://castbrick.com) API — send SMS, manage contacts and run broadcasts.

## Installation

```bash
go get github.com/IldySilva/castbrick-go
```

## Quick start

```go
import "github.com/IldySilva/castbrick-go/castbrick"

cb := castbrick.New("your_api_key")

result, err := cb.SMS.Send(ctx, castbrick.SendSmsOptions{
    To:      []string{"+244923000000"},
    Content: "Hello from CastBrick!",
})
```

## SMS

```go
// Send
result, err := cb.SMS.Send(ctx, castbrick.SendSmsOptions{
    To:       []string{"+244923000000"},
    Content:  "Your OTP is 1234",
    SenderID: "MyApp", // optional
})

// Get
msg, err := cb.SMS.Get(ctx, "message-id")

// List
page, err := cb.SMS.List(ctx, 1, 20)
fmt.Println(page.TotalCount)

// Cancel scheduled
err = cb.SMS.CancelScheduled(ctx, "message-id")
```

## Contacts

```go
// List (with optional search)
page, err := cb.Contacts.List(ctx, 1, 20, "john")

// Get
contact, err := cb.Contacts.Get(ctx, "contact-id")

// Create
created, err := cb.Contacts.Create(ctx, "", "+244923000000,+244912000000")

// Delete
err = cb.Contacts.Delete(ctx, "contact-id")

// Contact lists
list, err := cb.Contacts.CreateList(ctx, "VIP Customers")
err = cb.Contacts.AddToList(ctx, list.ID, contact.ID)
err = cb.Contacts.RemoveFromList(ctx, list.ID, contact.ID)
```

## Broadcasts

```go
// Create
id, err := cb.Broadcasts.Create(ctx, castbrick.CreateBroadcastOptions{
    Name:    "Black Friday",
    Message: "50% off everything today!",
})

// Send
err = cb.Broadcasts.Send(ctx, id)

// Update with schedule
t := time.Date(2026, 11, 28, 9, 0, 0, 0, time.UTC)
_, err = cb.Broadcasts.Update(ctx, id, castbrick.UpdateBroadcastOptions{
    Name:       "Black Friday",
    Message:    "50% off everything today!",
    ScheduleAt: &t,
})

// Other operations
err = cb.Broadcasts.Cancel(ctx, id)
newID, err := cb.Broadcasts.Duplicate(ctx, id)
err = cb.Broadcasts.Delete(ctx, id)
```

## Error handling

```go
_, err := cb.SMS.Send(ctx, ...)
var apiErr *castbrick.APIError
if errors.As(err, &apiErr) {
    fmt.Printf("%d: %s\n", apiErr.StatusCode, apiErr.Body)
}
```

## Custom HTTP client

```go
cb := castbrick.NewWithOptions("your_api_key", "https://api.castbrick.com/v1", &http.Client{
    Timeout: 10 * time.Second,
})
```

## License

MIT
