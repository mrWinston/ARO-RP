package oidcbuilder

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	mock_azblob "github.com/Azure/ARO-RP/pkg/util/mocks/azblob"
	utilerror "github.com/Azure/ARO-RP/test/util/error"
)

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

func TestEnsureOIDCDocs(t *testing.T) {
	ctx := context.Background()
	fakeContainerName := "fakeContainer"
	blobContainerURL := "fakeBlobContainerURL"
	endpointURL := "fakeEndPointURL"

	priKey, pubKey, err := CreateKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatal(err)
	}
	pubKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	incorrectlyEncodedPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyBytes,
	})

	invalidKey := []byte("Invalid Key")

	for _, tt := range []struct {
		name        string
		mocks       func(*mock_azblob.MockAZBlobClient)
		oidcbuilder *OIDCBuilder
		wantErr     string
	}{
		{
			name: "Success",
			oidcbuilder: &OIDCBuilder{
				privateKey:       priKey,
				publicKey:        pubKey,
				blobContainerURL: blobContainerURL,
				endpointURL:      endpointURL,
			},
			mocks: func(azblobClient *mock_azblob.MockAZBlobClient) {
				azblobClient.EXPECT().
					UploadBuffer(gomock.Any(), "", BodyKey, gomock.Any()).
					Return(nil)
				azblobClient.EXPECT().
					UploadBuffer(gomock.Any(), "", JWKSKey, gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "Fail -Invalid Public Key fails during decoding",
			oidcbuilder: &OIDCBuilder{
				privateKey:       priKey,
				publicKey:        invalidKey,
				blobContainerURL: blobContainerURL,
				endpointURL:      endpointURL,
			},
			wantErr: "Failed to decode PEM file",
		},
		{
			name: "Fail - Valid Public Key(PEM) but not expected type",
			oidcbuilder: &OIDCBuilder{
				privateKey:       priKey,
				publicKey:        incorrectlyEncodedPublicKey,
				blobContainerURL: blobContainerURL,
				endpointURL:      endpointURL,
			},
			wantErr: "Failed to parse key content: x509: failed to parse public key (use ParsePKCS1PublicKey instead for this key format)",
		},
		{
			name: "Fail - Error when uploading OIDC main configuration",
			oidcbuilder: &OIDCBuilder{
				privateKey:       priKey,
				publicKey:        pubKey,
				blobContainerURL: blobContainerURL,
				endpointURL:      endpointURL,
			},
			mocks: func(azblobClient *mock_azblob.MockAZBlobClient) {
				azblobClient.EXPECT().
					UploadBuffer(gomock.Any(), "", BodyKey, gomock.Any()).
					Return(errors.New("generic error"))
			},
			wantErr: "generic error",
		},
		{
			name: "Fail - Error when uploading JWKS",
			oidcbuilder: &OIDCBuilder{
				privateKey:       priKey,
				publicKey:        pubKey,
				blobContainerURL: blobContainerURL,
				endpointURL:      endpointURL,
			},
			mocks: func(azblobClient *mock_azblob.MockAZBlobClient) {
				azblobClient.EXPECT().
					UploadBuffer(gomock.Any(), "", BodyKey, gomock.Any()).
					Return(nil)
				azblobClient.EXPECT().
					UploadBuffer(gomock.Any(), "", JWKSKey, gomock.Any()).
					Return(errors.New("generic error"))
			},
			wantErr: "generic error",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			azBlobClient := mock_azblob.NewMockAZBlobClient(controller)

			if tt.mocks != nil {
				tt.mocks(azBlobClient)
			}

			err = tt.oidcbuilder.EnsureOIDCDocs(ctx, fakeContainerName, azBlobClient)
			utilerror.AssertErrorMessage(t, err, tt.wantErr)

			if tt.oidcbuilder.GetEndpointUrl() != tt.oidcbuilder.endpointURL {
				t.Fatalf("GetEndpointUrl doesn't match the original endpointURL - %s != %s (wanted)", tt.oidcbuilder.GetEndpointUrl(), tt.oidcbuilder.endpointURL)
			}

			if tt.oidcbuilder.GetPrivateKey() != string(tt.oidcbuilder.privateKey) {
				t.Fatalf("GetPrivateKey doesn't match the original endpointURL - %s != %s (wanted)", tt.oidcbuilder.GetPrivateKey(), string(tt.oidcbuilder.privateKey))
			}

			if tt.oidcbuilder.GetBlobContainerURL() != tt.oidcbuilder.blobContainerURL {
				t.Fatalf("GetBlobContainerURL doesn't match the original endpointURL - %s != %s (wanted)", tt.oidcbuilder.GetBlobContainerURL(), tt.oidcbuilder.blobContainerURL)
			}
		})
	}
}
