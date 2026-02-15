# Lab 12.3: Line Segment Intersection & Sweep Line

**Module**: Module 12 - Computational Geometry  
**Duration**: 4-5 hours  
**Difficulty**: Hard

---

Lab 12.3: Line Segment Intersection & Sweep Line

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

---

## Back to Module

[← Back to Module 12: Computational Geometry](../module_12/README.md)

## Navigation

- [Previous Lab](./lab_12_2.md) (if exists)
- [Next Lab](./lab_12_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 12, Lab 3 of 4*
