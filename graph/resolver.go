package graph

import (
	"sync"

	"gqlgen-subscription-sample/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ChannelsByMatID map[int64][]chan<- *model.SmartMat
	Mutex           sync.Mutex
}
