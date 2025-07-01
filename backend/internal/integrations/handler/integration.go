package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"

	metrics "github.com/elkarto91/operary/internal/handlers"
	"github.com/elkarto91/operary/internal/integrations/model"
)

// SendWebhook handles POST /integrations/webhook/send and dispatches an HTTP webhook.
func SendWebhook(w http.ResponseWriter, r *http.Request) {
	var req model.WebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	method := req.Method
	if method == "" {
		method = http.MethodPost
	}

	body, _ := json.Marshal(req.Payload)
	httpReq, err := http.NewRequest(method, req.URL, bytes.NewBuffer(body))
	if err != nil {
		metrics.WebhookFailures.Inc()
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}
	if httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		metrics.WebhookFailures.Inc()
		http.Error(w, "Webhook call failed", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		metrics.WebhookSuccess.Inc()
	} else {
		metrics.WebhookFailures.Inc()
	}

	w.WriteHeader(resp.StatusCode)
}

// SendEmail handles POST /integrations/email/send and dispatches an email via SMTP.
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var req model.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if len(req.To) == 0 || req.Subject == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	if host == "" || port == "" || user == "" || pass == "" {
		http.Error(w, "SMTP not configured", http.StatusInternalServerError)
		return
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", req.Subject, req.Body))
	if err := smtp.SendMail(addr, auth, user, req.To, msg); err != nil {
		metrics.EmailFailures.Inc()
		http.Error(w, "Failed to send email", http.StatusBadGateway)
		return
	}

	metrics.EmailSuccess.Inc()
	w.WriteHeader(http.StatusOK)
}
