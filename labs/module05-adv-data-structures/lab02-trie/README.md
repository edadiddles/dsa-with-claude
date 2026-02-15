# Lab 5.2: Trie (Prefix Tree)

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 5.2: Trie (Prefix Tree)

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

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_1.md) (if exists)
- [Next Lab](./lab_5_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 2 of 6*
