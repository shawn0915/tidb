// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"github.com/pingcap/tidb/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// planner core metrics vars
var (
	PlanCacheCounter     prometheus.Counter
	PlanCacheMissCounter prometheus.Counter

	PseudoEstimationNotAvailable prometheus.Counter
	PseudoEstimationOutdate      prometheus.Counter
)

func init() {
	InitMetricsVars()
}

// InitMetricsVars init planner core metrics vars.
func InitMetricsVars() {
	PlanCacheCounter = metrics.PlanCacheCounter.WithLabelValues("prepare")
	PlanCacheMissCounter = metrics.PlanCacheMissCounter.WithLabelValues("cache_miss")
	PseudoEstimationNotAvailable = metrics.PseudoEstimation.WithLabelValues("nodata")
	PseudoEstimationOutdate = metrics.PseudoEstimation.WithLabelValues("outdate")
}
