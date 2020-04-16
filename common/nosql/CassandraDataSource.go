package nosql

import (
	"github.com/gocql/gocql"
)

type CassandraDataSource interface {
	GetSession() *gocql.Session
}
