# Lab 0.1: Build Your Analysis Framework

**Module**: Module 0 - Foundations & Analysis  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 0.1: Build Your Analysis Framework

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

---

## Back to Module

[← Back to Module 0: Foundations & Analysis](..)

## Navigation

- [Next Lab](../lab02-solving-recurrences/README.md)
- [All Labs](../..)

---

*Part of the comprehensive DSA Curriculum*  
*Module 0, Lab 1 of 4*
