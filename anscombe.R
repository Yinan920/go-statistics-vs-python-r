# anscombe.R
#
# The Anscombe Quartet — Linear Regression with Execution Timing (R)
#
# Uses lm() to fit simple linear regression for each of the four Anscombe
# datasets and reports coefficients alongside execution time and memory.
#
# Reference:
#   Anscombe, F. J. 1973. Graphs in statistical analysis.
#   The American Statistician 27: 17-21.
#
# Run: Rscript anscombe.R

# ---------------------------------------------------------------------------
# Data
# ---------------------------------------------------------------------------

anscombe <- data.frame(
  x1 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x2 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x3 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x4 = c(8,  8,  8, 8,  8,  8, 8, 19,  8, 8, 8),
  y1 = c(8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24,  4.26, 10.84, 4.82, 5.68),
  y2 = c(9.14, 8.14,  8.74, 8.77, 9.26, 8.10, 6.13,  3.10,  9.13, 7.26, 4.74),
  y3 = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08,  5.39,  8.15, 6.42, 5.73),
  y4 = c(6.58, 5.76,  7.71, 8.84, 8.47, 7.04, 5.25, 12.50,  5.56, 7.91, 6.89)
)

# ---------------------------------------------------------------------------
# Helper
# ---------------------------------------------------------------------------

# fit_and_print fits an OLS model and prints the label, intercept, and slope.
fit_and_print <- function(label, x_col, y_col) {
  model <- lm(anscombe[[y_col]] ~ anscombe[[x_col]])
  coefs <- coef(model)
  cat(sprintf("  %-10s  Intercept: %.4f  Slope: %.4f\n",
              label, coefs[[1]], coefs[[2]]))
}

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

cat("============================================\n")
cat("  Anscombe Quartet: R Linear Regression\n")
cat("  Function: lm()\n")
cat("============================================\n")
cat(sprintf("  %-10s  %-12s  %-10s\n", "Dataset", "Intercept", "Slope"))
cat("  ", strrep("-", 38), "\n", sep = "")

mem_before <- gc(verbose = FALSE)
start_time <- proc.time()

fit_and_print("Set I",   "x1", "y1")
fit_and_print("Set II",  "x2", "y2")
fit_and_print("Set III", "x3", "y3")
fit_and_print("Set IV",  "x4", "y4")

elapsed  <- proc.time() - start_time
mem_after <- gc(verbose = FALSE)

cat("  ", strrep("-", 38), "\n", sep = "")
cat(sprintf("\nExecution time : %.4f ms\n", elapsed["elapsed"] * 1000))
cat(sprintf("Memory used    : %.1f MB\n",
            sum(mem_after[, "used"] - mem_before[, "used"]) / 1024))
