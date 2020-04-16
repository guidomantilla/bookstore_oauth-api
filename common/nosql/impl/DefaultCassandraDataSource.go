package impl

import (
	"github.com/gocql/gocql"
)

type DefaultCassandraDataSource struct {
	session *gocql.Session
}

func NewDefaultCassandraDataSource(username string, password string, keyspace string, url ...string) *DefaultCassandraDataSource {

	cluster := gocql.NewCluster(url...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	return &DefaultCassandraDataSource{
		session: session,
	}
}

func (cassandraDataSource *DefaultCassandraDataSource) GetSession() *gocql.Session {
	return cassandraDataSource.session
}
