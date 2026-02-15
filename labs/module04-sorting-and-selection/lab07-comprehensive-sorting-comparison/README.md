# Lab 4.7: Comprehensive Sorting Comparison

**Module**: Module 4 - Sorting & Selection  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 4.7: Comprehensive Sorting Comparison

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Create definitive comparison of all sorting algorithms
- Build intuition for algorithm selection
- Understand practical vs. theoretical performance

### Requirements

1. **Benchmark all implemented sorts on**:
   - Random data
   - Sorted data
   - Reverse sorted data
   - Nearly sorted data (with various noise levels)
   - Data with many duplicates
   - Data with specific distributions (Gaussian, uniform, Zipfian)

2. **Measure**:
   - Running time
   - Number of comparisons
   - Number of swaps/moves
   - Memory usage
   - Cache performance

3. **Create decision tree/matrix**:
   - Input size ranges
   - Data characteristics
   - Memory constraints
   - Stability requirements
   - Recommended algorithm

### Implementation Details

```python
class SortBenchmark:
    def __init__(self):
        self.algorithms = {
            'quicksort': quicksort,
            'mergesort': mergesort_top_down,
            'heapsort': heapsort,
            'timsort': timsort,
            'counting': counting_sort,
            'radix': radix_sort_lsd,
            'bucket': bucket_sort,
        }
    
    def benchmark_all(self, test_data, runs=10):
        results = {}
        
        for name, algorithm in self.algorithms.items():
            times = []
            comparisons = []
            swaps = []
            
            for _ in range(runs):
                data = test_data[:]
                
                start = time.time()
                algorithm(data)
                elapsed = time.time() - start
                
                times.append(elapsed)
            
            results[name] = {
                'mean_time': statistics.mean(times),
                'std_time': statistics.stdev(times) if len(times) > 1 else 0,
                'min_time': min(times),
                'max_time': max(times),
            }
        
        return results
    
    def test_all_patterns(self, size):
        patterns = {
            'random': generate_random(size),
            'sorted': generate_sorted(size),
            'reverse': generate_reverse_sorted(size),
            'nearly_sorted': generate_nearly_sorted(size, noise=0.1),
            'many_duplicates': generate_many_duplicates(size),
            'gaussian': generate_gaussian(size),
        }
        
        all_results = {}
        for pattern_name, data in patterns.items():
            all_results[pattern_name] = self.benchmark_all(data)
        
        return all_results

def generate_test_data(size, pattern):
    if pattern == 'random':
        return [random.randint(0, size) for _ in range(size)]
    elif pattern == 'sorted':
        return list(range(size))
    elif pattern == 'reverse':
        return list(range(size, 0, -1))
    elif pattern == 'nearly_sorted':
        data = list(range(size))
        # Add 10% noise
        for _ in range(size // 10):
            i, j = random.randint(0, size-1), random.randint(0, size-1)
            data[i], data[j] = data[j], data[i]
        return data
    elif pattern == 'many_duplicates':
        return [random.randint(0, size // 10) for _ in range(size)]
    elif pattern == 'gaussian':
        return [int(random.gauss(size/2, size/6)) for _ in range(size)]
    elif pattern == 'zipfian':
        # Implement Zipfian distribution
        pass

def create_decision_matrix():
    """
    Generate decision matrix for choosing sorting algorithm
    """
    decision_matrix = {
        'small_n_random': 'insertion_sort',
        'small_n_nearly_sorted': 'insertion_sort',
        'medium_n_random': 'quicksort',
        'medium_n_nearly_sorted': 'timsort',
        'large_n_random': 'quicksort',
        'large_n_nearly_sorted': 'timsort',
        'stability_required': 'mergesort or timsort',
        'limited_memory': 'heapsort',
        'known_range_integers': 'counting_sort or radix_sort',
        'uniform_distribution': 'bucket_sort',
        'external_data': 'external_mergesort',
    }
    
    return decision_matrix
```

### Deliverables
- `sort_comparison.{py,go,zig}` - Comprehensive benchmark suite
- `test_data_generator.{py,go,zig}` - Generate various input patterns
- `results_visualizer.{py,go,zig}` - Create comparison charts
- `sorting_guide.md` - Decision guide for choosing sort
- `performance_report.md` - Detailed analysis with graphs
- `decision_matrix.md` - Quick reference guide

### Success Criteria
- Clear performance patterns emerge
- Decision guide is actionable
- Understanding of why certain sorts excel in certain scenarios
- Can make informed sorting algorithm choices

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_6.md) (if exists)
- [Next Lab](./lab_4_8.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 7 of 7*
