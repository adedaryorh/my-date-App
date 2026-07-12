package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	NotificationTitle  string
	NotificationStatus string
)

const (
	NotificationTitleScholarshipUpdate NotificationTitle  = "Scholarship Update"
	NotificationStatusSent             NotificationStatus = "sent"
	NotificationStatusDelivered        NotificationStatus = "delivered"
	NotificationStatusFailed           NotificationStatus = "failed"
	NotificationStatusRead             NotificationStatus = "read"
)

type Notification struct {
	Id               uuid.UUID            `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	Owner            string               `json:"owner,omitempty"`
	OwnerId          uuid.UUID            `json:"owner_id,omitempty"`
	Status           NotificationStatus   `json:"status,omitempty"`
	Template         string               `json:"template"`
	Title            string               `json:"title"`
	NotificationType string               `json:"notification_type"`
	Content          string               `json:"content"`
	MetaData         NotificationMetaData `json:"meta_data,omitempty"`
	CreatedAt        time.Time            `json:"created_at,omitempty"`
	UpdatedAt        time.Time            `json:"updated_at,omitempty"`
}

type NotificationMetaData struct {
	Action    string    `json:"action,omitempty"`
	ObjectRef string    `json:"object_ref,omitempty"`
	ObjectId  uuid.UUID `json:"object_id,omitempty"`
}

func (n NotificationMetaData) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *NotificationMetaData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to byte failed")
	}
	return json.Unmarshal(b, &n)
}

var NotificationChannels = struct {
	Email    string
	SMS      string
	Push     string
	Whatsapp string
}{
	Email:    "email",
	SMS:      "sms",
	Push:     "push",
	Whatsapp: "whatsapp",
}

var NotificationTemplate = struct {
	AdminInvitation               string
	AdminInvitationAccepted       string
	ConfirmPhone                  string
	ConfirmEmail                  string
	PasswordReset                 string
	ConfirmBeneficiaryEmail       string
	SavingsAccountActivated       string
	TransactionCompleted          string
	ExpenseAlmostDue              string
	ExpensePastDueDate            string
	BudgetLimitExceeded           string
	MoneyRequestCreated           string
	ScholarshipInReview           string
	ScholarshipUnderConsideration string
	ScholarshipReviewerAssigned   string
	ScholarshipRejected           string
	ScholarshipShortlisted        string
	ScholarshipScheduledInterview string
	ScholarshipAwarded            string
	ScholarshipProgramFunded      string
}{
	AdminInvitation:               "admin.invitation",
	AdminInvitationAccepted:       "admin.invitation-accepted",
	ConfirmPhone:                  "confirm-phone",
	ConfirmEmail:                  "confirm-email",
	PasswordReset:                 "password-reset",
	ConfirmBeneficiaryEmail:       "beneficiary.confirm-email",
	SavingsAccountActivated:       "account.activated",
	TransactionCompleted:          "transaction.completed",
	ExpenseAlmostDue:              "expense.almost-due",
	ExpensePastDueDate:            "expense.past-due-date",
	BudgetLimitExceeded:           "budget.limit-exceeded",
	MoneyRequestCreated:           "money-request.created",
	ScholarshipInReview:           "scholarship.in-review",
	ScholarshipUnderConsideration: "scholarship.under-consideration",
	ScholarshipReviewerAssigned:   "scholarship.reviewer-assigned",
	ScholarshipRejected:           "scholarship.rejected",
	ScholarshipShortlisted:        "scholarship.shortlisted",
	ScholarshipScheduledInterview: "scholarship.scheduled-interview",
	ScholarshipAwarded:            "scholarship.awarded",
	ScholarshipProgramFunded:      "scholarship-program.funded",
}

var NotificationTemplates = map[string][]string{
	NotificationTemplate.AdminInvitation:               {NotificationChannels.Email},
	NotificationTemplate.AdminInvitationAccepted:       {NotificationChannels.Email},
	NotificationTemplate.PasswordReset:                 {NotificationChannels.Email},
	NotificationTemplate.ConfirmEmail:                  {NotificationChannels.Email},
	NotificationTemplate.ConfirmPhone:                  {NotificationChannels.SMS},
	NotificationTemplate.ScholarshipProgramFunded:      {NotificationChannels.Email},
	NotificationTemplate.ConfirmPhone:                  {NotificationChannels.SMS},
	NotificationTemplate.ConfirmBeneficiaryEmail:       {NotificationChannels.Email},
	NotificationTemplate.MoneyRequestCreated:           {NotificationChannels.Email},
	NotificationTemplate.BudgetLimitExceeded:           {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.SavingsAccountActivated:       {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ExpenseAlmostDue:              {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.TransactionCompleted:          {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ExpensePastDueDate:            {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipInReview:           {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipUnderConsideration: {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipRejected:           {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipShortlisted:        {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipReviewerAssigned:   {NotificationChannels.Email},
	NotificationTemplate.ScholarshipScheduledInterview: {NotificationChannels.Email, NotificationChannels.Push},
	NotificationTemplate.ScholarshipAwarded:            {NotificationChannels.Email, NotificationChannels.Push},
}

type NotificationOpts struct {
	Message          string    `json:"message"`
	Title            string    `json:"title"`
	NotificationType string    `json:"notification_type"`
	Action           string    `json:"action,omitempty"`
	ObjectRef        string    `json:"object_ref,omitempty"`
	ObjectId         uuid.UUID `json:"object_id,omitempty"`
}

type NotificationJob struct {
	From     string                 `json:"fromEmail,omitempty"`
	To       []string               `json:"toEmail,omitempty"`
	FromName string                 `json:"fromName,omitempty"`
	CcEmails []string               `json:"cc,omitempty"`
	Template string                 `json:"template,omitempty"`
	Content  map[string]interface{} `json:"content,omitempty"`
	Phone    string
}
