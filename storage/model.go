package storage

// UserProfile describes model for postgres
type UserProfile struct {
	Id   uint32 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
