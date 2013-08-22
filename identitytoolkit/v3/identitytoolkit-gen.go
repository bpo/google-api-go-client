// Package identitytoolkit provides access to the Google Identity Toolkit API.
//
// See https://developers.google.com/identity-toolkit/v3/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/identitytoolkit/v3"
//   ...
//   identitytoolkitService, err := identitytoolkit.New(oauthHttpClient)
package identitytoolkit

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

const apiId = "identitytoolkit:v3"
const apiName = "identitytoolkit"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Relyingparty = NewRelyingpartyService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Relyingparty *RelyingpartyService
}

func NewRelyingpartyService(s *Service) *RelyingpartyService {
	rs := &RelyingpartyService{s: s}
	return rs
}

type RelyingpartyService struct {
	s *Service
}

type CreateAuthUriResponse struct {
	// AuthUri: The URI used by the IDP to authenticate the user.
	AuthUri string `json:"authUri,omitempty"`

	// Kind: The fixed string identitytoolkit#CreateAuthUriResponse".
	Kind string `json:"kind,omitempty"`

	// Providers: Existing IDP's for the user.
	Providers []string `json:"providers,omitempty"`

	// Registered: Whether the user is registered if the identifier is an
	// email.
	Registered bool `json:"registered,omitempty"`
}

type DeleteAccountResponse struct {
	// Kind: The fixed string "identitytoolkit#DeleteAccountResponse".
	Kind string `json:"kind,omitempty"`
}

type GetAccountInfoResponse struct {
	// Kind: The fixed string "identitytoolkit#GetAccountInfoResponse".
	Kind string `json:"kind,omitempty"`

	// Users: The info of the users.
	Users []*GetAccountInfoResponseUsers `json:"users,omitempty"`
}

type GetAccountInfoResponseUsers struct {
	// DateOfBirth: The user's date of birth.
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email returned by the IdP. NOTE: The federated login user
	// may not own the email.
	Email string `json:"email,omitempty"`

	// Language: The language of the user.
	Language string `json:"language,omitempty"`

	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`

	// Password: The user's hashed password.
	Password string `json:"password,omitempty"`

	// PasswordUpdatedAt: The timestamp when the password was last updated.
	PasswordUpdatedAt int64 `json:"passwordUpdatedAt,omitempty,string"`

	// PhotoUrl: The URL of the user profile photo.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderUserInfo: The IDP of the user.
	ProviderUserInfo []*GetAccountInfoResponseUsersProviderUserInfo `json:"providerUserInfo,omitempty"`

	// Salt: The user's password salt.
	Salt string `json:"salt,omitempty"`

	// TimeZone: The time zone of the user.
	TimeZone string `json:"timeZone,omitempty"`

	// Version: Version of the user's password.
	Version int64 `json:"version,omitempty"`
}

type GetAccountInfoResponseUsersProviderUserInfo struct {
	// DisplayName: The user's display name at the IDP.
	DisplayName string `json:"displayName,omitempty"`

	// PhotoUrl: The user's photo url at the IDP.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name, e.g., google.com, aol.com, live.net and yahoo.com. For other
	// OpenID IdPs it's the OP identifier.
	ProviderId string `json:"providerId,omitempty"`
}

type GetOobConfirmationCodeResponse struct {
	// Kind: The fixed string
	// "identitytoolkit#GetOobConfirmationCodeResponse".
	Kind string `json:"kind,omitempty"`

	// OobCode: The code to be send to the user.
	OobCode string `json:"oobCode,omitempty"`
}

type IdentitytoolkitRelyingpartyCreateAuthUriRequest struct {
	// ClientId: The relying party OAuth client ID.
	ClientId string `json:"clientId,omitempty"`

	// Context: The opaque value used by the client to maintain context info
	// between the authentication request and the IDP callback.
	Context string `json:"context,omitempty"`

	// ContinueUri: The URI to which the IDP redirects the user after the
	// federated login flow.
	ContinueUri string `json:"continueUri,omitempty"`

	// Identifier: The email or federated ID of the user.
	Identifier string `json:"identifier,omitempty"`

	// OpenidRealm: Optional realm for OpenID protocol. The sub string
	// "scheme://domain:port" of the param "continueUri" is used if this is
	// not set.
	OpenidRealm string `json:"openidRealm,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name e.g. google.com, aol.com, live.net and yahoo.com. For other
	// OpenID IdPs it's the OP identifier.
	ProviderId string `json:"providerId,omitempty"`
}

type IdentitytoolkitRelyingpartyDeleteAccountRequest struct {
	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`
}

type IdentitytoolkitRelyingpartyGetAccountInfoRequest struct {
	// Email: The list of emails of the users to inquiry.
	Email []string `json:"email,omitempty"`

	// IdToken: The GITKit token of the authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// LocalId: The list of local ID's of the users to inquiry.
	LocalId []string `json:"localId,omitempty"`
}

type IdentitytoolkitRelyingpartyResetPasswordRequest struct {
	// Email: The email address of the user.
	Email string `json:"email,omitempty"`

	// NewPassword: The new password inputted by the user.
	NewPassword string `json:"newPassword,omitempty"`

	// OldPassword: The old password inputted by the user.
	OldPassword string `json:"oldPassword,omitempty"`

	// OobCode: The confirmation code.
	OobCode string `json:"oobCode,omitempty"`
}

type IdentitytoolkitRelyingpartySetAccountInfoRequest struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// EmailVerified: Mark the email as verified or not.
	EmailVerified bool `json:"emailVerified,omitempty"`

	// IdToken: The GITKit token of the authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`

	// OobCode: The out-of-band code of the change email request.
	OobCode string `json:"oobCode,omitempty"`

	// Password: The new password of the user.
	Password string `json:"password,omitempty"`

	// Provider: The associated IDPs of the user.
	Provider []string `json:"provider,omitempty"`

	// UpgradeToFederatedLogin: Mark the user to upgrade to federated login.
	UpgradeToFederatedLogin bool `json:"upgradeToFederatedLogin,omitempty"`
}

type IdentitytoolkitRelyingpartyUploadAccountRequest struct {
	HashAlgorithm string `json:"hashAlgorithm,omitempty"`

	MemoryCost int64 `json:"memoryCost,omitempty"`

	Rounds int64 `json:"rounds,omitempty"`

	SaltSeparator string `json:"saltSeparator,omitempty"`

	SignerKey string `json:"signerKey,omitempty"`

	// UserAccount: The account info to be stored.
	UserAccount []*Userinfo `json:"userAccount,omitempty"`
}

type IdentitytoolkitRelyingpartyVerifyAssertionRequest struct {
	// PendingIdToken: The GITKit token for the non-trusted IDP pending to
	// be confirmed by the user.
	PendingIdToken string `json:"pendingIdToken,omitempty"`

	// PostBody: The post body if the request is a HTTP POST.
	PostBody string `json:"postBody,omitempty"`

	// RequestUri: The URI to which the IDP redirects the user back. It may
	// contain federated login result params added by the IDP.
	RequestUri string `json:"requestUri,omitempty"`
}

type IdentitytoolkitRelyingpartyVerifyPasswordRequest struct {
	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// Password: The password inputed by the user.
	Password string `json:"password,omitempty"`

	// PendingIdToken: The GITKit token for the non-trusted IDP, which is to
	// be confirmed by the user.
	PendingIdToken string `json:"pendingIdToken,omitempty"`
}

type Relyingparty struct {
	// CaptchaResp: The recaptcha response from the user.
	CaptchaResp string `json:"captchaResp,omitempty"`

	// Challenge: The recaptcha challenge presented to the user.
	Challenge string `json:"challenge,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// IdToken: The user's Gitkit login token for email change.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#relyingparty".
	Kind string `json:"kind,omitempty"`

	// NewEmail: The new email if the code is for email change.
	NewEmail string `json:"newEmail,omitempty"`

	// RequestType: The request type.
	RequestType string `json:"requestType,omitempty"`

	// UserIp: The IP address of the user.
	UserIp string `json:"userIp,omitempty"`
}

type ResetPasswordResponse struct {
	// Email: The user's email.
	Email string `json:"email,omitempty"`

	// Kind: The fixed string "identitytoolkit#ResetPasswordResponse".
	Kind string `json:"kind,omitempty"`
}

type SetAccountInfoResponse struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// IdToken: The Gitkit id token to login the newly sign up user.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#SetAccountInfoResponse".
	Kind string `json:"kind,omitempty"`

	// Provider: The associated IDPs of the user.
	Provider []string `json:"provider,omitempty"`
}

type UploadAccountResponse struct {
	// Error: The error encountered while processing the account info.
	Error []*UploadAccountResponseError `json:"error,omitempty"`

	// Kind: The fixed string "identitytoolkit#UploadAccountResponse".
	Kind string `json:"kind,omitempty"`
}

type UploadAccountResponseError struct {
	// Index: The index of the malformed account, starting from 0.
	Index int64 `json:"index,omitempty"`

	// Message: Detailed error message for the account info.
	Message string `json:"message,omitempty"`
}

type Userinfo struct {
	// Email: email
	Email string `json:"email,omitempty"`

	// Kind: Identifies this object as a user info.
	Kind string `json:"kind,omitempty"`

	// LocalId: user's id at the site
	LocalId string `json:"localId,omitempty"`

	// Password: password
	Password string `json:"password,omitempty"`

	// Salt: salt
	Salt string `json:"salt,omitempty"`
}

type VerifyAssertionResponse struct {
	// Action: The action code.
	Action string `json:"action,omitempty"`

	// Context: The opaque value used by the client to maintain context info
	// between the authentication request and the IDP callback.
	Context string `json:"context,omitempty"`

	// DateOfBirth: The birth date of the IdP account.
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// DisplayName: The display name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email returned by the IdP. NOTE: The federated login user
	// may not own the email.
	Email string `json:"email,omitempty"`

	// EmailRecycled: It's true if the email is recycled.
	EmailRecycled bool `json:"emailRecycled,omitempty"`

	// EmailVerified: The value is true if the IDP is also the email
	// provider. It means the user owns the email.
	EmailVerified bool `json:"emailVerified,omitempty"`

	// FederatedId: The unique ID identifies the IdP account.
	FederatedId string `json:"federatedId,omitempty"`

	// FirstName: The first name of the user.
	FirstName string `json:"firstName,omitempty"`

	// FullName: The full name of the user.
	FullName string `json:"fullName,omitempty"`

	// IdToken: The ID token.
	IdToken string `json:"idToken,omitempty"`

	// InputEmail: It's the identifier param in the createAuthUri request if
	// the identifier is an email. It can be used to check whether the user
	// input email is different from the asserted email.
	InputEmail string `json:"inputEmail,omitempty"`

	// Kind: The fixed string "identitytoolkit#VerifyAssertionResponse".
	Kind string `json:"kind,omitempty"`

	// Language: The language preference of the user.
	Language string `json:"language,omitempty"`

	// LastName: The last name of the user.
	LastName string `json:"lastName,omitempty"`

	// LocalId: The RP local ID if it's already been mapped to the IdP
	// account identified by the federated ID.
	LocalId string `json:"localId,omitempty"`

	// NickName: The nick name of the user.
	NickName string `json:"nickName,omitempty"`

	// OauthRequestToken: The user approved request token for the OpenID
	// OAuth extension.
	OauthRequestToken string `json:"oauthRequestToken,omitempty"`

	// OauthScope: The scope for the OpenID OAuth extension.
	OauthScope string `json:"oauthScope,omitempty"`

	// OriginalEmail: The original email stored in the mapping storage. It's
	// returned when the federated ID is associated to a different email.
	OriginalEmail string `json:"originalEmail,omitempty"`

	// PhotoUrl: The URI of the public accessible profiel picture.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name e.g. google.com, aol.com, live.net and yahoo.com. If the
	// "providerId" param is set to OpenID OP identifer other than the
	// whilte listed IdPs the OP identifier is returned. If the "identifier"
	// param is federated ID in the createAuthUri request. The domain part
	// of the federated ID is returned.
	ProviderId string `json:"providerId,omitempty"`

	// TimeZone: The timezone of the user.
	TimeZone string `json:"timeZone,omitempty"`

	// VerifiedProvider: When action is 'map', contains the idps which can
	// be used for confirmation.
	VerifiedProvider []string `json:"verifiedProvider,omitempty"`
}

type VerifyPasswordResponse struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email returned by the IdP. NOTE: The federated login user
	// may not own the email.
	Email string `json:"email,omitempty"`

	// IdToken: The GITKit token for authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#VerifyPasswordResponse".
	Kind string `json:"kind,omitempty"`

	// LocalId: The RP local ID if it's already been mapped to the IdP
	// account identified by the federated ID.
	LocalId string `json:"localId,omitempty"`

	// Registered: Whether the email is registered.
	Registered bool `json:"registered,omitempty"`
}

// method id "identitytoolkit.relyingparty.createAuthUri":

type RelyingpartyCreateAuthUriCall struct {
	s                                               *Service
	identitytoolkitrelyingpartycreateauthurirequest *IdentitytoolkitRelyingpartyCreateAuthUriRequest
	opt_                                            map[string]interface{}
}

// CreateAuthUri: Creates the URI used by the IdP to authenticate the
// user.
func (r *RelyingpartyService) CreateAuthUri(identitytoolkitrelyingpartycreateauthurirequest *IdentitytoolkitRelyingpartyCreateAuthUriRequest) *RelyingpartyCreateAuthUriCall {
	c := &RelyingpartyCreateAuthUriCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartycreateauthurirequest = identitytoolkitrelyingpartycreateauthurirequest
	return c
}

func (c *RelyingpartyCreateAuthUriCall) Do() (*CreateAuthUriResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartycreateauthurirequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "createAuthUri")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(CreateAuthUriResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the URI used by the IdP to authenticate the user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.createAuthUri",
	//   "path": "createAuthUri",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyCreateAuthUriRequest"
	//   },
	//   "response": {
	//     "$ref": "CreateAuthUriResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.deleteAccount":

type RelyingpartyDeleteAccountCall struct {
	s                                               *Service
	identitytoolkitrelyingpartydeleteaccountrequest *IdentitytoolkitRelyingpartyDeleteAccountRequest
	opt_                                            map[string]interface{}
}

// DeleteAccount: Delete user account.
func (r *RelyingpartyService) DeleteAccount(identitytoolkitrelyingpartydeleteaccountrequest *IdentitytoolkitRelyingpartyDeleteAccountRequest) *RelyingpartyDeleteAccountCall {
	c := &RelyingpartyDeleteAccountCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartydeleteaccountrequest = identitytoolkitrelyingpartydeleteaccountrequest
	return c
}

func (c *RelyingpartyDeleteAccountCall) Do() (*DeleteAccountResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartydeleteaccountrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "deleteAccount")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(DeleteAccountResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete user account.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.deleteAccount",
	//   "path": "deleteAccount",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyDeleteAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "DeleteAccountResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.getAccountInfo":

type RelyingpartyGetAccountInfoCall struct {
	s                                                *Service
	identitytoolkitrelyingpartygetaccountinforequest *IdentitytoolkitRelyingpartyGetAccountInfoRequest
	opt_                                             map[string]interface{}
}

// GetAccountInfo: Returns the account info.
func (r *RelyingpartyService) GetAccountInfo(identitytoolkitrelyingpartygetaccountinforequest *IdentitytoolkitRelyingpartyGetAccountInfoRequest) *RelyingpartyGetAccountInfoCall {
	c := &RelyingpartyGetAccountInfoCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartygetaccountinforequest = identitytoolkitrelyingpartygetaccountinforequest
	return c
}

func (c *RelyingpartyGetAccountInfoCall) Do() (*GetAccountInfoResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartygetaccountinforequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "getAccountInfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(GetAccountInfoResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the account info.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.getAccountInfo",
	//   "path": "getAccountInfo",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyGetAccountInfoRequest"
	//   },
	//   "response": {
	//     "$ref": "GetAccountInfoResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.getOobConfirmationCode":

type RelyingpartyGetOobConfirmationCodeCall struct {
	s            *Service
	relyingparty *Relyingparty
	opt_         map[string]interface{}
}

// GetOobConfirmationCode: Get a code for user action confirmation.
func (r *RelyingpartyService) GetOobConfirmationCode(relyingparty *Relyingparty) *RelyingpartyGetOobConfirmationCodeCall {
	c := &RelyingpartyGetOobConfirmationCodeCall{s: r.s, opt_: make(map[string]interface{})}
	c.relyingparty = relyingparty
	return c
}

func (c *RelyingpartyGetOobConfirmationCodeCall) Do() (*GetOobConfirmationCodeResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.relyingparty)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "getOobConfirmationCode")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(GetOobConfirmationCodeResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a code for user action confirmation.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.getOobConfirmationCode",
	//   "path": "getOobConfirmationCode",
	//   "request": {
	//     "$ref": "Relyingparty"
	//   },
	//   "response": {
	//     "$ref": "GetOobConfirmationCodeResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.resetPassword":

type RelyingpartyResetPasswordCall struct {
	s                                               *Service
	identitytoolkitrelyingpartyresetpasswordrequest *IdentitytoolkitRelyingpartyResetPasswordRequest
	opt_                                            map[string]interface{}
}

// ResetPassword: Set account info for a user.
func (r *RelyingpartyService) ResetPassword(identitytoolkitrelyingpartyresetpasswordrequest *IdentitytoolkitRelyingpartyResetPasswordRequest) *RelyingpartyResetPasswordCall {
	c := &RelyingpartyResetPasswordCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartyresetpasswordrequest = identitytoolkitrelyingpartyresetpasswordrequest
	return c
}

func (c *RelyingpartyResetPasswordCall) Do() (*ResetPasswordResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartyresetpasswordrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "resetPassword")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(ResetPasswordResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Set account info for a user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.resetPassword",
	//   "path": "resetPassword",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyResetPasswordRequest"
	//   },
	//   "response": {
	//     "$ref": "ResetPasswordResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.setAccountInfo":

type RelyingpartySetAccountInfoCall struct {
	s                                                *Service
	identitytoolkitrelyingpartysetaccountinforequest *IdentitytoolkitRelyingpartySetAccountInfoRequest
	opt_                                             map[string]interface{}
}

// SetAccountInfo: Set account info for a user.
func (r *RelyingpartyService) SetAccountInfo(identitytoolkitrelyingpartysetaccountinforequest *IdentitytoolkitRelyingpartySetAccountInfoRequest) *RelyingpartySetAccountInfoCall {
	c := &RelyingpartySetAccountInfoCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartysetaccountinforequest = identitytoolkitrelyingpartysetaccountinforequest
	return c
}

func (c *RelyingpartySetAccountInfoCall) Do() (*SetAccountInfoResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartysetaccountinforequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "setAccountInfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(SetAccountInfoResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Set account info for a user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.setAccountInfo",
	//   "path": "setAccountInfo",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartySetAccountInfoRequest"
	//   },
	//   "response": {
	//     "$ref": "SetAccountInfoResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.uploadAccount":

type RelyingpartyUploadAccountCall struct {
	s                                               *Service
	identitytoolkitrelyingpartyuploadaccountrequest *IdentitytoolkitRelyingpartyUploadAccountRequest
	opt_                                            map[string]interface{}
}

// UploadAccount: Batch upload existing user accounts.
func (r *RelyingpartyService) UploadAccount(identitytoolkitrelyingpartyuploadaccountrequest *IdentitytoolkitRelyingpartyUploadAccountRequest) *RelyingpartyUploadAccountCall {
	c := &RelyingpartyUploadAccountCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartyuploadaccountrequest = identitytoolkitrelyingpartyuploadaccountrequest
	return c
}

func (c *RelyingpartyUploadAccountCall) Do() (*UploadAccountResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartyuploadaccountrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "uploadAccount")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(UploadAccountResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Batch upload existing user accounts.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.uploadAccount",
	//   "path": "uploadAccount",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyUploadAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "UploadAccountResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.verifyAssertion":

type RelyingpartyVerifyAssertionCall struct {
	s                                                 *Service
	identitytoolkitrelyingpartyverifyassertionrequest *IdentitytoolkitRelyingpartyVerifyAssertionRequest
	opt_                                              map[string]interface{}
}

// VerifyAssertion: Verifies the assertion returned by the IdP.
func (r *RelyingpartyService) VerifyAssertion(identitytoolkitrelyingpartyverifyassertionrequest *IdentitytoolkitRelyingpartyVerifyAssertionRequest) *RelyingpartyVerifyAssertionCall {
	c := &RelyingpartyVerifyAssertionCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartyverifyassertionrequest = identitytoolkitrelyingpartyverifyassertionrequest
	return c
}

func (c *RelyingpartyVerifyAssertionCall) Do() (*VerifyAssertionResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartyverifyassertionrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "verifyAssertion")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(VerifyAssertionResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Verifies the assertion returned by the IdP.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.verifyAssertion",
	//   "path": "verifyAssertion",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyVerifyAssertionRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyAssertionResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.verifyPassword":

type RelyingpartyVerifyPasswordCall struct {
	s                                                *Service
	identitytoolkitrelyingpartyverifypasswordrequest *IdentitytoolkitRelyingpartyVerifyPasswordRequest
	opt_                                             map[string]interface{}
}

// VerifyPassword: Verifies the user entered password.
func (r *RelyingpartyService) VerifyPassword(identitytoolkitrelyingpartyverifypasswordrequest *IdentitytoolkitRelyingpartyVerifyPasswordRequest) *RelyingpartyVerifyPasswordCall {
	c := &RelyingpartyVerifyPasswordCall{s: r.s, opt_: make(map[string]interface{})}
	c.identitytoolkitrelyingpartyverifypasswordrequest = identitytoolkitrelyingpartyverifypasswordrequest
	return c
}

func (c *RelyingpartyVerifyPasswordCall) Do() (*VerifyPasswordResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.identitytoolkitrelyingpartyverifypasswordrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/identitytoolkit/v3/relyingparty/", "verifyPassword")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(VerifyPasswordResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Verifies the user entered password.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.verifyPassword",
	//   "path": "verifyPassword",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyVerifyPasswordRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyPasswordResponse"
	//   }
	// }

}
