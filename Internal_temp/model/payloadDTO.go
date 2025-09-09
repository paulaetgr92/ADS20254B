package model

import (
	"time"

	"github.com/google/uuid"
)

type PayloadDTO struct {
	Status         string    `json:"status"`
	ID             uuid.UUID `json:"id"`
	UserID         int64     `json:"user_id"`
	UserNickname   string    `json:"user_nickname"`
	ExpiryAt       time.Time `json:"expiry_at"`
	AccessKey      int64     `json:"access_key"`
	AccessID       int64     `json:"access_id"`
	TenantID       uuid.UUID `json:"tenant_id"`
	UserOrgId      int64     `json:"user_org_id"`
	UserEmail      string    `json:"user_email"`
	OrganizationID int64     `json:"organization_id"`
	Document       string    `json:"document"`
	ActivationCode string    `json:"activationCode"`
}

type Payload struct {
	Who      string    `json:"who"`
	WhoID    string    `json:"who_id"`
	TenantID uuid.UUID `json:"tenant_id"`
	AccessID int64     `json:"access_id"`
}
