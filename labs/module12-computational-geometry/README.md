# Module 12: Computational Geometry

**Duration**: 2 weeks (Optional Advanced Module)  
**Difficulty**: Medium to Hard

## Overview

Computational geometry deals with algorithms for solving geometric problems. This module covers fundamental geometric algorithms used in computer graphics, GIS, robotics, CAD, and game development. While optional, these techniques are essential for many real-world applications.

**Why this matters**: Geometric algorithms power GPS navigation, computer graphics, collision detection in games, robot path planning, image processing, and CAD software. Understanding these fundamentals enables you to build mapping applications, graphics engines, and spatial analysis tools.

## Learning Objectives

- Master basic geometric primitives and operations
- Implement convex hull algorithms
- Solve line segment intersection problems
- Build spatial data structures (k-d trees, quad trees)
- Apply geometric algorithms to practical problems
- Understand computational complexity of geometric problems

## Topics Covered

- Geometric primitives (points, lines, polygons)
- Convex hull (Graham scan, Jarvis march, QuickHull)
- Line segment intersection
- Closest pair of points (from Module 10, review)
- Range searching and spatial data structures
- Polygon operations (area, containment, triangulation)
- Sweep line algorithms
- Voronoi diagrams and Delaunay triangulation (overview)

---

## Lab 12.1: Geometric Primitives & Basic Operations

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement fundamental geometric operations
- Handle floating-point precision issues
- Build foundation for advanced algorithms
- Apply to simple geometric problems

### Requirements

1. **Implement geometric primitives**:
   - **Point** class with basic operations
   - **Line** and **Line Segment** classes
   - **Polygon** class
   - **Vector** operations

2. **Basic operations**:
   - Distance calculations
   - Orientation test (left/right turn)
   - Line-line intersection
   - Point-in-polygon test
   - Polygon area

3. **Precision handling**:
   - Epsilon comparisons
   - Robust predicates

### Implementation Details

```python
import math

# EPSILON for floating-point comparisons
EPSILON = 1e-9

def equals(a, b):
    """Check if two floating-point numbers are equal"""
    return abs(a - b) < EPSILON

# POINT CLASS

class Point:
    """2D Point"""
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def __eq__(self, other):
        return equals(self.x, other.x) and equals(self.y, other.y)
    
    def __lt__(self, other):
        """For sorting points"""
        if not equals(self.x, other.x):
            return self.x < other.x
        return self.y < other.y
    
    def __repr__(self):
        return f"Point({self.x}, {self.y})"
    
    def distance_to(self, other):
        """Euclidean distance"""
        return math.sqrt((self.x - other.x)**2 + (self.y - other.y)**2)
    
    def manhattan_distance(self, other):
        """Manhattan distance"""
        return abs(self.x - other.x) + abs(self.y - other.y)
    
    def dot(self, other):
        """Dot product (treating points as vectors from origin)"""
        return self.x * other.x + self.y * other.y
    
    def cross(self, other):
        """Cross product magnitude (for 2D vectors)"""
        return self.x * other.y - self.y * other.x

# ORIENTATION TEST

def orientation(p, q, r):
    """
    Find orientation of ordered triplet (p, q, r).
    Returns:
        0 if p, q, r are collinear
        1 if clockwise
        2 if counterclockwise
    
    Uses cross product: (q-p) × (r-p)
    """
    val = (q.y - p.y) * (r.x - q.x) - (q.x - p.x) * (r.y - q.y)
    
    if equals(val, 0):
        return 0  # Collinear
    
    return 1 if val > 0 else 2  # Clock or counterclock wise

def ccw(p, q, r):
    """
    Counterclockwise test.
    Returns True if p->q->r makes a left turn (counterclockwise).
    """
    return (q.x - p.x) * (r.y - p.y) - (q.y - p.y) * (r.x - p.x) > EPSILON

# LINE SEGMENT

class LineSegment:
    """Line segment defined by two endpoints"""
    def __init__(self, p1, p2):
        self.p1 = p1
        self.p2 = p2
    
    def length(self):
        """Length of segment"""
        return self.p1.distance_to(self.p2)
    
    def contains_point(self, p):
        """Check if point lies on segment (assuming collinear)"""
        return (min(self.p1.x, self.p2.x) <= p.x <= max(self.p1.x, self.p2.x) and
                min(self.p1.y, self.p2.y) <= p.y <= max(self.p1.y, self.p2.y))
    
    def intersects(self, other):
        """Check if this segment intersects another segment"""
        p1, p2 = self.p1, self.p2
        p3, p4 = other.p1, other.p2
        
        # Find orientations
        o1 = orientation(p1, p2, p3)
        o2 = orientation(p1, p2, p4)
        o3 = orientation(p3, p4, p1)
        o4 = orientation(p3, p4, p2)
        
        # General case: different orientations
        if o1 != o2 and o3 != o4:
            return True
        
        # Special cases: collinear points
        if o1 == 0 and self.contains_point(p3):
            return True
        if o2 == 0 and self.contains_point(p4):
            return True
        if o3 == 0 and other.contains_point(p1):
            return True
        if o4 == 0 and other.contains_point(p2):
            return True
        
        return False
    
    def intersection_point(self, other):
        """
        Find intersection point of two line segments.
        Returns None if no intersection or infinite intersections.
        """
        p1, p2 = self.p1, self.p2
        p3, p4 = other.p1, other.p2
        
        x1, y1 = p1.x, p1.y
        x2, y2 = p2.x, p2.y
        x3, y3 = p3.x, p3.y
        x4, y4 = p4.x, p4.y
        
        denom = (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)
        
        if equals(denom, 0):
            return None  # Parallel or collinear
        
        t = ((x1 - x3) * (y3 - y4) - (y1 - y3) * (x3 - x4)) / denom
        u = -((x1 - x2) * (y1 - y3) - (y1 - y2) * (x1 - x3)) / denom
        
        if 0 <= t <= 1 and 0 <= u <= 1:
            x = x1 + t * (x2 - x1)
            y = y1 + t * (y2 - y1)
            return Point(x, y)
        
        return None

# POLYGON

class Polygon:
    """Polygon defined by list of vertices"""
    def __init__(self, vertices):
        self.vertices = vertices
    
    def area(self):
        """
        Calculate polygon area using shoelace formula.
        Works for simple polygons (no self-intersections).
        """
        n = len(self.vertices)
        area = 0
        
        for i in range(n):
            j = (i + 1) % n
            area += self.vertices[i].x * self.vertices[j].y
            area -= self.vertices[j].x * self.vertices[i].y
        
        return abs(area) / 2
    
    def perimeter(self):
        """Calculate perimeter"""
        n = len(self.vertices)
        perimeter = 0
        
        for i in range(n):
            j = (i + 1) % n
            perimeter += self.vertices[i].distance_to(self.vertices[j])
        
        return perimeter
    
    def contains_point(self, point):
        """
        Point-in-polygon test using ray casting algorithm.
        Cast ray from point to infinity and count intersections.
        Odd = inside, Even = outside.
        """
        n = len(self.vertices)
        inside = False
        
        p1 = self.vertices[0]
        for i in range(1, n + 1):
            p2 = self.vertices[i % n]
            
            if point.y > min(p1.y, p2.y):
                if point.y <= max(p1.y, p2.y):
                    if point.x <= max(p1.x, p2.x):
                        if p1.y != p2.y:
                            x_intersection = (point.y - p1.y) * (p2.x - p1.x) / (p2.y - p1.y) + p1.x
                        
                        if p1.x == p2.x or point.x <= x_intersection:
                            inside = not inside
            
            p1 = p2
        
        return inside
    
    def is_convex(self):
        """Check if polygon is convex"""
        n = len(self.vertices)
        
        if n < 3:
            return False
        
        sign = None
        
        for i in range(n):
            p1 = self.vertices[i]
            p2 = self.vertices[(i + 1) % n]
            p3 = self.vertices[(i + 2) % n]
            
            cross = (p2.x - p1.x) * (p3.y - p2.y) - (p2.y - p1.y) * (p3.x - p2.x)
            
            if not equals(cross, 0):
                if sign is None:
                    sign = cross > 0
                elif (cross > 0) != sign:
                    return False
        
        return True

# GEOMETRIC UTILITIES

def angle_between_points(p1, p2, p3):
    """
    Calculate angle at p2 formed by p1-p2-p3.
    Returns angle in radians.
    """
    v1_x = p1.x - p2.x
    v1_y = p1.y - p2.y
    v2_x = p3.x - p2.x
    v2_y = p3.y - p2.y
    
    dot = v1_x * v2_x + v1_y * v2_y
    det = v1_x * v2_y - v1_y * v2_x
    
    return math.atan2(det, dot)

def closest_point_on_segment(segment, point):
    """
    Find closest point on line segment to given point.
    """
    p1, p2 = segment.p1, segment.p2
    
    dx = p2.x - p1.x
    dy = p2.y - p1.y
    
    if equals(dx, 0) and equals(dy, 0):
        return p1
    
    t = ((point.x - p1.x) * dx + (point.y - p1.y) * dy) / (dx * dx + dy * dy)
    t = max(0, min(1, t))
    
    return Point(p1.x + t * dx, p1.y + t * dy)

def point_to_segment_distance(segment, point):
    """Distance from point to line segment"""
    closest = closest_point_on_segment(segment, point)
    return point.distance_to(closest)
```

### Test Cases
- Orientation tests with collinear points
- Line segment intersections (various cases)
- Point-in-polygon for concave and convex polygons
- Polygon area calculations
- Edge cases (degenerate polygons, vertical lines)

### Deliverables
- `geometry_primitives.{py,go,zig}` - Point, Line, Polygon classes
- `orientation.{py,go,zig}` - Orientation tests
- `intersection.{py,go,zig}` - Line segment intersection
- `polygon_operations.{py,go,zig}` - Area, containment, etc.
- `geometric_utilities.{py,go,zig}` - Helper functions
- `precision_handling.md` - Floating-point considerations
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All geometric operations work correctly
- Proper handling of floating-point precision
- Point-in-polygon works for complex shapes
- Understanding of orientation tests

---

## Lab 12.2: Convex Hull Algorithms

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

## Lab 12.3: Line Segment Intersection & Sweep Line

**Duration**: 4-5 hours  
**Difficulty**: Hard

### Objectives
- Implement sweep line algorithm
- Find all segment intersections efficiently
- Apply to practical problems
- Understand sweep line paradigm

### Requirements

1. **Implement sweep line for segments**:
   - **Bentley-Ottmann algorithm** (O((n + k) log n))
   - Event queue and sweep line status
   - Handle degenerate cases

2. **Applications**:
   - Map overlay
   - Polygon intersection
   - Visibility problems

### Implementation Details

```python
import heapq
from enum import Enum

class EventType(Enum):
    START = 1
    END = 2
    INTERSECTION = 3

class Event:
    """Event for sweep line algorithm"""
    def __init__(self, point, event_type, segment=None):
        self.point = point
        self.type = event_type
        self.segment = segment
    
    def __lt__(self, other):
        # Sort by x-coordinate, then by type
        if not equals(self.point.x, other.point.x):
            return self.point.x < other.point.x
        if not equals(self.point.y, other.point.y):
            return self.point.y < other.point.y
        return self.type.value < other.type.value

def bentley_ottmann(segments):
    """
    Bentley-Ottmann algorithm to find all segment intersections.
    Time: O((n + k) log n) where k is number of intersections
    """
    # Event queue
    events = []
    
    # Add segment endpoints
    for seg in segments:
        # Ensure p1.x <= p2.x
        if seg.p1.x > seg.p2.x or (equals(seg.p1.x, seg.p2.x) and seg.p1.y > seg.p2.y):
            seg.p1, seg.p2 = seg.p2, seg.p1
        
        heapq.heappush(events, Event(seg.p1, EventType.START, seg))
        heapq.heappush(events, Event(seg.p2, EventType.END, seg))
    
    # Sweep line status (segments currently intersecting sweep line)
    status = []
    intersections = []
    
    while events:
        event = heapq.heappop(events)
        
        if event.type == EventType.START:
            # Add segment to status
            status.append(event.segment)
            status.sort(key=lambda s: sweep_line_y(s, event.point.x))
            
            # Check for intersections with neighbors
            idx = status.index(event.segment)
            
            if idx > 0:
                check_intersection(status[idx-1], status[idx], events, intersections)
            if idx < len(status) - 1:
                check_intersection(status[idx], status[idx+1], events, intersections)
        
        elif event.type == EventType.END:
            # Remove segment from status
            idx = status.index(event.segment)
            
            # Check neighbors for intersection
            if 0 < idx < len(status) - 1:
                check_intersection(status[idx-1], status[idx+1], events, intersections)
            
            status.remove(event.segment)
        
        elif event.type == EventType.INTERSECTION:
            intersections.append(event.point)
            
            # Swap segments in status
            # (Simplified - full implementation more complex)
    
    return intersections

def sweep_line_y(segment, x):
    """Calculate y-coordinate where segment intersects vertical line at x"""
    p1, p2 = segment.p1, segment.p2
    
    if equals(p1.x, p2.x):
        return (p1.y + p2.y) / 2
    
    t = (x - p1.x) / (p2.x - p1.x)
    return p1.y + t * (p2.y - p1.y)

def check_intersection(seg1, seg2, events, intersections):
    """Check if two segments intersect and add event if they do"""
    intersection = seg1.intersection_point(seg2)
    
    if intersection:
        heapq.heappush(events, Event(intersection, EventType.INTERSECTION))

# SIMPLER ALL-PAIRS CHECK

def find_all_intersections_naive(segments):
    """
    Naive algorithm: check all pairs.
    Time: O(n²)
    """
    intersections = []
    n = len(segments)
    
    for i in range(n):
        for j in range(i + 1, n):
            point = segments[i].intersection_point(segments[j])
            if point:
                intersections.append((i, j, point))
    
    return intersections

# POLYGON INTERSECTION

def polygon_intersection_points(poly1, poly2):
    """
    Find all intersection points between two polygons.
    """
    intersections = []
    
    # Check all edge pairs
    for i in range(len(poly1.vertices)):
        seg1 = LineSegment(
            poly1.vertices[i],
            poly1.vertices[(i + 1) % len(poly1.vertices)]
        )
        
        for j in range(len(poly2.vertices)):
            seg2 = LineSegment(
                poly2.vertices[j],
                poly2.vertices[(j + 1) % len(poly2.vertices)]
            )
            
            point = seg1.intersection_point(seg2)
            if point:
                intersections.append(point)
    
    return intersections

# APPLICATIONS

def are_polygons_intersecting(poly1, poly2):
    """
    Check if two polygons intersect.
    """
    # Check if any vertices of one polygon are inside the other
    for v in poly1.vertices:
        if poly2.contains_point(v):
            return True
    
    for v in poly2.vertices:
        if poly1.contains_point(v):
            return True
    
    # Check if any edges intersect
    intersections = polygon_intersection_points(poly1, poly2)
    return len(intersections) > 0
```

### Test Cases
- Simple intersecting segments
- Many segments with no intersections
- Overlapping segments
- Degenerate cases
- Performance benchmarks

### Deliverables
- `sweep_line.{py,go,zig}` - Bentley-Ottmann algorithm
- `segment_intersection.{py,go,zig}` - All implementations
- `polygon_intersection.{py,go,zig}` - Polygon-polygon intersection
- `applications.{py,go,zig}` - Practical uses
- `sweep_line_visualizer.{py,go,zig}` - Visualize algorithm
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Sweep line finds all intersections
- Correct handling of degenerate cases
- Understanding of sweep line paradigm
- Performance matches O((n + k) log n)

---

## Lab 12.4: Spatial Data Structures

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

## Module Resources

### CLRS References
- Chapter 33: Computational Geometry

### Additional Reading
- "Computational Geometry: Algorithms and Applications" by de Berg et al.
- GIS algorithm references
- Computer graphics textbooks

### Key Takeaways

By the end of this module, you should:
1. Master geometric primitives and operations
2. Implement convex hull algorithms
3. Apply sweep line technique
4. Use spatial data structures effectively
5. Solve practical geometric problems
6. Understand precision issues in geometry

### Next Steps

Congratulations on completing the DSA curriculum! You've mastered:
- Fundamental data structures
- Core algorithms (sorting, searching, graphs)
- Advanced techniques (DP, greedy, backtracking)
- Specialized algorithms (string, geometry)

**Continue learning**:
- Implement algorithms in production systems
- Contribute to open-source projects
- Solve competitive programming problems
- Build real-world applications

---

## Tips for Success

1. **Visualize geometry** - Draw diagrams for every problem
2. **Handle precision carefully** - Use epsilon comparisons
3. **Test edge cases** - Collinear points, degenerate shapes
4. **Start with simple cases** - Build up to complex
5. **Use existing libraries** - But understand internals first

## Common Pitfalls

- **Floating-point comparison errors** - Use epsilon
- **Not handling degenerate cases** - Collinear, vertical lines
- **Incorrect orientation tests** - Sign errors
- **Off-by-one in polygon operations**
- **Not considering numerical stability**
- **Forgetting to handle edge cases** - Empty sets, single points
