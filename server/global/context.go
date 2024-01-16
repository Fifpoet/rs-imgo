package global

import "sync"

var UserContext = sync.Map{}

type UserState struct {
	Height  int
	Width   int
	CenterX int
	CenterY int
	Meta    interface{}
}

func (c *UserState) StoreUserState(id int, state UserState) {
	UserContext.LoadOrStore(id, state)
}

func (c *UserState) GetUserState() {

}
