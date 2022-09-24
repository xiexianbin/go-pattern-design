package bridge

import "testing"

func TestApp(t *testing.T) {
	appCorp := AppCorp{
		App: &ChatApp{},
	}

	appCorp.developApp()
}
