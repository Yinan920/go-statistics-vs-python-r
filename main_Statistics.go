// main.go
//
// Anscombe Quartet: Linear Regression in Go
//
// Uses the montanaflynn/stats package to compute OLS linear regression
// for each of the four Anscombe Quartet datasets, prints estimated
// intercept and slope, and reports execution time and memory usage.
//
// Reference:
//   Anscombe, F. J. 1973. Graphs in statistical analysis.
//   The American Statistician 27: 17–21.

package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/montanaflynn/stats"
)

// dataset holds one Anscombe Quartet data set: a label, x values, and y values.
type dataset struct {
	name string
	x    []float64
	y    []float64
}

// regressionResult stores the OLS slope and intercept for a single dataset.
type regressionResult struct {
	slope     float64
	intercept float64
}

// anscombeDatasets returns the four Anscombe Quartet datasets.
// Sets I–III share the same x values; Set IV uses a different x.
func anscombeDatasets() []dataset {
	sharedX := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	setIVx := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}

	return []dataset{
		{
			name: "Set I",
			x:    sharedX,
			y:    []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{
			name: "Set II",
			x:    sharedX,
			y:    []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
		},
		{
			name: "Set III",
			x:    sharedX,
			y:    []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{
			name: "Set IV",
			x:    setIVx,
			y:    []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89},
		},
	}
}

// toSeries converts parallel x and y slices into a stats.Series ([]stats.Coordinate).
func toSeries(xs, ys []float64) stats.Series {
	series := make(stats.Series, len(xs))
	for i := range xs {
		series[i] = stats.Coordinate{X: xs[i], Y: ys[i]}
	}
	return series
}

func linearRegression(xs, ys []float64) (regressionResult, error) {
	fitted, err := stats.LinearRegression(toSeries(xs, ys))
	if err != nil {
		return regressionResult{}, err
	}

	// Find the coordinates with the smallest and largest x values.
	// We cannot rely on order since the library preserves input order.
	minC, maxC := fitted[0], fitted[0]
	for _, c := range fitted {
		if c.X < minC.X {
			minC = c
		}
		if c.X > maxC.X {
			maxC = c
		}
	}

	if maxC.X == minC.X {
		return regressionResult{}, fmt.Errorf("all x values are identical; slope undefined")
	}

	slope := (maxC.Y - minC.Y) / (maxC.X - minC.X)
	intercept := minC.Y - slope*minC.X

	return regressionResult{slope: slope, intercept: intercept}, nil
}

// memoryUsed returns current heap memory allocated in bytes.
func memoryUsed() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	fmt.Println("============================================")
	fmt.Println("  Anscombe Quartet: Go Linear Regression")
	fmt.Println("  Package: github.com/montanaflynn/stats")
	fmt.Println("============================================")
	fmt.Printf("%-10s  %-12s  %-10s\n", "Dataset", "Intercept", "Slope")
	fmt.Println("--------------------------------------------")

	memBefore := memoryUsed()
	start := time.Now()

	for _, ds := range anscombeDatasets() {
		result, err := linearRegression(ds.x, ds.y)
		if err != nil {
			fmt.Printf("%-10s  error: %v\n", ds.name, err)
			continue
		}
		fmt.Printf("%-10s  %-12.4f  %-10.4f\n",
			ds.name, result.intercept, result.slope)
	}

	elapsed := time.Since(start)
	memAfter := memoryUsed()

	fmt.Println("--------------------------------------------")
	fmt.Printf("\nExecution time : %v\n", elapsed)
	fmt.Printf("Memory used    : %d bytes\n", memAfter-memBefore)
}
