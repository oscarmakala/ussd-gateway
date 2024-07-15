package domain

//func TestUssdSay(t *testing.T) {
//	interpreter := NewInterpreter(make(map[string]string))
//	ussdSayStep := domain.NewUssdSayStep("Menu", "Test")
//	ussdSayExt := ussdSayStep.render(*interpreter)
//	log.Println(ussdSayExt)
//	assert.EqualValues(t, ussdSayStep.Text, ussdSayExt.Text)
//}
//
//func TestUssdCollect(t *testing.T) {
//	interpreter := NewInterpreter(make(map[string]string))
//	step := domain.NewUssdCollectStep("Menu")
//	stepExt := step.render(*interpreter, "Menu01")
//	log.Println(stepExt)
//	assert.EqualValues(t, "Menu01.Menu.handle", stepExt.Action["target"])
//}
//
//func TestParsingTarget(t *testing.T) {
//	interpreter := NewInterpreter(make(map[string]string))
//	interpreter.dispatch("Menu01.Menu.handle")
//}
//
//func TestPopulatingVariables(t *testing.T) {
//	requestParams := make(map[string]string)
//	requestParams["menu"] = "testing"
//
//	interpreter := NewInterpreter(requestParams)
//	ussdSayStep := domain.NewUssdSayStep("Menu", "My name is $menu")
//
//	result := ussdSayStep.render(*interpreter)
//	assert.EqualValues(t, "My name is testing", result.Text)
//}
