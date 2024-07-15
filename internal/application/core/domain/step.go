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

type ExtStep struct {
}

func (s *Step) HandelAction(interpreter *Interpreter, targetModule Node) {

}

func (s *Step) Render(interpreter *Interpreter, name string) (*ExtStep, error) {
	return nil, nil
}

func (s *Step) Process(i *Interpreter) (string, error) {
	return "", nil
}
