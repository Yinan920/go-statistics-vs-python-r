// regression_test.go
//
// Unit tests and benchmarks for the Anscombe Quartet linear regression.
//
// All four Anscombe datasets are designed to yield nearly identical OLS
// coefficients: intercept ≈ 3.00, slope ≈ 0.50 (Anscombe 1973).
//
// Run unit tests:       go test -v
// Run benchmarks:       go test -bench=. -benchmem
// Run with coverage:    go test -v -cover

package main

import (
	"math"
	"testing"
)

// tolerance is the maximum acceptable deviation from the published
// reference coefficients for the Anscombe Quartet.
const tolerance = 0.001

// referenceCoefficients lists the known OLS intercepts and slopes for
// each Anscombe dataset, derived from Anscombe (1973).
var referenceCoefficients = []struct {
	name      string
	intercept float64
	slope     float64
}{
	{"Set I", 3.0001, 0.5001},
	{"Set II", 3.0009, 0.5000},
	{"Set III", 3.0025, 0.4997},
	{"Set IV", 3.0017, 0.4999},
}

// withinTolerance reports whether |got - want| <= tol.
func withinTolerance(got, want, tol float64) bool {
	return math.Abs(got-want) <= tol
}

// TestLinearRegressionAllSets verifies that the Go stats package produces
// slope and intercept values within tolerance of the published reference
// values for each of the four Anscombe datasets.
func TestLinearRegressionAllSets(t *testing.T) {
	datasets := anscombeDatasets()

	for i, ds := range datasets {
		ref := referenceCoefficients[i]

		t.Run(ds.name, func(t *testing.T) {
			result, err := linearRegression(ds.x, ds.y)
			if err != nil {
				t.Fatalf("linearRegression returned an unexpected error: %v", err)
			}

			if !withinTolerance(result.intercept, ref.intercept, tolerance) {
				t.Errorf("%s intercept: got %.4f, want %.4f (±%.3f)",
					ds.name, result.intercept, ref.intercept, tolerance)
			}

			if !withinTolerance(result.slope, ref.slope, tolerance) {
				t.Errorf("%s slope: got %.4f, want %.4f (±%.3f)",
					ds.name, result.slope, ref.slope, tolerance)
			}
		})
	}
}

// TestToSeries verifies that toSeries correctly converts x and y slices
// into a stats.Series with the expected length and coordinate values.
func TestToSeries(t *testing.T) {
	xs := []float64{1.0, 2.0, 3.0}
	ys := []float64{4.0, 5.0, 6.0}

	series := toSeries(xs, ys)

	if len(series) != len(xs) {
		t.Errorf("series length: got %d, want %d", len(series), len(xs))
	}
	for i, c := range series {
		if c.X != xs[i] || c.Y != ys[i] {
			t.Errorf("series[%d]: got (%.1f, %.1f), want (%.1f, %.1f)",
				i, c.X, c.Y, xs[i], ys[i])
		}
	}
}

// TestLinearRegressionKnownLine verifies regression on a perfect y = 2x + 1
// relationship, where the expected slope is exactly 2 and intercept is exactly 1.
func TestLinearRegressionKnownLine(t *testing.T) {
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{3, 5, 7, 9, 11} // y = 2x + 1

	result, err := linearRegression(xs, ys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !withinTolerance(result.slope, 2.0, tolerance) {
		t.Errorf("slope: got %.4f, want 2.0000", result.slope)
	}
	if !withinTolerance(result.intercept, 1.0, tolerance) {
		t.Errorf("intercept: got %.4f, want 1.0000", result.intercept)
	}
}

// BenchmarkLinearRegression measures execution time for running OLS regression
// across all four Anscombe datasets. The -benchmem flag adds memory allocation
// statistics to the output.
func BenchmarkLinearRegression(b *testing.B) {
	datasets := anscombeDatasets()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ds := range datasets {
			_, _ = linearRegression(ds.x, ds.y)
		}
	}
}
