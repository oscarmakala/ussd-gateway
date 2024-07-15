package domain

// StepKind represents the kind of step
type StepKind int

const (
	UssdCollect StepKind = iota
	UssdLanguage
	UssdSay
	ExternalService
)

type Step struct {
	Kind StepKind
	Name string
}

type StepExt struct {
}
