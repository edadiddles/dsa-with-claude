# Lab 12.4: Spatial Data Structures

**Module**: Module 12 - Computational Geometry  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 12.4: Spatial Data Structures

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement spatial data structures
- Perform range queries efficiently
- Apply to nearest neighbor search
- Understand space partitioning

### Requirements

1. **Implement spatial structures**:
   - **Quad Tree** (2D space partitioning)
   - **k-d Tree** (k-dimensional binary tree)
   - **R-Tree** (bounding rectangles)

2. **Operations**:
   - Insert/delete
   - Range query
   - Nearest neighbor
   - k-nearest neighbors

3. **Applications**:
   - Geographic information systems
   - Collision detection
   - Spatial indexing

### Implementation Details

```python
# QUAD TREE

class QuadTreeNode:
    """Node in quad tree"""
    def __init__(self, bounds, capacity=4):
        self.bounds = bounds  # (min_x, min_y, max_x, max_y)
        self.capacity = capacity
        self.points = []
        self.divided = False
        
        # Children (NW, NE, SW, SE)
        self.nw = None
        self.ne = None
        self.sw = None
        self.se = None
    
    def contains(self, point):
        """Check if point is within bounds"""
        min_x, min_y, max_x, max_y = self.bounds
        return (min_x <= point.x < max_x and 
                min_y <= point.y < max_y)
    
    def subdivide(self):
        """Subdivide node into four quadrants"""
        min_x, min_y, max_x, max_y = self.bounds
        mid_x = (min_x + max_x) / 2
        mid_y = (min_y + max_y) / 2
        
        self.nw = QuadTreeNode((min_x, mid_y, mid_x, max_y), self.capacity)
        self.ne = QuadTreeNode((mid_x, mid_y, max_x, max_y), self.capacity)
        self.sw = QuadTreeNode((min_x, min_y, mid_x, mid_y), self.capacity)
        self.se = QuadTreeNode((mid_x, min_y, max_x, mid_y), self.capacity)
        
        self.divided = True
    
    def insert(self, point):
        """Insert point into quad tree"""
        if not self.contains(point):
            return False
        
        if len(self.points) < self.capacity:
            self.points.append(point)
            return True
        
        if not self.divided:
            self.subdivide()
        
        return (self.nw.insert(point) or self.ne.insert(point) or
                self.sw.insert(point) or self.se.insert(point))
    
    def query_range(self, range_bounds):
        """Find all points in rectangular range"""
        found = []
        
        # Check if range intersects bounds
        if not ranges_intersect(self.bounds, range_bounds):
            return found
        
        # Check points in this node
        for point in self.points:
            if point_in_range(point, range_bounds):
                found.append(point)
        
        # Recursively check children
        if self.divided:
            found.extend(self.nw.query_range(range_bounds))
            found.extend(self.ne.query_range(range_bounds))
            found.extend(self.sw.query_range(range_bounds))
            found.extend(self.se.query_range(range_bounds))
        
        return found

def ranges_intersect(bounds1, bounds2):
    """Check if two rectangular ranges intersect"""
    min_x1, min_y1, max_x1, max_y1 = bounds1
    min_x2, min_y2, max_x2, max_y2 = bounds2
    
    return not (max_x1 < min_x2 or max_x2 < min_x1 or
                max_y1 < min_y2 or max_y2 < min_y1)

def point_in_range(point, bounds):
    """Check if point is in rectangular range"""
    min_x, min_y, max_x, max_y = bounds
    return min_x <= point.x <= max_x and min_y <= point.y <= max_y

# K-D TREE

class KDNode:
    """Node in k-d tree"""
    def __init__(self, point, left=None, right=None):
        self.point = point
        self.left = left
        self.right = right

class KDTree:
    """k-dimensional binary search tree (2D for this implementation)"""
    def __init__(self, points):
        self.root = self.build(points, depth=0)
    
    def build(self, points, depth):
        """Build k-d tree recursively"""
        if not points:
            return None
        
        # Select axis based on depth (0 for x, 1 for y)
        axis = depth % 2
        
        # Sort and choose median
        points.sort(key=lambda p: p.x if axis == 0 else p.y)
        median = len(points) // 2
        
        return KDNode(
            points[median],
            self.build(points[:median], depth + 1),
            self.build(points[median + 1:], depth + 1)
        )
    
    def nearest_neighbor(self, target):
        """Find nearest neighbor to target point"""
        best = [None, float('inf')]
        
        def search(node, depth):
            if node is None:
                return
            
            # Calculate distance to current node
            dist = target.distance_to(node.point)
            
            if dist < best[1]:
                best[0] = node.point
                best[1] = dist
            
            # Determine which subtree to search first
            axis = depth % 2
            target_val = target.x if axis == 0 else target.y
            node_val = node.point.x if axis == 0 else node.point.y
            
            if target_val < node_val:
                first, second = node.left, node.right
            else:
                first, second = node.right, node.left
            
            search(first, depth + 1)
            
            # Check if we need to search other subtree
            if abs(target_val - node_val) < best[1]:
                search(second, depth + 1)
        
        search(self.root, 0)
        return best[0]
    
    def k_nearest_neighbors(self, target, k):
        """Find k nearest neighbors"""
        import heapq
        
        # Max heap of (negative distance, point)
        heap = []
        
        def search(node, depth):
            if node is None:
                return
            
            dist = target.distance_to(node.point)
            
            if len(heap) < k:
                heapq.heappush(heap, (-dist, node.point))
            elif dist < -heap[0][0]:
                heapq.heapreplace(heap, (-dist, node.point))
            
            axis = depth % 2
            target_val = target.x if axis == 0 else target.y
            node_val = node.point.x if axis == 0 else node.point.y
            
            if target_val < node_val:
                first, second = node.left, node.right
            else:
                first, second = node.right, node.left
            
            search(first, depth + 1)
            
            if len(heap) < k or abs(target_val - node_val) < -heap[0][0]:
                search(second, depth + 1)
        
        search(self.root, 0)
        
        return [point for _, point in sorted(heap, reverse=True)]
```

### Test Cases
- Insert many points
- Range queries (various sizes)
- Nearest neighbor searches
- k-nearest neighbors
- Performance benchmarks

### Deliverables
- `quadtree.{py,go,zig}` - Quad tree implementation
- `kdtree.{py,go,zig}` - k-d tree with NN search
- `rtree.{py,go,zig}` - R-tree (optional)
- `spatial_queries.{py,go,zig}` - Range and NN queries
- `performance_comparison.{py,go,zig}` - Compare structures
- `applications.{py,go,zig}` - GIS, collision detection
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Efficient range queries
- Nearest neighbor search works correctly
- Understanding of space partitioning
- Can choose appropriate structure

---

---

## Back to Module

[← Back to Module 12: Computational Geometry](../module_12/README.md)

## Navigation

- [Previous Lab](./lab_12_3.md) (if exists)
- [Next Lab](./lab_12_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 12, Lab 4 of 4*
