package evaluator

type ScopeDefinition struct {
	ArbitraryAccess string `json:"arbitraryAccess"`
	Function        string `json:"fn"`
	Scope           string `json:"sc"`
	Service         string `json:"serviceName"`
	ServicePath     string `json:"servicePath"`
}
