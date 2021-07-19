package eventlog

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	ID uint `gorm:"primaryKey"`

	Type string `json:"type"`

	Payload string `json:"payload"`

	Version float64 `json:"version"`

	Topic string `json:"topic"`

	Cid string `json:"cid"`

	Emitter string `json:"emitter"`

	NewAccounts pq.StringArray `gorm:"type:text[]"`

	Timestamp uint64 `json:"timestamp"`

	BlockNumber uint64 `json:"blockNumber"`

	Depth uint64 `json:"depth"`

	TxHash string `json:"txHash"`

	CreatedAt time.Time

	UpdatedAt time.Time
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Event{})
}
