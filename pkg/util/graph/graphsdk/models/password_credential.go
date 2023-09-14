package models

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"

	i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PasswordCredential
type PasswordCredential struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewPasswordCredential instantiates a new passwordCredential and sets the default values.
func NewPasswordCredential() *PasswordCredential {
	m := &PasswordCredential{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	return m
}

// CreatePasswordCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPasswordCredential(), nil
}

// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PasswordCredential) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetCustomKeyIdentifier gets the customKeyIdentifier property value. Do not use.
func (m *PasswordCredential) GetCustomKeyIdentifier() []byte {
	val, err := m.GetBackingStore().Get("customKeyIdentifier")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]byte)
	}
	return nil
}

// GetDisplayName gets the displayName property value. Friendly name for the password. Optional.
func (m *PasswordCredential) GetDisplayName() *string {
	val, err := m.GetBackingStore().Get("displayName")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetEndDateTime gets the endDateTime property value. The date and time at which the password expires represented using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *PasswordCredential) GetEndDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	val, err := m.GetBackingStore().Get("endDateTime")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	}
	return nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordCredential) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["customKeyIdentifier"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetByteArrayValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCustomKeyIdentifier(val)
		}
		return nil
	}
	res["displayName"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisplayName(val)
		}
		return nil
	}
	res["endDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEndDateTime(val)
		}
		return nil
	}
	res["hint"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHint(val)
		}
		return nil
	}
	res["keyId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetUUIDValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyId(val)
		}
		return nil
	}
	res["@odata.type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOdataType(val)
		}
		return nil
	}
	res["secretText"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSecretText(val)
		}
		return nil
	}
	res["startDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartDateTime(val)
		}
		return nil
	}
	return res
}

// GetHint gets the hint property value. Contains the first three characters of the password. Read-only.
func (m *PasswordCredential) GetHint() *string {
	val, err := m.GetBackingStore().Get("hint")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetKeyId gets the keyId property value. The unique identifier for the password.
func (m *PasswordCredential) GetKeyId() *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID {
	val, err := m.GetBackingStore().Get("keyId")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
	}
	return nil
}

// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordCredential) GetOdataType() *string {
	val, err := m.GetBackingStore().Get("odataType")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetSecretText gets the secretText property value. Read-only; Contains the strong passwords generated by Azure AD that are 16-64 characters in length. The generated password value is only returned during the initial POST request to addPassword. There is no way to retrieve this password in the future.
func (m *PasswordCredential) GetSecretText() *string {
	val, err := m.GetBackingStore().Get("secretText")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetStartDateTime gets the startDateTime property value. The date and time at which the password becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *PasswordCredential) GetStartDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	val, err := m.GetBackingStore().Get("startDateTime")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	}
	return nil
}

// Serialize serializes information the current object
func (m *PasswordCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteByteArrayValue("customKeyIdentifier", m.GetCustomKeyIdentifier())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("displayName", m.GetDisplayName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("hint", m.GetHint())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteUUIDValue("keyId", m.GetKeyId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("@odata.type", m.GetOdataType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("secretText", m.GetSecretText())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PasswordCredential) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetCustomKeyIdentifier sets the customKeyIdentifier property value. Do not use.
func (m *PasswordCredential) SetCustomKeyIdentifier(value []byte) {
	err := m.GetBackingStore().Set("customKeyIdentifier", value)
	if err != nil {
		panic(err)
	}
}

// SetDisplayName sets the displayName property value. Friendly name for the password. Optional.
func (m *PasswordCredential) SetDisplayName(value *string) {
	err := m.GetBackingStore().Set("displayName", value)
	if err != nil {
		panic(err)
	}
}

// SetEndDateTime sets the endDateTime property value. The date and time at which the password expires represented using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *PasswordCredential) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	err := m.GetBackingStore().Set("endDateTime", value)
	if err != nil {
		panic(err)
	}
}

// SetHint sets the hint property value. Contains the first three characters of the password. Read-only.
func (m *PasswordCredential) SetHint(value *string) {
	err := m.GetBackingStore().Set("hint", value)
	if err != nil {
		panic(err)
	}
}

// SetKeyId sets the keyId property value. The unique identifier for the password.
func (m *PasswordCredential) SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
	err := m.GetBackingStore().Set("keyId", value)
	if err != nil {
		panic(err)
	}
}

// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordCredential) SetOdataType(value *string) {
	err := m.GetBackingStore().Set("odataType", value)
	if err != nil {
		panic(err)
	}
}

// SetSecretText sets the secretText property value. Read-only; Contains the strong passwords generated by Azure AD that are 16-64 characters in length. The generated password value is only returned during the initial POST request to addPassword. There is no way to retrieve this password in the future.
func (m *PasswordCredential) SetSecretText(value *string) {
	err := m.GetBackingStore().Set("secretText", value)
	if err != nil {
		panic(err)
	}
}

// SetStartDateTime sets the startDateTime property value. The date and time at which the password becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *PasswordCredential) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	err := m.GetBackingStore().Set("startDateTime", value)
	if err != nil {
		panic(err)
	}
}

// PasswordCredentialable
type PasswordCredentialable interface {
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetCustomKeyIdentifier() []byte
	GetDisplayName() *string
	GetEndDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetHint() *string
	GetKeyId() *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
	GetOdataType() *string
	GetSecretText() *string
	GetStartDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetCustomKeyIdentifier(value []byte)
	SetDisplayName(value *string)
	SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetHint(value *string)
	SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
	SetOdataType(value *string)
	SetSecretText(value *string)
	SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
}
