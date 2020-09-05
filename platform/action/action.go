package action

// Action is a operation on resource like (Read message, Write report, Write post)
type Action interface {
	ID() string
	Name() string
	Resource() string
	Operation() string
	Rename(name string) string
}

type action struct {
	id        string
	name      string
	resource  string
	operation string
}

func (a action) ID() string {
	return a.id
}

func (a action) Name() string {
	return a.name
}

func (a action) Resource() string {
	return a.resource
}

func (a action) Operation() string {
	return a.operation
}

func (a action) Rename(name string) action {
	a.name = name
	return a
}

func BuildMakeAction(uuid func() string) func(m map[string]string) action {
	return func(m map[string]string) action {
		return action{
			id:   uuid(),
			name: m["name"],
		}
	}
}

// func (a action) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]interface{}{
// 		"id":        a.ID(),
// 		"name":      a.Name(),
// 		"resource":  a.Resource(),
// 		"operation": a.Operation(),
// 	})
// }

// var (
// 	actions []action
// )

// func InitActions() {
// 	action1 := action{resource: "report", operation: "read", name: "hey"}
// 	actions = []action{action1}
// }

// func GetActions() []action {
// 	return actions
// }
