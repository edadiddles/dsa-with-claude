# Lab 7.5: Route Planning Application

**Module**: Module 7 - Graph Algorithms II - Shortest Paths  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 7.5: Route Planning Application

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Build practical route planning system
- Implement realistic constraints
- Optimize for real-world usage

### Requirements

1. **Features**:
   - Multi-criteria routing (shortest, fastest, avoid tolls)
   - Turn restrictions
   - Time-dependent edge weights (traffic patterns)
   - Intermediate waypoints

2. **Optimizations**:
   - A* with good heuristics
   - Bidirectional search
   - Caching common routes

3. **Realistic constraints**:
   - One-way streets
   - Road closures
   - Vehicle restrictions

### Implementation Details

```python
class RouteOptions:
    def __init__(self):
        self.optimize_for = "time"  # or "distance", "fuel"
        self.avoid_tolls = False
        self.avoid_highways = False
        self.avoid_ferries = False
        self.waypoints = []
        self.departure_time = None  # For time-dependent routing

class RoadNetwork:
    """Enhanced graph for road networks"""
    def __init__(self, num_vertices):
        self.graph = AdjacencyList(num_vertices, directed=True)
        self.coordinates = {}  # vertex -> (lat, lon)
        self.road_types = {}  # (u,v) -> type (highway, toll, etc.)
        self.speed_limits = {}  # (u,v) -> speed limit
        self.turn_restrictions = set()  # set of (u,v,w) where turn from u->v->w is forbidden
    
    def add_road(self, u, v, distance, road_type="street", speed_limit=30):
        """Add a road segment"""
        self.graph.add_edge(u, v, distance)
        self.road_types[(u, v)] = road_type
        self.speed_limits[(u, v)] = speed_limit
    
    def set_coordinates(self, vertex, lat, lon):
        """Set GPS coordinates for vertex"""
        self.coordinates[vertex] = (lat, lon)
    
    def add_turn_restriction(self, u, v, w):
        """Forbid turn from u->v->w"""
        self.turn_restrictions.add((u, v, w))

def multi_criteria_route(network, source, target, options):
    """
    Find route optimizing for different criteria.
    """
    def get_edge_cost(u, v):
        distance = network.graph.get_weight(u, v)
        road_type = network.road_types.get((u, v), "street")
        speed_limit = network.speed_limits.get((u, v), 30)
        
        # Base cost
        if options.optimize_for == "distance":
            cost = distance
        elif options.optimize_for == "time":
            cost = distance / speed_limit  # Time = distance / speed
        elif options.optimize_for == "fuel":
            # Simplified fuel model
            cost = distance * (1.2 if road_type == "highway" else 1.0)
        else:
            cost = distance
        
        # Apply penalties
        if options.avoid_tolls and road_type == "toll":
            cost *= 10  # Heavy penalty
        
        if options.avoid_highways and road_type == "highway":
            cost *= 5
        
        if options.avoid_ferries and road_type == "ferry":
            cost *= 20
        
        return cost
    
    # Create custom graph with modified costs
    custom_graph = AdjacencyList(network.graph.num_vertices, directed=True)
    for u in range(network.graph.num_vertices):
        for v in network.graph.get_neighbors(u):
            cost = get_edge_cost(u, v)
            custom_graph.add_edge(u, v, cost)
    
    # Use A* with haversine heuristic
    def heuristic(v, target):
        if v not in network.coordinates or target not in network.coordinates:
            return 0
        return haversine_distance(
            network.coordinates[v],
            network.coordinates[target]
        )
    
    path, cost = a_star(custom_graph, source, target, heuristic)
    
    # Check turn restrictions
    if path and len(path) >= 3:
        for i in range(len(path) - 2):
            if (path[i], path[i+1], path[i+2]) in network.turn_restrictions:
                # This path has a forbidden turn, need to find alternative
                # In practice, we'd remove this turn and re-route
                pass
    
    return path, cost

def route_with_waypoints(network, start, waypoints, end, options):
    """
    Find route visiting all waypoints in order.
    """
    full_route = []
    total_cost = 0
    
    current = start
    for waypoint in waypoints:
        segment, cost = multi_criteria_route(network, current, waypoint, options)
        
        if segment is None:
            return None, float('inf')
        
        # Append segment (excluding last vertex to avoid duplicates)
        if full_route:
            full_route.extend(segment[1:])
        else:
            full_route.extend(segment)
        
        total_cost += cost
        current = waypoint
    
    # Final segment to end
    segment, cost = multi_criteria_route(network, current, end, options)
    
    if segment is None:
        return None, float('inf')
    
    full_route.extend(segment[1:])
    total_cost += cost
    
    return full_route, total_cost

def time_dependent_routing(network, source, target, departure_time, traffic_patterns):
    """
    Route planning with time-dependent edge weights (traffic).
    traffic_patterns[(u,v)] is a function: time -> travel_time
    """
    def get_travel_time(u, v, current_time):
        if (u, v) in traffic_patterns:
            return traffic_patterns[(u, v)](current_time)
        
        # Default: use speed limit
        distance = network.graph.get_weight(u, v)
        speed = network.speed_limits.get((u, v), 30)
        return distance / speed
    
    # Modified Dijkstra with time tracking
    dist = [float('inf')] * network.graph.num_vertices
    arrival_time = [float('inf')] * network.graph.num_vertices
    parent = [-1] * network.graph.num_vertices
    
    dist[source] = 0
    arrival_time[source] = departure_time
    
    pq = [(0, source, departure_time)]
    visited = set()
    
    while pq:
        d, u, time = heapq.heappop(pq)
        
        if u == target:
            path = reconstruct_path(parent, source, target)
            return path, dist[target], arrival_time[target]
        
        if u in visited:
            continue
        
        visited.add(u)
        
        for v in network.graph.get_neighbors(u):
            travel_time = get_travel_time(u, v, time)
            new_arrival = time + travel_time
            
            if dist[u] + travel_time < dist[v]:
                dist[v] = dist[u] + travel_time
                arrival_time[v] = new_arrival
                parent[v] = u
                heapq.heappush(pq, (dist[v], v, new_arrival))
    
    return None, float('inf'), None

def haversine_distance(coord1, coord2):
    """
    Calculate great-circle distance between two points on Earth.
    coord1, coord2 are (lat, lon) in degrees.
    Returns distance in kilometers.
    """
    import math
    
    lat1, lon1 = coord1
    lat2, lon2 = coord2
    
    # Convert to radians
    lat1, lon1 = math.radians(lat1), math.radians(lon1)
    lat2, lon2 = math.radians(lat2), math.radians(lon2)
    
    # Haversine formula
    dlat = lat2 - lat1
    dlon = lon2 - lon1
    a = math.sin(dlat/2)**2 + math.cos(lat1) * math.cos(lat2) * math.sin(dlon/2)**2
    c = 2 * math.asin(math.sqrt(a))
    
    # Earth radius in kilometers
    r = 6371
    
    return c * r

class RouteCache:
    """Cache common routes to avoid recomputation"""
    def __init__(self, max_size=1000):
        self.cache = {}
        self.max_size = max_size
    
    def get(self, source, target, options_hash):
        key = (source, target, options_hash)
        return self.cache.get(key)
    
    def put(self, source, target, options_hash, route, cost):
        if len(self.cache) >= self.max_size:
            # Simple LRU: remove oldest
            self.cache.pop(next(iter(self.cache)))
        
        key = (source, target, options_hash)
        self.cache[key] = (route, cost)
    
    def hash_options(self, options):
        """Create hash of options for cache key"""
        return (
            options.optimize_for,
            options.avoid_tolls,
            options.avoid_highways,
            options.avoid_ferries
        )
```

### Test Cases
- Simple point-to-point routing
- Multi-waypoint routes
- Routes with turn restrictions
- Time-dependent routing with traffic
- Compare optimizations (distance vs. time)

### Deliverables
- `route_planner.{py,go,zig}` - Full routing system
- `multi_criteria.{py,go,zig}` - Support multiple optimization criteria
- `waypoint_routing.{py,go,zig}` - Routes with intermediate stops
- `time_dependent.{py,go,zig}` - Traffic-aware routing
- `road_network_loader.{py,go,zig}` - Load real road data
- `route_visualizer.{py,go,zig}` - Visualize routes on map
- `web_interface/` - Simple web UI for route planning (optional)

### Success Criteria
- Handles realistic constraints
- Fast enough for interactive use (< 1 second for typical queries)
- Correct routes for complex queries
- Understanding of practical routing challenges

---

---

## Back to Module

[← Back to Module 7: Graph Algorithms II - Shortest Paths](../module_07/README.md)

## Navigation

- [Previous Lab](./lab_7_4.md) (if exists)
- [Next Lab](./lab_7_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 7, Lab 5 of 5*
