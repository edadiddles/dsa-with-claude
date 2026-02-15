# Module 5: Advanced Data Structures

**Duration**: 3 weeks  
**Difficulty**: Medium to Hard

## Overview

Specialized data structures solve specific problem domains far more efficiently than general-purpose structures. This module introduces powerful structures that may seem esoteric at first but become indispensable once you recognize their use cases.

**Why this matters**: These structures appear in production systems everywhere - Union-Find in network connectivity and image processing, tries in autocomplete and spell checking, segment trees in databases and competitive programming, Bloom filters in caches and CDNs. Learning these expands your problem-solving toolkit dramatically.

## Learning Objectives

- Implement Union-Find with path compression and union by rank
- Build prefix trees (tries) for string operations
- Master segment trees for range queries
- Understand Fenwick trees and their elegant bit manipulation
- Implement probabilistic structures (skip lists, Bloom filters)
- Recognize when specialized structures dramatically outperform general ones

## Topics Covered

- Disjoint set (Union-Find) with optimizations
- Tries and compressed tries (Patricia trees)
- Segment trees with lazy propagation
- Fenwick trees (Binary Indexed Trees)
- Skip lists as probabilistic alternatives to balanced trees
- Bloom filters and space-efficient probabilistic membership

---

## Lab 5.1: Union-Find (Disjoint Set)

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement Union-Find with optimizations
- Understand amortized analysis and inverse Ackermann function
- Apply to connectivity problems

### Requirements

1. **Implement Union-Find with multiple strategies**:
   - **Quick Find**: O(1) find, O(n) union
   - **Quick Union**: O(n) find, O(n) union
   - **Union by Rank**: O(log n) both operations
   - **Path Compression**: Near-constant amortized time
   - **Both Optimizations**: O(α(n)) amortized (α = inverse Ackermann)

2. **Operations**:
   - `make_set(x)` - Create singleton set
   - `find(x)` - Find representative with path compression
   - `union(x, y)` - Merge sets with union by rank
   - `connected(x, y)` - Check if in same set

3. **Applications**:
   - Detect cycles in undirected graph
   - Connected components
   - Kruskal's MST algorithm
   - Image segmentation (group similar pixels)

### Implementation Details

```python
class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n
        self.num_sets = n
    
    def find(self, x):
        """Find with path compression"""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # Path compression
        return self.parent[x]
    
    def union(self, x, y):
        """Union by rank"""
        root_x = self.find(x)
        root_y = self.find(y)
        
        if root_x == root_y:
            return False  # Already in same set
        
        # Attach smaller tree under larger tree
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1
        
        self.num_sets -= 1
        return True
    
    def connected(self, x, y):
        return self.find(x) == self.find(y)
    
    def count_sets(self):
        return self.num_sets

class UnionFindQuickFind:
    """Quick Find variant - O(1) find, O(n) union"""
    def __init__(self, n):
        self.id = list(range(n))
        self.num_sets = n
    
    def find(self, x):
        return self.id[x]
    
    def union(self, x, y):
        id_x = self.id[x]
        id_y = self.id[y]
        
        if id_x == id_y:
            return False
        
        # Change all elements with id_y to id_x
        for i in range(len(self.id)):
            if self.id[i] == id_y:
                self.id[i] = id_x
        
        self.num_sets -= 1
        return True

def detect_cycle_undirected(edges, n):
    """Detect cycle in undirected graph using Union-Find"""
    uf = UnionFind(n)
    
    for u, v in edges:
        if uf.connected(u, v):
            return True  # Cycle detected
        uf.union(u, v)
    
    return False

def count_connected_components(edges, n):
    """Count connected components"""
    uf = UnionFind(n)
    
    for u, v in edges:
        uf.union(u, v)
    
    return uf.count_sets()

def image_segmentation(image, threshold):
    """Group similar pixels using Union-Find"""
    height, width = len(image), len(image[0])
    uf = UnionFind(height * width)
    
    def pixel_to_id(r, c):
        return r * width + c
    
    # Union adjacent similar pixels
    for r in range(height):
        for c in range(width):
            current_id = pixel_to_id(r, c)
            
            # Check right neighbor
            if c + 1 < width and abs(image[r][c] - image[r][c+1]) <= threshold:
                uf.union(current_id, pixel_to_id(r, c+1))
            
            # Check bottom neighbor
            if r + 1 < height and abs(image[r][c] - image[r+1][c]) <= threshold:
                uf.union(current_id, pixel_to_id(r+1, c))
    
    # Build segments
    segments = {}
    for r in range(height):
        for c in range(width):
            root = uf.find(pixel_to_id(r, c))
            if root not in segments:
                segments[root] = []
            segments[root].append((r, c))
    
    return list(segments.values())
```

### Test Cases
- Union all elements into one set
- Create multiple disjoint sets
- Pathological tree (chain)
- Random union operations
- Verify path compression effect
- Benchmark different variants

### Deliverables
- `union_find.{py,go,zig}` - All variants
- `amortized_analysis.md` - Proof of O(α(n))
- `applications.{py,go,zig}` - Cycle detection, connected components, image segmentation
- `benchmarks.{py,go,zig}` - Compare variants
- `visualization.{py,go,zig}` - Visualize tree structure and compression

### Success Criteria
- Optimized version significantly faster than naive
- Understanding of inverse Ackermann function
- Can apply to graph problems
- Path compression demonstrably improves performance

---

## Lab 5.2: Trie (Prefix Tree)

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement trie for string operations
- Understand space-time tradeoffs
- Build autocomplete and dictionary applications

### Requirements

1. **Implement trie with**:
   - `insert(word)` - Add word to trie
   - `search(word)` - Check if exact word exists
   - `starts_with(prefix)` - Check if prefix exists
   - `delete(word)` - Remove word
   - `get_all_words()` - Return all stored words
   - `count_words()` - Total words in trie

2. **Optimizations**:
   - Compressed trie (Patricia trie)
   - Ternary search tree
   - Array-based vs. hashmap-based children

3. **Applications**:
   - Autocomplete system
   - Spell checker
   - IP routing (longest prefix match)
   - Word games (Boggle solver)

### Implementation Details

```python
class TrieNode:
    def __init__(self):
        self.children = {}  # or array of size 26 for lowercase English
        self.is_end_of_word = False
        self.word_count = 0  # Number of words ending here

class Trie:
    def __init__(self):
        self.root = TrieNode()
        self.total_words = 0
    
    def insert(self, word):
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        
        if not node.is_end_of_word:
            self.total_words += 1
        node.is_end_of_word = True
        node.word_count += 1
    
    def search(self, word):
        node = self._find_node(word)
        return node is not None and node.is_end_of_word
    
    def starts_with(self, prefix):
        return self._find_node(prefix) is not None
    
    def _find_node(self, prefix):
        node = self.root
        for char in prefix:
            if char not in node.children:
                return None
            node = node.children[char]
        return node
    
    def delete(self, word):
        def _delete_helper(node, word, index):
            if index == len(word):
                if not node.is_end_of_word:
                    return False
                node.is_end_of_word = False
                node.word_count -= 1
                return len(node.children) == 0
            
            char = word[index]
            if char not in node.children:
                return False
            
            should_delete_child = _delete_helper(node.children[char], word, index + 1)
            
            if should_delete_child:
                del node.children[char]
                return len(node.children) == 0 and not node.is_end_of_word
            
            return False
        
        if _delete_helper(self.root, word, 0):
            self.total_words -= 1
    
    def autocomplete(self, prefix, max_results=10):
        node = self._find_node(prefix)
        if node is None:
            return []
        
        results = []
        self._dfs_collect_words(node, prefix, results, max_results)
        return results
    
    def _dfs_collect_words(self, node, current_word, results, max_results):
        if len(results) >= max_results:
            return
        
        if node.is_end_of_word:
            results.append(current_word)
        
        for char, child in sorted(node.children.items()):
            self._dfs_collect_words(child, current_word + char, results, max_results)
    
    def get_all_words(self):
        return self.autocomplete("", max_results=float('inf'))
    
    def count_words(self):
        return self.total_words

class CompressedTrie:
    """Patricia trie - stores edge labels as strings"""
    class Node:
        def __init__(self):
            self.children = {}  # edge_label -> Node
            self.is_end = False
    
    def __init__(self):
        self.root = self.Node()
    
    def insert(self, word):
        node = self.root
        i = 0
        
        while i < len(word):
            found = False
            for edge_label in node.children:
                # Find common prefix
                j = 0
                while j < len(edge_label) and i + j < len(word) and edge_label[j] == word[i + j]:
                    j += 1
                
                if j > 0:
                    if j == len(edge_label):
                        # Full edge match, continue
                        node = node.children[edge_label]
                        i += j
                        found = True
                        break
                    else:
                        # Partial match, split edge
                        self._split_edge(node, edge_label, j, word[i:])
                        return
            
            if not found:
                # No matching edge, create new
                node.children[word[i:]] = self.Node()
                node.children[word[i:]].is_end = True
                return
        
        node.is_end = True

class TernarySearchTree:
    """Space-efficient alternative to trie"""
    class Node:
        def __init__(self, char):
            self.char = char
            self.is_end = False
            self.left = None   # Less than
            self.middle = None # Equal (next char)
            self.right = None  # Greater than
    
    def __init__(self):
        self.root = None
    
    def insert(self, word):
        self.root = self._insert(self.root, word, 0)
    
    def _insert(self, node, word, index):
        if index >= len(word):
            return node
        
        char = word[index]
        
        if node is None:
            node = self.Node(char)
        
        if char < node.char:
            node.left = self._insert(node.left, word, index)
        elif char > node.char:
            node.right = self._insert(node.right, word, index)
        else:
            if index == len(word) - 1:
                node.is_end = True
            else:
                node.middle = self._insert(node.middle, word, index + 1)
        
        return node

def spell_checker(dictionary, word):
    """Suggest corrections for misspelled word"""
    trie = Trie()
    for w in dictionary:
        trie.insert(w)
    
    suggestions = []
    
    # Check exact match
    if trie.search(word):
        return [word]
    
    # Generate candidates (edit distance 1)
    candidates = set()
    
    # Deletions
    for i in range(len(word)):
        candidates.add(word[:i] + word[i+1:])
    
    # Insertions
    for i in range(len(word) + 1):
        for c in 'abcdefghijklmnopqrstuvwxyz':
            candidates.add(word[:i] + c + word[i:])
    
    # Replacements
    for i in range(len(word)):
        for c in 'abcdefghijklmnopqrstuvwxyz':
            candidates.add(word[:i] + c + word[i+1:])
    
    # Transpositions
    for i in range(len(word) - 1):
        candidates.add(word[:i] + word[i+1] + word[i] + word[i+2:])
    
    # Check which candidates are valid
    for candidate in candidates:
        if trie.search(candidate):
            suggestions.append(candidate)
    
    return suggestions[:10]

def boggle_solver(board, dictionary):
    """Find all words in Boggle board"""
    trie = Trie()
    for word in dictionary:
        trie.insert(word)
    
    rows, cols = len(board), len(board[0])
    found_words = set()
    
    def dfs(r, c, node, path, visited):
        if r < 0 or r >= rows or c < 0 or c >= cols or (r, c) in visited:
            return
        
        char = board[r][c]
        if char not in node.children:
            return
        
        next_node = node.children[char]
        path += char
        visited.add((r, c))
        
        if next_node.is_end_of_word and len(path) >= 3:
            found_words.add(path)
        
        # Explore all 8 directions
        for dr in [-1, 0, 1]:
            for dc in [-1, 0, 1]:
                if dr != 0 or dc != 0:
                    dfs(r + dr, c + dc, next_node, path, visited.copy())
    
    # Try starting from each cell
    for r in range(rows):
        for c in range(cols):
            dfs(r, c, trie.root, "", set())
    
    return sorted(found_words)
```

### Test Cases
- Insert dictionary of 100K words
- Search for existing and non-existing words
- Autocomplete with various prefixes
- Delete words and verify trie structure
- Memory usage analysis
- Spell checker with real misspellings

### Deliverables
- `trie.{py,go,zig}` - Standard trie
- `compressed_trie.{py,go,zig}` - Patricia trie
- `ternary_search_tree.{py,go,zig}` - TST implementation
- `autocomplete.{py,go,zig}` - Full autocomplete system
- `spell_checker.{py,go,zig}` - Spell checking application
- `boggle_solver.{py,go,zig}` - Word game solver
- `memory_comparison.md` - Space analysis
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Fast prefix searches
- Autocomplete works efficiently
- Understanding of trie space usage
- Can choose between trie variants

---

## Lab 5.3: Segment Tree

**Duration**: 5-6 hours  
**Difficulty**: Hard

### Objectives
- Implement segment tree for range queries
- Support point updates and range updates
- Understand lazy propagation

### Requirements

1. **Implement segment tree supporting**:
   - `build(array)` - O(n) construction
   - `query(left, right)` - O(log n) range query
   - `update(index, value)` - O(log n) point update
   - Operations: sum, min, max, GCD, etc.

2. **Implement lazy propagation for**:
   - Range updates: add value to range
   - Range set: set all values in range
   - More complex updates

3. **Additional features**:
   - Persistent segment tree
   - 2D segment tree (for matrix range queries)
   - Dynamic segment tree (for large coordinate spaces)

### Implementation Details

```python
class SegmentTree:
    def __init__(self, array, operation='sum'):
        self.n = len(array)
        self.tree = [0] * (4 * self.n)
        self.operation = operation
        self.identity = self._get_identity()
        self._build(array, 0, 0, self.n - 1)
    
    def _get_identity(self):
        if self.operation == 'sum':
            return 0
        elif self.operation == 'min':
            return float('inf')
        elif self.operation == 'max':
            return float('-inf')
        elif self.operation == 'gcd':
            return 0
    
    def _combine(self, left_val, right_val):
        if self.operation == 'sum':
            return left_val + right_val
        elif self.operation == 'min':
            return min(left_val, right_val)
        elif self.operation == 'max':
            return max(left_val, right_val)
        elif self.operation == 'gcd':
            import math
            return math.gcd(left_val, right_val)
    
    def _build(self, array, node, start, end):
        if start == end:
            self.tree[node] = array[start]
        else:
            mid = (start + end) // 2
            left_child = 2 * node + 1
            right_child = 2 * node + 2
            
            self._build(array, left_child, start, mid)
            self._build(array, right_child, mid + 1, end)
            
            self.tree[node] = self._combine(
                self.tree[left_child],
                self.tree[right_child]
            )
    
    def query(self, left, right):
        return self._query(0, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        # No overlap
        if start > right or end < left:
            return self.identity
        
        # Complete overlap
        if start >= left and end <= right:
            return self.tree[node]
        
        # Partial overlap
        mid = (start + end) // 2
        left_child = 2 * node + 1
        right_child = 2 * node + 2
        
        left_result = self._query(left_child, start, mid, left, right)
        right_result = self._query(right_child, mid + 1, end, left, right)
        
        return self._combine(left_result, right_result)
    
    def update(self, index, value):
        self._update(0, 0, self.n - 1, index, value)
    
    def _update(self, node, start, end, index, value):
        if start == end:
            self.tree[node] = value
        else:
            mid = (start + end) // 2
            left_child = 2 * node + 1
            right_child = 2 * node + 2
            
            if index <= mid:
                self._update(left_child, start, mid, index, value)
            else:
                self._update(right_child, mid + 1, end, index, value)
            
            self.tree[node] = self._combine(
                self.tree[left_child],
                self.tree[right_child]
            )

class LazySegmentTree:
    """Segment tree with lazy propagation for range updates"""
    def __init__(self, array):
        self.n = len(array)
        self.tree = [0] * (4 * self.n)
        self.lazy = [0] * (4 * self.n)
        self._build(array, 0, 0, self.n - 1)
    
    def _build(self, array, node, start, end):
        if start == end:
            self.tree[node] = array[start]
        else:
            mid = (start + end) // 2
            self._build(array, 2*node+1, start, mid)
            self._build(array, 2*node+2, mid+1, end)
            self.tree[node] = self.tree[2*node+1] + self.tree[2*node+2]
    
    def _push(self, node, start, end):
        """Push lazy updates down"""
        if self.lazy[node] != 0:
            self.tree[node] += (end - start + 1) * self.lazy[node]
            
            if start != end:
                self.lazy[2*node+1] += self.lazy[node]
                self.lazy[2*node+2] += self.lazy[node]
            
            self.lazy[node] = 0
    
    def range_update(self, left, right, value):
        """Add value to all elements in [left, right]"""
        self._range_update(0, 0, self.n - 1, left, right, value)
    
    def _range_update(self, node, start, end, left, right, value):
        self._push(node, start, end)
        
        if start > right or end < left:
            return
        
        if start >= left and end <= right:
            self.lazy[node] += value
            self._push(node, start, end)
            return
        
        mid = (start + end) // 2
        self._range_update(2*node+1, start, mid, left, right, value)
        self._range_update(2*node+2, mid+1, end, left, right, value)
        
        self._push(2*node+1, start, mid)
        self._push(2*node+2, mid+1, end)
        self.tree[node] = self.tree[2*node+1] + self.tree[2*node+2]
    
    def query(self, left, right):
        return self._query(0, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        if start > right or end < left:
            return 0
        
        self._push(node, start, end)
        
        if start >= left and end <= right:
            return self.tree[node]
        
        mid = (start + end) // 2
        return (self._query(2*node+1, start, mid, left, right) +
                self._query(2*node+2, mid+1, end, left, right))

class SegmentTree2D:
    """2D segment tree for matrix range queries"""
    def __init__(self, matrix):
        self.rows = len(matrix)
        self.cols = len(matrix[0])
        self.tree = [[0] * (4 * self.cols) for _ in range(4 * self.rows)]
        self._build_rows(matrix, 0, 0, self.rows - 1)
    
    def _build_rows(self, matrix, node, start, end):
        if start == end:
            self._build_cols(matrix[start], node, 0, 0, self.cols - 1)
        else:
            mid = (start + end) // 2
            self._build_rows(matrix, 2*node+1, start, mid)
            self._build_rows(matrix, 2*node+2, mid+1, end)
            
            # Merge column segment trees
            for col_node in range(4 * self.cols):
                self.tree[node][col_node] = (
                    self.tree[2*node+1][col_node] + 
                    self.tree[2*node+2][col_node]
                )
    
    def _build_cols(self, row, row_node, col_node, start, end):
        if start == end:
            self.tree[row_node][col_node] = row[start]
        else:
            mid = (start + end) // 2
            self._build_cols(row, row_node, 2*col_node+1, start, mid)
            self._build_cols(row, row_node, 2*col_node+2, mid+1, end)
            self.tree[row_node][col_node] = (
                self.tree[row_node][2*col_node+1] + 
                self.tree[row_node][2*col_node+2]
            )
```

### Test Cases
- Range sum queries on random arrays
- Verify against naive O(n) queries
- Point updates and re-query
- Range updates with lazy propagation
- 2D queries on matrices
- Stress test with 100K operations

### Deliverables
- `segment_tree.{py,go,zig}` - Basic implementation
- `lazy_segment_tree.{py,go,zig}` - With lazy propagation
- `segment_tree_2d.{py,go,zig}` - 2D version
- `applications.{py,go,zig}` - Range queries in practice
- `complexity_analysis.md` - Proof of O(log n) operations
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct range queries and updates
- Lazy propagation works efficiently
- Understanding of when segment trees are needed
- Can extend to other operations (min, max, GCD)

---

## Lab 5.4: Fenwick Tree (Binary Indexed Tree)

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement Fenwick tree for prefix sums
- Understand bit manipulation tricks
- Compare with segment tree

### Requirements

1. **Implement Fenwick tree with**:
   - `build(array)` - O(n) construction
   - `prefix_sum(index)` - Sum of elements [0, index]
   - `range_sum(left, right)` - Sum of elements [left, right]
   - `update(index, delta)` - Add delta to element at index

2. **Understand the bit manipulation**:
   - `i & (-i)` gives lowest set bit
   - Parent: `i + (i & -i)`
   - Next: `i - (i & -i)`

3. **Extensions**:
   - 2D Fenwick tree (for matrix prefix sums)
   - Range update, point query variant
   - Binary search on Fenwick tree

### Implementation Details

```python
class FenwickTree:
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (n + 1)  # 1-indexed
    
    def update(self, index, delta):
        """Add delta to element at index (0-indexed input)"""
        index += 1  # Convert to 1-indexed
        while index <= self.n:
            self.tree[index] += delta
            index += index & (-index)  # Move to parent
    
    def prefix_sum(self, index):
        """Sum of elements [0, index] (0-indexed input)"""
        index += 1  # Convert to 1-indexed
        result = 0
        while index > 0:
            result += self.tree[index]
            index -= index & (-index)  # Move to next node
        return result
    
    def range_sum(self, left, right):
        """Sum of elements [left, right] (0-indexed)"""
        if left == 0:
            return self.prefix_sum(right)
        return self.prefix_sum(right) - self.prefix_sum(left - 1)
    
    def build(self, array):
        """Build from array in O(n)"""
        for i, val in enumerate(array):
            self.update(i, val)
    
    def build_optimized(self, array):
        """O(n) build by propagating upward"""
        self.tree = [0] + array[:]  # 1-indexed with values
        
        for i in range(1, self.n + 1):
            parent = i + (i & -i)
            if parent <= self.n:
                self.tree[parent] += self.tree[i]

class FenwickTree2D:
    def __init__(self, rows, cols):
        self.rows = rows
        self.cols = cols
        self.tree = [[0] * (cols + 1) for _ in range(rows + 1)]
    
    def update(self, row, col, delta):
        """Add delta to element at (row, col)"""
        row += 1
        col += 1
        
        i = row
        while i <= self.rows:
            j = col
            while j <= self.cols:
                self.tree[i][j] += delta
                j += j & (-j)
            i += i & (-i)
    
    def prefix_sum(self, row, col):
        """Sum of rectangle from (0,0) to (row, col)"""
        row += 1
        col += 1
        result = 0
        
        i = row
        while i > 0:
            j = col
            while j > 0:
                result += self.tree[i][j]
                j -= j & (-j)
            i -= i & (-i)
        
        return result
    
    def range_sum(self, r1, c1, r2, c2):
        """Sum of rectangle from (r1,c1) to (r2,c2)"""
        return (self.prefix_sum(r2, c2) - 
                self.prefix_sum(r1-1, c2) - 
                self.prefix_sum(r2, c1-1) + 
                self.prefix_sum(r1-1, c1-1))

class RangeUpdateFenwick:
    """Fenwick tree for range updates and point queries"""
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (n + 1)
    
    def range_update(self, left, right, delta):
        """Add delta to all elements in [left, right]"""
        self._update(left, delta)
        self._update(right + 1, -delta)
    
    def _update(self, index, delta):
        index += 1
        while index <= self.n:
            self.tree[index] += delta
            index += index & (-index)
    
    def point_query(self, index):
        """Get value at index"""
        index += 1
        result = 0
        while index > 0:
            result += self.tree[index]
            index -= index & (-index)
        return result

def binary_search_fenwick(fenwick, target):
    """Find smallest index where prefix_sum >= target"""
    # Works if all elements are non-negative
    index = 0
    power = 1
    
    # Find highest power of 2 <= n
    while power * 2 <= fenwick.n:
        power *= 2
    
    current_sum = 0
    
    while power > 0:
        if index + power <= fenwick.n and current_sum + fenwick.tree[index + power] < target:
            index += power
            current_sum += fenwick.tree[index]
        power //= 2
    
    return index  # 0-indexed result
```

### Test Cases
- Build from array, verify prefix sums
- Random updates and queries
- Compare with naive O(n) prefix sum
- 2D matrix operations
- Verify correctness against segment tree
- Performance benchmarks

### Deliverables
- `fenwick_tree.{py,go,zig}` - 1D implementation
- `fenwick_tree_2d.{py,go,zig}` - 2D implementation
- `range_update_fenwick.{py,go,zig}` - Range update variant
- `bit_manipulation_explanation.md` - Explain the magic
- `fenwick_vs_segment.md` - Comparison
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct prefix sum and update operations
- Understanding of bit manipulation
- Knowledge of when Fenwick tree is preferable to segment tree (simpler, less memory)
- Can explain the `i & (-i)` trick

---

## Lab 5.5: Skip List

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement probabilistic balanced structure
- Understand randomized data structures
- Compare with balanced BST

### Requirements

1. **Implement skip list with**:
   - `insert(value)` - O(log n) expected
   - `search(value)` - O(log n) expected
   - `delete(value)` - O(log n) expected
   - Random level generation (p = 0.5 typically)

2. **Features**:
   - Configurable probability p
   - Maximum level limit
   - Iterator support (in-order traversal)

3. **Analysis**:
   - Measure actual height distribution
   - Compare search cost with BST
   - Study effect of different p values

### Implementation Details

```python
import random

class SkipNode:
    def __init__(self, value, level):
        self.value = value
        self.forward = [None] * (level + 1)

class SkipList:
    def __init__(self, max_level=16, p=0.5):
        self.max_level = max_level
        self.p = p
        self.header = SkipNode(None, max_level)
        self.level = 0
    
    def random_level(self):
        level = 0
        while random.random() < self.p and level < self.max_level:
            level += 1
        return level
    
    def search(self, value):
        current = self.header
        
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
        
        current = current.forward[0]
        return current and current.value == value
    
    def insert(self, value):
        update = [None] * (self.max_level + 1)
        current = self.header
        
        # Find position to insert
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
            update[i] = current
        
        current = current.forward[0]
        
        # Check if already exists
        if current and current.value == value:
            return  # Duplicate
        
        # Generate random level for new node
        new_level = self.random_level()
        
        if new_level > self.level:
            for i in range(self.level + 1, new_level + 1):
                update[i] = self.header
            self.level = new_level
        
        new_node = SkipNode(value, new_level)
        
        # Insert node
        for i in range(new_level + 1):
            new_node.forward[i] = update[i].forward[i]
            update[i].forward[i] = new_node
    
    def delete(self, value):
        update = [None] * (self.max_level + 1)
        current = self.header
        
        # Find position
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
            update[i] = current
        
        current = current.forward[0]
        
        if current and current.value == value:
            # Remove node
            for i in range(self.level + 1):
                if update[i].forward[i] != current:
                    break
                update[i].forward[i] = current.forward[i]
            
            # Update level
            while self.level > 0 and self.header.forward[self.level] is None:
                self.level -= 1
    
    def __iter__(self):
        current = self.header.forward[0]
        while current:
            yield current.value
            current = current.forward[0]
    
    def range_query(self, start, end):
        """Find all values in [start, end]"""
        result = []
        current = self.header
        
        # Navigate to start
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < start:
                current = current.forward[i]
        
        current = current.forward[0]
        
        # Collect values in range
        while current and current.value <= end:
            result.append(current.value)
            current = current.forward[0]
        
        return result
    
    def get_height(self):
        return self.level + 1
    
    def measure_level_distribution(self):
        """Measure how many nodes at each level"""
        counts = [0] * (self.level + 1)
        current = self.header.forward[0]
        
        while current:
            for i in range(len(current.forward)):
                if current.forward[i] is not None or i == 0:
                    counts[i] += 1
            current = current.forward[0]
        
        return counts

def analyze_skip_list_performance(p_values, sizes):
    """Analyze performance with different p values"""
    results = {}
    
    for p in p_values:
        results[p] = {}
        for size in sizes:
            skip_list = SkipList(p=p)
            values = list(range(size))
            random.shuffle(values)
            
            # Measure insertion time
            import time
            start = time.time()
            for v in values:
                skip_list.insert(v)
            insert_time = time.time() - start
            
            # Measure search time
            start = time.time()
            for v in values:
                skip_list.search(v)
            search_time = time.time() - start
            
            results[p][size] = {
                'height': skip_list.get_height(),
                'insert_time': insert_time,
                'search_time': search_time,
                'level_dist': skip_list.measure_level_distribution()
            }
    
    return results
```

### Test Cases
- Insert random values, verify search
- Delete values, verify structure
- Measure level distribution
- Compare with AVL/Red-Black tree
- Test different p values (0.25, 0.5, 0.75)

### Deliverables
- `skip_list.{py,go,zig}` - Full implementation
- `probability_analysis.md` - Study of p values
- `height_distribution.{py,go,zig}` - Measure actual heights
- `comparison.md` - Skip list vs. BST
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct operations
- Understanding of probabilistic guarantees
- Knowledge of skip list advantages (simpler than balanced BST, good cache locality)
- Can explain expected height formula

---

## Lab 5.6: Bloom Filter

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

## Module Resources

### CLRS References
- Chapter 21: Data Structures for Disjoint Sets (Union-Find)
- Various advanced sources for other structures

### Additional Reading
- "Advanced Data Structures" by Brass
- Original papers: Bloom filter, skip list
- Patricia trie and suffix tree papers

### Key Takeaways

By the end of this module, you should:
1. Recognize when specialized structures dramatically outperform general ones
2. Understand amortized analysis for Union-Find
3. Know when to use tries for string problems
4. Master range query structures (segment tree, Fenwick tree)
5. Appreciate probabilistic data structures

### Next Module

**Module 6: Graph Algorithms I - Fundamentals** - You'll implement graph representations and core graph algorithms like BFS, DFS, topological sort, and minimum spanning trees.

---

## Tips for Success

1. **Visualize the structures** - They're often tree-based or array-based
2. **Understand the invariants** - Each structure maintains specific properties
3. **Test edge cases** - Empty structures, single elements, duplicates
4. **Measure performance** - Verify theoretical complexity empirically
5. **Build intuition** - When would you reach for each structure?

## Common Pitfalls

- **Union-Find**: Forgetting path compression or union by rank
- **Trie**: Not handling deletion correctly (need to check if node can be removed)
- **Segment Tree**: Off-by-one errors in range calculations
- **Fenwick Tree**: Forgetting 1-indexed arrays
- **Skip List**: Not understanding probabilistic height distribution
- **Bloom Filter**: Choosing poor hash functions leading to correlations
