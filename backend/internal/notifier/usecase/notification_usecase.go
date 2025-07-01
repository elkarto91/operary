package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/elkarto91/operary/internal/notifier/model"
	"github.com/elkarto91/operary/internal/notifier/repo"
)

// NotificationRequest is the payload for sending a manual notification.
type NotificationRequest struct {
	UserID  string `json:"user_id"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// Send dispatches a notification via the requested channel and logs the result.
func Send(req NotificationRequest) (model.Notification, error) {
	if req.UserID == "" || req.Channel == "" {
		return model.Notification{}, errors.New("missing user or channel")
	}
	n := model.Notification{
		UserID:    req.UserID,
		Channel:   req.Channel,
		Message:   req.Message,
		Timestamp: time.Now(),
	}
	var err error
	switch req.Channel {
	case "email":
		err = sendEmail(req.UserID, req.Message)
	case "push":
		err = sendPush(req.UserID, req.Message)
	default:
		err = fmt.Errorf("unsupported channel: %s", req.Channel)
	}
	if err != nil {
		n.Status = "failed"
	} else {
		n.Status = "sent"
	}
	repo.Insert(n)
	return n, err
}

// ListUserNotifications returns recent notifications for the given user.
func ListUserNotifications(userID string, limit int) ([]model.Notification, error) {
	return repo.ListByUser(userID, limit)
}

// sendEmail delivers an email using SMTP settings from environment.
func sendEmail(to, msg string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	if host == "" || user == "" {
		return errors.New("smtp not configured")
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)
	body := []byte("To: " + to + "\r\nSubject: Operary Alert\r\n\r\n" + msg)
	return smtp.SendMail(addr, auth, user, []string{to}, body)
}

// sendPush sends a push notification via Firebase Cloud Messaging.
// user token is expected in `to` parameter.
func sendPush(to, msg string) error {
	key := os.Getenv("FCM_SERVER_KEY")
	if key == "" {
		return errors.New("fcm not configured")
	}
	payload := map[string]interface{}{
		"to": to,
		"notification": map[string]string{
			"title": "Operary",
			"body":  msg,
		},
	}
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(b))
	req.Header.Set("Authorization", "key="+key)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("fcm error: %s", resp.Status)
	}
	return nil
}

// TriggerFromEvent allows other modules to send notifications automatically.
func TriggerFromEvent(eventType, userID, message string) {
	req := NotificationRequest{UserID: userID, Channel: "push", Message: message}
	Send(req)
}
