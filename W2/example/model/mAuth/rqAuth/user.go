package rqAuth

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/W2/example/conf"
	"github.com/kokizzu/gotro/X"
	"golang.org/x/crypto/bcrypt"
)

func (s *Users) FindOffsetLimit(offset, limit uint32) (res []*Users) {
	query := `
SELECT ` + s.sqlSelectAllFields() + `
FROM ` + s.sqlTableName() + `
ORDER BY ` + s.sqlId() + `
LIMIT ` + X.ToS(limit) + `
OFFSET ` + X.ToS(offset)
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	s.Adapter.QuerySql(query, func(row []interface{}) {
		obj := &Users{}
		obj.FromArray(row)
		obj.CensorFields()
		res = append(res, obj)
	})
	return
}

func (s *Users) CensorFields() {
	s.Password = ``
	s.SecretCode = ``
}

func (s *Users) CheckPassword(currentPassword string) bool {
	hash := []byte(s.Password)
	pass := []byte(currentPassword)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return !L.IsError(err, `bcrypt.CompareHashAndPassword`)
}

// add more custom methods here
