# Lab 11.5: Advanced String Problems

**Module**: Module 11 - String Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 11.5: Advanced String Problems

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

---

## Back to Module

[← Back to Module 11: String Algorithms](../module_11/README.md)

## Navigation

- [Previous Lab](./lab_11_4.md) (if exists)
- [Next Lab](./lab_11_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 11, Lab 5 of 5*
