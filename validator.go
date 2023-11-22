package evaluator

import (
	"fmt"

	jira "github.com/DMXMax/jira-subsystem"
)

type EvaluationAction int8

func (a EvaluationAction) done() bool {
	return (a == ActionFail || a == ActionManual)
}

/*
rules of the game

A client in the ClientAllowList is allowed unless it is covered by the ClientDenyList
A scope in the ScopeAllowList is allowed unless it is covered by the ScopeDenyList

if scope in ScopeDenylist or client in ClientDenyList
	deny

if scope in ScopeAllowList and client in ClientAllowList
	allow
if not,
	deny

*/

// clients that match any entry in this list require manual approvals
type ClientDenyList []Validator

func (cdl *ClientDenyList) Evaluate(t *jira.ValidatedTicket, res *EvaluationResults) error {
	for _, req := range t.ScopeRequests {
		for _, rule := range *cdl {

			if rule.Validate(req.Client) { //Being inside a deny list means the client must be manually approved
				res.SetAction(ActionManual)
				res.addMessage(fmt.Sprintf("Requests for client %v are autmatically set to manual approval", req.Client))
			}

		}

	}
	return nil
}

// clients not in any of these validators require manual approval
type ClientAllowList []Validator

func (cal *ClientAllowList) Evaluate(t *jira.ValidatedTicket, res *EvaluationResults) error {
	for _, req := range t.ScopeRequests {
		for _, rule := range *cal {

			if rule.Validate(req.Client) { //Being inside a allow list means the client must be manually approved
				res.SetAction(ActionApproveAppsec)
			}
		}

		// If we are not explicitly allowed to be automated, the request goes to a manual state
		if res.action == ActionUnset {
			res.SetAction(ActionManual)
			res.addMessage(fmt.Sprintf("Client %v is not in the allow list so is set to manual approval", req.Client))
		}

	}
	return nil

}

type ScopeDenyList []Validator

func (sdl *ScopeDenyList) Evaluate(t *jira.ValidatedTicket, res *EvaluationResults) error {
	for _, req := range t.ScopeRequests {
		for _, rule := range *sdl {

			if rule.Validate(req.Scope) { //Being inside a deny list means the client must be manually approved
				res.SetAction(ActionManual)
				res.addMessage(fmt.Sprintf("Requests for scope %v are autmatically set to manual approval", req.Scope))
			}

		}

	}
	return nil
}

type ScopeAllowList []Validator

func (sal *ScopeAllowList) Evaluate(t *jira.ValidatedTicket, res *EvaluationResults) error {
	for _, req := range t.ScopeRequests {
		for _, rule := range *sal {

			if rule.Validate(req.Client) { //Being inside a allow list means the client must be manually approved
				res.SetAction(ActionApproveAppsec)
			}
		}

		// If we are not explicitly allowed to be automated, the request goes to a manual state
		if res.action == ActionUnset {
			res.SetAction(ActionManual)
			res.addMessage(fmt.Sprintf("Scope %v not in the allow list. Setting to Manual Approval", req.Scope))
		}

	}
	return nil

}

var cal = ClientAllowList{
	&RegexValidator{
		Filter:      ".+",
		FailMessage: "%s is not in the client allow list.",
	},
}

var cdl = ClientDenyList{}

var sdl = ScopeDenyList{
	&RegexValidator{
		Filter: "vault.+",
	},
}

var sal = ScopeAllowList{}
