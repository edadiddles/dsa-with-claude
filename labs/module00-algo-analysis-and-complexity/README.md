# Module 0: Foundations & Analysis

**Duration**: 2 weeks  
**Difficulty**: Medium to Hard

## Overview

You can't evaluate solutions without being able to analyze them rigorously. This module builds the analytical foundation you'll use throughout the entire curriculum.

**Why this matters**: Understanding how to measure, analyze, and compare algorithms is fundamental to everything that follows. You'll build tools that you'll use in every subsequent module.

## Learning Objectives

- Master mathematical foundations for algorithm analysis
- Build empirical performance analysis tools
- Understand the gap between theoretical and practical complexity
- Develop intuition for when theory diverges from practice

## Topics Covered

- Asymptotic notation (O, Ω, Θ)
- Recurrence relations and solving techniques
- Empirical performance measurement
- Cache effects and real hardware behavior
- Statistical analysis of algorithm performance

---

## Lab 0.1: Build Your Analysis Framework

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Create reusable timing and profiling infrastructure
- Understand measurement variance and statistical significance
- Build visualization tools for performance data

### Requirements

1. **Implement a Timer class/module** that:
   - Measures wall-clock time, CPU time, and user/system time
   - Supports multiple timing runs with statistical aggregation (mean, median, std dev)
   - Handles warmup runs to account for JIT compilation/caching
   - Exports results to CSV/JSON for further analysis

2. **Create a Profiler wrapper** that:
   - Tracks function calls and execution time
   - Measures memory allocation (language-specific)
   - Counts basic operations (comparisons, assignments, etc.)
   - Generates call graphs or flame graphs

3. **Build visualization tools**:
   - Plot execution time vs. input size
   - Generate comparison charts for multiple algorithms
   - Create interactive notebooks/scripts for exploration

### Implementation Details

```python
timer.measure(function, input_sizes, runs=10, warmup=3)
# Returns: {size: {mean, median, std_dev, min, max}}

profiler.profile(function, *args)
# Returns: {time, memory, operations, call_tree}
```

### Test Cases
- Time simple operations (array access, hash lookup) across input sizes
- Verify statistical properties (std dev should decrease with more runs)
- Compare timer accuracy against known benchmarks

### Deliverables
- `timer.{py,go,zig}` - Timing infrastructure
- `profiler.{py,go,zig}` - Profiling tools
- `visualize.{py,go,zig}` - Plotting utilities
- `README.md` - Usage documentation with examples
- `examples/` - Demonstration of tools on simple algorithms

### Success Criteria
- Can measure sub-millisecond operations accurately
- Statistical analysis shows appropriate variance
- Visualizations clearly show algorithmic trends (O(n), O(n²), etc.)

---

## Lab 0.2: Solving Recurrences

**Duration**: 2-3 hours  
**Difficulty**: Medium

### Objectives
- Solve recurrences using substitution, recursion tree, and Master theorem
- Verify solutions empirically
- Understand the gap between theory and practice

### Requirements

1. **Implement 5 recursive algorithms** with different recurrence relations:
   - T(n) = T(n-1) + O(1) - Linear recursion
   - T(n) = 2T(n/2) + O(n) - Divide and conquer
   - T(n) = T(n/2) + O(1) - Binary search pattern
   - T(n) = 2T(n/2) + O(1) - Tree recursion
   - T(n) = T(n-1) + O(n) - Quadratic recursion

2. **For each algorithm**:
   - Derive the closed-form solution mathematically
   - Implement the recursive function
   - Measure actual runtime empirically
   - Compare theoretical vs. empirical complexity

3. **Create a recurrence solver tool** that:
   - Takes a recurrence relation as input
   - Applies Master theorem when applicable
   - Generates recursion tree visualization
   - Estimates the solution

### Test Cases
- Verify solutions match known complexity classes
- Test edge cases (n=0, n=1, large n)
- Compare recursive vs. iterative implementations

### Deliverables
- `recurrence_examples.{py,go,zig}` - Implementations of all 5 patterns
- `recurrence_solver.{py,go,zig}` - Tool for solving recurrences
- `analysis.md` - Mathematical derivations and empirical results
- Charts showing theoretical vs. actual performance

### Success Criteria
- All recurrences solved correctly
- Empirical data matches theoretical predictions within margin of error
- Can explain when/why empirical results diverge from theory

---

## Lab 0.3: Cache Effects & Real Hardware

**Duration**: 3-4 hours  
**Difficulty**: Hard

### Objectives
- Observe how cache hierarchy affects algorithm performance
- Understand why O(n) notation doesn't tell the whole story
- Learn cache-oblivious algorithm design

### Requirements

1. **Implement matrix operations** with different access patterns:
   - Row-major traversal
   - Column-major traversal
   - Blocked/tiled multiplication
   - Cache-oblivious algorithms

2. **Measure and analyze**:
   - Cache hit/miss rates (use performance counters if available)
   - Impact of data structure layout on performance
   - Effects of memory alignment
   - TLB misses and page faults

3. **Create experiments demonstrating**:
   - Locality of reference (spatial and temporal)
   - False sharing in concurrent programs
   - Prefetching effects
   - Cache-friendly vs. cache-hostile data structures

### Implementation Details

```python
matrix_multiply_naive(A, B, n)
matrix_multiply_blocked(A, B, n, block_size)
matrix_multiply_cache_oblivious(A, B, n)

measure_cache_performance(function, matrix_size)
# Returns: {time, cache_hits, cache_misses, miss_rate}
```

### Test Cases
- Matrices of increasing size (64×64, 128×128, ..., 4096×4096)
- Different block sizes for tiled multiplication
- Comparison with BLAS library implementations

### Deliverables
- `matrix_ops.{py,go,zig}` - Various matrix multiplication implementations
- `cache_analysis.{py,go,zig}` - Cache performance measurement tools
- `results.md` - Detailed analysis of cache effects
- Performance graphs showing cache impact

### Success Criteria
- Clear demonstration of 2-10x speedup from cache-friendly access
- Accurate measurement of cache metrics
- Understanding of when cache effects dominate asymptotic complexity

---

## Lab 0.4: Micro-Benchmarking Framework

**Duration**: 4-5 hours  
**Difficulty**: Hard

### Objectives
- Build production-quality benchmarking infrastructure
- Understand sources of measurement noise
- Create repeatable, reliable benchmarks

### Requirements

1. **Develop a comprehensive benchmarking framework** that:
   - Isolates benchmarks from system noise
   - Handles CPU frequency scaling
   - Detects and reports outliers
   - Supports comparative benchmarking (A/B testing)
   - Generates confidence intervals

2. **Implement noise reduction techniques**:
   - Pin to specific CPU cores
   - Disable frequency scaling
   - Run at elevated priority
   - Disable hyperthreading effects
   - Control for background processes

3. **Create benchmark suites for**:
   - Algorithmic complexity (big-O verification)
   - Constant factor comparisons
   - Regression testing
   - Cross-language/cross-implementation comparisons

### Implementation Details

```python
benchmark_suite = BenchmarkSuite()
benchmark_suite.add("algorithm_name", function, input_generator)
benchmark_suite.configure(runs=100, warmup=10, outlier_threshold=2.0)
results = benchmark_suite.run()
results.compare("algo1", "algo2")
results.export("report.html")
```

### Test Cases
- Verify reproducibility (same results across runs)
- Detect performance regressions
- Statistical significance testing

### Deliverables
- `benchmark_framework.{py,go,zig}` - Core benchmarking library
- `benchmark_suite.{py,go,zig}` - Test suite infrastructure
- `report_generator.{py,go,zig}` - HTML/Markdown report generation
- `docs/` - Complete documentation and best practices guide
- `examples/` - Sample benchmark suites

### Success Criteria
- Coefficient of variation < 2% for stable benchmarks
- Can detect 5% performance differences reliably
- Framework used successfully in subsequent modules

---

## Module Resources

### CLRS References
- Chapter 1: The Role of Algorithms in Computing
- Chapter 2: Getting Started (2.2 Analyzing algorithms)
- Chapter 3: Growth of Functions
- Chapter 4: Divide-and-Conquer (4.3-4.6 recurrences)

### Additional Reading
- "Systems Performance" by Brendan Gregg (cache analysis)
- "The Art of Computer Programming" Vol 1 (mathematical foundations)
- Performance analysis papers for your chosen language

### Tools You'll Build
All tools from this module will be used throughout the curriculum:
- Timer and profiler (every module)
- Benchmark framework (every module)
- Visualization tools (every module)
- Recurrence solver (divide-and-conquer algorithms)

### Next Module
Once you complete Module 0, you'll move to **Module 1: Elementary Data Structures**, where you'll apply these analysis tools to compare different data structure implementations.

---

## Tips for Success

1. **Don't skip this module** - The tools you build here are essential for all future work
2. **Invest time in good visualizations** - They'll help you understand results quickly
3. **Test your tools thoroughly** - Incorrect measurements will mislead you later
4. **Document everything** - Future you will thank present you
5. **Start simple** - Get basic timing working before adding fancy features

## Common Pitfalls

- **Forgetting warmup runs** - JIT compilation can skew first measurements
- **Not accounting for variance** - Single measurements are unreliable
- **Ignoring cache effects** - Real performance often differs from theory
- **Over-engineering** - Start with simple tools, refine as needed
- **Not validating measurements** - Always sanity-check against known results
