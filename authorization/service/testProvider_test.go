package service

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/plgd-dev/cloud/authorization/provider"
	"github.com/plgd-dev/kit/codec/json"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

var jwkPrivateKey *ecdsa.PrivateKey
var jwkKeyID = `QkY4MzFGMTdFMzMyN0NGQjEyOUFFMzE5Q0ZEMUYzQUQxNkNENTlEMg`
var jwkKey jwk.Key
var clientID = "test"
var UserToken = ""
var DeviceAccessToken = "123"
var DeviceUserID = "1"
var DeviceExpiresIn = time.Minute * 60 * 24 * 30

func init() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("failed to generate key: %s", err)
	}
	jwkPrivateKey = privateKey
	key, err := jwk.New(&jwkPrivateKey.PublicKey)
	if err != nil {
		log.Fatalf("failed to create JWK: %s", err)
	}
	key.Set(jwk.KeyIDKey, jwkKeyID)
	key.Set(jwk.AlgorithmKey, jwa.ES256.String())
	jwkKey = key

	token, err := generateToken(false)
	if err != nil {
		log.Fatal(err)
	}
	UserToken = token.AccessToken
}

func generateToken(isService bool) (*provider.Token, error) {
	t := provider.Token{
		RefreshToken: "refresh-token",
		Expiry:       time.Now().Add(time.Hour * 24 * 365),
		Owner:        DeviceUserID,
	}
	token := jwt.New()
	if !isService {
		token.Set(jwt.SubjectKey, t.Owner)
	}
	token.Set(jwt.AudienceKey, []string{"https://127.0.0.1", "https://localhost"})
	token.Set(jwt.IssuedAtKey, time.Now().Unix())
	token.Set(jwt.ExpirationKey, t.Expiry.Unix())
	token.Set(`scope`, []string{"openid", "r:deviceinformation:*", "r:resources:*", "w:resources:*", "w:subscriptions:*"})
	token.Set(`client_id`, clientID)
	token.Set(`email`, `test@test.com`)
	token.Set(jwt.IssuerKey, "https://localhost/")
	buf, err := json.Encode(token)
	if err != nil {
		return nil, fmt.Errorf("failed to encode token: %s", err)
	}

	hdr := jws.NewHeaders()
	hdr.Set(jws.AlgorithmKey, jwa.ES256.String())
	hdr.Set(jws.TypeKey, `JWT`)
	hdr.Set(jws.KeyIDKey, jwkKeyID)
	payload, err := jws.Sign(buf, jwa.ES256, jwkPrivateKey, jws.WithHeaders(hdr))
	if err != nil {
		return nil, fmt.Errorf("failed to create UserToken: %s", err)
	}
	t.AccessToken = string(payload)
	return &t, nil
}

// NewTestProvider creates GitHub Oauth client.
func NewTestProvider() *TestProvider {
	return &TestProvider{}
}

// TestProvider basic configuration.
type TestProvider struct {
}

// GetProviderName provides provider name.
func (p *TestProvider) GetProviderName() string {
	return clientID
}

// Exchange Auth Code for Access Token via OAuth.
func (p *TestProvider) Exchange(ctx context.Context, authorizationProvider, authorizationCode string) (*provider.Token, error) {
	return &provider.Token{
		Owner:        DeviceUserID,
		AccessToken:  DeviceAccessToken,
		Expiry:       time.Now().Add(DeviceExpiresIn),
		RefreshToken: "refresh-token",
	}, nil
}

// Refresh gets new Access Token via OAuth.
func (p *TestProvider) Refresh(ctx context.Context, refreshToken string) (*provider.Token, error) {
	return &provider.Token{
		Owner:        DeviceUserID,
		AccessToken:  DeviceAccessToken,
		Expiry:       time.Now().Add(DeviceExpiresIn),
		RefreshToken: "refresh-token",
	}, nil
}

// AuthCodeURL returns URL for redirecting.
func (p *TestProvider) AuthCodeURL(csrfToken string) string {
	return "redirect-url"
}

func (p *TestProvider) HandleAuthorizationCode(ctx *fasthttp.RequestCtx) {
	uri := ctx.QueryArgs().Peek("redirect_uri")
	if len(uri) > 0 {
		state := ctx.QueryArgs().Peek("state")
		u, err := url.Parse(string(uri))
		if err != nil {
			setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
			return
		}
		q, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
			return
		}
		q.Add("state", string(state))
		q.Add("code", DeviceAccessToken)
		u.RawQuery = q.Encode()
		ctx.Redirect(u.String(), fasthttp.StatusTemporaryRedirect)
		return
	}
	resp := map[string]interface{}{
		"code": DeviceAccessToken,
	}
	data, err := json.Encode(resp)
	if err != nil {
		setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	r := &ctx.Response
	r.Header.SetContentType("application/json")
	r.SetStatusCode(fasthttp.StatusOK)
	r.SetBodyString(string(data))
}

func (p *TestProvider) HandleAccessToken(ctx *fasthttp.RequestCtx) {
	clientID := string(ctx.QueryArgs().Peek("ClientId"))
	var isService bool
	if clientID == "service" {
		isService = true
	}
	token, err := generateToken(isService)
	if err != nil {
		setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	resp := map[string]interface{}{
		"access_token": token.AccessToken,
		"expires_in":   int64(time.Until(token.Expiry).Seconds()),
		"scope":        "openid",
		"token_type":   "Bearer",
	}
	data, err := json.Encode(resp)
	if err != nil {
		setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	r := &ctx.Response
	r.Header.SetContentType("application/json")
	r.SetStatusCode(fasthttp.StatusOK)
	r.SetBodyString(string(data))
}

func (p *TestProvider) HandleJWKs(ctx *fasthttp.RequestCtx) {
	resp := map[string]interface{}{
		"keys": []jwk.Key{
			jwkKey,
		},
	}
	data, err := json.Encode(resp)
	if err != nil {
		setErrorResponse(&ctx.Response, fasthttp.StatusInternalServerError, err.Error())
		return
	}

	r := &ctx.Response
	r.Header.SetContentType("application/json")
	r.SetStatusCode(fasthttp.StatusOK)
	r.SetBodyString(string(data))
}
