# Lab 6.3: Topological Sort

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 6.3: Topological Sort

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement topological sorting for DAGs
- Detect cycles in directed graphs
- Apply to dependency resolution problems

### Requirements

1. **Implement two algorithms**:
   - **Kahn's algorithm**: BFS-based using in-degrees
   - **DFS-based**: Using finish times from DFS

2. **Cycle detection** for directed graphs

3. **Applications**:
   - Course prerequisite ordering
   - Build system dependencies
   - Task scheduling
   - Package manager dependency resolution

### Implementation Details

```python
from collections import deque

def topological_sort_kahn(graph):
    """Kahn's algorithm for topological sorting"""
    # Calculate in-degrees
    in_degree = [0] * graph.num_vertices
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            in_degree[v] += 1
    
    # Start with vertices having in-degree 0
    queue = deque([v for v in range(graph.num_vertices) if in_degree[v] == 0])
    topo_order = []
    
    while queue:
        u = queue.popleft()
        topo_order.append(u)
        
        # Remove this vertex from graph (conceptually)
        for v in graph.get_neighbors(u):
            in_degree[v] -= 1
            if in_degree[v] == 0:
                queue.append(v)
    
    # If not all vertices processed, there's a cycle
    if len(topo_order) != graph.num_vertices:
        return None  # Cycle detected
    
    return topo_order

def topological_sort_dfs(graph):
    """DFS-based topological sorting"""
    visited = set()
    rec_stack = set()  # For cycle detection
    topo_order = []
    
    def dfs(v):
        visited.add(v)
        rec_stack.add(v)
        
        for neighbor in graph.get_neighbors(v):
            if neighbor not in visited:
                if not dfs(neighbor):
                    return False  # Cycle detected
            elif neighbor in rec_stack:
                return False  # Back edge (cycle)
        
        rec_stack.remove(v)
        topo_order.append(v)
        return True
    
    for vertex in range(graph.num_vertices):
        if vertex not in visited:
            if not dfs(vertex):
                return None  # Cycle detected
    
    return list(reversed(topo_order))

def all_topological_sorts(graph):
    """Generate all possible topological orderings"""
    in_degree = [0] * graph.num_vertices
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            in_degree[v] += 1
    
    result = []
    
    def backtrack(current_order, remaining_in_degree):
        if len(current_order) == graph.num_vertices:
            result.append(current_order[:])
            return
        
        # Try all vertices with in-degree 0
        for v in range(graph.num_vertices):
            if remaining_in_degree[v] == 0 and v not in current_order:
                # Temporarily use this vertex
                new_in_degree = remaining_in_degree[:]
                new_in_degree[v] = -1  # Mark as used
                
                # Decrease in-degree of neighbors
                for neighbor in graph.get_neighbors(v):
                    new_in_degree[neighbor] -= 1
                
                current_order.append(v)
                backtrack(current_order, new_in_degree)
                current_order.pop()
    
    backtrack([], in_degree[:])
    return result

def has_cycle_directed(graph):
    """Check if directed graph has a cycle"""
    WHITE, GRAY, BLACK = 0, 1, 2
    color = [WHITE] * graph.num_vertices
    
    def dfs(v):
        color[v] = GRAY
        
        for neighbor in graph.get_neighbors(v):
            if color[neighbor] == GRAY:
                return True  # Back edge
            if color[neighbor] == WHITE and dfs(neighbor):
                return True
        
        color[v] = BLACK
        return False
    
    for v in range(graph.num_vertices):
        if color[v] == WHITE:
            if dfs(v):
                return True
    
    return False

class DependencyResolver:
    """Resolve dependencies using topological sort"""
    def __init__(self):
        self.packages = {}
        self.graph = None
        self.pkg_to_id = {}
        self.id_to_pkg = {}
    
    def add_package(self, package, dependencies):
        """Add a package with its dependencies"""
        self.packages[package] = dependencies
    
    def build_graph(self):
        """Build dependency graph"""
        # Create package ID mapping
        all_packages = set(self.packages.keys())
        for deps in self.packages.values():
            all_packages.update(deps)
        
        self.pkg_to_id = {pkg: i for i, pkg in enumerate(sorted(all_packages))}
        self.id_to_pkg = {i: pkg for pkg, i in self.pkg_to_id.items()}
        
        # Build graph
        self.graph = AdjacencyList(len(all_packages), directed=True)
        
        for package, dependencies in self.packages.items():
            pkg_id = self.pkg_to_id[package]
            for dep in dependencies:
                dep_id = self.pkg_to_id[dep]
                # Edge from dependency to package (dep must come first)
                self.graph.add_edge(dep_id, pkg_id)
    
    def get_install_order(self):
        """Get valid installation order"""
        self.build_graph()
        
        order_ids = topological_sort_kahn(self.graph)
        if order_ids is None:
            raise ValueError("Circular dependency detected!")
        
        return [self.id_to_pkg[i] for i in order_ids]

def course_schedule(num_courses, prerequisites):
    """
    Determine if all courses can be finished given prerequisites.
    prerequisites = [(course, prereq), ...]
    """
    graph = AdjacencyList(num_courses, directed=True)
    
    for course, prereq in prerequisites:
        graph.add_edge(prereq, course)
    
    order = topological_sort_kahn(graph)
    return order is not None  # Can finish if no cycle

def find_course_order(num_courses, prerequisites):
    """Find valid order to take all courses"""
    graph = AdjacencyList(num_courses, directed=True)
    
    for course, prereq in prerequisites:
        graph.add_edge(prereq, course)
    
    return topological_sort_kahn(graph)
```

### Test Cases
- DAG with single valid ordering
- DAG with multiple valid orderings
- Graph with cycle (should return None)
- Complete DAG (many orderings)
- Real dependency graphs (npm packages, courses)
- Empty graph, single vertex

### Deliverables
- `topological_sort.{py,go,zig}` - Both algorithms
- `cycle_detection.{py,go,zig}` - Directed graph cycle detection
- `dependency_resolver.{py,go,zig}` - Practical application
- `all_toposorts.{py,go,zig}` - Find all valid orderings
- `comparison.md` - Kahn's vs. DFS-based
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Both algorithms produce valid topological orderings
- Cycle detection works correctly
- Understanding of when topological sort is applicable
- Can solve real dependency problems

---

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_2.md) (if exists)
- [Next Lab](./lab_6_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 3 of 6*
