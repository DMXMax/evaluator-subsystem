package evaluator

import jira "github.com/DMXMax/jira-subsystem"

const (
	ActionUnset         EvaluationAction = iota //Initial state. If the Evaluator returns this state, this is an error
	ActionFail                                  // Failure in the system or the scope is syntatically wrong. The system cannot find the scope
	ActionManual                                // Both Appsec and Platsec manual approval is required
	ActionApproveAppsec                         // Platsec manual approval is required
	ActionApproveBoth                           // Neither is required to provide manual approval
)

type EvaluationResults struct {
	EvaluatorVersion string
	Ticket           jira.ValidatedTicket
	action           EvaluationAction
	Messages         []string
}

func (r *EvaluationResults) addMessage(str string) {
	r.Messages = append(r.Messages, str)

}

// handles the rules of setting actions.
// Manual Action is sticky and can only be set to Error
func (r *EvaluationResults) SetAction(a EvaluationAction) {
	if a == ActionUnset || r.action == ActionFail {
		return
	}

	if r.action == ActionManual && a != ActionFail {
		return
	}

	if a == ActionApproveBoth && r.action != ActionUnset {
		return
	}

	r.action = a

}

func (r *EvaluationResults) GetAction() EvaluationAction {
	return r.action
}

func (r *EvaluationResults) Done() bool {
	return r.action.done()
}
