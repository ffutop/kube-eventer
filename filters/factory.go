package filters

import (
	"fmt"
	"github.com/AliyunContainerService/kube-eventer/common/flags"
	"github.com/AliyunContainerService/kube-eventer/core"
	"github.com/AliyunContainerService/kube-eventer/filters/reason"
	"k8s.io/klog"
)

type FilterFactory struct {
}

func NewFilterFactory() *FilterFactory {
	return &FilterFactory{}
}

func (factory *FilterFactory) Build(uri flags.Uri) (core.EventFilter, error) {
	switch uri.Key {
	case "reason":
		return reason.CreateReasonFilter(&uri.Val)
	default:
		return nil, fmt.Errorf("Filter not recognized: %s", uri.Key)
	}
}

func (factory *FilterFactory) BuildAll(uris flags.Uris) []core.EventFilter {
	result := make([]core.EventFilter, 0, len(uris))
	for _, uri := range uris {
		filter, err := factory.Build(uri)
		if err != nil {
			klog.Errorf("Failed to create %v filter: %v", uri, err)
			continue
		}
		result = append(result, filter)
	}
	return result
}
