package infra

import (
	"log"
	"sync"
)

var userMap = sync.Map{}

type UserState struct {
	Height  int
	Width   int
	CenterX int
	CenterY int
	Scale   int
	Meta    interface{}
}

func StoreUserState(id int, state UserState) {
	if v, ok := userMap.Load(id); ok {
		//check state, update cache
		old := v.(UserState)
		UpdateUserCache(old.CenterX, old.CenterY, old.Height, old.Width, old.Scale, state.CenterX, state.CenterY, state.Height, state.Width, state.Scale)
	} else {
		//init cache
		InitUserCache(state.CenterX, state.CenterY, state.Height, state.Width, state.Scale)
		log.Printf("StoreUserState： user{%v} init, state {%v}", id, state)
	}
	userMap.Store(id, state)
}

func InitUserCache(x, y, h, w, s int) {

}

func UpdateUserCache(X0, y0, h0, w0, s0, x, y, h, w, s int) {

}
