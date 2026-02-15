# Lab 5.6: Bloom Filter

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 5.6: Bloom Filter

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement probabilistic set membership
- Understand false positive rates
- Apply to practical problems

### Requirements

1. **Implement Bloom filter with**:
   - `add(item)` - Add item to set
   - `contains(item)` - Check membership (may have false positives)
   - Configurable size m and number of hash functions k

2. **Hash functions**:
   - Use multiple hash functions
   - Or use double hashing: h_i(x) = (h1(x) + i*h2(x)) mod m

3. **Analysis**:
   - Calculate optimal k given m and n
   - Measure actual false positive rate
   - Study effect of load factor

4. **Applications**:
   - Cache filter (avoid expensive lookups)
   - Spell checker
   - Distributed systems (reduce network lookups)

### Implementation Details

```python
import hashlib
import math

class BloomFilter:
    def __init__(self, expected_elements, false_positive_rate):
        # Calculate optimal m and k
        self.m = self._optimal_m(expected_elements, false_positive_rate)
        self.k = self._optimal_k(self.m, expected_elements)
        self.bit_array = [False] * self.m
        self.num_elements = 0
    
    def _optimal_m(self, n, p):
        """m = -(n * ln(p)) / (ln(2)^2)"""
        return int(-(n * math.log(p)) / (math.log(2) ** 2))
    
    def _optimal_k(self, m, n):
        """k = (m / n) * ln(2)"""
        return int((m / n) * math.log(2))
    
    def _hash(self, item, seed):
        """Generate hash using seed"""
        h = hashlib.md5((str(item) + str(seed)).encode())
        return int(h.hexdigest(), 16) % self.m
    
    def add(self, item):
        """Add item to filter"""
        for i in range(self.k):
            index = self._hash(item, i)
            self.bit_array[index] = True
        self.num_elements += 1
    
    def contains(self, item):
        """Check if item might be in set"""
        for i in range(self.k):
            index = self._hash(item, i)
            if not self.bit_array[index]:
                return False  # Definitely not in set
        return True  # Might be in set (possible false positive)
    
    def expected_fpr(self):
        """Calculate expected false positive rate"""
        # FPR = (1 - e^(-kn/m))^k
        return (1 - math.exp(-self.k * self.num_elements / self.m)) ** self.k
    
    def bits_set(self):
        """Count number of bits set to True"""
        return sum(self.bit_array)
    
    def load_factor(self):
        """Fraction of bits set"""
        return self.bits_set() / self.m

class ScalableBloomFilter:
    """Bloom filter that grows as needed"""
    def __init__(self, initial_capacity, fpr, growth_factor=2):
        self.filters = []
        self.initial_capacity = initial_capacity
        self.fpr = fpr
        self.growth_factor = growth_factor
        self._add_filter()
    
    def _add_filter(self):
        capacity = self.initial_capacity * (self.growth_factor ** len(self.filters))
        self.filters.append(BloomFilter(capacity, self.fpr))
    
    def add(self, item):
        # Add to most recent filter
        if self.filters[-1].expected_fpr() > self.fpr:
            self._add_filter()
        self.filters[-1].add(item)
    
    def contains(self, item):
        # Check all filters
        return any(f.contains(item) for f in self.filters)

class CountingBloomFilter:
    """Bloom filter that supports deletion"""
    def __init__(self, expected_elements, false_positive_rate):
        self.m = int(-(expected_elements * math.log(false_positive_rate)) / 
                     (math.log(2) ** 2))
        self.k = int((self.m / expected_elements) * math.log(2))
        self.counters = [0] * self.m
    
    def _hash(self, item, seed):
        h = hashlib.md5((str(item) + str(seed)).encode())
        return int(h.hexdigest(), 16) % self.m
    
    def add(self, item):
        for i in range(self.k):
            index = self._hash(item, i)
            self.counters[index] += 1
    
    def remove(self, item):
        # Only safe if item was actually added
        for i in range(self.k):
            index = self._hash(item, i)
            if self.counters[index] > 0:
                self.counters[index] -= 1
    
    def contains(self, item):
        for i in range(self.k):
            index = self._hash(item, i)
            if self.counters[index] == 0:
                return False
        return True

def web_cache_filter_example():
    """Use Bloom filter to avoid cache misses"""
    cache = {}
    bloom = BloomFilter(expected_elements=10000, false_positive_rate=0.01)
    
    def get(url):
        # Check Bloom filter first
        if not bloom.contains(url):
            # Definitely not in cache
            result = expensive_fetch(url)
            cache[url] = result
            bloom.add(url)
            return result
        else:
            # Might be in cache (check actual cache)
            if url in cache:
                return cache[url]  # Cache hit
            else:
                # False positive
                result = expensive_fetch(url)
                cache[url] = result
                return result
    
    return get

def spell_checker_bloom():
    """Fast spell checking with Bloom filter"""
    # Load dictionary
    dictionary_bloom = BloomFilter(expected_elements=100000, 
                                    false_positive_rate=0.001)
    
    with open('/usr/share/dict/words') as f:
        for word in f:
            dictionary_bloom.add(word.strip().lower())
    
    def is_word(word):
        return dictionary_bloom.contains(word.lower())
    
    return is_word
```

### Test Cases
- Add 10K elements, test false positive rate
- Vary m and k, measure FPR
- Test with different hash functions
- Overload filter and measure FPR increase
- Compare with set for space usage

### Deliverables
- `bloom_filter.{py,go,zig}` - Full implementation
- `scalable_bloom.{py,go,zig}` - Growing Bloom filter
- `counting_bloom.{py,go,zig}` - With deletion support
- `fpr_analysis.{py,go,zig}` - Measure false positive rates
- `optimal_parameters.md` - Guide for choosing m and k
- `applications.{py,go,zig}` - Practical use cases
- `benchmarks.{py,go,zig}` - Performance vs. hash set

### Success Criteria
- False positive rate matches theoretical predictions
- Understanding of space savings
- Knowledge of when Bloom filters are appropriate
- Can calculate optimal parameters

---

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_5.md) (if exists)
- [Next Lab](./lab_5_7.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 6 of 6*
