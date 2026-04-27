# Anscombe Quartet — Linear Regression in Go, Python, and R

**MSDS 431 · Week 4 Assignment**

This repository compares OLS linear regression results and execution performance across **Go**, **Python**, and **R** using the four [Anscombe Quartet](https://en.wikipedia.org/wiki/Anscombe%27s_quartet) datasets (Anscombe 1973). The goal is to verify that the Go [`montanaflynn/stats`](https://github.com/montanaflynn/stats) package produces coefficients consistent with established statistical tools.

---

## Repository Structure

```
anscombe-quartet/
├── main_Statistics.go   # Go program: OLS regression + timing + memory
├── regression_test.go   # Go unit tests and benchmarks
├── go.mod               # Go module definition
├── go.sum               # Go dependency checksums (auto-generated)
├── anscombe.py          # Python program: statsmodels OLS + timing
├── anscombe.R           # R program: lm() OLS + timing
└── README.md            # This file
```

---

## Prerequisites

| Tool | Version | Install |
|------|---------|---------|
| Go   | ≥ 1.21  | https://go.dev/dl/ |
| Python | ≥ 3.9 | https://python.org |
| R    | ≥ 4.0   | https://cran.r-project.org |

Python packages required:
```bash
pip install pandas statsmodels
```

---

## Running the Go Program

```bash
# 1. Clone the repository
git clone https://github.com/<your-username>/anscombe-quartet.git
cd anscombe-quartet

# 2. Download Go dependencies
go mod tidy

# 3. Run the program
go run main.go

# 4. Build a standalone executable (macOS)
go build -o anscombe
./anscombe

# 4b. Build for Windows
GOOS=windows go build -o anscombe.exe
```

**Expected output:**
```
============================================
  Anscombe Quartet: Go Linear Regression
  Package: github.com/montanaflynn/stats
============================================
Dataset     Intercept     Slope
--------------------------------------------
Set I       3.0001        0.5001
Set II      3.0009        0.5000
Set III     3.0025        0.4997
Set IV      3.0017        0.4999
--------------------------------------------

Execution time : <measured value>
Memory used    : <measured value> bytes
```

---

## Running Go Unit Tests and Benchmarks

```bash
# Run all unit tests (verbose output)
go test -v

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem

# Run tests with coverage report
go test -v -cover
```

Unit tests verify that slope and intercept for each dataset are within **±0.001** of the published reference values from Anscombe (1973). The benchmark (`BenchmarkLinearRegression`) reports nanoseconds per operation and bytes allocated per operation.

---

## Running the Python Program

```bash
python anscombe.py
```

**Expected output:**
```
============================================
  Anscombe Quartet: Python Linear Regression
  Package: statsmodels OLS
============================================
  Dataset     Intercept     Slope
  --------------------------------------
  Set I       3.0001        0.5001
  Set II      3.0009        0.5000
  Set III     3.0025        0.4997
  Set IV      3.0017        0.4999
  --------------------------------------

Execution time : <measured value> ms
Peak memory    : <measured value> bytes
```

---

## Running the R Program

```bash
Rscript anscombe.R
```

**Expected output:**
```
============================================
  Anscombe Quartet: R Linear Regression
  Function: lm()
============================================
  Dataset     Intercept     Slope
  --------------------------------------
  Set I       3.0001        0.5001
  Set II      3.0009        0.5000
  Set III     3.0025        0.4997
  Set IV      3.0017        0.4999
  --------------------------------------

Execution time : <measured value> ms
Memory used    : <measured value> MB
```

---

## Regression Results — Comparison Table

All three languages produce virtually identical OLS coefficients, confirming the `montanaflynn/stats` package is numerically accurate.

| Dataset  | Language | Intercept | Slope  |
|----------|----------|-----------|--------|
| Set I    | Go       | 3.0001    | 0.5001 |
| Set I    | Python   | 3.0001    | 0.5001 |
| Set I    | R        | 3.0001    | 0.5001 |
| Set II   | Go       | 3.0009    | 0.5000 |
| Set II   | Python   | 3.0009    | 0.5000 |
| Set II   | R        | 3.0009    | 0.5000 |
| Set III  | Go       | 3.0025    | 0.4997 |
| Set III  | Python   | 3.0025    | 0.4997 |
| Set III  | R        | 3.0025    | 0.4997 |
| Set IV   | Go       | 3.0017    | 0.4999 |
| Set IV   | Python   | 3.0017    | 0.4999 |
| Set IV   | R        | 3.0017    | 0.4999 |

---

## Performance Comparison

| Language | Package / Function    | Execution Time | Peak Memory  |
|----------|-----------------------|----------------|--------------|
| Go       | `montanaflynn/stats`  | 31.334 µs      | 3,456 bytes  |
| Python   | `statsmodels` OLS     | 8.4924 ms      | 60,617 bytes |
| R        | `lm()`                | 8.0000 ms      | 210.3 MB     |

Go benchmark (`go test -bench=. -benchmem`, Apple M4):
BenchmarkLinearRegression-10: 2,785,522 iterations · 425.6 ns/op · 2,496 B/op · 16 allocs/op
---

## Recommendation to Management

### Summary

The Go `montanaflynn/stats` package produces OLS regression coefficients that are numerically identical to those from Python (`statsmodels`) and R (`lm()`), to four decimal places. From a **correctness standpoint**, Go is a viable alternative for linear regression.

### Data Scientists' Concerns

While Go can perform basic statistical computations correctly, data scientists have the following concerns about adopting it as their primary tool:

1. **Limited statistical depth.** The `montanaflynn/stats` package covers only a narrow set of statistical methods (descriptive statistics, basic regression). Python and R offer mature ecosystems — `statsmodels`, `scikit-learn`, `scipy`, `tidymodels`, `caret` — with thousands of models, diagnostics, and visualizations.

2. **No native data visualization.** R (`ggplot2`) and Python (`matplotlib`, `seaborn`) have rich, production-ready plotting libraries. Go has no comparable ecosystem for statistical graphics, making exploratory data analysis far more difficult.

3. **No interactive computing environment.** Jupyter notebooks (Python) and RStudio (R) provide the interactive, literate programming workflows that data scientists rely on for exploration and communication. Go lacks an equivalent.

4. **Smaller community for statistics.** The number of contributors and users for `montanaflynn/stats` (~28 contributors, ~15,000 users as of mid-2023) is orders of magnitude smaller than NumPy, pandas, or R's base packages, meaning fewer community resources, tutorials, and vetted implementations.

5. **Productivity cost of transition.** Retraining data scientists in Go would involve significant upfront investment with uncertain return. Python and R are the industry standard for data science for strong reasons.

### Recommendation

We recommend a **hybrid approach**:
- Use **Go** for backend web servers, database layers, and distributed cloud services — areas where Go excels.
- Retain **Python and R** for statistical analysis, machine learning, and data visualization — the domains where their ecosystems are unmatched.

This approach lets software engineers and data scientists each use the language best suited to their work, while sharing data and results through APIs and file formats (JSON, CSV, Parquet). Forcing a single language across both disciplines is likely to reduce productivity and increase risk without meaningful benefit.

---

## References

- Anscombe, F. J. 1973. Graphs in statistical analysis. *The American Statistician* 27: 17–21.
- Miller, T. W. 2015. *Modeling Techniques in Predictive Analytics*. Pearson FT Press.
- Bates, K., & LaNou, J. 2023. *Learning Go*. O'Reilly.
- Bodner, J. 2024. *Learning Go*, 2nd ed. O'Reilly.
- Montana Flynn. `stats` package. https://github.com/montanaflynn/stats
