package evaluator

import (
	"fmt"

	jira "github.com/DMXMax/jira-subsystem"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var evlog zerolog.Logger

const version = "0.0.1"

// ValidateAndRetrieveTicket retrieves a ticket from Jira and validates the ticket is correct.
// If the ticket is not correct, ValidateAndRetrieveTicket returns JiraNotFound in its error
func EvaluateTicket(t jira.ValidatedTicket) (EvaluationResults, error) {

	evlog.Trace().Interface("EvaluateTicket", t).Send()

	results := EvaluationResults{
		EvaluatorVersion: version,
		Ticket:           t,
	}

	// check for syntactic correctness in the scope
	results = wellFormed(t, results)

	if results.Done() {
		return results, nil
	}

	if t.IsTest {
		// Scopes in test are not evaluated further
		results.addMessage("Scopes for Test and Development are automatically approved without evaluation and added to corp-overrides")
		results.SetAction(ActionApproveBoth)

		return results, nil
	}

	// check the client deny list. Update results
	cdl.Evaluate(&t, &results)

	// check the scope deny list. updte results
	sdl.Evaluate(&t, &results)

	if results.Done() {
		return results, nil
	}
	//check the client allow list. Update results
	cal.Evaluate(&t, &results)
	// check the scope allow list. Update results
	sal.Evaluate(&t, &results)

	//
	return results, nil
}

func wellFormed(t jira.ValidatedTicket, results EvaluationResults) EvaluationResults {
	for _, r := range t.ScopeRequests {
		if r.Scope != cleanScope(r.Scope) {
			evlog.Trace().Str("scope", r.Scope).Msg("ill-formed scope")
			results.addMessage(fmt.Sprintf("%s is not a well-formed scope", r.Scope))
			results.SetAction(ActionFail)

		}
	}

	return results
}

func init() {
	evlog = log.With().
		Str("component", "Evaluator").
		Logger()

}
