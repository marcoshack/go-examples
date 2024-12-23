package timeutil

import (
	"sort"
	"time"

	"github.com/rs/zerolog"
)

type LatencyStats struct {
	durations                           []time.Duration
	Min, Max, Avg, P95, P99, TM95, TM99 time.Duration
	Size                                int
}

func NewLatencyStats(durations []time.Duration) *LatencyStats {
	size := len(durations)
	if size == 0 {
		return &LatencyStats{}
	}
	stats := LatencyStats{
		durations: durations,
		Size:      size,
	}

	var totalLatency time.Duration
	for _, latency := range durations {
		totalLatency += latency
	}
	stats.Avg = totalLatency / time.Duration(size)

	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})
	stats.Min = durations[0]
	stats.Max = durations[size-1]
	stats.P95 = stats.P(95)
	stats.P99 = stats.P(99)
	stats.TM95 = stats.TM(95)
	stats.TM99 = stats.TM(99)
	return &stats
}

func (s *LatencyStats) TM(percentile int) time.Duration {
	if s.Size == 0 {
		return 0
	}
	if s.Size == 1 {
		return s.durations[0]
	}
	truncatedLatencies := s.durations[:int(float64(s.Size)*float64(percentile)/100)]
	var totalTruncatedLatency time.Duration
	for _, latency := range truncatedLatencies {
		totalTruncatedLatency += latency
	}
	return totalTruncatedLatency / time.Duration(len(truncatedLatencies))
}

func (s *LatencyStats) P(percentile int) time.Duration {
	return s.durations[int(float64(s.Size)*float64(percentile)/100)]
}

func (s *LatencyStats) MarshalZerologObject(e *zerolog.Event) {
	e.Int("size", s.Size).
		Str("min", s.Min.String()).
		Str("max", s.Max.String()).
		Str("avg", s.Avg.String()).
		Str("p95", s.P95.String()).
		Str("p99", s.P99.String()).
		Str("tm95", s.TM95.String()).
		Str("tm99", s.TM99.String())
}
