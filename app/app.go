package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore("", []byte("ssomething-very-secret"))
	Store.Options = &sessions.Options{
		MaxAge: 86400,
	}
	gob.Register(map[string]interface{}{})
	return nil
}
