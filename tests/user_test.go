package controllers

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yomraliahmet/fiber-api/controllers"
)

func TestTopla(t *testing.T) {
	var toplam = controllers.Topla(2, 2)
	/*
		if toplam != 4 {
			t.Error("Beklenen sonuç 4, elde edilen sonuç ", toplam)
		}
	*/
	assert.Equal(t, 4, toplam)
}

func TestMahmut(t *testing.T) {
	var name string = "Mahmut"
	assert.Equal(t, "Mahmut", name)
}

func TestHomePageIsRunning(t *testing.T) {

	e := httpexpect.New(t, baseUrl())
	e.GET("/").
		Expect().
		Status(http.StatusOK)
}

func baseUrl() string {

	const BASE_URL = "http://127.0.0.1"
	const PORT = "8080"

	url := BASE_URL + ":" + PORT

	return url
}
