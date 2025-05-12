package output

import (
	"backend/internal/core/dto"
)

type MailerRepository interface {
	SendMail(mailer dto.MailDTO) error
	SendInvitation(email string, tenant string, acceptURL string, inviter string, role string, receiverProfile string, tenantProfile string) error
}
