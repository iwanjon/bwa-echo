package transaction

import (
	"bwastartupecho/campaign"
	"bwastartupecho/user"
	"time"
	// "github.com/leekchan/accounting"
)

type Transaction struct {
	ID         string
	CampaignID string
	UserID     string
	Amount     int
	Status     string
	Code       string
	// PaymentURL string `gorm:"-"`
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// func (t Transaction) AmountFormatIDR() string {
// 	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
// 	return ac.FormatMoney(t.Amount)
// }
