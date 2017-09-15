package main

import (
	"testing"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCheckCommandReturnDifferentResponse(t *testing.T) {
	request := models.CheckRequest{}

	response := CheckCommand(&request)
	assert.Equal(t, 1, len(response))

	time.Sleep(1 * time.Millisecond)

	anotherResponse := CheckCommand(&request)
	assert.Equal(t, 1, len(anotherResponse))

	responseTime := response[0].Timestamp.UnixNano()
	anotherResponseTime := anotherResponse[0].Timestamp.UnixNano()
	assert.NotEqual(t, responseTime, anotherResponseTime)
}

func TestCheckCommandReturnPreviousVersion(t *testing.T) {
	version := models.Version{Timestamp: time.Now()}
	request := models.CheckRequest{Version: version}

	time.Sleep(1 * time.Millisecond)

	response := CheckCommand(&request)
	assert.Equal(t, 2, len(response))

	requestTime := request.Version.Timestamp.UnixNano()
	responseTime := response[0].Timestamp.UnixNano()
	assert.Equal(t, requestTime, responseTime)
}

func TestCheckCommandResponseTimeIsGreaterThanRequestTime(t *testing.T) {
	version := models.Version{Timestamp: time.Now()}
	request := models.CheckRequest{Version: version}

	time.Sleep(1 * time.Millisecond)

	response := CheckCommand(&request)

	requestTime := request.Version.Timestamp.UnixNano()
	responseTime := response[1].Timestamp.UnixNano()
	assert.True(t, responseTime > requestTime)
}
