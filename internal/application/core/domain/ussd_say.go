package domain

type UssdSayStep struct {
	Step
	Text string
}
type UssdSayExt struct {
	Text string
}

func NewUssdSayStep(name string, text string) UssdSayStep {
	return UssdSayStep{
		Step: Step{
			Name: name,
			Kind: UssdCollect,
		},
		Text: text,
	}
}

func (u *UssdSayStep) Render(interpreter *Interpreter, containerModule string) (UssdSayExt, error) {
	ext := UssdSayExt{}
	ext.Text = interpreter.populateVariables(u.Text)
	return ext, nil
}
