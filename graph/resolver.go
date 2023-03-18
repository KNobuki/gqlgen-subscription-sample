package graph

import (
	"gqlgen-subscription-sample/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	channelsByMatID map[int64][]chan<- *model.SmartMat
	mutex           sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		channelsByMatID: make(map[int64][]chan<- *model.SmartMat),
		mutex:           sync.Mutex{},
	}
}
