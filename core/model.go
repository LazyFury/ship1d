package core

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime       `json:"created_at" format:"2006-01-02 15:04:05"`
	UpdatedAt LocalTime       `json:"updated_at" format:"2006-01-02 15:04:05"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type LocalTime time.Time

var _ JSON = (*LocalTime)(nil)
var _ VALUE = (*LocalTime)(nil)

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	if tTime.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
func (l *LocalTime) UnmarshalJSON(b []byte) error {
	tTime, err := time.Parse("2006-01-02 15:04:05", string(b))
	if err != nil {
		return err
	}
	*l = LocalTime(tTime)
	return nil
}

func (l *LocalTime) Scan(v interface{}) error {
	switch v := v.(type) {
	case time.Time:
		*l = LocalTime(v)
		return nil
	default:
		return nil
	}
}

func (l LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(l)
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(l), nil
}
