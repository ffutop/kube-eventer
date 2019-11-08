package filters

import "github.com/AliyunContainerService/kube-eventer/core"

type FilterManager struct {
	filters []*core.EventFilter
}

func (manager *FilterManager) Filter(batch *core.EventBatch) *core.EventBatch {
	for _, filter := range manager.filters {
	}
}
