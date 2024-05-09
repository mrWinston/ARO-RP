package common

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/go-autorest/autorest"
)

var ClientOptions = arm.ClientOptions{
	ClientOptions: azcore.ClientOptions{
		Retry: policy.RetryOptions{
			TryTimeout:  10 * time.Minute,
			ShouldRetry: shouldRetry,
		},
	},
}

// shouldRetry checks if the response is retriable.
func shouldRetry(resp *http.Response, err error) bool {
	if err != nil {
		return false
	}
	// Don't retry if successful
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return false
	}
	// Retry if the status code is retriable
	for _, sc := range autorest.StatusCodesForRetry {
		if resp.StatusCode == sc {
			return true
		}
	}

	// Check if the body contains the certain strings that can be retried.
	var b []byte
	_, err = resp.Body.Read(b)
	if err != nil {
		return true
	}
	body := string(b)
	return strings.Contains(body, "AADSTS7000215") ||
		strings.Contains(body, "AADSTS7000216") ||
		strings.Contains(body, "AuthorizationFailed")
}
