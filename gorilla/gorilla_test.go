package gorilla_test

import (
	"testing"

	"github.com/gorilla/sessions"
	"github.com/simonedbarber/session/test"
	"github.com/simonedbarber/session/gorilla"
)

func TestAll(t *testing.T) {
	engine := sessions.NewCookieStore([]byte("something-very-secret"))
	manager := gorilla.New("_session", engine)
	test.TestAll(manager, t)
}
