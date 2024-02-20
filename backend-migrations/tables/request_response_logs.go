package tables

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type RequestResponseLogs struct {
	ID                 int       `gorm:"column:request_response_logs_id;primaryKey;autoIncrement"`
	PID                string    `gorm:"column:request_response_logs_pid;unique;not null;type:varchar(40)"`
	TraceID            string    `gorm:"column:trace_id;not null;type:varchar(40)"`
	UserPID            string    `gorm:"column:user_pid;not null;type:varchar(40)"`
	Method             string    `gorm:"column:method;type:varchar(10)"`
	ResponseCode       int       `gorm:"column:response_code"`
	RequestBodyLength  int       `gorm:"column:request_body_length"`
	ResponseBodyLength int       `gorm:"column:response_body_length"`
	EndPoint           string    `gorm:"column:end_point"`
	HostURL            string    `gorm:"column:host_url"`
	ClientIP           string    `gorm:"column:client_ip"` //
	RemoteIP           string    `gorm:"column:remote_ip"` //
	Params             string    `gorm:"column:params"`
	QueryParams        string    `gorm:"column:query_params"`
	ContentType        string    `gorm:"column:content_type"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoCreateTime:true"`
	// NOTE: below should be understood carefully
	RequestBody  ReqResJson `gorm:"column:request_body;type:json"`
	ResponseBody ReqResJson `gorm:"column:response_body;type:json"`
}

// ReqResJson interface for Json Fields
type ReqResJson []interface{}

// TODO: didn't got the below functions
func (a *ReqResJson) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ReqResJson) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
