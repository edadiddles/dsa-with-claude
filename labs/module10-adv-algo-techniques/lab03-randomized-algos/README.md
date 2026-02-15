# Lab 10.3: Randomized Algorithms

**Module**: Module 10 - Advanced Algorithm Techniques  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 10.3: Randomized Algorithms

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Understand randomized algorithm paradigm
- Implement Monte Carlo and Las Vegas algorithms
- Analyze expected running time
- Apply randomization to practical problems

### Requirements

1. **Implement randomized algorithms**:
   - **Randomized Quicksort**
   - **Randomized Selection** (QuickSelect)
   - **Monte Carlo Primality Testing** (Miller-Rabin)
   - **Randomized Min-Cut** (Karger's algorithm)
   - **Skip List** (from Module 5, review)
   - **Bloom Filter** (from Module 5, review)

2. **Analysis**:
   - Expected time complexity
   - Probability of correctness
   - Worst-case vs. expected case

3. **Comparison**:
   - Deterministic vs. randomized
   - Las Vegas vs. Monte Carlo

### Implementation Details

```python
import random

# RANDOMIZED QUICKSORT

def randomized_quicksort(arr):
    """
    Quicksort with random pivot selection.
    Expected O(n log n), worst case O(n^2) but unlikely.
    """
    if len(arr) <= 1:
        return arr
    
    # Random pivot
    pivot_idx = random.randint(0, len(arr) - 1)
    pivot = arr[pivot_idx]
    
    # Partition
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    
    return randomized_quicksort(left) + middle + randomized_quicksort(right)

# RANDOMIZED SELECTION (QuickSelect)

def randomized_select(arr, k):
    """
    Find k-th smallest element using randomized selection.
    Expected O(n), worst case O(n^2).
    Las Vegas algorithm: always correct, randomized time.
    """
    if len(arr) == 1:
        return arr[0]
    
    # Random pivot
    pivot_idx = random.randint(0, len(arr) - 1)
    pivot = arr[pivot_idx]
    
    # Partition
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    
    if k < len(left):
        return randomized_select(left, k)
    elif k < len(left) + len(middle):
        return pivot
    else:
        return randomized_select(right, k - len(left) - len(middle))

# MILLER-RABIN PRIMALITY TEST

def miller_rabin(n, k=5):
    """
    Miller-Rabin primality test.
    Monte Carlo algorithm: randomized, small error probability.
    
    Parameters:
    - n: number to test
    - k: number of rounds (higher k = lower error probability)
    
    Error probability: at most 1/(4^k)
    """
    if n < 2:
        return False
    if n == 2 or n == 3:
        return True
    if n % 2 == 0:
        return False
    
    # Write n-1 as 2^r * d
    r, d = 0, n - 1
    while d % 2 == 0:
        r += 1
        d //= 2
    
    # Witness loop
    for _ in range(k):
        a = random.randint(2, n - 2)
        x = pow(a, d, n)
        
        if x == 1 or x == n - 1:
            continue
        
        for _ in range(r - 1):
            x = pow(x, 2, n)
            if x == n - 1:
                break
        else:
            return False
    
    return True

def generate_prime(bits):
    """Generate random prime with given bit length"""
    while True:
        candidate = random.getrandbits(bits)
        # Make odd
        candidate |= 1
        
        if miller_rabin(candidate, k=10):
            return candidate

# KARGER'S MIN-CUT ALGORITHM

def karger_min_cut(graph):
    """
    Karger's randomized min-cut algorithm.
    Find minimum cut in undirected graph.
    
    Monte Carlo: may give wrong answer, repeat for better probability.
    Running O(n^2 log n) times gives high probability of correctness.
    """
    # Graph representation: adjacency list with parallel edges
    # Copy graph to avoid modifying original
    import copy
    g = copy.deepcopy(graph)
    
    def contract_edge(g, u, v):
        """Contract edge (u,v) by merging v into u"""
        # Move all edges from v to u
        for neighbor in g[v]:
            if neighbor != u:
                g[u].append(neighbor)
                # Update neighbor's edge
                for i, n in enumerate(g[neighbor]):
                    if n == v:
                        g[neighbor][i] = u
        
        # Remove v
        del g[v]
        
        # Remove self-loops at u
        g[u] = [n for n in g[u] if n != u]
    
    # Contract until 2 vertices remain
    while len(g) > 2:
        # Choose random edge
        u = random.choice(list(g.keys()))
        v = random.choice(g[u])
        
        contract_edge(g, u, v)
    
    # Count edges between remaining vertices
    vertices = list(g.keys())
    return len(g[vertices[0]])

def karger_min_cut_repeated(graph, iterations=None):
    """
    Repeat Karger's algorithm multiple times.
    Choose minimum cut found.
    
    Running O(n^2 log n) times gives probability > 1 - 1/n of correctness.
    """
    n = len(graph)
    if iterations is None:
        iterations = n * n * int(math.log(n) + 1)
    
    min_cut = float('inf')
    
    for _ in range(iterations):
        cut_size = karger_min_cut(graph)
        min_cut = min(min_cut, cut_size)
    
    return min_cut

# RESERVOIR SAMPLING

def reservoir_sampling(stream, k):
    """
    Select k random elements from stream of unknown length.
    Each element has equal probability k/n of being selected.
    """
    reservoir = []
    
    for i, element in enumerate(stream):
        if i < k:
            reservoir.append(element)
        else:
            # Random index in [0, i]
            j = random.randint(0, i)
            if j < k:
                reservoir[j] = element
    
    return reservoir

# RANDOMIZED ROUNDING

def randomized_rounding(fractional_solution):
    """
    Convert fractional solution to integer solution randomly.
    Used in approximation algorithms.
    """
    integer_solution = []
    
    for x in fractional_solution:
        if random.random() < x:
            integer_solution.append(1)
        else:
            integer_solution.append(0)
    
    return integer_solution

# LAS VEGAS VS MONTE CARLO

class AlgorithmClassifier:
    """Classify randomized algorithms"""
    
    @staticmethod
    def is_las_vegas(algorithm):
        """
        Las Vegas: Always correct, randomized running time.
        Examples: Randomized Quicksort, QuickSelect
        """
        return {
            'always_correct': True,
            'randomized_time': True,
            'examples': ['Randomized Quicksort', 'QuickSelect']
        }
    
    @staticmethod
    def is_monte_carlo(algorithm):
        """
        Monte Carlo: Randomized correctness, fixed time.
        Examples: Miller-Rabin, Karger's Min-Cut
        """
        return {
            'randomized_correctness': True,
            'fixed_time': True,
            'error_probability': 'can be made arbitrarily small',
            'examples': ['Miller-Rabin', 'Karger Min-Cut']
        }

# PROBABILISTIC ANALYSIS

def analyze_expected_time(algorithm, inputs, trials=1000):
    """
    Empirically measure expected running time.
    """
    import time
    
    times = []
    
    for _ in range(trials):
        test_input = random.choice(inputs)
        
        start = time.time()
        algorithm(test_input)
        end = time.time()
        
        times.append(end - start)
    
    return {
        'mean': sum(times) / len(times),
        'min': min(times),
        'max': max(times),
        'median': sorted(times)[len(times)//2]
    }

def probability_of_correctness(monte_carlo_algorithm, test_cases, trials=1000):
    """
    Empirically measure probability of correctness.
    """
    correct = 0
    
    for test_input, expected_output in test_cases:
        for _ in range(trials):
            result = monte_carlo_algorithm(test_input)
            if result == expected_output:
                correct += 1
    
    return correct / (len(test_cases) * trials)
```

### Test Cases
- Randomized quicksort on various inputs
- QuickSelect for finding median
- Miller-Rabin on known primes and composites
- Karger's min-cut on small graphs
- Measure probability of correctness

### Deliverables
- `randomized_quicksort.{py,go,zig}` - Implementation
- `randomized_select.{py,go,zig}` - QuickSelect
- `miller_rabin.{py,go,zig}` - Primality testing
- `karger_mincut.{py,go,zig}` - Min-cut algorithm
- `reservoir_sampling.{py,go,zig}` - Streaming algorithm
- `probabilistic_analysis.{py,go,zig}` - Analysis tools
- `randomized_vs_deterministic.md` - Comparison
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Understanding of Las Vegas vs. Monte Carlo
- Can analyze expected running time
- Probability of correctness measurable
- Know when randomization helps

---

---

## Back to Module

[← Back to Module 10: Advanced Algorithm Techniques](../module_10/README.md)

## Navigation

- [Previous Lab](./lab_10_2.md) (if exists)
- [Next Lab](./lab_10_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 10, Lab 3 of 5*
