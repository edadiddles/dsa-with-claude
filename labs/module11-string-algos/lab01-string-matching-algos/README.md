# Lab 11.1: String Matching Algorithms

**Module**: Module 11 - String Algorithms  
**Duration**: 5-6 hours  
**Difficulty**: Medium

---

Lab 11.1: String Matching Algorithms

**Duration**: 5-6 hours  
**Difficulty**: Medium

### Objectives
- Implement classic string matching algorithms
- Understand preprocessing and search phases
- Compare algorithm performance
- Apply to pattern matching problems

### Requirements

1. **Implement string matching algorithms**:
   - **Naive algorithm** (baseline)
   - **Knuth-Morris-Pratt (KMP)**
   - **Boyer-Moore**
   - **Rabin-Karp** (rolling hash)
   - **Z-algorithm**

2. **Analysis**:
   - Preprocessing time
   - Search time
   - Space complexity
   - Best/worst cases

3. **Applications**:
   - Find all occurrences
   - Count matches
   - Multiple pattern matching

### Implementation Details

```python
# NAIVE STRING MATCHING

def naive_string_match(text, pattern):
    """
    Naive string matching: O(nm) worst case.
    Check every position in text.
    """
    n, m = len(text), len(pattern)
    matches = []
    
    for i in range(n - m + 1):
        # Check if pattern matches at position i
        match = True
        for j in range(m):
            if text[i + j] != pattern[j]:
                match = False
                break
        
        if match:
            matches.append(i)
    
    return matches

# KNUTH-MORRIS-PRATT (KMP)

def compute_lps(pattern):
    """
    Compute Longest Proper Prefix which is also Suffix array.
    LPS[i] = length of longest proper prefix of pattern[0..i]
    which is also a suffix of pattern[0..i].
    
    Time: O(m)
    """
    m = len(pattern)
    lps = [0] * m
    length = 0  # Length of previous longest prefix suffix
    i = 1
    
    while i < m:
        if pattern[i] == pattern[length]:
            length += 1
            lps[i] = length
            i += 1
        else:
            if length != 0:
                length = lps[length - 1]
            else:
                lps[i] = 0
                i += 1
    
    return lps

def kmp_search(text, pattern):
    """
    KMP algorithm: O(n + m) time.
    Never backtracks in text, uses LPS array.
    """
    n, m = len(text), len(pattern)
    
    if m == 0:
        return []
    
    # Preprocessing: compute LPS array
    lps = compute_lps(pattern)
    
    matches = []
    i = 0  # Index for text
    j = 0  # Index for pattern
    
    while i < n:
        if pattern[j] == text[i]:
            i += 1
            j += 1
        
        if j == m:
            # Found match
            matches.append(i - j)
            j = lps[j - 1]
        elif i < n and pattern[j] != text[i]:
            if j != 0:
                j = lps[j - 1]
            else:
                i += 1
    
    return matches

# BOYER-MOORE

def bad_character_heuristic(pattern):
    """
    Preprocessing for bad character heuristic.
    Returns last occurrence of each character in pattern.
    """
    m = len(pattern)
    bad_char = {}
    
    for i in range(m):
        bad_char[pattern[i]] = i
    
    return bad_char

def good_suffix_heuristic(pattern):
    """
    Preprocessing for good suffix heuristic.
    More complex than bad character.
    """
    m = len(pattern)
    suffix = [0] * m
    good_suffix = [0] * m
    
    # Compute suffix array
    suffix[m - 1] = m
    g = m - 1
    f = 0
    
    for i in range(m - 2, -1, -1):
        if i > g and suffix[i + m - 1 - f] < i - g:
            suffix[i] = suffix[i + m - 1 - f]
        else:
            if i < g:
                g = i
            f = i
            while g >= 0 and pattern[g] == pattern[g + m - 1 - f]:
                g -= 1
            suffix[i] = f - g
    
    # Compute good suffix shift
    for i in range(m):
        good_suffix[i] = m
    
    j = 0
    for i in range(m - 1, -1, -1):
        if suffix[i] == i + 1:
            while j < m - 1 - i:
                if good_suffix[j] == m:
                    good_suffix[j] = m - 1 - i
                j += 1
    
    for i in range(m - 1):
        good_suffix[m - 1 - suffix[i]] = m - 1 - i
    
    return good_suffix

def boyer_moore_search(text, pattern):
    """
    Boyer-Moore algorithm: O(n/m) best case, O(nm) worst case.
    Scans pattern from right to left.
    Uses bad character and good suffix heuristics.
    """
    n, m = len(text), len(pattern)
    
    if m == 0:
        return []
    
    # Preprocessing
    bad_char = bad_character_heuristic(pattern)
    good_suffix = good_suffix_heuristic(pattern)
    
    matches = []
    s = 0  # Shift of pattern with respect to text
    
    while s <= n - m:
        j = m - 1
        
        # Scan pattern from right to left
        while j >= 0 and pattern[j] == text[s + j]:
            j -= 1
        
        if j < 0:
            # Match found
            matches.append(s)
            s += good_suffix[0] if s + m < n else 1
        else:
            # Mismatch: use both heuristics
            bad_char_shift = j - bad_char.get(text[s + j], -1)
            good_suffix_shift = good_suffix[j]
            s += max(bad_char_shift, good_suffix_shift)
    
    return matches

# RABIN-KARP (Rolling Hash)

def rabin_karp_search(text, pattern, prime=101):
    """
    Rabin-Karp algorithm: O(n + m) expected, O(nm) worst case.
    Uses rolling hash for fast comparison.
    """
    n, m = len(text), len(pattern)
    d = 256  # Number of characters in alphabet
    
    if m == 0:
        return []
    
    matches = []
    pattern_hash = 0
    text_hash = 0
    h = 1
    
    # h = d^(m-1) % prime
    for i in range(m - 1):
        h = (h * d) % prime
    
    # Calculate hash for pattern and first window of text
    for i in range(m):
        pattern_hash = (d * pattern_hash + ord(pattern[i])) % prime
        text_hash = (d * text_hash + ord(text[i])) % prime
    
    # Slide pattern over text
    for i in range(n - m + 1):
        # Check if hash values match
        if pattern_hash == text_hash:
            # Verify actual match (hash collision possible)
            if text[i:i + m] == pattern:
                matches.append(i)
        
        # Calculate hash for next window
        if i < n - m:
            text_hash = (d * (text_hash - ord(text[i]) * h) + ord(text[i + m])) % prime
            
            # Ensure positive hash
            if text_hash < 0:
                text_hash += prime
    
    return matches

# Z-ALGORITHM

def compute_z_array(s):
    """
    Compute Z-array: Z[i] = length of longest substring
    starting from s[i] which is also a prefix of s.
    
    Time: O(n)
    """
    n = len(s)
    z = [0] * n
    z[0] = n
    
    l, r = 0, 0
    
    for i in range(1, n):
        if i > r:
            l, r = i, i
            while r < n and s[r - l] == s[r]:
                r += 1
            z[i] = r - l
            r -= 1
        else:
            k = i - l
            if z[k] < r - i + 1:
                z[i] = z[k]
            else:
                l = i
                while r < n and s[r - l] == s[r]:
                    r += 1
                z[i] = r - l
                r -= 1
    
    return z

def z_algorithm_search(text, pattern):
    """
    Z-algorithm for pattern matching: O(n + m).
    Concatenate pattern$text and compute Z-array.
    """
    # Concatenate with separator
    concat = pattern + '$' + text
    z = compute_z_array(concat)
    
    m = len(pattern)
    matches = []
    
    for i in range(m + 1, len(concat)):
        if z[i] == m:
            matches.append(i - m - 1)
    
    return matches

# PERFORMANCE COMPARISON

def compare_string_matching_algorithms(text, pattern):
    """
    Compare performance of different algorithms.
    """
    import time
    
    algorithms = {
        'Naive': naive_string_match,
        'KMP': kmp_search,
        'Boyer-Moore': boyer_moore_search,
        'Rabin-Karp': rabin_karp_search,
        'Z-Algorithm': z_algorithm_search
    }
    
    results = {}
    
    for name, func in algorithms.items():
        start = time.time()
        matches = func(text, pattern)
        elapsed = time.time() - start
        
        results[name] = {
            'time': elapsed,
            'matches': len(matches),
            'positions': matches[:5]  # First 5 matches
        }
    
    return results

# APPLICATIONS

def find_all_anagrams(text, pattern):
    """
    Find all starting positions of anagrams of pattern in text.
    Uses sliding window with character count.
    """
    from collections import Counter
    
    n, m = len(text), len(pattern)
    if m > n:
        return []
    
    pattern_count = Counter(pattern)
    window_count = Counter(text[:m])
    
    matches = []
    
    if window_count == pattern_count:
        matches.append(0)
    
    for i in range(m, n):
        # Add new character
        window_count[text[i]] += 1
        
        # Remove old character
        window_count[text[i - m]] -= 1
        if window_count[text[i - m]] == 0:
            del window_count[text[i - m]]
        
        # Check if anagram
        if window_count == pattern_count:
            matches.append(i - m + 1)
    
    return matches

def longest_common_prefix(strs):
    """
    Find longest common prefix of array of strings.
    """
    if not strs:
        return ""
    
    # Binary search on length
    min_len = min(len(s) for s in strs)
    
    low, high = 0, min_len
    
    while low <= high:
        mid = (low + high) // 2
        
        if all(s[:mid] == strs[0][:mid] for s in strs):
            low = mid + 1
        else:
            high = mid - 1
    
    return strs[0][:high]
```

### Test Cases
- Simple patterns and texts
- Pattern not in text
- Multiple occurrences
- Pattern at boundaries
- Long texts and patterns
- Performance benchmarking

### Deliverables
- `naive_matching.{py,go,zig}` - Baseline implementation
- `kmp.{py,go,zig}` - KMP with LPS array
- `boyer_moore.{py,go,zig}` - Both heuristics
- `rabin_karp.{py,go,zig}` - Rolling hash
- `z_algorithm.{py,go,zig}` - Z-array implementation
- `performance_comparison.{py,go,zig}` - Benchmark all algorithms
- `algorithm_analysis.md` - Complexity comparison
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All algorithms find correct matches
- Understanding of preprocessing trade-offs
- KMP never backtracks in text
- Boyer-Moore achieves sublinear best case
- Performance matches theoretical analysis

---

---

## Back to Module

[← Back to Module 11: String Algorithms](../module_11/README.md)

## Navigation

- [Previous Lab](./lab_11_0.md) (if exists)
- [Next Lab](./lab_11_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 11, Lab 1 of 5*
