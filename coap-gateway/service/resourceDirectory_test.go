package service_test

import (
	"context"
	"testing"

	oauthTest "github.com/plgd-dev/cloud/authorization/provider"
	"github.com/plgd-dev/cloud/coap-gateway/uri"
	testCfg "github.com/plgd-dev/cloud/test/config"
	"github.com/plgd-dev/go-coap/v2/message"
	coapCodes "github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/tcp"
	"github.com/stretchr/testify/require"
)

type TestResource struct {
	DeviceID string `json:"di"`
	//Eps interface{} `json:"eps"`
	Href       string   `json:"href"`
	ID         string   `json:"id"`
	Interfaces []string `json:"if"`
	InstanceID uint64   `json:"-"`
	//P             interface{} `json:"p"`
	ResourceTypes []string `json:"rt"`
	Type          []string `json:"type"`
}

type TestWkRD struct {
	DeviceID         string         `json:"di"`
	Links            []TestResource `json:"links"`
	TimeToLive       int            `json:"ttl"`
	TimeToLiveLegacy int            `json:"lt"`
}

var tblResourceDirectory = []testEl{
	{"BadRequest0", input{coapCodes.POST, `{ "di":"` + CertIdentity + `" }`, nil}, output{coapCodes.BadRequest, `empty links`, nil}},
	{"BadRequest1", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":"abc" }`, nil}, output{coapCodes.BadRequest, `cannot publish resource: cbor: cannot unmarshal`, nil}},
	{"BadRequest2", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ "abc" ]}`, nil}, output{coapCodes.BadRequest, `cannot publish resource: cbor: cannot unmarshal`, nil}},
	{"BadRequest4", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "href":"` + TestAResourceHref + `" } ], "ttl":-1}`, nil}, output{coapCodes.BadRequest, `invalid TimeToLive`, nil}},
	{"BadRequest5", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "di":"` + CertIdentity + `", "href":"" } ], "ttl":12345}`, nil}, output{coapCodes.BadRequest, `empty links`, nil}},
	{"BadRequest6", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "href":"" } ], "ttl":12345}`, nil}, output{coapCodes.BadRequest, `empty links`, nil}},
	{"Changed0", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "di":"` + CertIdentity + `", "href":"` + TestAResourceHref + `" } ], "ttl":12345}`, nil},
		output{coapCodes.Changed, TestWkRD{
			DeviceID:         CertIdentity,
			TimeToLive:       12345,
			TimeToLiveLegacy: 12345,
			Links: []TestResource{
				{
					DeviceID: CertIdentity,
					Href:     TestAResourceHref,
					ID:       TestAResourceId,
				},
			},
		}, nil}},

	{"Changed1", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "di":"` + CertIdentity + `", "href":"/b" } ], "ttl":12345}`, nil},
		output{coapCodes.Changed, TestWkRD{
			DeviceID:         CertIdentity,
			TimeToLive:       12345,
			TimeToLiveLegacy: 12345,
			Links: []TestResource{
				{
					DeviceID: CertIdentity,
					Href:     "/b",
					ID:       "1f36abb2-c5f8-556e-bf74-3b34ed66a2b4",
				},
			},
		}, nil}},
	{"Changed2", input{coapCodes.POST, `{ "di":"` + CertIdentity + `", "links":[ { "di":"` + CertIdentity + `", "href":"/b" } , { "di":"` + CertIdentity + `", "href":"/c" }], "ttl":12345}`, nil},
		output{coapCodes.Changed, TestWkRD{
			DeviceID:         CertIdentity,
			TimeToLive:       12345,
			TimeToLiveLegacy: 12345,
			Links: []TestResource{
				{
					DeviceID: CertIdentity,
					Href:     "/b",
					ID:       "1f36abb2-c5f8-556e-bf74-3b34ed66a2b4",
				},
				{
					DeviceID: CertIdentity,
					Href:     "/c",
					ID:       "41529a9c-b80f-5487-82da-da4a476402ae",
				},
			},
		}, nil}},
}

func TestResourceDirectoryPostHandler(t *testing.T) {
	shutdown := setUp(t)
	defer shutdown()

	co := testCoapDial(t, testCfg.GW_HOST)
	if co == nil {
		return
	}
	defer co.Close()

	signUpEl := testEl{"signUp", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "authprovider": "` + oauthTest.NewTestProvider().GetProviderName() + `"}`, nil}, output{coapCodes.Changed, TestCoapSignUpResponse{RefreshToken: "refresh-token", UserID: AuthorizationUserId}, nil}}
	t.Run(signUpEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignUp, signUpEl, co)
	})
	signInEl := testEl{"signIn", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "uid":"` + AuthorizationUserId + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "login": true }`, nil}, output{coapCodes.Changed, TestCoapSignInResponse{}, nil}}
	t.Run(signInEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignIn, signInEl, co)
	})

	for _, test := range tblResourceDirectory {
		tf := func(t *testing.T) {
			testPostHandler(t, uri.ResourceDirectory, test, co)
		}
		t.Run(test.name, tf)
	}
}

func TestResourceDirectoryDeleteHandler(t *testing.T) {
	//set counter 0, when other test run with this that it can be modified
	deletetblResourceDirectory := []testEl{
		{"NotExist1", input{coapCodes.DELETE, ``, []string{"di=c", "ins=5"}}, output{coapCodes.BadRequest, `cannot found resources for the DELETE request parameters`, nil}},                 // Non-existent device ID.
		{"NotExist2", input{coapCodes.DELETE, ``, []string{"ins=4"}}, output{coapCodes.BadRequest, `cannot parse queries: deviceID not found`, nil}},                                         // Device ID empty.
		{"NotExist3", input{coapCodes.DELETE, ``, []string{`di=` + CertIdentity, "ins=999"}}, output{coapCodes.BadRequest, `cannot found resources for the DELETE request parameters`, nil}}, // Instance ID non-existent.
		{"Exist1", input{coapCodes.DELETE, ``, []string{`di=` + CertIdentity}}, output{coapCodes.Deleted, nil, nil}},                                                                         // If instanceIDs empty, all instances for a given device ID should be unpublished.
		{"NotExist4", input{coapCodes.DELETE, ``, []string{`di=` + CertIdentity}}, output{coapCodes.BadRequest, `cannot found resources for the DELETE request parameters`, nil}},
	}

	shutdown := setUp(t)
	defer shutdown()

	co := testCoapDial(t, testCfg.GW_HOST)
	if co == nil {
		return
	}
	defer co.Close()

	signUpEl := testEl{"signUp", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "authprovider": "` + oauthTest.NewTestProvider().GetProviderName() + `"}`, nil}, output{coapCodes.Changed, TestCoapSignUpResponse{RefreshToken: "refresh-token", UserID: AuthorizationUserId}, nil}}
	t.Run(signUpEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignUp, signUpEl, co)
	})
	signInEl := testEl{"signIn", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "uid":"` + AuthorizationUserId + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "login": true }`, nil}, output{coapCodes.Changed, TestCoapSignInResponse{}, nil}}
	t.Run(signInEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignIn, signInEl, co)
	})

	// Publish resources first!
	for _, test := range tblResourceDirectory {
		tf := func(t *testing.T) {
			testPostHandler(t, uri.ResourceDirectory, test, co)
		}
		t.Run(test.name, tf)
	}

	//delete resources
	for _, test := range deletetblResourceDirectory {
		tf := func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), TestExchangeTimeout)
			defer cancel()
			req, err := tcp.NewDeleteRequest(ctx, uri.ResourceDirectory)
			require.NoError(t, err)
			for _, q := range test.in.queries {
				req.AddOptionString(message.URIQuery, q)
			}
			resp, err := co.Do(req)
			if err != nil {
				t.Fatalf("Cannot send/retrieve msg: %v", err)
			}
			testValidateResp(t, test, resp)
		}
		t.Run(test.name, tf)
	}
}

type TestGetSelector struct {
	Selector uint64 `json:"sel"`
}

func TestResourceDirectoryGetSelector(t *testing.T) {
	tbl := []testEl{
		{"GetSelector", input{coapCodes.GET, ``, []string{}}, output{coapCodes.Content, TestGetSelector{}, nil}},
	}

	shutdown := setUp(t)
	defer shutdown()

	co := testCoapDial(t, testCfg.GW_HOST)
	if co == nil {
		return
	}
	defer co.Close()

	signUpEl := testEl{"signUp", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "authprovider": "` + oauthTest.NewTestProvider().GetProviderName() + `"}`, nil}, output{coapCodes.Changed, TestCoapSignUpResponse{RefreshToken: "refresh-token", UserID: AuthorizationUserId}, nil}}
	t.Run(signUpEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignUp, signUpEl, co)
	})
	signInEl := testEl{"signIn", input{coapCodes.POST, `{"di": "` + CertIdentity + `", "uid":"` + AuthorizationUserId + `", "accesstoken":"` + oauthTest.DeviceAccessToken + `", "login": true }`, nil}, output{coapCodes.Changed, TestCoapSignInResponse{}, nil}}
	t.Run(signInEl.name, func(t *testing.T) {
		testPostHandler(t, uri.SignIn, signInEl, co)
	})

	for _, test := range tbl {
		tf := func(t *testing.T) {
			req, err := tcp.NewGetRequest(co.Context(), uri.ResourceDirectory)
			require.NoError(t, err)
			for _, q := range test.in.queries {
				req.AddOptionString(message.URIQuery, q)
			}
			resp, err := co.Do(req)
			require.NoError(t, err)
			testValidateResp(t, test, resp)
		}
		t.Run(test.name, tf)
	}
}
