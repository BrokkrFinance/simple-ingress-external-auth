package prometheus_test

import (
	"context"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"

	metricsprometheus "github.com/slok/simple-ingress-external-auth/internal/metrics/prometheus"
)

func TestRecorder(t *testing.T) {
	tests := map[string]struct {
		measure    func(r metricsprometheus.Recorder)
		expMetrics string
	}{
		"Measure token reviews.": {
			measure: func(r metricsprometheus.Recorder) {
				r.TokenReview(context.TODO(), true, true, "")
				r.TokenReview(context.TODO(), true, true, "")
				r.TokenReview(context.TODO(), false, false, "")
				r.TokenReview(context.TODO(), true, false, "something")
				r.TokenReview(context.TODO(), true, false, "otherthing")
				r.TokenReview(context.TODO(), true, false, "otherthing")
			},
			expMetrics: `
				# HELP simple_ingress_external_auth_token_reviews_total The number of token reviews.
				# TYPE simple_ingress_external_auth_token_reviews_total counter
				simple_ingress_external_auth_token_reviews_total{invalid_reason="",success="false",valid="false"} 1
				simple_ingress_external_auth_token_reviews_total{invalid_reason="",success="true",valid="true"} 2
				simple_ingress_external_auth_token_reviews_total{invalid_reason="otherthing",success="true",valid="false"} 2
				simple_ingress_external_auth_token_reviews_total{invalid_reason="something",success="true",valid="false"} 1
			`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			reg := prometheus.NewRegistry()
			rec := metricsprometheus.NewRecorder(reg)

			test.measure(rec)

			// Check metrics.
			err := testutil.GatherAndCompare(reg, strings.NewReader(test.expMetrics))
			assert.NoError(err)
		})
	}
}
