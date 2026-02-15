# Lab 10.5: Algorithm Design Patterns

**Module**: Module 10 - Advanced Algorithm Techniques  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 10.5: Algorithm Design Patterns

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Recognize common algorithm design patterns
- Apply patterns to new problems
- Build problem-solving toolkit

### Requirements

1. **Master design patterns**:
   - Two pointers
   - Sliding window
   - Prefix sum / difference array
   - Monotonic stack/queue
   - Top-K elements
   - Union-Find applications

2. **Problem-solving framework**:
   - Pattern recognition
   - Template application
   - Adaptation to variants

3. **Practice problems**:
   - Solve 20+ problems using patterns
   - Categorize by pattern
   - Build intuition

### Implementation Details

```python
# TWO POINTERS

def two_sum_sorted(arr, target):
    """
    Find two numbers that sum to target in sorted array.
    Pattern: Two pointers (one from each end).
    """
    left, right = 0, len(arr) - 1
    
    while left < right:
        current_sum = arr[left] + arr[right]
        
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1
    
    return None

def remove_duplicates(arr):
    """
    Remove duplicates from sorted array in-place.
    Pattern: Two pointers (fast and slow).
    """
    if not arr:
        return 0
    
    slow = 0
    
    for fast in range(1, len(arr)):
        if arr[fast] != arr[slow]:
            slow += 1
            arr[slow] = arr[fast]
    
    return slow + 1

# SLIDING WINDOW

def max_sum_subarray(arr, k):
    """
    Maximum sum of subarray of size k.
    Pattern: Sliding window.
    """
    if len(arr) < k:
        return None
    
    # Initial window
    window_sum = sum(arr[:k])
    max_sum = window_sum
    
    # Slide window
    for i in range(k, len(arr)):
        window_sum += arr[i] - arr[i - k]
        max_sum = max(max_sum, window_sum)
    
    return max_sum

def longest_substring_k_distinct(s, k):
    """
    Longest substring with at most k distinct characters.
    Pattern: Sliding window with hashmap.
    """
    from collections import defaultdict
    
    char_count = defaultdict(int)
    left = 0
    max_length = 0
    
    for right in range(len(s)):
        char_count[s[right]] += 1
        
        while len(char_count) > k:
            char_count[s[left]] -= 1
            if char_count[s[left]] == 0:
                del char_count[s[left]]
            left += 1
        
        max_length = max(max_length, right - left + 1)
    
    return max_length

# PREFIX SUM

def range_sum_query(arr):
    """
    Precompute prefix sums for O(1) range queries.
    Pattern: Prefix sum.
    """
    prefix = [0]
    for num in arr:
        prefix.append(prefix[-1] + num)
    
    def query(left, right):
        return prefix[right + 1] - prefix[left]
    
    return query

def subarray_sum_equals_k(arr, k):
    """
    Count subarrays with sum equal to k.
    Pattern: Prefix sum with hashmap.
    """
    from collections import defaultdict
    
    count = 0
    prefix_sum = 0
    sum_count = defaultdict(int)
    sum_count[0] = 1
    
    for num in arr:
        prefix_sum += num
        
        # Check if prefix_sum - k exists
        count += sum_count[prefix_sum - k]
        
        sum_count[prefix_sum] += 1
    
    return count

# MONOTONIC STACK

def next_greater_element(arr):
    """
    For each element, find next greater element.
    Pattern: Monotonic stack.
    """
    result = [-1] * len(arr)
    stack = []
    
    for i in range(len(arr)):
        while stack and arr[i] > arr[stack[-1]]:
            idx = stack.pop()
            result[idx] = arr[i]
        stack.append(i)
    
    return result

def largest_rectangle_histogram(heights):
    """
    Largest rectangle in histogram.
    Pattern: Monotonic stack.
    """
    stack = []
    max_area = 0
    
    for i in range(len(heights)):
        while stack and heights[i] < heights[stack[-1]]:
            h_idx = stack.pop()
            h = heights[h_idx]
            w = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, h * w)
        
        stack.append(i)
    
    while stack:
        h_idx = stack.pop()
        h = heights[h_idx]
        w = len(heights) if not stack else len(heights) - stack[-1] - 1
        max_area = max(max_area, h * w)
    
    return max_area

# TOP K ELEMENTS

def top_k_frequent(nums, k):
    """
    Find k most frequent elements.
    Pattern: Heap / Bucket sort.
    """
    from collections import Counter
    import heapq
    
    count = Counter(nums)
    
    # Method 1: Min heap of size k
    heap = []
    for num, freq in count.items():
        heapq.heappush(heap, (freq, num))
        if len(heap) > k:
            heapq.heappop(heap)
    
    return [num for freq, num in heap]

def kth_largest(arr, k):
    """
    Find kth largest element.
    Pattern: QuickSelect or heap.
    """
    import heapq
    return heapq.nlargest(k, arr)[-1]

# PATTERN TEMPLATES

class AlgorithmPatterns:
    """Collection of algorithm pattern templates"""
    
    @staticmethod
    def two_pointer_template():
        """
        Template for two pointer problems:
        
        left, right = 0, len(arr) - 1
        
        while left < right:
            if condition:
                # Found answer
                return
            elif need_larger:
                left += 1
            else:
                right -= 1
        """
        pass
    
    @staticmethod
    def sliding_window_template():
        """
        Template for sliding window:
        
        left = 0
        for right in range(len(arr)):
            # Add arr[right] to window
            
            while window_invalid:
                # Remove arr[left] from window
                left += 1
            
            # Update result
        """
        pass
    
    @staticmethod
    def monotonic_stack_template():
        """
        Template for monotonic stack:
        
        stack = []
        for i in range(len(arr)):
            while stack and condition(arr[i], arr[stack[-1]]):
                # Process stack.pop()
            
            stack.append(i)
        """
        pass

# PROBLEM RECOGNITION GUIDE

def identify_pattern(problem_description):
    """
    Identify which pattern to use:
    
    Two Pointers:
    - Sorted array
    - Find pair/triplet
    - Opposite ends moving toward middle
    
    Sliding Window:
    - Contiguous subarray/substring
    - All subarrays of size k
    - Longest/shortest with property
    
    Prefix Sum:
    - Range sum queries
    - Subarray sum equals k
    - Multiple queries on static array
    
    Monotonic Stack:
    - Next greater/smaller element
    - Largest rectangle
    - Stock span problem
    
    Top K:
    - K largest/smallest
    - K most frequent
    - Heap or QuickSelect
    """
    pass
```

### Test Cases
- Problems for each pattern
- Variations requiring adaptation
- Combined patterns

### Deliverables
- `two_pointers.{py,go,zig}` - 10+ problems
- `sliding_window.{py,go,zig}` - 10+ problems
- `prefix_sum.{py,go,zig}` - 10+ problems
- `monotonic_stack.{py,go,zig}` - 10+ problems
- `top_k.{py,go,zig}` - 10+ problems
- `pattern_templates.{py,go,zig}` - Reusable templates
- `pattern_recognition_guide.md` - How to identify patterns
- `problem_catalog.md` - 50+ problems categorized

### Success Criteria
- Can recognize patterns in new problems
- Fluent with pattern templates
- Can combine multiple patterns
- Build problem-solving intuition

---

---

## Back to Module

[← Back to Module 10: Advanced Algorithm Techniques](../module_10/README.md)

## Navigation

- [Previous Lab](./lab_10_4.md) (if exists)
- [Next Lab](./lab_10_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 10, Lab 5 of 5*
