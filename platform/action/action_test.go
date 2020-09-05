package action

import "testing"

func TestRename(t *testing.T) {
	const newName = "Hello"
	const oldName = "Hi"
	data := map[string]string{"name": oldName}
	action := MakeAction(data)
	action = action.Rename(newName)
	if action.Name() != newName {
		t.Errorf("Action not renamed")
	}
}
