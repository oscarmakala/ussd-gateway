package domain

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"sort"
	"strings"
)

const (
	StickyPrefix = "sticky_"
	ModulePrefix = "module_"
)

type ExtResponse struct {
}

type ProjectIndex struct {
	DefaultTarget string
	NodeNames     []NodeName
}

type Interpreter struct {
	Variables      map[string]string
	NodeNames      []NodeName
	TargetParam    string
	AppResult      ExtResponse
	ProjectOptions ProjectIndex
}

type VariableInText struct {
	VariableName string
	Position     int
}

type Target struct {
	NodeName string
	StepName string
	Action   string
}

func NewInterpreter(requestParams map[string]string, projectOptions ProjectIndex) *Interpreter {
	return &Interpreter{
		Variables:      processRequestParameters(requestParams),
		ProjectOptions: projectOptions,
	}
}

func processRequestParameters(requestParams map[string]string) map[string]string {
	variables := make(map[string]string)
	for key, value := range requestParams {
		if strings.HasPrefix(key, StickyPrefix) || strings.HasPrefix(key, ModulePrefix) {
			variables[fmt.Sprintf("%s%s", StickyPrefix, key)] = value
		} else {
			variables[key] = value
		}
	}
	return variables
}

func (i *Interpreter) Interpret() ExtResponse {
	i.NodeNames = i.ProjectOptions.NodeNames
	if i.TargetParam == "" {
		i.TargetParam = i.ProjectOptions.DefaultTarget
	}
	i.dispatch(i.TargetParam)
	return i.AppResult
}

func (i *Interpreter) dispatch(targetParam string) {
	target := i.parseTarget(targetParam)
	log.Infof("dispatch: %v", target)

	if target.Action != nil {

	}
}

func (i *Interpreter) parseTarget(targetParam string) Target {
	target := Target{}
	pattern := regexp.MustCompile(`^([^.]+)(.([^.]+))?(.([^.]+))?`)
	matcher := pattern.FindStringSubmatch(targetParam)
	if len(matcher) >= 1 {
		target.NodeName = matcher[1]
	}
	if len(matcher) >= 3 {
		target.StepName = matcher[3]
	}
	if len(matcher) >= 5 {
		target.Action = matcher[5]
	}
	return target
}
func (i *Interpreter) interpret(module Node, startingStepName string, prependStep string, originModule Node) {
}
func (i *Interpreter) populateVariables(sourceText string) string {
	if sourceText == "" {
		return ""
	}

	pattern := regexp.MustCompile(`\$([A-Za-z]+[A-Za-z0-9_]*)`)
	matches := pattern.FindAllStringSubmatchIndex(sourceText, -1)

	var variablesInText []VariableInText
	for _, match := range matches {
		variableName := sourceText[match[2]:match[3]]
		position := match[0]
		variablesInText = append(variablesInText, VariableInText{variableName, position})
	}

	// Sort variablesInText in reverse order by position
	sort.Slice(variablesInText, func(i, j int) bool {
		return variablesInText[i].Position > variablesInText[j].Position
	})

	var buffer strings.Builder
	buffer.WriteString(sourceText)
	for _, v := range variablesInText {
		replaceValue := ""
		if val, ok := i.Variables[v.VariableName]; ok {
			replaceValue = val
		} else if _, ok := i.Variables[ModulePrefix+v.VariableName]; ok {
			replaceValue = val
		} else if _, ok := i.Variables[StickyPrefix+v.VariableName]; ok {
			replaceValue = val
		}

		startPos := v.Position
		endPos := v.Position + len(v.VariableName) + 1 // +1 is for the $ character
		bufferStr := buffer.String()
		buffer.Reset()
		buffer.WriteString(bufferStr[:startPos])
		buffer.WriteString(replaceValue)
		buffer.WriteString(bufferStr[endPos:])
	}
	return buffer.String()
}

func (i *Interpreter) buildAction(pairs map[string]string) {
	// append sticky parameters and module-scoped variables
	for key := range i.Variables {
		if strings.HasPrefix(key, StickyPrefix) || strings.HasPrefix(key, ModulePrefix) {
			pairs[key] = i.Variables[key]
		}
	}
}
