# Lab 6.5: Minimum Spanning Trees

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 6.5: Minimum Spanning Trees

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement Kruskal's and Prim's algorithms for MST
- Use Union-Find for Kruskal's
- Understand cut property and MST uniqueness

### Requirements

1. **Implement Kruskal's Algorithm**:
   - Sort edges by weight
   - Use Union-Find to detect cycles
   - Greedy edge selection

2. **Implement Prim's Algorithm**:
   - Start from arbitrary vertex
   - Use min-heap for edge selection
   - Grow single tree

3. **Optimizations**:
   - Prim's with Fibonacci heap (theoretical)
   - Parallel Kruskal's (Borůvka's algorithm)

4. **Applications**:
   - Network design (minimize cable cost)
   - Clustering (remove heaviest edges)
   - Approximation algorithms

### Implementation Details

```python
import heapq
from union_find import UnionFind  # From Module 5

def kruskal_mst(graph):
    """Kruskal's algorithm for MST"""
    # Get all edges and sort by weight
    edges = graph.get_all_edges()
    edges.sort(key=lambda e: e[2])  # Sort by weight
    
    uf = UnionFind(graph.num_vertices)
    mst = []
    total_weight = 0
    
    for u, v, weight in edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, weight))
            total_weight += weight
            
            # Early termination: MST has V-1 edges
            if len(mst) == graph.num_vertices - 1:
                break
    
    return mst, total_weight

def prim_mst(graph, start=0):
    """Prim's algorithm for MST"""
    visited = set()
    min_heap = [(0, start, -1)]  # (weight, vertex, parent)
    mst = []
    total_weight = 0
    
    while min_heap and len(visited) < graph.num_vertices:
        weight, u, parent = heapq.heappop(min_heap)
        
        if u in visited:
            continue
        
        visited.add(u)
        
        if parent != -1:
            mst.append((parent, u, weight))
            total_weight += weight
        
        # Add all edges from u to unvisited vertices
        for v in graph.get_neighbors(u):
            if v not in visited:
                edge_weight = graph.get_weight(u, v)
                heapq.heappush(min_heap, (edge_weight, v, u))
    
    return mst, total_weight

def boruvka_mst(graph):
    """
    Borůvka's algorithm - can be parallelized.
    Faster than Kruskal's for dense graphs.
    """
    uf = UnionFind(graph.num_vertices)
    mst = []
    total_weight = 0
    
    while uf.num_sets > 1:
        # Find cheapest edge for each component
        cheapest = [None] * graph.num_vertices
        
        for u in range(graph.num_vertices):
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                root_u = uf.find(u)
                root_v = uf.find(v)
                
                if root_u != root_v:
                    if cheapest[root_u] is None or weight < cheapest[root_u][2]:
                        cheapest[root_u] = (u, v, weight)
                    if cheapest[root_v] is None or weight < cheapest[root_v][2]:
                        cheapest[root_v] = (v, u, weight)
        
        # Add cheapest edges
        for i in range(graph.num_vertices):
            if cheapest[i] is not None:
                u, v, weight = cheapest[i]
                if uf.find(u) != uf.find(v):
                    uf.union(u, v)
                    mst.append((u, v, weight))
                    total_weight += weight
    
    return mst, total_weight

def verify_mst(graph, mst):
    """Verify that MST is valid"""
    if len(mst) != graph.num_vertices - 1:
        return False, "Wrong number of edges"
    
    # Check connectivity using MST edges
    mst_graph = AdjacencyList(graph.num_vertices, directed=False)
    for u, v, _ in mst:
        mst_graph.add_edge(u, v)
    
    # BFS to check if all vertices reachable
    from collections import deque
    visited = set([0])
    queue = deque([0])
    
    while queue:
        u = queue.popleft()
        for v in mst_graph.get_neighbors(u):
            if v not in visited:
                visited.add(v)
                queue.append(v)
    
    if len(visited) != graph.num_vertices:
        return False, "MST not connected"
    
    # Check no cycles (implicitly satisfied if V-1 edges and connected)
    return True, "Valid MST"

def mst_second_best(graph):
    """Find second-best MST"""
    mst, min_weight = kruskal_mst(graph)
    
    second_best_weight = float('inf')
    second_best_mst = None
    
    # Try removing each MST edge and replacing with another
    for i, (u_remove, v_remove, w_remove) in enumerate(mst):
        # Build MST without this edge
        edges = graph.get_all_edges()
        edges = [(u, v, w) for u, v, w in edges if not (u == u_remove and v == v_remove)]
        edges.sort(key=lambda e: e[2])
        
        uf = UnionFind(graph.num_vertices)
        current_mst = []
        current_weight = 0
        
        for u, v, weight in edges:
            if uf.find(u) != uf.find(v):
                uf.union(u, v)
                current_mst.append((u, v, weight))
                current_weight += weight
                
                if len(current_mst) == graph.num_vertices - 1:
                    break
        
        if len(current_mst) == graph.num_vertices - 1:
            if current_weight < second_best_weight and current_weight > min_weight:
                second_best_weight = current_weight
                second_best_mst = current_mst
    
    return second_best_mst, second_best_weight

def clustering_with_mst(graph, k):
    """
    Use MST for clustering: remove k-1 heaviest edges to get k clusters.
    """
    mst, _ = kruskal_mst(graph)
    
    # Sort MST edges by weight (descending)
    mst.sort(key=lambda e: e[2], reverse=True)
    
    # Remove k-1 heaviest edges
    kept_edges = mst[k-1:]
    
    # Find connected components in remaining graph
    uf = UnionFind(graph.num_vertices)
    for u, v, _ in kept_edges:
        uf.union(u, v)
    
    # Build clusters
    clusters = {}
    for v in range(graph.num_vertices):
        root = uf.find(v)
        if root not in clusters:
            clusters[root] = []
        clusters[root].append(v)
    
    return list(clusters.values())

def maximum_spanning_tree(graph):
    """Find maximum spanning tree (negate weights)"""
    # Create graph with negated weights
    edges = graph.get_all_edges()
    negated_graph = EdgeList(graph.num_vertices, directed=False)
    
    for u, v, weight in edges:
        negated_graph.add_edge(u, v, -weight)
    
    mst, neg_weight = kruskal_mst(negated_graph)
    
    # Restore original weights
    mst_with_orig_weights = [(u, v, -w) for u, v, w in mst]
    return mst_with_orig_weights, -neg_weight
```

### Test Cases
- Complete graphs
- Sparse graphs
- Graphs with equal edge weights
- Graphs with unique weights
- Disconnected graphs
- Large random graphs

### Deliverables
- `kruskal.{py,go,zig}` - Kruskal's algorithm
- `prim.{py,go,zig}` - Prim's algorithm
- `boruvka.{py,go,zig}` - Borůvka's algorithm
- `mst_verification.{py,go,zig}` - Verify MST correctness
- `mst_applications.{py,go,zig}` - Clustering, second-best MST
- `algorithm_comparison.md` - Performance comparison
- `benchmarks.{py,go,zig}` - Benchmark all three algorithms

### Success Criteria
- All algorithms find MST of same weight
- Kruskal's and Prim's produce valid MSTs
- Understanding of when each algorithm is preferable
- Can apply to real network design problems

---

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_4.md) (if exists)
- [Next Lab](./lab_6_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 5 of 6*
