package reason

import (
	"fmt"
	"github.com/AliyunContainerService/kube-eventer/core"
	kube_api "k8s.io/api/core/v1"
	"net/url"
)

type ReasonFilter struct {
	Includes []string `json:"includes":["aaa", "bbb", "ccc"]`
	Excludes []string `json:"excludes":["xxx", "yyy", "zzz"]`

	includesMap map[string]struct{}
	excludesMap map[string]struct{}
}

func CreateReasonFilter(url *url.URL) (*ReasonFilter, error) {
	rf := &ReasonFilter{
		Includes: nil,
		Excludes: nil,
	}
	opt := url.Query()
	if opt["includes"] != nil {
		includes := opt["includes"]
		rf.Includes = includes
		for _, include := range includes {
			rf.includesMap[include] = struct{}{}
		}
	}

	if opt["excludes"] != nil {
		excludes := opt["excludes"]
		rf.Excludes = excludes
		for _, exclude := range excludes {
			rf.excludesMap[exclude] = struct{}{}
		}
	}

	if rf.Includes == nil && rf.Excludes == nil {
		return nil, fmt.Errorf("at least one parameter \"includes\" or \"excludes\" is needed.")
	}

	if rf.Includes != nil && rf.Excludes != nil {
		return nil, fmt.Errorf("only one parameter is allowed at the same time.")
	}

	return rf, nil
}

func (filter *ReasonFilter) Name() string {
	return "ReasonFilter"
}

func (filter *ReasonFilter) Filter(batch *core.EventBatch) *core.EventBatch {
	// only one list (includes / excludes) will work
	newEvents := make([]*kube_api.Event, len(batch.Events))
	if len(filter.Includes) != 0 {
		for _, event := range batch.Events {
			_, ok := filter.includesMap[event.Reason]
			if ok == true {
				newEvents = append(newEvents, event)
			}
		}
		return &core.EventBatch{
			Timestamp: batch.Timestamp,
			Events:    newEvents,
		}
	} else if len(filter.Excludes) != 0 {
		for _, event := range batch.Events {
			_, ok := filter.excludesMap[event.Reason]
			if ok == false {
				newEvents = append(newEvents, event)
			}
		}
		return &core.EventBatch{
			Timestamp: batch.Timestamp,
			Events:    newEvents,
		}
	}
	// can not be empty at the same time
	return nil
}
