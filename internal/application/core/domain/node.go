package domain

import "errors"

type NodeName struct {
	Name  string
	Label string
}

type Node struct {
	Name  string
	Steps []Step
}

func NewNode(name string) Node {
	return Node{
		Name:  name,
		Steps: []Step{},
	}
}

func (n *Node) getStepByName(stepName string) (*Step, error) {
	if stepName == "" {
		return nil, errors.New(`stepName shouldn't be null`)
	}
	for _, step := range n.Steps {
		if stepName == step.Name {
			return &step, nil
		}
	}
	return nil, nil
}

func (n *Node) getStepNames() []string {
	var names = make([]string, 0)
	for _, step := range n.Steps {
		names = append(names, step.Name)
	}
	return names
}
