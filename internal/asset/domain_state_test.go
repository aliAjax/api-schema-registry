package asset

import "testing"

func TestTerminalStatusClassification(t *testing.T) {
	if !Withdrawn.Terminal() || !Deprecated.Terminal() || Published.Terminal() {
		t.Fatal("terminal status classification is inconsistent")
	}
}
