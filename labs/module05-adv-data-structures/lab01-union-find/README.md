# Lab 5.1: Union-Find (Disjoint Set)

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 5.1: Union-Find (Disjoint Set)

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

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_0.md) (if exists)
- [Next Lab](./lab_5_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 1 of 6*
