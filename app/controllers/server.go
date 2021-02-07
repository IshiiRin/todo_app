package controllers

import (
	"net/http"
	"todo_app/config"
)

//StartMainServer サーバーを立ち上げる
func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
