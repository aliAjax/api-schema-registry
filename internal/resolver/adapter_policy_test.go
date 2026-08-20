package resolver

import "testing"

func TestRemotePolicyRejectsUnlistedHost(t *testing.T) {
	if (RemotePolicy{AllowHosts: []string{"registry.internal"}}).Allows("evil.example", false) {
		t.Fatal("unlisted remote host was allowed")
	}
}
