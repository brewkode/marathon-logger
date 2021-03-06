package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanAppName(t *testing.T) {
	var testCaseOne, testCaseTwo, testCaseThree TaskInfo
	testCaseOne.App = "/test.aayush.http"
	assert.Equal(t, testCaseOne.CleanAppName(), "test.aayush.http", "They should be equal") // Dummy string
	testCaseTwo.App = "/test/aayush/http"
	assert.Equal(t, testCaseTwo.CleanAppName(), "test.aayush.http", "They should be equal")
	testCaseThree.App = "/test.aayush.http/"
	assert.Equal(t, testCaseThree.CleanAppName(), "test.aayush.http.", "They should be equal")
}

func TestRenderRsyslogTemplate(t *testing.T) {

	hostname, err := os.Hostname()
	var rsyslog Rsyslog
	label := map[string]string{
		"logs.enabled": "true",
	}
	taskInfo := TaskInfo{
		App:      "/test.aayush.http",
		Hostname: hostname,
		Labels:   label,
		TaskID:   "abcdefghij",
		CWD:      "/foo/bar",
		FileName: "test_file_name.txt",
	}

	expected := `
######################################
# Created via marathon-logger,
# PLEASE DON'T EDIT THIS FILE MANUALLY
# Name - /test.aayush.http
# File - test_file_name.txt
######################################

module(load="imfile")

input(type="imfile"
			File="/foo/bar/test_file_name.txt"
			Tag="test.aayush.http	abcdefghij"
			statefile="abcdefghij"
      Severity="info")
`
	template, err := rsyslog.render(taskInfo)
	assert.NoError(t, err)
	assert.Equal(t, expected, template)
}
