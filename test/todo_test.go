package main

import (
	"bytes"
	"fmt"
	"github.com/daniramdani/todo/test/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIndexPatientRelative testing for IndexPatientRelative
func TestGetTodoDetail(t *testing.T) {
	body := bytes.NewBufferString(``)
	resp, err := integration.DoRequest("GET", "/v1/todo", body)
	assert.NoError(t, err)
	fmt.Println(resp.Body.String())
	assert.Equal(t, 200, resp.Code)
}

/*// Test testing for VerifyPhone
func TestVerifyPhone(t *testing.T) {
	body := bytes.NewBufferString(`{"phone": "+966123123123123123123","phone_verification_token": "3333"}`)
	resp, err := integration.DoRequest("PUT", "/v1/account/verify_phone", body, token)
	assert.NoError(t, err)
	fmt.Println(resp.Body.String())
	assert.Equal(t, 200, resp.Code)
	body = bytes.NewBufferString(`{"phone": "+966123123123123123123","phone_verification_token": "1234"}`)
	resp, err = integration.DoRequest("PUT", "/v1/account/verify_phone", body, token)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.Code)
}

// TestIndexPatientRelative testing for SendVerificationToken
func TestSendVerificationToken(t *testing.T) {
	body := bytes.NewBufferString(`{"phone": "6123123123123123123","country_code": "+96"}`)
	resp, err := integration.DoRequest("PUT", "/v1/account/send_phone_verification_token", body, token)
	assert.NoError(t, err)
	fmt.Println(resp.Body.String())
	assert.Equal(t, 200, resp.Code)
	body = bytes.NewBufferString(`{"phone": "6123123123123123123"`)
	resp, err = integration.DoRequest("PUT", "/v1/account/send_phone_verification_token", body, token)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.Code)
}
*/