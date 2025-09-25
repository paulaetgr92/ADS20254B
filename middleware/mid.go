package middleware

import (
	"time"

	"awesomeProject/Internal_temp/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetPayloadToken(c echo.Context) model.PayloadDTO {
	strID, _ := c.Get("token_id").(uuid.UUID)
	strUserID, _ := c.Get("token_user_id").(int64)
	strUserNickname, _ := c.Get("token_user_nickname").(string)
	strExpiryAt, _ := c.Get("token_expiry_at").(time.Time)
	strAccessKey, _ := c.Get("token_access_key").(int64)
	strAccessID, _ := c.Get("token_access_ID").(int64)
	strTenantID, _ := c.Get("token_tenant_id").(string)
	strUserOrgId, _ := c.Get("token_user_org_id").(int64)
	strUserEmail, _ := c.Get("token_user_email").(string)
	strOrganizationID, _ := c.Get("token_organization_id").(int64)
	strDocument, _ := c.Get("token_document").(string)

	tenantID, _ := uuid.Parse(strTenantID)

	return model.PayloadDTO{
		ID:             strID,
		UserID:         int64(strUserID),
		UserNickname:   strUserNickname,
		ExpiryAt:       strExpiryAt,
		AccessKey:      strAccessKey,
		AccessID:       strAccessID,
		TenantID:       tenantID,
		UserOrgId:      strUserOrgId,
		UserEmail:      strUserEmail,
		OrganizationID: strOrganizationID,
		Document:       strDocument,
	}
}
