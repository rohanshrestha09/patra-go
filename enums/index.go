package enums

import "database/sql/driver"

type Provider string

const (
	Facebook Provider = "FACEBOOK"
	Google   Provider = "GOOGLE"
	EMAIL    Provider = "EMAIL"
)

func (self *Provider) Scan(value interface{}) error {
	*self = Provider(value.([]byte))
	return nil
}

func (self Provider) Value() (driver.Value, error) {
	return string(self), nil
}

type Model string

const (
	USER    Model = "user"
	PROJECT Model = "project"
	REQUEST Model = "request"
)

type ProjectStatus string

const (
	OPEN        ProjectStatus = "OPEN"
	IN_PROGRESS ProjectStatus = "IN_PROGRESS"
	COMPLETED   ProjectStatus = "COMPLETED"
)

type RequestStatus string

const (
	PENDING  RequestStatus = "PENDING"
	APPROVED RequestStatus = "APPROVED"
	REJECTED RequestStatus = "REJECTED"
)

type FileType string

const (
	IMAGE FileType = "image"
)

type FileDIR string

const (
	USER_DIR    FileDIR = "users/"
	PROJECT_DIR FileDIR = "projects/"
)
