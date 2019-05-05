package cmd

import (
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Unixtime struct {
	time.Time
}

func IntToUnixtime(timestamp int) Unixtime {
	return Unixtime{time.Unix(int64(timestamp), 0)}
}

func (t *Unixtime) MarshalJSON() ([]byte, error) {
	timestamp := fmt.Sprint(t.Unix())
	return []byte(timestamp), nil
}

func (t *Unixtime) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(timestamp), 0)

	return nil
}

func (t Unixtime) String() string {
	return t.Format(time.RFC3339)
}

type User struct {
	ID        string
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	ProjectID uuid.UUID  `gorm:"primary_key;type:char(36);"`
	RoleID    uuid.UUID  `gorm:"type:char(36);"`
	Name      string     `json:"user_name"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
}

type UserShowRequest struct {
	ID string
}

type UserShowResponse struct {
	User User `json:"user"`
}
