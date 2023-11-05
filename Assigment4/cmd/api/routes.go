package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/sequrity", app.listGamesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sequrity", app.createGameHandler)
	router.HandlerFunc(http.MethodGet, "/v1/sequrity/:id", app.showGameHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/sequrity/:id", app.updateGameHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/sequrity/:id", app.deleteGameHandler)
	return app.recoverPanic(app.rateLimit(router))
