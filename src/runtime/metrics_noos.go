// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

var timeHistBuckets  []float64

type metricKind int

const (
	// These values must be kept identical to their corresponding Kind* values
	// in the runtime/metrics package.
	metricKindBad metricKind = iota
	metricKindUint64
	metricKindFloat64
	metricKindFloat64Histogram
)


type metricValue struct {
	kind    metricKind
	scalar  uint64         // contains scalar values for scalar Kinds.
	pointer unsafe.Pointer // contains non-scalar values.
}

type metricFloat64Histogram struct {
	counts  []uint64
	buckets []float64
}

func (v *metricValue) float64HistOrInit(buckets []float64) *metricFloat64Histogram {
	var hist *metricFloat64Histogram
	if v.kind == metricKindFloat64Histogram && v.pointer != nil {
		hist = (*metricFloat64Histogram)(v.pointer)
	} else {
		v.kind = metricKindFloat64Histogram
		hist = new(metricFloat64Histogram)
		v.pointer = unsafe.Pointer(hist)
	}
	hist.buckets = buckets
	if len(hist.counts) != len(hist.buckets)-1 {
		hist.counts = make([]uint64, len(buckets)-1)
	}
	return hist
}

//go:linkname godebug_registerMetric internal/godebug.registerMetric
func godebug_registerMetric(name string, read func() uint64) {}