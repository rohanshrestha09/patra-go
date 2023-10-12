package enums

// import "database/sql/driver"

type Provider string

const (
	FACEBOOK Provider = "FACEBOOK"
	GOOGLE   Provider = "GOOGLE"
	EMAIL    Provider = "EMAIL"
)

// func (self *Provider) Scan(value interface{}) error {
// 	*self = Provider(value.([]byte))
// 	return nil
// }

// func (self Provider) Value() (driver.Value, error) {
// 	return string(self), nil
// }

type Association string

const (
	FOLLOWING Association = "Following"
)

func (a Association) Value() string {
	return string(a)
}

type SearchColumn string

const (
	NAME_COLUMN  SearchColumn = "name"
	EMAIL_COLUMN SearchColumn = "email"
)

func (s SearchColumn) Value() string {
	return string(s)
}

type FileType string

const (
	IMAGE FileType = "image"
)

type FileDIR string

const (
	USER_DIR    FileDIR = "users/"
	PROJECT_DIR FileDIR = "projects/"
)
