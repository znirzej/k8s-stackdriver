package translator

import (
	"time"

	stackdriver "google.golang.org/api/monitoring/v3"
	//	"k8s.io/component-base/metrics"
	//	"k8s.io/component-base/metrics/legacyregistry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	summaryRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stackdriver_api_request_duration",
			Help:    "The Stackdriver api request latencies in seconds.",
			Buckets: []float64{0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.25, 1.5, 1.75, 2, 2.5, 3, 3.5, 4, 4.5, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30, 40, 50, 60},
		},
		[]string{"node"},
	)
)

func Do(r *stackdriver.ProjectsTimeSeriesListCall, resource string) (*stackdriver.ListTimeSeriesResponse, error) {
	startTime := myClock.Now()
	defer func() {
		summaryRequestLatency.WithLabelValues(resource).Observe(float64(myClock.Since(startTime)) / float64(time.Second))
	}()
	return r.Do()
}

var myClock clock = &realClock{}