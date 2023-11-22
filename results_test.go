package evaluator

import (
	"testing"
)

func TestResults(t *testing.T) {
	results1 := new(EvaluationResults)

	if results1.GetAction() != ActionUnset {
		t.Fatalf("results should start out as %v. This one had %v", ActionUnset, results1.action)
	}

	results1.SetAction(ActionFail)

	if results1.GetAction() != ActionFail {
		t.Fatalf("Set Test: results should be %v, but this result had %v", ActionFail, results1.action)
	} else {
		t.Log("pass setting failure test")
	}

	results1.SetAction(ActionApproveAppsec) // this should fail. The action is already in a failed state

	if results1.GetAction() != ActionFail {
		t.Fatalf("Set Test: results should be %v, but this result had %v", ActionFail, results1.action)
	} else {
		t.Log("passed setting after failure test")
	}

	results2 := new(EvaluationResults)

	results2.SetAction(ActionManual)

	if results2.GetAction() != ActionManual {
		t.Fatalf("Set Test: results should be %v, but this result had %v", ActionManual, results1.action)
	} else {
		t.Log("passed setting ActionManual")
	}

	results2.SetAction(ActionFail)

	if results1.GetAction() != ActionFail {
		t.Fatalf("Set Test: results should be %v, but this result had %v", ActionFail, results1.action)
	} else {
		t.Log("passed setting fail after manual test")
	}

	results2.SetAction(ActionManual)

	if results1.GetAction() != ActionFail {
		t.Fatalf("Set Test: results should be %v, but this result had %v", ActionFail, results1.action)
	} else {
		t.Log("passed fail sticky after manual test")
	}

}

func TestDone(t *testing.T) {
	results := new(EvaluationResults)
	if results.Done() != false {
		t.Fatalf("These results should be marked Done() but Done() returned %v", results.Done())
	} else {
		t.Logf("results Done()  on action %v returned %v", results.GetAction(), results.Done())
	}

	results.SetAction(ActionApproveAppsec)
	if results.Done() != false {
		t.Fatalf("These results should be marked Done() but Done() returned %v", results.Done())
	} else {
		t.Logf("results Done()  on action %v returned %v", results.GetAction(), results.Done())
	}

	results.SetAction(ActionApproveBoth)
	if results.Done() != false {
		t.Fatalf("These results should be marked Done() but Done() returned %v", results.Done())
	} else {
		t.Logf("results Done()  on action %v returned %v", results.GetAction(), results.Done())
	}

	results.SetAction(ActionManual)

	if results.Done() != true {
		t.Fatalf("These results should be marked Done() but Done() returned %v", results.Done())
	} else {
		t.Logf("results Done()  on action %v returned %v", results.GetAction(), results.Done())
	}

	results.SetAction(ActionFail)

	if results.Done() != true {
		t.Fatalf("These results should be marked Done() but Done() returned %v", results.Done())
	} else {
		t.Logf("results Done()  on action %v returned %v", results.GetAction(), results.Done())
	}
}
