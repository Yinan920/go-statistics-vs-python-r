# anscombe.py
#
# The Anscombe Quartet — Linear Regression with Execution Timing (Python)
#
# Uses statsmodels OLS to fit simple linear regression for each of the
# four Anscombe datasets and reports coefficients alongside execution time.
#
# Reference:
#   Anscombe, F. J. 1973. Graphs in statistical analysis.
#   The American Statistician 27: 17–21.
#
# Run: python anscombe.py

from __future__ import division, print_function
import time
import tracemalloc

import pandas as pd
import statsmodels.api as sm

# ---------------------------------------------------------------------------
# Data
# ---------------------------------------------------------------------------

anscombe = pd.DataFrame({
    'x1': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x2': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x3': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x4': [8,  8,  8, 8,  8,  8, 8, 19,  8, 8, 8],
    'y1': [8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24,  4.26, 10.84, 4.82, 5.68],
    'y2': [9.14, 8.14,  8.74, 8.77, 9.26, 8.10, 6.13,  3.10,  9.13, 7.26, 4.74],
    'y3': [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08,  5.39,  8.15, 6.42, 5.73],
    'y4': [6.58, 5.76,  7.71, 8.84, 8.47, 7.04, 5.25, 12.50,  5.56, 7.91, 6.89],
})

# ---------------------------------------------------------------------------
# Helper
# ---------------------------------------------------------------------------

def fit_and_print(label: str, x_col: str, y_col: str) -> None:
    """Fit OLS regression and print the intercept and slope."""
    design = sm.add_constant(anscombe[x_col])
    model = sm.OLS(anscombe[y_col], design).fit()
    intercept, slope = model.params
    print(f"  {label:<10}  Intercept: {intercept:.4f}  Slope: {slope:.4f}")

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

print("============================================")
print("  Anscombe Quartet: Python Linear Regression")
print("  Package: statsmodels OLS")
print("============================================")
print(f"  {'Dataset':<10}  {'Intercept':<12}  {'Slope':<10}")
print("  " + "-" * 38)

tracemalloc.start()
start = time.perf_counter()

fit_and_print("Set I",   'x1', 'y1')
fit_and_print("Set II",  'x2', 'y2')
fit_and_print("Set III", 'x3', 'y3')
fit_and_print("Set IV",  'x4', 'y4')

elapsed = time.perf_counter() - start
_, peak_memory = tracemalloc.get_traced_memory()
tracemalloc.stop()

print("  " + "-" * 38)
print(f"\nExecution time : {elapsed * 1000:.4f} ms")
print(f"Peak memory    : {peak_memory} bytes")
