package storage

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

func NewMongoConnection(addr, dbname, username, password string) (*mgo.Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(addr, ","),
		Database: dbname,
		Username: username,
		Password: password,
		Timeout:  5 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	return session, nil
}
