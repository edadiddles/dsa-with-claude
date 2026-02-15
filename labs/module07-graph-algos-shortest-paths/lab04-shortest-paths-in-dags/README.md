# Lab 7.4: Shortest Paths in DAGs

**Module**: Module 7 - Graph Algorithms II - Shortest Paths  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 7.4: Shortest Paths in DAGs

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Exploit DAG structure for linear-time shortest paths
- Implement critical path method
- Apply to scheduling problems

### Requirements

1. **Implement DAG shortest paths**:
   - Topological sort
   - Single-pass relaxation
   - Longest paths (negate weights)

2. **Applications**:
   - Critical Path Method (CPM) for project scheduling
   - PERT charts
   - Longest increasing subsequence (using DAG)

### Implementation Details

```python
from topological_sort import topological_sort_dfs

def dag_shortest_paths(graph, source):
    """
    Shortest paths in DAG - O(V + E) time.
    Works for any edge weights (including negative).
    """
    # Get topological ordering
    topo_order = topological_sort_dfs(graph)
    
    if topo_order is None:
        return None, None  # Not a DAG
    
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    # Process vertices in topological order
    for u in topo_order:
        if dist[u] != float('inf'):
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                if dist[u] + weight < dist[v]:
                    dist[v] = dist[u] + weight
                    parent[v] = u
    
    return dist, parent

def dag_longest_paths(graph, source):
    """
    Longest paths in DAG.
    Negate weights and find shortest paths.
    """
    # Create graph with negated weights
    n = graph.num_vertices
    negated = AdjacencyList(n, directed=True)
    
    for u in range(n):
        for v in graph.get_neighbors(u):
            negated.add_edge(u, v, -graph.get_weight(u, v))
    
    dist, parent = dag_shortest_paths(negated, source)
    
    if dist is None:
        return None, None
    
    # Negate distances back
    dist = [-d if d != float('inf') else float('-inf') for d in dist]
    
    return dist, parent

class CriticalPathMethod:
    """
    Critical Path Method for project scheduling.
    Tasks are vertices, dependencies are edges, weights are durations.
    """
    def __init__(self):
        self.tasks = []
        self.graph = None
        self.task_to_id = {}
    
    def add_task(self, task_name, duration, dependencies=None):
        """Add a task with duration and dependencies"""
        if dependencies is None:
            dependencies = []
        self.tasks.append({
            'name': task_name,
            'duration': duration,
            'dependencies': dependencies
        })
    
    def build_graph(self):
        """Build DAG from tasks"""
        # Create mapping
        self.task_to_id = {task['name']: i for i, task in enumerate(self.tasks)}
        n = len(self.tasks)
        
        self.graph = AdjacencyList(n, directed=True)
        
        # Add edges from dependencies to tasks
        for i, task in enumerate(self.tasks):
            for dep in task['dependencies']:
                dep_id = self.task_to_id[dep]
                # Edge weight is the duration of the dependency
                self.graph.add_edge(dep_id, i, self.tasks[dep_id]['duration'])
    
    def find_critical_path(self):
        """Find critical path and project duration"""
        self.build_graph()
        
        # Add virtual start and end nodes
        n = self.graph.num_vertices
        start = n
        end = n + 1
        
        augmented = AdjacencyList(n + 2, directed=True)
        
        # Copy edges
        for u in range(n):
            for v in self.graph.get_neighbors(u):
                augmented.add_edge(u, v, self.graph.get_weight(u, v))
        
        # Connect start to all tasks with no dependencies
        for i, task in enumerate(self.tasks):
            if not task['dependencies']:
                augmented.add_edge(start, i, 0)
        
        # Connect all tasks to end
        for i in range(n):
            augmented.add_edge(i, end, self.tasks[i]['duration'])
        
        # Find longest path from start to end
        longest_dist, parent = dag_longest_paths(augmented, start)
        
        if longest_dist is None:
            return None, "Cycle detected in dependencies!"
        
        project_duration = longest_dist[end]
        
        # Reconstruct critical path
        critical_path = []
        current = end
        while parent[current] != -1:
            if current < n:  # Exclude virtual nodes in output
                critical_path.append(self.tasks[current]['name'])
            current = parent[current]
        
        critical_path = list(reversed(critical_path))
        
        # Calculate earliest and latest start times
        earliest_start = [longest_dist[i] for i in range(n)]
        
        # Latest start times (work backwards)
        latest_finish = [project_duration - self.tasks[i]['duration'] for i in range(n)]
        
        # Calculate slack
        slack = {}
        for i, task in enumerate(self.tasks):
            slack[task['name']] = latest_finish[i] - earliest_start[i]
        
        return {
            'duration': project_duration,
            'critical_path': critical_path,
            'earliest_start': {self.tasks[i]['name']: earliest_start[i] for i in range(n)},
            'slack': slack
        }

def longest_increasing_subsequence_dag(arr):
    """
    Find longest increasing subsequence using DAG approach.
    Build DAG where edge (i,j) exists if arr[i] < arr[j] and i < j.
    """
    n = len(arr)
    graph = AdjacencyList(n, directed=True)
    
    # Build DAG
    for i in range(n):
        for j in range(i + 1, n):
            if arr[i] < arr[j]:
                graph.add_edge(i, j, 1)  # Weight 1 for path length
    
    # Find longest path starting from any vertex
    max_length = 0
    best_path = []
    
    for start in range(n):
        dist, parent = dag_longest_paths(graph, start)
        
        for end in range(n):
            if dist[end] > max_length:
                max_length = dist[end]
                # Reconstruct path
                path = []
                current = end
                while current != -1:
                    path.append(current)
                    current = parent[current]
                best_path = list(reversed(path))
    
    # Convert indices to actual values
    lis = [arr[i] for i in best_path]
    
    return lis, len(lis)
```

### Test Cases
- Simple DAGs with unique longest path
- DAGs with multiple longest paths
- Project scheduling examples with realistic dependencies
- LIS test cases
- Compare DAG algorithm with Bellman-Ford (should be much faster)

### Deliverables
- `dag_shortest_paths.{py,go,zig}` - Linear-time algorithm
- `critical_path.{py,go,zig}` - CPM implementation
- `lis_dag.{py,go,zig}` - LIS using DAG
- `project_scheduler.{py,go,zig}` - Practical scheduling tool
- `performance_comparison.md` - DAG vs. general shortest paths
- `examples/` - Real project scheduling examples

### Success Criteria
- Linear-time complexity for DAGs
- Correct critical path identification
- Understanding of DAG optimizations
- Can apply to real scheduling problems

---

---

## Back to Module

[← Back to Module 7: Graph Algorithms II - Shortest Paths](../module_07/README.md)

## Navigation

- [Previous Lab](./lab_7_3.md) (if exists)
- [Next Lab](./lab_7_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 7, Lab 4 of 5*
