# Lab 6.6: Graph Visualization & Analysis Tools

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 6.6: Graph Visualization & Analysis Tools

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Build tools to visualize graphs
- Aid in debugging graph algorithms
- Compute graph statistics

### Requirements

1. **Visualization capabilities**:
   - Display graph structure (vertices and edges)
   - Highlight paths, trees, components
   - Animate algorithms (BFS, DFS, MST)
   - Show vertex/edge weights and labels

2. **Output formats**:
   - ASCII art for terminal
   - DOT format for Graphviz
   - Interactive HTML/JavaScript
   - Export to image formats

3. **Analysis tools**:
   - Compute graph statistics (density, diameter, etc.)
   - Find central vertices (degree centrality)
   - Identify bridges and articulation points

### Implementation Details

```python
def graph_to_dot(graph, highlight_edges=None, highlight_vertices=None, labels=None):
    """Convert graph to DOT format for Graphviz"""
    dot = "digraph G {\n" if graph.directed else "graph G {\n"
    edge_symbol = "->" if graph.directed else "--"
    
    # Add vertices
    for v in range(graph.num_vertices):
        color = "red" if highlight_vertices and v in highlight_vertices else "black"
        label = labels.get(v, str(v)) if labels else str(v)
        dot += f'  {v} [label="{label}", color={color}];\n'
    
    # Add edges
    seen_edges = set()
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            if not graph.directed and (v, u) in seen_edges:
                continue
            
            weight = graph.get_weight(u, v)
            color = "red" if highlight_edges and (u, v) in highlight_edges else "black"
            dot += f'  {u} {edge_symbol} {v} [label="{weight}", color={color}];\n'
            seen_edges.add((u, v))
    
    dot += "}\n"
    return dot

def ascii_graph(graph, max_width=80):
    """Simple ASCII representation"""
    lines = []
    for u in range(graph.num_vertices):
        neighbors = graph.get_neighbors(u)
        if graph.directed:
            line = f"{u} -> {neighbors}"
        else:
            line = f"{u}: {neighbors}"
        lines.append(line)
    return "\n".join(lines)

def graph_statistics(graph):
    """Compute various graph statistics"""
    stats = {}
    
    # Basic counts
    stats['vertices'] = graph.num_vertices
    stats['edges'] = len(graph.get_all_edges())
    
    # Density
    max_edges = graph.num_vertices * (graph.num_vertices - 1)
    if not graph.directed:
        max_edges //= 2
    stats['density'] = stats['edges'] / max_edges if max_edges > 0 else 0
    
    # Degree distribution
    degrees = [graph.degree(u) for u in range(graph.num_vertices)]
    stats['avg_degree'] = sum(degrees) / len(degrees) if degrees else 0
    stats['max_degree'] = max(degrees) if degrees else 0
    stats['min_degree'] = min(degrees) if degrees else 0
    
    # Components
    if not graph.directed:
        stats['num_components'] = count_connected_components(graph)
    
    # Diameter (for connected graphs)
    if not graph.directed and stats.get('num_components', 1) == 1:
        stats['diameter'] = compute_diameter(graph)
    
    return stats

def compute_diameter(graph):
    """Compute graph diameter (longest shortest path)"""
    diameter = 0
    
    for source in range(graph.num_vertices):
        result = bfs(graph, source)
        max_dist = max(result['distance'].values()) if result['distance'] else 0
        diameter = max(diameter, max_dist)
    
    return diameter

def find_bridges(graph):
    """
    Find bridges (edges whose removal disconnects the graph).
    Uses Tarjan's bridge-finding algorithm.
    """
    visited = set()
    discovery = {}
    low = {}
    parent = {}
    bridges = []
    time = [0]
    
    def dfs(u):
        visited.add(u)
        discovery[u] = low[u] = time[0]
        time[0] += 1
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                parent[v] = u
                dfs(v)
                
                low[u] = min(low[u], low[v])
                
                # If lowest vertex reachable from v is below u, then (u,v) is a bridge
                if low[v] > discovery[u]:
                    bridges.append((u, v))
            
            elif v != parent.get(u):
                low[u] = min(low[u], discovery[v])
    
    for v in range(graph.num_vertices):
        if v not in visited:
            parent[v] = None
            dfs(v)
    
    return bridges

def find_articulation_points(graph):
    """
    Find articulation points (vertices whose removal disconnects the graph).
    """
    visited = set()
    discovery = {}
    low = {}
    parent = {}
    articulation_points = set()
    time = [0]
    
    def dfs(u):
        children = 0
        visited.add(u)
        discovery[u] = low[u] = time[0]
        time[0] += 1
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                children += 1
                parent[v] = u
                dfs(v)
                
                low[u] = min(low[u], low[v])
                
                # u is an articulation point in two cases:
                # 1) u is root of DFS tree and has two or more children
                if parent.get(u) is None and children > 1:
                    articulation_points.add(u)
                
                # 2) u is not root and low value of one of its children is >= discovery of u
                if parent.get(u) is not None and low[v] >= discovery[u]:
                    articulation_points.add(u)
            
            elif v != parent.get(u):
                low[u] = min(low[u], discovery[v])
    
    for v in range(graph.num_vertices):
        if v not in visited:
            parent[v] = None
            dfs(v)
    
    return list(articulation_points)

def animate_bfs(graph, start):
    """Generate steps for BFS animation"""
    steps = []
    visited = set()
    queue = deque([start])
    visited.add(start)
    
    steps.append({
        'type': 'start',
        'vertex': start,
        'queue': list(queue),
        'visited': list(visited)
    })
    
    while queue:
        u = queue.popleft()
        steps.append({
            'type': 'visit',
            'vertex': u,
            'queue': list(queue),
            'visited': list(visited)
        })
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                visited.add(v)
                queue.append(v)
                steps.append({
                    'type': 'discover',
                    'vertex': v,
                    'from': u,
                    'queue': list(queue),
                    'visited': list(visited)
                })
    
    return steps

def export_to_html(graph, algorithm_name="Graph", animation_steps=None):
    """Export interactive graph visualization to HTML"""
    html = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>{algorithm_name}</title>
        <script src="https://d3js.org/d3.v7.min.js"></script>
        <style>
            .node {{ fill: #69b3a2; stroke: #333; stroke-width: 2px; }}
            .node.visited {{ fill: #ff6b6b; }}
            .link {{ stroke: #999; stroke-width: 2px; }}
            .link.highlighted {{ stroke: #ff6b6b; stroke-width: 4px; }}
        </style>
    </head>
    <body>
        <svg width="800" height="600"></svg>
        <script>
            // D3.js visualization code here
            // This is a simplified example
            const nodes = {[{{"id": i}} for i in range(graph.num_vertices)]};
            const edges = {graph.get_all_edges()};
            
            // Visualization logic...
        </script>
    </body>
    </html>
    """
    return html
```

### Test Cases
- Small graphs (easy to verify visually)
- Large graphs (stress test layout)
- Different graph types (trees, DAGs, cyclic)
- Animation of all algorithms

### Deliverables
- `graph_visualizer.{py,go,zig}` - Core visualization library
- `ascii_viz.{py,go,zig}` - Terminal-based visualization
- `dot_export.{py,go,zig}` - Graphviz export
- `graph_stats.{py,go,zig}` - Statistical analysis
- `algorithm_animator.{py,go,zig}` - Animate graph algorithms
- `bridge_articulation.{py,go,zig}` - Find bridges and articulation points
- `interactive_viz.html` - Web-based interactive visualizer (optional)

### Success Criteria
- Clear, readable visualizations
- Supports all graph types implemented
- Useful for debugging and learning
- Statistics are accurate

---

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_5.md) (if exists)
- [Next Lab](./lab_6_7.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 6 of 6*
