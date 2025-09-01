package middleware

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"reflect"
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

func WhoMapper[T any](data T, payload model.PayloadDTO) (T, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return data, fmt.Errorf("input data must be a struct")
	}

	whoField := v.FieldByName("who")
	if !whoField.IsValid() {
		return data, fmt.Errorf("struct doesn't have a field named 'who'")
	}

	whoField.Set(reflect.ValueOf(payload.UserNickname))

	return data, nil
}

func DownloadAndConvertToBase64(imageURL string) (string, error) {
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "image/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("resposta inválida ao baixar imagem: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler dados da imagem: %w", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str, nil
}
