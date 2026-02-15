# Lab 12.2: Convex Hull Algorithms

**Module**: Module 12 - Computational Geometry  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 12.2: Convex Hull Algorithms

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement multiple convex hull algorithms
- Understand different approaches and complexities
- Apply to practical problems
- Visualize hulls

### Requirements

1. **Implement convex hull algorithms**:
   - **Graham Scan** (O(n log n))
   - **Jarvis March** (Gift Wrapping, O(nh))
   - **QuickHull** (O(n log n) expected)
   - **Chan's Algorithm** (O(n log h), optimal)

2. **Analysis**:
   - Compare running times
   - Understand when each is best
   - Handle degenerate cases

3. **Applications**:
   - Collision detection
   - Pattern recognition
   - Geographic analysis

### Implementation Details

```python
# GRAHAM SCAN

def graham_scan(points):
    """
    Graham Scan algorithm for convex hull.
    Time: O(n log n)
    
    Steps:
    1. Find lowest point (pivot)
    2. Sort by polar angle from pivot
    3. Process points maintaining convex hull
    """
    if len(points) < 3:
        return points
    
    # Find pivot (lowest y, then leftmost)
    pivot = min(points, key=lambda p: (p.y, p.x))
    
    # Sort by polar angle
    def polar_angle_key(p):
        if p == pivot:
            return -math.pi, 0
        angle = math.atan2(p.y - pivot.y, p.x - pivot.x)
        dist = pivot.distance_to(p)
        return angle, dist
    
    sorted_points = sorted(points, key=polar_angle_key)
    
    # Build convex hull
    hull = []
    
    for p in sorted_points:
        # Remove points that make clockwise turn
        while len(hull) >= 2 and not ccw(hull[-2], hull[-1], p):
            hull.pop()
        hull.append(p)
    
    return hull

# JARVIS MARCH (Gift Wrapping)

def jarvis_march(points):
    """
    Jarvis March (Gift Wrapping) algorithm.
    Time: O(nh) where h is number of hull points
    
    Repeatedly find leftmost point from current hull point.
    """
    if len(points) < 3:
        return points
    
    # Find leftmost point
    start = min(points, key=lambda p: (p.x, p.y))
    
    hull = []
    current = start
    
    while True:
        hull.append(current)
        
        # Find most counterclockwise point
        next_point = points[0]
        
        for p in points[1:]:
            if p == current:
                continue
            
            if next_point == current:
                next_point = p
            else:
                # Check if p is more counterclockwise than next_point
                cross = (next_point.x - current.x) * (p.y - current.y) - \
                        (next_point.y - current.y) * (p.x - current.x)
                
                if cross < 0:  # p is more counterclockwise
                    next_point = p
                elif equals(cross, 0):  # Collinear, choose farther
                    if current.distance_to(p) > current.distance_to(next_point):
                        next_point = p
        
        current = next_point
        
        if current == start:
            break
    
    return hull

# QUICKHULL

def quickhull(points):
    """
    QuickHull algorithm.
    Average time: O(n log n), Worst case: O(n²)
    
    Similar to quicksort: divide and conquer.
    """
    if len(points) < 3:
        return points
    
    # Find extreme points
    min_point = min(points, key=lambda p: p.x)
    max_point = max(points, key=lambda p: p.x)
    
    # Divide points into two sets
    def find_hull(p1, p2, points_set):
        if not points_set:
            return []
        
        # Find farthest point from line p1-p2
        max_dist = 0
        farthest = None
        
        for p in points_set:
            dist = abs((p2.y - p1.y) * p.x - (p2.x - p1.x) * p.y + 
                      p2.x * p1.y - p2.y * p1.x)
            dist /= math.sqrt((p2.y - p1.y)**2 + (p2.x - p1.x)**2)
            
            if dist > max_dist:
                max_dist = dist
                farthest = p
        
        if farthest is None:
            return []
        
        # Divide points
        left_set = [p for p in points_set if ccw(p1, farthest, p)]
        right_set = [p for p in points_set if ccw(farthest, p2, p)]
        
        # Recursively find hull
        return (find_hull(p1, farthest, left_set) + 
                [farthest] + 
                find_hull(farthest, p2, right_set))
    
    # Split points
    upper_points = [p for p in points if ccw(min_point, max_point, p)]
    lower_points = [p for p in points if ccw(max_point, min_point, p)]
    
    # Build hull
    hull = ([min_point] + 
            find_hull(min_point, max_point, upper_points) + 
            [max_point] + 
            find_hull(max_point, min_point, lower_points))
    
    return hull

# CHAN'S ALGORITHM

def chans_algorithm(points):
    """
    Chan's Algorithm: O(n log h) where h is hull size.
    Combines Jarvis March and Graham Scan.
    """
    if len(points) < 3:
        return points
    
    n = len(points)
    
    for m in [2**i for i in range(1, int(math.log2(n)) + 2)]:
        # Divide points into groups of size m
        groups = [points[i:i+m] for i in range(0, n, m)]
        
        # Compute convex hull of each group using Graham Scan
        mini_hulls = [graham_scan(group) for group in groups]
        
        # Use Jarvis March on mini hulls
        hull = []
        current = min(points, key=lambda p: (p.x, p.y))
        
        for _ in range(m):
            hull.append(current)
            
            # Find next point using mini hulls
            next_point = None
            
            for mini_hull in mini_hulls:
                # Binary search in this mini hull
                candidate = find_tangent(mini_hull, current)
                
                if next_point is None:
                    next_point = candidate
                else:
                    if ccw(current, next_point, candidate):
                        next_point = candidate
            
            if next_point == hull[0]:
                return hull
            
            current = next_point
    
    return hull

def find_tangent(hull, point):
    """Find tangent point from point to convex hull"""
    # Simplified version
    best = hull[0]
    
    for p in hull:
        if ccw(point, best, p):
            best = p
    
    return best

# COMPARISON AND ANALYSIS

def compare_convex_hull_algorithms(points):
    """Compare different convex hull algorithms"""
    import time
    
    algorithms = [
        ("Graham Scan", graham_scan),
        ("Jarvis March", jarvis_march),
        ("QuickHull", quickhull)
    ]
    
    results = {}
    
    for name, algo in algorithms:
        start = time.time()
        hull = algo(points[:])
        elapsed = time.time() - start
        
        results[name] = {
            'time': elapsed,
            'hull_size': len(hull),
            'hull_points': hull
        }
    
    return results

# APPLICATIONS

def is_inside_convex_hull(hull, point):
    """Check if point is inside convex hull"""
    n = len(hull)
    
    # Check if point is on same side of all edges
    sign = None
    
    for i in range(n):
        j = (i + 1) % n
        
        cross = (hull[j].x - hull[i].x) * (point.y - hull[i].y) - \
                (hull[j].y - hull[i].y) * (point.x - hull[i].x)
        
        if not equals(cross, 0):
            if sign is None:
                sign = cross > 0
            elif (cross > 0) != sign:
                return False
    
    return True

def convex_hull_diameter(hull):
    """
    Find diameter of convex hull (farthest pair of points).
    Uses rotating calipers: O(n)
    """
    n = len(hull)
    
    if n < 2:
        return 0
    
    max_dist = 0
    j = 1
    
    for i in range(n):
        while True:
            next_j = (j + 1) % n
            
            # Calculate cross product to determine rotation
            v1_x = hull[(i + 1) % n].x - hull[i].x
            v1_y = hull[(i + 1) % n].y - hull[i].y
            v2_x = hull[next_j].x - hull[j].x
            v2_y = hull[next_j].y - hull[j].y
            
            cross = v1_x * v2_y - v1_y * v2_x
            
            if cross <= 0:
                break
            
            j = next_j
        
        dist = hull[i].distance_to(hull[j])
        max_dist = max(max_dist, dist)
    
    return max_dist
```

### Test Cases
- Random point sets
- Collinear points
- Points on circle
- Degenerate cases
- Performance benchmarks

### Deliverables
- `graham_scan.{py,go,zig}` - Graham Scan implementation
- `jarvis_march.{py,go,zig}` - Gift wrapping
- `quickhull.{py,go,zig}` - QuickHull algorithm
- `chans_algorithm.{py,go,zig}` - Optimal output-sensitive
- `hull_applications.{py,go,zig}` - Diameter, containment
- `hull_visualizer.{py,go,zig}` - Visualize hulls
- `performance_comparison.{py,go,zig}` - Benchmark algorithms
- `algorithm_guide.md` - When to use each
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All algorithms produce correct convex hulls
- Understanding of time complexities
- Can choose appropriate algorithm
- Applications work correctly

---

---

## Back to Module

[← Back to Module 12: Computational Geometry](../module_12/README.md)

## Navigation

- [Previous Lab](./lab_12_1.md) (if exists)
- [Next Lab](./lab_12_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 12, Lab 2 of 4*
