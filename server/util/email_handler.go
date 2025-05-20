package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-playground/pkg/logging"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// EmailRequest represents the email request data structure
type EmailRequest struct {
	To          string            `json:"to"`
	Subject     string            `json:"subject"`
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	SenderName  string            `json:"sender_name"`
	Attachments map[string]string `json:"attachments,omitempty"`
}

// EmailHandler manages email sending operations
type EmailHandler struct {
	n8nWebhookURL string
	template      *template.Template
	client        *http.Client
	logger        zerolog.Logger
	mu            sync.RWMutex
}

// NewEmailHandler creates a new instance of EmailHandler
func NewEmailHandler(n8nWebhookURL string) (*EmailHandler, error) {
	logger := logging.GetLogger()

	// Parse email template
	tmplPath := filepath.Join("server", "templates", "email_template.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		logger.Error().
			Err(err).
			Str("template_path", tmplPath).
			Msg("Failed to parse email template")
		return nil, fmt.Errorf("failed to parse email template: %w", err)
	}

	return &EmailHandler{
		n8nWebhookURL: n8nWebhookURL,
		template:      tmpl,
		client:        &http.Client{},
		logger:       logger,
	}, nil
}

// SendEmail sends an email using the n8n webhook
func (h *EmailHandler) SendEmail(req *EmailRequest) error {
	// Validate request
	if err := h.validateRequest(req); err != nil {
		return err
	}

	h.logger.Info().
		Str("to", req.To).
		Str("subject", req.Subject).
		Msg("Sending email")

	// Render HTML template
	var htmlContent bytes.Buffer
	err := h.template.Execute(&htmlContent, map[string]interface{}{
		"Title":      req.Title,
		"Content":    template.HTML(req.Content), // Allow HTML in content
		"SenderName": req.SenderName,
	})
	if err != nil {
		h.logger.Error().
			Err(err).
			Msg("Failed to render email template")
		return fmt.Errorf("failed to render email template: %w", err)
	}

	// Prepare webhook payload
	webhookData := map[string]interface{}{
		"to":      req.To,
		"subject": req.Subject,
		"body":    htmlContent.String(),
	}

	// Add attachments if any
	if len(req.Attachments) > 0 {
		webhookData["attachments"] = req.Attachments
	}

	// Convert to JSON
	jsonData, err := json.Marshal(webhookData)
	if err != nil {
		h.logger.Error().
			Err(err).
			Msg("Failed to marshal webhook data")
		return fmt.Errorf("failed to marshal webhook data: %w", err)
	}

	// Send request to n8n webhook
	resp, err := h.client.Post(h.n8nWebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		h.logger.Error().
			Err(err).
			Msg("Failed to send email via n8n webhook")
		return fmt.Errorf("failed to send email via n8n webhook: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		h.logger.Error().
			Int("status_code", resp.StatusCode).
			Msg("n8n webhook returned non-200 status")
		return fmt.Errorf("n8n webhook returned status %d", resp.StatusCode)
	}

	h.logger.Info().
		Str("to", req.To).
		Msg("Email sent successfully")

	return nil
}

// validateRequest validates the email request
func (h *EmailHandler) validateRequest(req *EmailRequest) error {
	if req.To == "" {
		return fmt.Errorf("recipient email is required")
	}
	if !isValidEmail(req.To) {
		return fmt.Errorf("invalid recipient email format")
	}
	if req.Subject == "" {
		return fmt.Errorf("subject is required")
	}
	if req.Content == "" {
		return fmt.Errorf("content is required")
	}
	if req.Title == "" {
		req.Title = req.Subject // Use subject as title if not provided
	}
	if req.SenderName == "" {
		req.SenderName = "System" // Use default sender name
	}
	return nil
}

// GetTemplate returns a copy of the email template
func (h *EmailHandler) GetTemplate() *template.Template {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.template
}

// UpdateTemplate updates the email template
func (h *EmailHandler) UpdateTemplate(tmpl *template.Template) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.template = tmpl
}