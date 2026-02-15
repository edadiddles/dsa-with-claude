# Lab 4.6: Linear-Time Sorts

**Module**: Module 4 - Sorting & Selection  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 4.6: Linear-Time Sorts

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement non-comparison sorts
- Understand when O(n) sorting is possible
- Recognize problem constraints enabling linear time

### Requirements

1. **Implement Counting Sort**:
   - For integers in known range [0, k]
   - Time: O(n + k), Space: O(k)
   - Stable version

2. **Implement Radix Sort**:
   - LSD (Least Significant Digit) version
   - MSD (Most Significant Digit) version
   - Works for integers and strings
   - Time: O(d(n + k)) where d = number of digits

3. **Implement Bucket Sort**:
   - For uniformly distributed data
   - Expected O(n) time
   - Choose appropriate number of buckets

4. **Analysis**:
   - When can these be used?
   - What constraints must hold?
   - Comparison with O(n log n) sorts

### Implementation Details

```python
def counting_sort(array, max_val):
    """Stable counting sort for integers in [0, max_val]"""
    count = [0] * (max_val + 1)
    
    # Count occurrences
    for x in array:
        count[x] += 1
    
    # Cumulative counts for stable sort
    for i in range(1, len(count)):
        count[i] += count[i-1]
    
    # Build output array
    output = [0] * len(array)
    for x in reversed(array):  # Reverse to maintain stability
        output[count[x] - 1] = x
        count[x] -= 1
    
    return output

def counting_sort_by_digit(array, digit, base=10):
    """Helper for radix sort - sort by specific digit"""
    count = [0] * base
    
    for x in array:
        digit_value = (x // (base ** digit)) % base
        count[digit_value] += 1
    
    for i in range(1, base):
        count[i] += count[i-1]
    
    output = [0] * len(array)
    for x in reversed(array):
        digit_value = (x // (base ** digit)) % base
        output[count[digit_value] - 1] = x
        count[digit_value] -= 1
    
    return output

def radix_sort_lsd(array, num_digits=None, base=10):
    """Least Significant Digit radix sort"""
    if not array:
        return array
    
    if num_digits is None:
        num_digits = len(str(max(array)))
    
    for digit in range(num_digits):
        array = counting_sort_by_digit(array, digit, base)
    
    return array

def radix_sort_msd(array, digit=None, base=10):
    """Most Significant Digit radix sort"""
    if not array or len(array) <= 1:
        return array
    
    if digit is None:
        digit = len(str(max(array))) - 1
    
    if digit < 0:
        return array
    
    # Partition by digit
    buckets = [[] for _ in range(base)]
    for x in array:
        digit_value = (x // (base ** digit)) % base
        buckets[digit_value].append(x)
    
    # Recursively sort each bucket
    result = []
    for bucket in buckets:
        result.extend(radix_sort_msd(bucket, digit - 1, base))
    
    return result

def bucket_sort(array, num_buckets=None):
    """Bucket sort for uniformly distributed data"""
    if not array:
        return array
    
    min_val, max_val = min(array), max(array)
    
    if num_buckets is None:
        num_buckets = len(array)
    
    bucket_range = (max_val - min_val) / num_buckets
    if bucket_range == 0:
        bucket_range = 1
    
    # Create buckets
    buckets = [[] for _ in range(num_buckets)]
    
    # Distribute elements
    for x in array:
        bucket_idx = int((x - min_val) / bucket_range)
        if bucket_idx == num_buckets:  # Handle max_val
            bucket_idx -= 1
        buckets[bucket_idx].append(x)
    
    # Sort each bucket and concatenate
    result = []
    for bucket in buckets:
        # Use insertion sort for small buckets
        insertion_sort(bucket)
        result.extend(bucket)
    
    return result

def radix_sort_strings(strings):
    """Radix sort for strings (MSD)"""
    def get_char(s, pos):
        return ord(s[pos]) if pos < len(s) else -1
    
    def radix_msd_helper(strings, pos, result):
        if not strings or pos >= max(len(s) for s in strings):
            result.extend(strings)
            return
        
        # Bucket by character at position
        buckets = {}
        for s in strings:
            char = get_char(s, pos)
            if char not in buckets:
                buckets[char] = []
            buckets[char].append(s)
        
        # Process buckets in order
        for char in sorted(buckets.keys()):
            if char == -1:
                result.extend(buckets[char])
            else:
                radix_msd_helper(buckets[char], pos + 1, result)
    
    result = []
    radix_msd_helper(strings, 0, result)
    return result
```

### Test Cases
- Counting sort: Small range (k << n), large range (k >> n)
- Radix sort: Integers, strings, floating-point (with tricks)
- Bucket sort: Uniform distribution, skewed distribution
- Edge cases: empty arrays, single elements, all duplicates

### Deliverables
- `counting_sort.{py,go,zig}`
- `radix_sort.{py,go,zig}` - Both LSD and MSD
- `bucket_sort.{py,go,zig}`
- `applicability_guide.md` - When to use each
- `benchmarks.{py,go,zig}` - Compare with comparison sorts
- `string_radix_sort.{py,go,zig}` - For string sorting

### Success Criteria
- All sorts correct and stable where required
- Understanding of problem constraints for linear time
- Knowledge of tradeoffs (time vs. space, range vs. size)
- Can identify when linear-time sorting is possible

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_5.md) (if exists)
- [Next Lab](./lab_4_7.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 6 of 7*
