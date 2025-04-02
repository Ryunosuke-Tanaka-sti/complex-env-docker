package models

import (
	"time"
)

// Userモデルの定義
type User struct {
	ID        uint      `gorm:"primaryKey"`                        // 主キー
	Name      string    `gorm:"type:varchar(100);not null"`        // 名前（必須）
	Email     string    `gorm:"type:varchar(100);unique;not null"` // メールアドレス（ユニーク制約付き）
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`         // 作成日時
}
