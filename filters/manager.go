package filters

import "github.com/AliyunContainerService/kube-eventer/core"

type FilterManager struct {
	filters []core.EventFilter
}

func NewFilterManager(filters []core.EventFilter) (core.EventFilter, error) {
	filterManager := &FilterManager{filters: filters}
	return filterManager, nil
}

func (manager *FilterManager) Name() string {
	return "Filter Manager"
}

func (manager *FilterManager) Filter(batch *core.EventBatch) *core.EventBatch {
	for _, filter := range manager.filters {
		batch = filter.Filter(batch)
	}
	return batch
}
