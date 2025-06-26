package entities

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
	Nick      string `gorm:"uniqueIndex"`
	Pass      string
	Role      string
}
