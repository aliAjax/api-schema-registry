package compatibility

import "testing"

func TestCompareReportsRemovedProperty(t *testing.T) {
	old := map[string]any{"type": "object", "properties": map[string]any{"id": map[string]any{"type": "string"}}}
	newDoc := map[string]any{"type": "object", "properties": map[string]any{}}
	result := Compare(old, newDoc, Backward)
	if result.Compatible || len(result.Differences) != 1 {
		t.Fatalf("result = %#v", result)
	}
}

func TestPolicyCheckRejectsBreakingResult(t *testing.T) {
	if err := (Policy{Mode: "backward"}).Check(Result{Compatible: false}); err == nil {
		t.Fatal("breaking result was accepted")
	}
}

func TestSummaryIncludesDifferenceCount(t *testing.T) {
	summary := Summary(Result{Compatible: false, Differences: []Difference{{Code: "removed"}}})
	if summary["difference_count"] != 1 {
		t.Fatalf("summary = %#v", summary)
	}
}
