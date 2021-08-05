package service_test

import (
	"testing"

	"github.com/plgd-dev/cloud/coap-gateway/uri"
	testCfg "github.com/plgd-dev/cloud/test/config"
	coapCodes "github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/stretchr/testify/require"
)

func TestSignOffHandler(t *testing.T) {
	shutdown := setUp(t)
	defer shutdown()

	co := testCoapDial(t, testCfg.GW_HOST)
	require.NotNil(t, co)
	signUpResp := testSignUp(t, CertIdentity, co)
	err := co.Close()
	require.NoError(t, err)

	tbl := []testEl{
		{"Bad query", input{coapCodes.DELETE, `{}`, []string{"di=%"}}, output{coapCodes.BadOption, "invalid URL escape", nil}, true},
		{"Bad request (no userId)", input{coapCodes.DELETE, `{}`, []string{"di=" + CertIdentity}}, output{coapCodes.BadRequest, "invalid user id", nil}, true},
		{"Bad request (invalid userId)", input{coapCodes.DELETE, `{}`, []string{"di=" + CertIdentity, "accesstoken=" + signUpResp.AccessToken, "uid=0"}}, output{coapCodes.InternalServerError, "invalid ownerClaim", nil}, true},
		{"Bad request (missing access token)", input{coapCodes.DELETE, `{}`, []string{"di=" + CertIdentity, "uid=0"}}, output{coapCodes.BadRequest, `invalid access token`, nil}, true},
		{"Deleted0", input{coapCodes.DELETE, `{}`, []string{"di=" + CertIdentity, "accesstoken=" + signUpResp.AccessToken, "uid=" + signUpResp.UserID}}, output{coapCodes.Deleted, nil, nil}, false},
	}

	for _, test := range tbl {
		tf := func(t *testing.T) {
			co := testCoapDial(t, testCfg.GW_HOST)
			require.NotNil(t, co)
			defer func() {
				_ = co.Close()
			}()

			// delete record for signUp
			testPostHandler(t, uri.SignUp, test, co)
		}
		t.Run(test.name, tf)
	}
}

func TestSignOffWithSignInHandler(t *testing.T) {
	shutdown := setUp(t)
	defer shutdown()

	tbl := []testEl{
		{"Deleted", input{coapCodes.DELETE, `{}`, nil}, output{coapCodes.Deleted, nil, nil}, false},
	}

	for _, test := range tbl {
		tf := func(t *testing.T) {
			co := testCoapDial(t, testCfg.GW_HOST)
			require.NotNil(t, co)
			testSignUpIn(t, CertIdentity, co)
			defer func() {
				_ = co.Close()
			}()
			// delete record for signUp
			testPostHandler(t, uri.SignUp, test, co)
		}
		t.Run(test.name, tf)
	}
}
