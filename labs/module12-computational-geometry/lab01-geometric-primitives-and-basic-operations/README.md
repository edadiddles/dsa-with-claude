# Lab 12.1: Geometric Primitives & Basic Operations

**Module**: Module 12 - Computational Geometry  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 12.1: Geometric Primitives & Basic Operations

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

---

## Back to Module

[← Back to Module 12: Computational Geometry](../module_12/README.md)

## Navigation

- [Previous Lab](./lab_12_0.md) (if exists)
- [Next Lab](./lab_12_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 12, Lab 1 of 4*
