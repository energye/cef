//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	cefTypes "github.com/energye/cef/147/types"
)

// ICEFOAuth2Helper Parent: IObject
type ICEFOAuth2Helper interface {
	IObject
	ParseTokenExchangeResponse(response string) bool         // function
	ParseRefreshTokenResponse(response string) bool          // function
	ParseCodeRequestResponse(uRL string) bool                // function
	Initialize()                                             // procedure
	AuthEndpoint() string                                    // property AuthEndpoint Getter
	SetAuthEndpoint(value string)                            // property AuthEndpoint Setter
	TokenEndpoint() string                                   // property TokenEndpoint Getter
	SetTokenEndpoint(value string)                           // property TokenEndpoint Setter
	ClientID() string                                        // property ClientID Getter
	SetClientID(value string)                                // property ClientID Setter
	ClientSecret() string                                    // property ClientSecret Getter
	SetClientSecret(value string)                            // property ClientSecret Setter
	RedirectHost() string                                    // property RedirectHost Getter
	SetRedirectHost(value string)                            // property RedirectHost Setter
	RedirectPort() int32                                     // property RedirectPort Getter
	SetRedirectPort(value int32)                             // property RedirectPort Setter
	ChallengeMethod() cefTypes.TOAuthChallengeMethod         // property ChallengeMethod Getter
	SetChallengeMethod(value cefTypes.TOAuthChallengeMethod) // property ChallengeMethod Setter
	Scope() string                                           // property Scope Getter
	SetScope(value string)                                   // property Scope Setter
	Error() string                                           // property Error Getter
	ErrorDescription() string                                // property ErrorDescription Getter
	AccessToken() string                                     // property AccessToken Getter
	IDToken() string                                         // property IDToken Getter
	RefreshToken() string                                    // property RefreshToken Getter
	TokenExpiry() int32                                      // property TokenExpiry Getter
	TokenType() string                                       // property TokenType Getter
	CodeVerifier() string                                    // property CodeVerifier Getter
	CodeChallenge() string                                   // property CodeChallenge Getter
	RedirectURI() string                                     // property RedirectURI Getter
	AuthCodeURI() string                                     // property AuthCodeURI Getter
	TokeExchangeParams() string                              // property TokeExchangeParams Getter
	RefreshParams() string                                   // property RefreshParams Getter
	ValidState() bool                                        // property ValidState Getter
}

type TCEFOAuth2Helper struct {
	TObject
}

func (m *TCEFOAuth2Helper) ParseTokenExchangeResponse(response string) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOAuth2HelperAPI().SysCallN(1, m.Instance(), api.PasStr(response))
	return api.GoBool(r)
}

func (m *TCEFOAuth2Helper) ParseRefreshTokenResponse(response string) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOAuth2HelperAPI().SysCallN(2, m.Instance(), api.PasStr(response))
	return api.GoBool(r)
}

func (m *TCEFOAuth2Helper) ParseCodeRequestResponse(uRL string) bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOAuth2HelperAPI().SysCallN(3, m.Instance(), api.PasStr(uRL))
	return api.GoBool(r)
}

func (m *TCEFOAuth2Helper) Initialize() {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(4, m.Instance())
}

func (m *TCEFOAuth2Helper) AuthEndpoint() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetAuthEndpoint(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) TokenEndpoint() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetTokenEndpoint(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) ClientID() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetClientID(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) ClientSecret() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetClientSecret(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) RedirectHost() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetRedirectHost(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) RedirectPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFOAuth2HelperAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCEFOAuth2Helper) SetRedirectPort(value int32) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCEFOAuth2Helper) ChallengeMethod() cefTypes.TOAuthChallengeMethod {
	if !m.IsValid() {
		return 0
	}
	r := cEFOAuth2HelperAPI().SysCallN(11, 0, m.Instance())
	return cefTypes.TOAuthChallengeMethod(r)
}

func (m *TCEFOAuth2Helper) SetChallengeMethod(value cefTypes.TOAuthChallengeMethod) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCEFOAuth2Helper) Scope() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) SetScope(value string) {
	if !m.IsValid() {
		return
	}
	cEFOAuth2HelperAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCEFOAuth2Helper) Error() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) ErrorDescription() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) AccessToken() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) IDToken() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) RefreshToken() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) TokenExpiry() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cEFOAuth2HelperAPI().SysCallN(18, m.Instance())
	return int32(r)
}

func (m *TCEFOAuth2Helper) TokenType() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) CodeVerifier() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) CodeChallenge() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) RedirectURI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) AuthCodeURI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) TokeExchangeParams() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) RefreshParams() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	cEFOAuth2HelperAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCEFOAuth2Helper) ValidState() bool {
	if !m.IsValid() {
		return false
	}
	r := cEFOAuth2HelperAPI().SysCallN(26, m.Instance())
	return api.GoBool(r)
}

// NewOAuth2Helper class constructor
func NewOAuth2Helper() ICEFOAuth2Helper {
	r := cEFOAuth2HelperAPI().SysCallN(0)
	return AsCEFOAuth2Helper(r)
}

var (
	cEFOAuth2HelperOnce   base.Once
	cEFOAuth2HelperImport *imports.Imports = nil
)

func cEFOAuth2HelperAPI() *imports.Imports {
	cEFOAuth2HelperOnce.Do(func() {
		cEFOAuth2HelperImport = api.NewDefaultImports()
		cEFOAuth2HelperImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCEFOAuth2Helper_Create", 0), // constructor NewOAuth2Helper
			/* 1 */ imports.NewTable("TCEFOAuth2Helper_ParseTokenExchangeResponse", 0), // function ParseTokenExchangeResponse
			/* 2 */ imports.NewTable("TCEFOAuth2Helper_ParseRefreshTokenResponse", 0), // function ParseRefreshTokenResponse
			/* 3 */ imports.NewTable("TCEFOAuth2Helper_ParseCodeRequestResponse", 0), // function ParseCodeRequestResponse
			/* 4 */ imports.NewTable("TCEFOAuth2Helper_Initialize", 0), // procedure Initialize
			/* 5 */ imports.NewTable("TCEFOAuth2Helper_AuthEndpoint", 0), // property AuthEndpoint
			/* 6 */ imports.NewTable("TCEFOAuth2Helper_TokenEndpoint", 0), // property TokenEndpoint
			/* 7 */ imports.NewTable("TCEFOAuth2Helper_ClientID", 0), // property ClientID
			/* 8 */ imports.NewTable("TCEFOAuth2Helper_ClientSecret", 0), // property ClientSecret
			/* 9 */ imports.NewTable("TCEFOAuth2Helper_RedirectHost", 0), // property RedirectHost
			/* 10 */ imports.NewTable("TCEFOAuth2Helper_RedirectPort", 0), // property RedirectPort
			/* 11 */ imports.NewTable("TCEFOAuth2Helper_ChallengeMethod", 0), // property ChallengeMethod
			/* 12 */ imports.NewTable("TCEFOAuth2Helper_Scope", 0), // property Scope
			/* 13 */ imports.NewTable("TCEFOAuth2Helper_Error", 0), // property Error
			/* 14 */ imports.NewTable("TCEFOAuth2Helper_ErrorDescription", 0), // property ErrorDescription
			/* 15 */ imports.NewTable("TCEFOAuth2Helper_AccessToken", 0), // property AccessToken
			/* 16 */ imports.NewTable("TCEFOAuth2Helper_IDToken", 0), // property IDToken
			/* 17 */ imports.NewTable("TCEFOAuth2Helper_RefreshToken", 0), // property RefreshToken
			/* 18 */ imports.NewTable("TCEFOAuth2Helper_TokenExpiry", 0), // property TokenExpiry
			/* 19 */ imports.NewTable("TCEFOAuth2Helper_TokenType", 0), // property TokenType
			/* 20 */ imports.NewTable("TCEFOAuth2Helper_CodeVerifier", 0), // property CodeVerifier
			/* 21 */ imports.NewTable("TCEFOAuth2Helper_CodeChallenge", 0), // property CodeChallenge
			/* 22 */ imports.NewTable("TCEFOAuth2Helper_RedirectURI", 0), // property RedirectURI
			/* 23 */ imports.NewTable("TCEFOAuth2Helper_AuthCodeURI", 0), // property AuthCodeURI
			/* 24 */ imports.NewTable("TCEFOAuth2Helper_TokeExchangeParams", 0), // property TokeExchangeParams
			/* 25 */ imports.NewTable("TCEFOAuth2Helper_RefreshParams", 0), // property RefreshParams
			/* 26 */ imports.NewTable("TCEFOAuth2Helper_ValidState", 0), // property ValidState
		}
	})
	return cEFOAuth2HelperImport
}
