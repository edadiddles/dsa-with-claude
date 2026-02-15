# Module 11: String Algorithms

**Duration**: 2 weeks  
**Difficulty**: Medium to Hard

## Overview

String algorithms are fundamental to text processing, bioinformatics, data compression, and search engines. This module covers pattern matching, suffix structures, and advanced string manipulation techniques that power real-world applications from DNA sequencing to web search.

**Why this matters**: String processing appears everywhere - search engines use pattern matching, DNA analysis uses string alignment, compilers use parsing, and data compression uses string statistics. Efficient string algorithms can make the difference between instant search and hours of processing.

## Learning Objectives

- Master string matching algorithms (KMP, Boyer-Moore, Rabin-Karp)
- Build and use suffix arrays and suffix trees
- Implement text compression algorithms
- Solve sequence alignment problems
- Apply string algorithms to practical problems
- Understand computational complexity of string problems

## Topics Covered

- Naive string matching and optimizations
- Knuth-Morris-Pratt (KMP) algorithm
- Boyer-Moore algorithm
- Rabin-Karp rolling hash
- Suffix arrays and LCP arrays
- Suffix trees and Ukkonen's algorithm
- Aho-Corasick multi-pattern matching
- Edit distance and sequence alignment
- String compression (LZ77, LZW, Burrows-Wheeler)

---

## Lab 11.1: String Matching Algorithms

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

## Lab 11.2: Suffix Arrays & Suffix Trees

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

## Lab 11.3: Text Compression

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement compression algorithms
- Understand compression techniques
- Measure compression ratios
- Apply to real files

### Requirements

1. **Implement compression algorithms**:
   - **Run-Length Encoding (RLE)**
   - **LZ77**
   - **LZW (Lempel-Ziv-Welch)**
   - **Burrows-Wheeler Transform**
   - **Huffman coding** (from Module 9, review)

2. **Analysis**:
   - Compression ratio
   - Compression/decompression speed
   - Dictionary size

3. **Applications**:
   - File compression utility
   - Compare algorithms on different data types

### Implementation Details

```python
# RUN-LENGTH ENCODING

def rle_encode(data):
    """
    Run-length encoding: compress consecutive repeated characters.
    Good for data with many runs.
    """
    if not data:
        return ""
    
    encoded = []
    count = 1
    prev = data[0]
    
    for i in range(1, len(data)):
        if data[i] == prev:
            count += 1
        else:
            encoded.append(f"{count}{prev}")
            prev = data[i]
            count = 1
    
    encoded.append(f"{count}{prev}")
    
    return ''.join(encoded)

def rle_decode(encoded):
    """Decode run-length encoded data"""
    decoded = []
    i = 0
    
    while i < len(encoded):
        # Read count
        count = 0
        while i < len(encoded) and encoded[i].isdigit():
            count = count * 10 + int(encoded[i])
            i += 1
        
        # Read character
        if i < len(encoded):
            char = encoded[i]
            decoded.append(char * count)
            i += 1
    
    return ''.join(decoded)

# LZ77

def lz77_encode(data, window_size=4096, lookahead_size=18):
    """
    LZ77 compression: find longest match in sliding window.
    Output: (offset, length, next_char) tuples.
    """
    encoded = []
    i = 0
    
    while i < len(data):
        # Find longest match in window
        best_offset = 0
        best_length = 0
        
        window_start = max(0, i - window_size)
        
        for j in range(window_start, i):
            length = 0
            while (i + length < len(data) and 
                   length < lookahead_size and
                   data[j + length] == data[i + length]):
                length += 1
            
            if length > best_length:
                best_length = length
                best_offset = i - j
        
        # Output match or literal
        if best_length > 0:
            next_char = data[i + best_length] if i + best_length < len(data) else ''
            encoded.append((best_offset, best_length, next_char))
            i += best_length + (1 if next_char else 0)
        else:
            encoded.append((0, 0, data[i]))
            i += 1
    
    return encoded

def lz77_decode(encoded):
    """Decode LZ77 compressed data"""
    decoded = []
    
    for offset, length, char in encoded:
        if length > 0:
            # Copy from window
            start = len(decoded) - offset
            for i in range(length):
                decoded.append(decoded[start + i])
        
        if char:
            decoded.append(char)
    
    return ''.join(decoded)

# LZW (Lempel-Ziv-Welch)

def lzw_encode(data):
    """
    LZW compression: build dictionary of sequences.
    Output: sequence of codes.
    """
    # Initialize dictionary with single characters
    dict_size = 256
    dictionary = {chr(i): i for i in range(dict_size)}
    
    result = []
    w = ""
    
    for c in data:
        wc = w + c
        if wc in dictionary:
            w = wc
        else:
            result.append(dictionary[w])
            dictionary[wc] = dict_size
            dict_size += 1
            w = c
    
    if w:
        result.append(dictionary[w])
    
    return result

def lzw_decode(encoded):
    """Decode LZW compressed data"""
    # Initialize dictionary
    dict_size = 256
    dictionary = {i: chr(i) for i in range(dict_size)}
    
    result = []
    w = chr(encoded[0])
    result.append(w)
    
    for code in encoded[1:]:
        if code in dictionary:
            entry = dictionary[code]
        elif code == dict_size:
            entry = w + w[0]
        else:
            raise ValueError("Invalid LZW code")
        
        result.append(entry)
        
        dictionary[dict_size] = w + entry[0]
        dict_size += 1
        
        w = entry
    
    return ''.join(result)

# BURROWS-WHEELER TRANSFORM

def burrows_wheeler_transform(text):
    """
    Burrows-Wheeler Transform: reversible permutation.
    Makes similar characters cluster together.
    """
    # Add end-of-string marker
    text += '$'
    
    # Generate all rotations
    rotations = [text[i:] + text[:i] for i in range(len(text))]
    
    # Sort rotations
    rotations.sort()
    
    # Take last column
    bwt = ''.join(rotation[-1] for rotation in rotations)
    
    # Find index of original string
    original_index = rotations.index(text)
    
    return bwt, original_index

def inverse_burrows_wheeler(bwt, original_index):
    """
    Inverse Burrows-Wheeler Transform.
    Reconstruct original string.
    """
    n = len(bwt)
    
    # Build first column (sorted BWT)
    first = sorted(bwt)
    
    # Build next array: next[i] = index in BWT of character
    # that follows first[i] in original text
    next_array = [0] * n
    count = {}
    
    for i, char in enumerate(bwt):
        count[char] = count.get(char, 0)
        
        # Find position of this occurrence in first
        j = 0
        occurrences = 0
        while occurrences < count[char] or first[j] != char:
            if first[j] == char:
                occurrences += 1
            j += 1
        
        next_array[j] = i
        count[char] += 1
    
    # Reconstruct original text
    result = []
    idx = original_index
    
    for _ in range(n):
        result.append(first[idx])
        idx = next_array[idx]
    
    return ''.join(result[1:])  # Remove '$'

# COMPRESSION UTILITY

class CompressionAnalyzer:
    """Analyze compression performance"""
    
    @staticmethod
    def compression_ratio(original_size, compressed_size):
        """Calculate compression ratio"""
        return (1 - compressed_size / original_size) * 100 if original_size > 0 else 0
    
    @staticmethod
    def compare_algorithms(data):
        """Compare different compression algorithms"""
        import sys
        
        algorithms = {
            'RLE': (rle_encode, rle_decode),
            'LZ77': (lz77_encode, lz77_decode),
            'LZW': (lzw_encode, lzw_decode),
        }
        
        results = {}
        original_size = len(data)
        
        for name, (encode, decode) in algorithms.items():
            try:
                compressed = encode(data)
                
                # Estimate compressed size
                if isinstance(compressed, str):
                    compressed_size = len(compressed)
                else:
                    compressed_size = sys.getsizeof(compressed)
                
                # Verify decompression
                decompressed = decode(compressed)
                correct = (decompressed == data)
                
                results[name] = {
                    'compressed_size': compressed_size,
                    'ratio': CompressionAnalyzer.compression_ratio(original_size, compressed_size),
                    'correct': correct
                }
            except Exception as e:
                results[name] = {'error': str(e)}
        
        return results
```

### Test Cases
- Highly repetitive data (good for RLE)
- Natural text (good for LZW)
- Binary data
- Already compressed data (negative compression)
- Various file types

### Deliverables
- `rle.{py,go,zig}` - Run-length encoding
- `lz77.{py,go,zig}` - LZ77 implementation
- `lzw.{py,go,zig}` - LZW implementation
- `burrows_wheeler.{py,go,zig}` - BWT
- `compression_utility.{py,go,zig}` - File compression tool
- `compression_analysis.md` - Algorithm comparison
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All algorithms compress and decompress correctly
- Compression ratios measured accurately
- Understanding of when each algorithm works best
- Practical compression tool works on files

---

## Lab 11.4: Sequence Alignment & Edit Distance

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement sequence alignment algorithms
- Understand edit distance variants
- Apply to bioinformatics problems
- Optimize alignment algorithms

### Requirements

1. **Implement alignment algorithms**:
   - **Edit distance** (Levenshtein)
   - **Needleman-Wunsch** (global alignment)
   - **Smith-Waterman** (local alignment)
   - **Longest Common Subsequence**

2. **Variants**:
   - Different scoring schemes
   - Gap penalties
   - Affine gap costs

3. **Applications**:
   - DNA sequence alignment
   - Spell checking
   - Diff utilities

### Implementation Details

```python
# EDIT DISTANCE (Review from Module 8)

def edit_distance(s1, s2):
    """
    Levenshtein distance: minimum edit operations.
    Operations: insert, delete, substitute (each cost 1).
    """
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1]
            else:
                dp[i][j] = 1 + min(
                    dp[i-1][j],      # Delete
                    dp[i][j-1],      # Insert
                    dp[i-1][j-1]     # Substitute
                )
    
    return dp[m][n]

# NEEDLEMAN-WUNSCH (Global Alignment)

def needleman_wunsch(seq1, seq2, match=1, mismatch=-1, gap=-1):
    """
    Global sequence alignment.
    Used in bioinformatics for DNA/protein alignment.
    """
    m, n = len(seq1), len(seq2)
    
    # Initialize DP table
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize first row and column with gap penalties
    for i in range(1, m + 1):
        dp[i][0] = i * gap
    for j in range(1, n + 1):
        dp[0][j] = j * gap
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            
            dp[i][j] = max(
                dp[i-1][j-1] + score_match,  # Match/mismatch
                dp[i-1][j] + gap,             # Gap in seq2
                dp[i][j-1] + gap              # Gap in seq1
            )
    
    # Traceback to get alignment
    align1, align2 = [], []
    i, j = m, n
    
    while i > 0 or j > 0:
        if i > 0 and j > 0:
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            if dp[i][j] == dp[i-1][j-1] + score_match:
                align1.append(seq1[i-1])
                align2.append(seq2[j-1])
                i -= 1
                j -= 1
                continue
        
        if i > 0 and dp[i][j] == dp[i-1][j] + gap:
            align1.append(seq1[i-1])
            align2.append('-')
            i -= 1
        else:
            align1.append('-')
            align2.append(seq2[j-1])
            j -= 1
    
    return ''.join(reversed(align1)), ''.join(reversed(align2)), dp[m][n]

# SMITH-WATERMAN (Local Alignment)

def smith_waterman(seq1, seq2, match=2, mismatch=-1, gap=-1):
    """
    Local sequence alignment: find best matching substring.
    Used for finding similar regions in sequences.
    """
    m, n = len(seq1), len(seq2)
    
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    max_score = 0
    max_pos = (0, 0)
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            
            dp[i][j] = max(
                0,                              # Start new alignment
                dp[i-1][j-1] + score_match,    # Match/mismatch
                dp[i-1][j] + gap,               # Gap in seq2
                dp[i][j-1] + gap                # Gap in seq1
            )
            
            if dp[i][j] > max_score:
                max_score = dp[i][j]
                max_pos = (i, j)
    
    # Traceback from max position
    align1, align2 = [], []
    i, j = max_pos
    
    while i > 0 and j > 0 and dp[i][j] > 0:
        if seq1[i-1] == seq2[j-1]:
            score_match = match
        else:
            score_match = mismatch
        
        if dp[i][j] == dp[i-1][j-1] + score_match:
            align1.append(seq1[i-1])
            align2.append(seq2[j-1])
            i -= 1
            j -= 1
        elif dp[i][j] == dp[i-1][j] + gap:
            align1.append(seq1[i-1])
            align2.append('-')
            i -= 1
        else:
            align1.append('-')
            align2.append(seq2[j-1])
            j -= 1
    
    return ''.join(reversed(align1)), ''.join(reversed(align2)), max_score

# AFFINE GAP PENALTIES

def affine_gap_alignment(seq1, seq2, match=1, mismatch=-1, gap_open=-2, gap_extend=-1):
    """
    Sequence alignment with affine gap penalties.
    Gap cost = gap_open + k * gap_extend for gap of length k.
    More biologically realistic.
    """
    m, n = len(seq1), len(seq2)
    
    # Three DP tables
    M = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Match/mismatch
    X = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Gap in seq1
    Y = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Gap in seq2
    
    # Initialize
    M[0][0] = 0
    for i in range(1, m + 1):
        Y[i][0] = gap_open + i * gap_extend
    for j in range(1, n + 1):
        X[0][j] = gap_open + j * gap_extend
    
    # Fill tables
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score = match if seq1[i-1] == seq2[j-1] else mismatch
            
            M[i][j] = score + max(M[i-1][j-1], X[i-1][j-1], Y[i-1][j-1])
            X[i][j] = max(
                M[i][j-1] + gap_open + gap_extend,
                X[i][j-1] + gap_extend
            )
            Y[i][j] = max(
                M[i-1][j] + gap_open + gap_extend,
                Y[i-1][j] + gap_extend
            )
    
    return max(M[m][n], X[m][n], Y[m][n])

# DIFF UTILITY

def diff_lines(file1_lines, file2_lines):
    """
    Compute diff between two files using LCS.
    Similar to Unix diff command.
    """
    m, n = len(file1_lines), len(file2_lines)
    
    # Compute LCS
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if file1_lines[i-1] == file2_lines[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    # Traceback to generate diff
    diff = []
    i, j = m, n
    
    while i > 0 or j > 0:
        if i > 0 and j > 0 and file1_lines[i-1] == file2_lines[j-1]:
            diff.append(('  ', file1_lines[i-1]))
            i -= 1
            j -= 1
        elif j > 0 and (i == 0 or dp[i][j-1] >= dp[i-1][j]):
            diff.append(('+', file2_lines[j-1]))
            j -= 1
        else:
            diff.append(('-', file1_lines[i-1]))
            i -= 1
    
    return list(reversed(diff))
```

### Test Cases
- DNA sequences
- Protein sequences
- Similar strings with typos
- File comparison
- Performance on long sequences

### Deliverables
- `edit_distance.{py,go,zig}` - Various edit distances
- `needleman_wunsch.{py,go,zig}` - Global alignment
- `smith_waterman.{py,go,zig}` - Local alignment
- `affine_gaps.{py,go,zig}` - Affine gap penalties
- `diff_utility.{py,go,zig}` - File diff tool
- `bioinformatics_applications.md` - DNA/protein analysis
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All alignment algorithms produce correct results
- Understanding of scoring schemes
- Can apply to real biological sequences
- Diff utility matches expected output

---

## Lab 11.5: Advanced String Problems

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Solve challenging string problems
- Apply multiple string techniques
- Build problem-solving skills

### Requirements

1. **Solve advanced problems**:
   - Aho-Corasick multi-pattern matching
   - Palindrome problems
   - String to string transformations
   - Cyclic string problems

2. **Combine techniques**:
   - Use multiple algorithms together
   - Choose appropriate technique for problem

### Implementation Details

```python
# AHO-CORASICK ALGORITHM

class AhoCorasickNode:
    def __init__(self):
        self.children = {}
        self.fail = None
        self.output = []

class AhoCorasick:
    """
    Aho-Corasick algorithm for multi-pattern matching.
    Finds all occurrences of multiple patterns in O(n + m + z) time.
    n = text length, m = total pattern length, z = matches
    """
    def __init__(self, patterns):
        self.root = AhoCorasickNode()
        self.build(patterns)
    
    def build(self, patterns):
        """Build trie and failure links"""
        # Build trie
        for pattern in patterns:
            node = self.root
            for char in pattern:
                if char not in node.children:
                    node.children[char] = AhoCorasickNode()
                node = node.children[char]
            node.output.append(pattern)
        
        # Build failure links (BFS)
        from collections import deque
        queue = deque()
        
        for child in self.root.children.values():
            child.fail = self.root
            queue.append(child)
        
        while queue:
            node = queue.popleft()
            
            for char, child in node.children.items():
                queue.append(child)
                
                # Find failure link
                fail = node.fail
                while fail and char not in fail.children:
                    fail = fail.fail
                
                child.fail = fail.children[char] if fail and char in fail.children else self.root
                child.output += child.fail.output
    
    def search(self, text):
        """Find all pattern occurrences in text"""
        matches = []
        node = self.root
        
        for i, char in enumerate(text):
            while node and char not in node.children:
                node = node.fail
            
            if node is None:
                node = self.root
                continue
            
            node = node.children[char]
            
            for pattern in node.output:
                matches.append((i - len(pattern) + 1, pattern))
        
        return matches

# PALINDROME PROBLEMS

def longest_palindromic_substring(s):
    """
    Find longest palindromic substring.
    Manacher's algorithm: O(n).
    """
    # Transform string to avoid even/odd length issues
    t = '#'.join('^{}$'.format(s))
    n = len(t)
    p = [0] * n  # p[i] = radius of palindrome centered at i
    center = 0
    right = 0
    
    for i in range(1, n - 1):
        # Mirror of i
        mirror = 2 * center - i
        
        if i < right:
            p[i] = min(right - i, p[mirror])
        
        # Expand around i
        while t[i + p[i] + 1] == t[i - p[i] - 1]:
            p[i] += 1
        
        # Update center
        if i + p[i] > right:
            center, right = i, i + p[i]
    
    # Find longest palindrome
    max_len = 0
    center_idx = 0
    for i in range(n):
        if p[i] > max_len:
            max_len = p[i]
            center_idx = i
    
    start = (center_idx - max_len) // 2
    return s[start:start + max_len]

def count_palindromic_substrings(s):
    """Count all palindromic substrings"""
    count = 0
    
    def expand_around_center(left, right):
        nonlocal count
        while left >= 0 and right < len(s) and s[left] == s[right]:
            count += 1
            left -= 1
            right += 1
    
    for i in range(len(s)):
        # Odd length palindromes
        expand_around_center(i, i)
        # Even length palindromes
        expand_around_center(i, i + 1)
    
    return count

# STRING TRANSFORMATIONS

def min_insertions_for_palindrome(s):
    """
    Minimum insertions to make string a palindrome.
    Answer = n - LPS where LPS = longest palindromic subsequence.
    """
    n = len(s)
    
    # LPS using LCS with reverse
    rev_s = s[::-1]
    
    # LCS DP
    dp = [[0] * (n + 1) for _ in range(n + 1)]
    
    for i in range(1, n + 1):
        for j in range(1, n + 1):
            if s[i-1] == rev_s[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    lps = dp[n][n]
    return n - lps

# CYCLIC STRING PROBLEMS

def is_rotation(s1, s2):
    """
    Check if s2 is rotation of s1.
    Trick: s2 is substring of s1 + s1.
    """
    if len(s1) != len(s2):
        return False
    
    return s2 in s1 + s1

def lexicographically_smallest_rotation(s):
    """
    Find lexicographically smallest rotation of string.
    Booth's algorithm: O(n).
    """
    s = s + s
    n = len(s) // 2
    
    f = [-1] * len(s)
    k = 0
    
    for j in range(1, len(s)):
        i = f[j - k - 1]
        while i != -1 and s[j] != s[k + i + 1]:
            if s[j] < s[k + i + 1]:
                k = j - i - 1
            i = f[i]
        
        if i == -1 and s[j] != s[k + i + 1]:
            if s[j] < s[k + i + 1]:
                k = j
            f[j - k] = -1
        else:
            f[j - k] = i + 1
    
    return s[k:k + n]
```

### Test Cases
- Multiple patterns in text
- Palindrome finding
- String transformations
- Cyclic strings

### Deliverables
- `aho_corasick.{py,go,zig}` - Multi-pattern matching
- `palindrome_problems.{py,go,zig}` - Various palindrome algorithms
- `string_transformations.{py,go,zig}` - Transformation problems
- `cyclic_strings.{py,go,zig}` - Rotation problems
- `advanced_problems.{py,go,zig}` - Collection of hard problems
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Aho-Corasick finds all pattern occurrences
- Palindrome algorithms work correctly
- Understanding of advanced techniques
- Can solve complex string problems

---

## Module Resources

### CLRS References
- Chapter 32: String Matching

### Additional Reading
- "Algorithms on Strings, Trees, and Sequences" by Gusfield
- Suffix tree and suffix array papers
- Bioinformatics algorithms texts

### Key Takeaways

By the end of this module, you should:
1. Master string matching algorithms
2. Build and use suffix structures
3. Implement compression algorithms
4. Solve sequence alignment problems
5. Apply string algorithms to real problems
6. Understand computational complexity of string operations

### Next Module

**Module 12: Computational Geometry** - Master geometric algorithms for points, lines, and polygons.

---

## Tips for Success

1. **Visualize patterns** - Draw suffix trees, KMP tables
2. **Test on real data** - DNA sequences, text files
3. **Understand preprocessing** - Most algorithms have this phase
4. **Practice pattern matching** - Core to many problems
5. **Use existing libraries** - But understand internals first

## Common Pitfalls

- **Off-by-one errors** in string indexing
- **Not handling edge cases** (empty strings, single characters)
- **Inefficient naive approaches** - Always optimize
- **Incorrect suffix array construction**
- **Forgetting to handle special characters**
- **Not testing decompression** for compression algorithms
