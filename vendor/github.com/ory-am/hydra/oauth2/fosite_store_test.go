package oauth2

import (
	"net/url"
	"os"
	"testing"
	"time"

	"fmt"
	"github.com/Sirupsen/logrus"
	c "github.com/ory-am/common/pkg"
	"github.com/ory-am/fosite"
	"github.com/ory-am/hydra/client"
	"github.com/ory-am/hydra/integration"
	"github.com/ory-am/hydra/pkg"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	r "gopkg.in/dancannon/gorethink.v2"
)

var rethinkManager *FositeRehinkDBStore
var clientManagers = map[string]pkg.FositeStorer{}
var clientManager = &client.MemoryManager{
	Clients: map[string]client.Client{"foobar": {ID: "foobar"}},
	Hasher:  &fosite.BCrypt{},
}

func init() {
	clientManagers["memory"] = &FositeMemoryStore{
		AuthorizeCodes: make(map[string]fosite.Requester),
		IDSessions:     make(map[string]fosite.Requester),
		AccessTokens:   make(map[string]fosite.Requester),
		RefreshTokens:  make(map[string]fosite.Requester),
	}
}

func TestMain(m *testing.M) {
	connectToPG()
	connectToRethinkDB()
	connectToMySQL()
	connectToRedis()

	s := m.Run()
	integration.KillAll()
	os.Exit(s)
}

func connectToPG() {
	var db = integration.ConnectToPostgres()
	s := &FositeSQLStore{DB: db, Manager: clientManager}
	if err := s.CreateSchemas(); err != nil {
		logrus.Fatalf("Could not create postgres schema: %v", err)
	}

	clientManagers["postgres"] = s
}

func connectToRethinkDB() {
	var session = integration.ConnectToRethinkDB("hydra", "hydra_authorize_code", "hydra_id_sessions", "hydra_access_token", "hydra_refresh_token")

	rethinkManager = &FositeRehinkDBStore{
		Session:             session,
		AuthorizeCodesTable: r.Table("hydra_authorize_code"),
		IDSessionsTable:     r.Table("hydra_id_sessions"),
		AccessTokensTable:   r.Table("hydra_access_token"),
		RefreshTokensTable:  r.Table("hydra_refresh_token"),
		AuthorizeCodes:      make(RDBItems),
		IDSessions:          make(RDBItems),
		AccessTokens:        make(RDBItems),
		RefreshTokens:       make(RDBItems),
	}
	rethinkManager.Watch(context.Background())
	clientManagers["rethink"] = rethinkManager
}

func connectToMySQL() {
	var db = integration.ConnectToMySQL()
	s := &FositeSQLStore{DB: db, Manager: clientManager}
	if err := s.CreateSchemas(); err != nil {
		logrus.Fatalf("Could not create postgres schema: %v", err)
	}

	clientManagers["mysql"] = s
}

func connectToRedis() {
	var db = integration.ConnectToRedis()
	clientManagers["redis"] = &FositeRedisStore{DB: db}
}

var defaultRequest = fosite.Request{
	RequestedAt:   time.Now().Round(time.Second),
	Client:        &client.Client{ID: "foobar"},
	Scopes:        fosite.Arguments{"fa", "ba"},
	GrantedScopes: fosite.Arguments{"fa", "ba"},
	Form:          url.Values{"foo": []string{"bar", "baz"}},
	Session:       &fosite.DefaultSession{Subject: "bar"},
}

func TestColdStartRethinkManager(t *testing.T) {
	ctx := context.Background()
	m := rethinkManager
	id := uuid.New()

	err := m.CreateAuthorizeCodeSession(ctx, id, &defaultRequest)
	pkg.AssertError(t, false, err)
	err = m.CreateAccessTokenSession(ctx, "12345", &fosite.Request{
		RequestedAt: time.Now().Round(time.Second),
		Client:      &client.Client{ID: "baz"},
	})
	pkg.AssertError(t, false, err)

	err = m.CreateAccessTokenSession(ctx, id, &defaultRequest)
	pkg.AssertError(t, false, err)

	_, err = m.GetAuthorizeCodeSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, false, err)
	_, err = m.GetAccessTokenSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, false, err)

	delete(rethinkManager.AuthorizeCodes, id)
	delete(rethinkManager.AccessTokens, id)
	delete(rethinkManager.AccessTokens, "12345")

	_, err = m.GetAuthorizeCodeSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, true, err)
	_, err = m.GetAccessTokenSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, true, err)

	err = rethinkManager.ColdStart()
	pkg.AssertError(t, false, err)

	_, err = m.GetAuthorizeCodeSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, false, err)

	s1, err := m.GetAccessTokenSession(ctx, id, &fosite.DefaultSession{})
	pkg.AssertError(t, false, err)
	s2, err := m.GetAccessTokenSession(ctx, "12345", &fosite.DefaultSession{})
	pkg.AssertError(t, false, err)
	assert.NotEqual(t, s1, s2)
}

func TestCreateImplicitAccessTokenSession(t *testing.T) {
	ctx := context.Background()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {

			_, err := m.GetAccessTokenSession(ctx, "implicit-4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			err = m.CreateImplicitAccessTokenSession(ctx, "implicit-4321", &defaultRequest)
			assert.Nil(t, err)

			res, err := m.GetAccessTokenSession(ctx, "implicit-4321", &fosite.DefaultSession{})
			require.Nil(t, err)
			c.AssertObjectKeysEqual(t, &defaultRequest, res, "Scopes", "GrantedScopes", "Form", "Session")

			err = m.DeleteAccessTokenSession(ctx, "implicit-4321")
			assert.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetAccessTokenSession(ctx, "implicit-4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)
		})
	}
}
func TestCreateGetDeleteAuthorizeCodes(t *testing.T) {
	ctx := context.Background()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
			_, err := m.GetAuthorizeCodeSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			err = m.CreateAuthorizeCodeSession(ctx, "4321", &defaultRequest)
			require.Nil(t, err)

			res, err := m.GetAuthorizeCodeSession(ctx, "4321", &fosite.DefaultSession{})
			require.Nil(t, err)
			c.AssertObjectKeysEqual(t, &defaultRequest, res, "Scopes", "GrantedScopes", "Form", "Session")

			err = m.DeleteAuthorizeCodeSession(ctx, "4321")
			require.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetAuthorizeCodeSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)
		})
	}
}

func TestCreateGetDeleteAccessTokenSession(t *testing.T) {
	ctx := context.Background()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
			_, err := m.GetAccessTokenSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			err = m.CreateAccessTokenSession(ctx, "4321", &defaultRequest)
			require.Nil(t, err)

			res, err := m.GetAccessTokenSession(ctx, "4321", &fosite.DefaultSession{})
			require.Nil(t, err)
			c.AssertObjectKeysEqual(t, &defaultRequest, res, "Scopes", "GrantedScopes", "Form", "Session")

			err = m.DeleteAccessTokenSession(ctx, "4321")
			require.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetAccessTokenSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)
		})
	}
}

func TestCreateGetDeleteOpenIDConnectSession(t *testing.T) {
	ctx := context.Background()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
			_, err := m.GetOpenIDConnectSession(ctx, "4321", &fosite.Request{})
			assert.NotNil(t, err)

			err = m.CreateOpenIDConnectSession(ctx, "4321", &defaultRequest)
			require.Nil(t, err)

			res, err := m.GetOpenIDConnectSession(ctx, "4321", &fosite.Request{Session: &fosite.DefaultSession{}})
			require.Nil(t, err)
			c.AssertObjectKeysEqual(t, &defaultRequest, res, "Scopes", "GrantedScopes", "Form", "Session")

			err = m.DeleteOpenIDConnectSession(ctx, "4321")
			require.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetOpenIDConnectSession(ctx, "4321", &fosite.Request{})
			assert.NotNil(t, err)
		})
	}
}

func TestCreateGetDeleteRefreshTokenSession(t *testing.T) {
	ctx := context.Background()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
			_, err := m.GetRefreshTokenSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			err = m.CreateRefreshTokenSession(ctx, "4321", &defaultRequest)
			require.Nil(t, err)

			res, err := m.GetRefreshTokenSession(ctx, "4321", &fosite.DefaultSession{})
			require.Nil(t, err)
			c.AssertObjectKeysEqual(t, &defaultRequest, res, "Scopes", "GrantedScopes", "Form", "Session")

			err = m.DeleteRefreshTokenSession(ctx, "4321")
			require.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetRefreshTokenSession(ctx, "4321", &fosite.DefaultSession{})
			assert.NotNil(t, err)
		})
	}
}

func TestRevokeRefreshToken(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
			_, err := m.GetRefreshTokenSession(ctx, "1111", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			err = m.CreateRefreshTokenSession(ctx, "1111", &fosite.Request{ID: id, Client: &client.Client{ID: "foobar"}, RequestedAt: time.Now().Round(time.Second)})
			require.Nil(t, err)

			err = m.CreateRefreshTokenSession(ctx, "1122", &fosite.Request{ID: id, Client: &client.Client{ID: "foobar"}, RequestedAt: time.Now().Round(time.Second)})
			require.Nil(t, err)

			_, err = m.GetRefreshTokenSession(ctx, "1111", &fosite.DefaultSession{})
			require.Nil(t, err)

			err = m.RevokeRefreshToken(ctx, id)
			require.Nil(t, err)

			time.Sleep(100 * time.Millisecond)

			_, err = m.GetRefreshTokenSession(ctx, "1111", &fosite.DefaultSession{})
			assert.NotNil(t, err)

			_, err = m.GetRefreshTokenSession(ctx, "1122", &fosite.DefaultSession{})
			assert.NotNil(t, err)
		})
	}
}
