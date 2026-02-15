# Lab 11.2: Suffix Arrays & Suffix Trees

**Module**: Module 11 - String Algorithms  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 11.2: Suffix Arrays & Suffix Trees

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Build suffix arrays efficiently
- Construct suffix trees (Ukkonen's algorithm)
- Use suffix structures for pattern matching
- Apply to bioinformatics problems

### Requirements

1. **Implement suffix structures**:
   - **Suffix array** (naive and optimized)
   - **LCP array** (Longest Common Prefix)
   - **Suffix tree** (Ukkonen's algorithm)
   - **Enhanced suffix array**

2. **Applications**:
   - Pattern matching in O(m log n)
   - Longest repeated substring
   - Longest common substring of two strings
   - Finding all repeats

3. **Optimizations**:
   - Linear-time construction
   - Space-efficient representations

### Implementation Details

```python
# SUFFIX ARRAY - Naive Construction

def build_suffix_array_naive(text):
    """
    Build suffix array: O(n^2 log n) time.
    Suffix array[i] = starting position of i-th suffix in sorted order.
    """
    n = len(text)
    suffixes = [(text[i:], i) for i in range(n)]
    
    # Sort suffixes
    suffixes.sort()
    
    # Extract indices
    suffix_array = [suffix[1] for suffix in suffixes]
    
    return suffix_array

# SUFFIX ARRAY - Optimized Construction

def build_suffix_array_optimized(text):
    """
    Build suffix array in O(n log^2 n) time.
    Uses doubling technique.
    """
    n = len(text)
    text += '$'  # Add terminator
    
    # Initial ranking based on first character
    rank = [ord(c) for c in text]
    
    # Temporary rank array
    k = 1
    while k < n:
        # Sort based on pairs (rank[i], rank[i+k])
        pairs = [(rank[i], rank[i + k] if i + k < n else -1, i) 
                for i in range(n)]
        pairs.sort()
        
        # Update ranks
        new_rank = [0] * n
        for i in range(n):
            if i > 0 and pairs[i][:2] == pairs[i-1][:2]:
                new_rank[pairs[i][2]] = new_rank[pairs[i-1][2]]
            else:
                new_rank[pairs[i][2]] = i
        
        rank = new_rank
        k *= 2
    
    # Build suffix array from ranks
    suffix_array = [0] * n
    for i in range(n):
        suffix_array[rank[i]] = i
    
    return suffix_array

# LCP ARRAY

def build_lcp_array(text, suffix_array):
    """
    Build LCP (Longest Common Prefix) array.
    LCP[i] = length of longest common prefix between
    suffix_array[i] and suffix_array[i-1].
    
    Time: O(n)
    """
    n = len(text)
    rank = [0] * n
    
    for i in range(n):
        rank[suffix_array[i]] = i
    
    lcp = [0] * n
    h = 0
    
    for i in range(n):
        if rank[i] > 0:
            j = suffix_array[rank[i] - 1]
            
            while i + h < n and j + h < n and text[i + h] == text[j + h]:
                h += 1
            
            lcp[rank[i]] = h
            
            if h > 0:
                h -= 1
    
    return lcp

# PATTERN MATCHING WITH SUFFIX ARRAY

def pattern_match_suffix_array(text, pattern, suffix_array):
    """
    Pattern matching using suffix array: O(m log n).
    Binary search for pattern range.
    """
    n = len(text)
    m = len(pattern)
    
    # Binary search for first occurrence
    left, right = 0, n - 1
    
    while left < right:
        mid = (left + right) // 2
        suffix = text[suffix_array[mid]:]
        
        if suffix < pattern:
            left = mid + 1
        else:
            right = mid
    
    start = left
    
    # Binary search for last occurrence
    left, right = 0, n - 1
    
    while left < right:
        mid = (left + right + 1) // 2
        suffix = text[suffix_array[mid]:]
        
        if suffix[:m] <= pattern:
            left = mid
        else:
            right = mid - 1
    
    end = left
    
    # Check if pattern found
    if text[suffix_array[start]:].startswith(pattern):
        return suffix_array[start:end + 1]
    
    return []

# LONGEST REPEATED SUBSTRING

def longest_repeated_substring(text):
    """
    Find longest repeated substring using suffix array and LCP.
    """
    suffix_array = build_suffix_array_optimized(text)
    lcp = build_lcp_array(text, suffix_array)
    
    if not lcp:
        return ""
    
    max_lcp = max(lcp)
    max_idx = lcp.index(max_lcp)
    
    return text[suffix_array[max_idx]:suffix_array[max_idx] + max_lcp]

# LONGEST COMMON SUBSTRING OF TWO STRINGS

def longest_common_substring(s1, s2):
    """
    Find longest common substring of two strings.
    Concatenate s1#s2 and use suffix array.
    """
    # Concatenate with separator
    text = s1 + '#' + s2
    n1 = len(s1)
    
    suffix_array = build_suffix_array_optimized(text)
    lcp = build_lcp_array(text, suffix_array)
    
    max_lcp = 0
    lcs_start = 0
    
    for i in range(1, len(lcp)):
        # Check if suffixes are from different strings
        idx1 = suffix_array[i - 1]
        idx2 = suffix_array[i]
        
        if (idx1 < n1) != (idx2 < n1):
            if lcp[i] > max_lcp:
                max_lcp = lcp[i]
                lcs_start = min(idx1, idx2)
    
    return text[lcs_start:lcs_start + max_lcp]

# SUFFIX TREE - Ukkonen's Algorithm (Simplified)

class SuffixTreeNode:
    def __init__(self):
        self.children = {}
        self.start = None
        self.end = None
        self.suffix_link = None

class SuffixTree:
    """
    Suffix tree using Ukkonen's algorithm.
    Linear time construction: O(n).
    """
    def __init__(self, text):
        self.text = text + '$'
        self.root = SuffixTreeNode()
        self.build()
    
    def build(self):
        """Build suffix tree incrementally"""
        # Simplified implementation
        # Full Ukkonen's algorithm is complex
        pass
    
    def search(self, pattern):
        """
        Search for pattern in suffix tree: O(m).
        """
        node = self.root
        
        for char in pattern:
            if char not in node.children:
                return False
            node = node.children[char]
        
        return True
    
    def count_occurrences(self, pattern):
        """Count number of times pattern appears"""
        # Navigate to pattern node
        # Count leaves in subtree
        pass

# APPLICATIONS

def find_all_repeats(text, min_length=2):
    """
    Find all repeated substrings of minimum length.
    Uses suffix array and LCP.
    """
    suffix_array = build_suffix_array_optimized(text)
    lcp = build_lcp_array(text, suffix_array)
    
    repeats = set()
    
    for i in range(1, len(lcp)):
        if lcp[i] >= min_length:
            substr = text[suffix_array[i]:suffix_array[i] + lcp[i]]
            repeats.add(substr)
    
    return list(repeats)

def distinct_substrings_count(text):
    """
    Count number of distinct substrings.
    Uses suffix array: answer = n(n+1)/2 - sum(LCP)
    """
    n = len(text)
    suffix_array = build_suffix_array_optimized(text)
    lcp = build_lcp_array(text, suffix_array)
    
    total_substrings = n * (n + 1) // 2
    duplicate_substrings = sum(lcp)
    
    return total_substrings - duplicate_substrings
```

### Test Cases
- Simple strings
- Strings with many repeats
- DNA sequences
- Performance on large texts
- Edge cases (empty, single character)

### Deliverables
- `suffix_array.{py,go,zig}` - Both naive and optimized
- `lcp_array.{py,go,zig}` - Linear-time construction
- `suffix_tree.{py,go,zig}` - Ukkonen's algorithm
- `pattern_matching.{py,go,zig}` - Using suffix structures
- `applications.{py,go,zig}` - LRS, LCS, repeats
- `complexity_analysis.md` - Time and space analysis
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Suffix array construction is correct
- LCP array matches expected values
- Pattern matching works efficiently
- Applications produce correct results
- Understanding of suffix structure benefits

---

---

## Back to Module

[← Back to Module 11: String Algorithms](../module_11/README.md)

## Navigation

- [Previous Lab](./lab_11_1.md) (if exists)
- [Next Lab](./lab_11_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 11, Lab 2 of 5*
