// Code generated by protoc-gen-goten-resource
// Resource: TimeSerie
// DO NOT EDIT!!!

package time_serie

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	bucket "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/bucket"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &bucket.Bucket{}
	_ = &common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type TimeSerieList []*TimeSerie

func (l TimeSerieList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*TimeSerie))
}

func (l TimeSerieList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(TimeSerieList)...)
}

func (l TimeSerieList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSerieList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l TimeSerieList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*TimeSerie)
}

func (l TimeSerieList) Length() int {
	return len(l)
}
