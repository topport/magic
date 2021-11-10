// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	crypto "github.com/topport/magic/pkg/crypto"
	identity "github.com/topport/magic/pkg/identity"

	mock "github.com/stretchr/testify/mock"

	types "github.com/topport/magic/pkg/types"
)

// KeyStore is an autogenerated mock type for the KeyStore type
type KeyStore struct {
	mock.Mock
}

// Addresses provides a mock function with given fields:
func (_m *KeyStore) Addresses() (types.AddressSet, error) {
	ret := _m.Called()

	var r0 types.AddressSet
	if rf, ok := ret.Get(0).(func() types.AddressSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AddressSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *KeyStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DefaultPublicIdentity provides a mock function with given fields:
func (_m *KeyStore) DefaultPublicIdentity() (identity.Identity, error) {
	ret := _m.Called()

	var r0 identity.Identity
	if rf, ok := ret.Get(0).(func() identity.Identity); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(identity.Identity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtraUserData provides a mock function with given fields: key
func (_m *KeyStore) ExtraUserData(key string) (interface{}, bool, error) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Identities provides a mock function with given fields:
func (_m *KeyStore) Identities() ([]identity.Identity, error) {
	ret := _m.Called()

	var r0 []identity.Identity
	if rf, ok := ret.Get(0).(func() []identity.Identity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]identity.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdentityExists provides a mock function with given fields: address
func (_m *KeyStore) IdentityExists(address types.Address) (bool, error) {
	ret := _m.Called(address)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Address) bool); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdentityWithAddress provides a mock function with given fields: address
func (_m *KeyStore) IdentityWithAddress(address types.Address) (identity.Identity, error) {
	ret := _m.Called(address)

	var r0 identity.Identity
	if rf, ok := ret.Get(0).(func(types.Address) identity.Identity); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(identity.Identity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalSymEncKey provides a mock function with given fields:
func (_m *KeyStore) LocalSymEncKey() crypto.SymEncKey {
	ret := _m.Called()

	var r0 crypto.SymEncKey
	if rf, ok := ret.Get(0).(func() crypto.SymEncKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.SymEncKey)
		}
	}

	return r0
}

// Mnemonic provides a mock function with given fields:
func (_m *KeyStore) Mnemonic() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIdentity provides a mock function with given fields: public
func (_m *KeyStore) NewIdentity(public bool) (identity.Identity, error) {
	ret := _m.Called(public)

	var r0 identity.Identity
	if rf, ok := ret.Get(0).(func(bool) identity.Identity); ok {
		r0 = rf(public)
	} else {
		r0 = ret.Get(0).(identity.Identity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenMessageFrom provides a mock function with given fields: usingIdentity, senderPublicKey, msgEncrypted
func (_m *KeyStore) OpenMessageFrom(usingIdentity types.Address, senderPublicKey *crypto.AsymEncPubkey, msgEncrypted []byte) ([]byte, error) {
	ret := _m.Called(usingIdentity, senderPublicKey, msgEncrypted)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.Address, *crypto.AsymEncPubkey, []byte) []byte); ok {
		r0 = rf(usingIdentity, senderPublicKey, msgEncrypted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address, *crypto.AsymEncPubkey, []byte) error); ok {
		r1 = rf(usingIdentity, senderPublicKey, msgEncrypted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicIdentities provides a mock function with given fields:
func (_m *KeyStore) PublicIdentities() ([]identity.Identity, error) {
	ret := _m.Called()

	var r0 []identity.Identity
	if rf, ok := ret.Get(0).(func() []identity.Identity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]identity.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveExtraUserData provides a mock function with given fields: key, value
func (_m *KeyStore) SaveExtraUserData(key string, value interface{}) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SealMessageFor provides a mock function with given fields: usingIdentity, recipientPubKey, msg
func (_m *KeyStore) SealMessageFor(usingIdentity types.Address, recipientPubKey *crypto.AsymEncPubkey, msg []byte) ([]byte, error) {
	ret := _m.Called(usingIdentity, recipientPubKey, msg)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.Address, *crypto.AsymEncPubkey, []byte) []byte); ok {
		r0 = rf(usingIdentity, recipientPubKey, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address, *crypto.AsymEncPubkey, []byte) error); ok {
		r1 = rf(usingIdentity, recipientPubKey, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignHash provides a mock function with given fields: usingIdentity, data
func (_m *KeyStore) SignHash(usingIdentity types.Address, data types.Hash) ([]byte, error) {
	ret := _m.Called(usingIdentity, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.Address, types.Hash) []byte); ok {
		r0 = rf(usingIdentity, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address, types.Hash) error); ok {
		r1 = rf(usingIdentity, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields: password, userMnemonic
func (_m *KeyStore) Unlock(password string, userMnemonic string) error {
	ret := _m.Called(password, userMnemonic)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(password, userMnemonic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifySignature provides a mock function with given fields: usingIdentity, hash, signature
func (_m *KeyStore) VerifySignature(usingIdentity types.Address, hash types.Hash, signature []byte) (bool, error) {
	ret := _m.Called(usingIdentity, hash, signature)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Address, types.Hash, []byte) bool); ok {
		r0 = rf(usingIdentity, hash, signature)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Address, types.Hash, []byte) error); ok {
		r1 = rf(usingIdentity, hash, signature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
