package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	InjectivePeggoPendingBatchSize = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "injective_peggo_pending_batch_size",
			Help: "Pending batch size",
		},
		[]string{"chain", "address"},
	)

	InjectivePeggoPendingValsetSize = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "injective_peggo_pending_valset_size",
			Help: "Pending valset size",
		},
		[]string{"chain", "address"},
	)

	InjectivePeggoLastEventNonce = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "injective_peggo_last_event_nonce",
			Help: "Last event nonce",
		},
		[]string{"chain", "address"},
	)

	InjectivePeggoLastObservedNonce = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "injective_peggo_last_observed_nonce",
			Help: "Last observed nonce",
		},
		[]string{"chain", "address"},
	)

	chain   string
	address string
)

func Initialize(c, a string) {
	chain = c
	address = a
}

func SetInjectivePeggoPendingBatchSize(batchSize int) {
	InjectivePeggoPendingBatchSize.With(prometheus.Labels{"chain": chain, "address": address}).Add(float64(batchSize))
}

func SetInjectivePeggoPendingValsetSize(valsetSize int) {
	InjectivePeggoPendingValsetSize.With(prometheus.Labels{"chain": chain, "address": address}).Add(float64(valsetSize))
}

func SetInjectivePeggoLastEventNonce(eventNonce uint64) {
	InjectivePeggoLastEventNonce.With(prometheus.Labels{"chain": chain, "address": address}).Set(float64(eventNonce))
}

func SetInjectivePeggoLastObservedNonce(observedNonce uint64) {
	InjectivePeggoLastObservedNonce.With(prometheus.Labels{"chain": chain, "address": address}).Set(float64(observedNonce))
}
