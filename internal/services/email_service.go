package services

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

type EmailService struct {
	config EmailConfig
}

func NewEmailService(config EmailConfig) *EmailService {
	return &EmailService{config: config}
}

func (e *EmailService) SendVerificationEmail(toEmail, verificationLink string) error {
	subject := "Verify your password"
	body := fmt.Sprintf(`
		<h2>Email Verification</h2>
		<p>Click the link below to verify your email:</p>
		<a href="%s">Verify Email</a>
		<p>This link expires in 24 hours.</p>
	`, verificationLink)

	return e.sendMail(toEmail, subject, body)
}

func (e *EmailService) SendPasswordResetEmail(toEmail, resetLink string) error {
	subject := "Reset your password"
	body := fmt.Sprintf(`
		<h2>Password Reset</h2>
		<p>Click the link below to reset your password:</p>
		<a href="%s">Reset Password</a>
		<p>This link expires in 1 hour.</p>
	`, resetLink)

	return e.sendMail(toEmail, subject, body)
}

func (e *EmailService) sendMail(to, subject, body string) error {
	message := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		e.config.From,
		to,
		subject,
		body,
	)

	auth := smtp.PlainAuth("", e.config.User, e.config.Password, e.config.Host)
	addr := e.config.Host + ":" + e.config.Port

	return smtp.SendMail(addr, auth, e.config.From, []string{to}, []byte(message))
}
