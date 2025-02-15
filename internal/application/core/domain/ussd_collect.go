package domain

import (
	"fmt"
)

type UssdCollectStep struct {
	Step
	Messages []UssdSayStep
}

type UssdCollectExt struct {
	Action   map[string]string
	Messages []UssdSayExt
}

func NewUssdCollectStep(name string) UssdCollectStep {
	return UssdCollectStep{
		Step: Step{
			Name: name,
			Kind: UssdCollect,
		},
		Messages: make([]UssdSayStep, 0),
	}
}

func (u *UssdCollectStep) Render(interpreter *Interpreter, containerModule string) (UssdCollectExt, error) {
	ext := UssdCollectExt{
		Messages: make([]UssdSayExt, 0),
	}

	newTarget := fmt.Sprintf("%s.%s.handle", containerModule, u.Name)

	pairs := make(map[string]string)
	pairs["target"] = newTarget

	interpreter.buildAction(pairs)
	ext.Action = pairs

	for _, message := range u.Messages {
		messageExt, _ := message.Render(interpreter, containerModule)
		ext.Messages = append(ext.Messages, messageExt)
	}
	return ext, nil
}
