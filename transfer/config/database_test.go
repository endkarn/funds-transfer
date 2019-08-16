package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetURI_Should_Be_Username_karnwat_Password_1234_Host_127_0_0_1_Port_3306_Database_funds_db(t *testing.T) {
	expectedResult := "karnwat:1234@tcp(127.0.0.1:3306)/funds_db"
	config := Config{
		Username: "karnwat",
		Password: "1234",
		Host:     "127.0.0.1",
		Database: "funds_db",
		Port:     "3306",
	}

	actualResult := config.GetURI()

	assert.Equal(t,expectedResult,actualResult)
}