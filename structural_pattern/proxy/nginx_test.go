package proxy

import (
	"testing"
)

func TestNginx(t *testing.T) {
	nginxServer := NewNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	t.Logf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	t.Logf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	t.Logf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "POST")
	t.Logf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "GET")
	t.Logf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
