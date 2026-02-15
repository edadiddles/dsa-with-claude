# Lab 10.2: Backtracking & Branch-and-Bound

**Module**: Module 10 - Advanced Algorithm Techniques  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 10.2: Backtracking & Branch-and-Bound

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Master backtracking technique
- Implement constraint satisfaction problems
- Apply branch-and-bound optimization
- Understand pruning strategies

### Requirements

1. **Implement classic backtracking problems**:
   - **N-Queens**
   - **Sudoku Solver**
   - **Graph Coloring**
   - **Subset Sum**
   - **Hamiltonian Path**
   - **Traveling Salesman (exact)**

2. **Optimizations**:
   - Constraint propagation
   - Intelligent ordering
   - Pruning strategies
   - Branch-and-bound for optimization

3. **Analysis**:
   - Measure search space reduction
   - Compare with/without pruning
   - Understand exponential complexity

### Implementation Details

```python
# N-QUEENS

def solve_n_queens(n):
    """
    Solve N-Queens problem using backtracking.
    Place n queens on n×n board so no two attack each other.
    """
    solutions = []
    board = [-1] * n  # board[row] = column of queen in that row
    
    def is_safe(row, col):
        """Check if queen can be placed at (row, col)"""
        for prev_row in range(row):
            prev_col = board[prev_row]
            
            # Check column conflict
            if prev_col == col:
                return False
            
            # Check diagonal conflict
            if abs(prev_row - row) == abs(prev_col - col):
                return False
        
        return True
    
    def backtrack(row):
        """Try placing queen in row"""
        if row == n:
            # Found complete solution
            solutions.append(board[:])
            return
        
        for col in range(n):
            if is_safe(row, col):
                board[row] = col
                backtrack(row + 1)
                board[row] = -1  # Backtrack
    
    backtrack(0)
    return solutions

def solve_n_queens_optimized(n):
    """
    Optimized N-Queens using bit manipulation.
    Tracks attacked columns and diagonals with bits.
    """
    solutions = []
    
    def backtrack(row, cols, diag1, diag2, board):
        if row == n:
            solutions.append(board[:])
            return
        
        # Try each column
        available = ((1 << n) - 1) & ~(cols | diag1 | diag2)
        
        while available:
            # Get rightmost available position
            col = available & -available
            col_idx = (col - 1).bit_length() - 1
            
            board.append(col_idx)
            
            backtrack(
                row + 1,
                cols | col,
                (diag1 | col) << 1,
                (diag2 | col) >> 1,
                board
            )
            
            board.pop()
            available &= available - 1
    
    backtrack(0, 0, 0, 0, [])
    return solutions

# SUDOKU SOLVER

def solve_sudoku(board):
    """
    Solve 9×9 Sudoku using backtracking.
    board is 9×9 list of lists, 0 represents empty cell.
    """
    def is_valid(row, col, num):
        """Check if num can be placed at (row, col)"""
        # Check row
        if num in board[row]:
            return False
        
        # Check column
        if num in [board[i][col] for i in range(9)]:
            return False
        
        # Check 3×3 box
        box_row, box_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(box_row, box_row + 3):
            for j in range(box_col, box_col + 3):
                if board[i][j] == num:
                    return False
        
        return True
    
    def find_empty():
        """Find next empty cell"""
        for i in range(9):
            for j in range(9):
                if board[i][j] == 0:
                    return i, j
        return None
    
    def solve():
        """Recursive backtracking solver"""
        empty = find_empty()
        
        if empty is None:
            return True  # Solved
        
        row, col = empty
        
        for num in range(1, 10):
            if is_valid(row, col, num):
                board[row][col] = num
                
                if solve():
                    return True
                
                board[row][col] = 0  # Backtrack
        
        return False
    
    solve()
    return board

def solve_sudoku_optimized(board):
    """
    Optimized Sudoku with constraint propagation.
    """
    # Track possible values for each cell
    possible = [[set(range(1, 10)) if board[i][j] == 0 else set() 
                 for j in range(9)] for i in range(9)]
    
    # Initialize constraints
    for i in range(9):
        for j in range(9):
            if board[i][j] != 0:
                propagate_constraints(board, possible, i, j, board[i][j])
    
    def propagate_constraints(board, possible, row, col, num):
        """Remove num from possible values in row, col, box"""
        # Row and column
        for k in range(9):
            possible[row][k].discard(num)
            possible[k][col].discard(num)
        
        # Box
        box_row, box_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(box_row, box_row + 3):
            for j in range(box_col, box_col + 3):
                possible[i][j].discard(num)
    
    def solve():
        # Find cell with minimum remaining values (MRV heuristic)
        min_choices = 10
        best_cell = None
        
        for i in range(9):
            for j in range(9):
                if board[i][j] == 0 and len(possible[i][j]) < min_choices:
                    min_choices = len(possible[i][j])
                    best_cell = (i, j)
        
        if best_cell is None:
            return True  # Solved
        
        if min_choices == 0:
            return False  # Dead end
        
        row, col = best_cell
        
        for num in list(possible[row][col]):
            board[row][col] = num
            old_possible = [row[:] for row in possible]
            
            propagate_constraints(board, possible, row, col, num)
            
            if solve():
                return True
            
            board[row][col] = 0
            possible[:] = old_possible
        
        return False
    
    solve()
    return board

# GRAPH COLORING

def graph_coloring(graph, k):
    """
    Color graph with k colors using backtracking.
    graph is adjacency list.
    """
    n = len(graph)
    colors = [-1] * n
    
    def is_safe(node, color):
        """Check if node can be colored with color"""
        for neighbor in graph[node]:
            if colors[neighbor] == color:
                return False
        return True
    
    def backtrack(node):
        if node == n:
            return True  # All nodes colored
        
        for color in range(k):
            if is_safe(node, color):
                colors[node] = color
                
                if backtrack(node + 1):
                    return True
                
                colors[node] = -1  # Backtrack
        
        return False
    
    if backtrack(0):
        return colors
    return None

# SUBSET SUM

def subset_sum_backtrack(nums, target):
    """
    Find subset that sums to target using backtracking.
    Returns all solutions.
    """
    solutions = []
    
    def backtrack(index, current_sum, subset):
        if current_sum == target:
            solutions.append(subset[:])
            return
        
        if index >= len(nums) or current_sum > target:
            return
        
        # Include current number
        subset.append(nums[index])
        backtrack(index + 1, current_sum + nums[index], subset)
        subset.pop()
        
        # Exclude current number
        backtrack(index + 1, current_sum, subset)
    
    backtrack(0, 0, [])
    return solutions

# TRAVELING SALESMAN (EXACT)

def tsp_backtracking(dist_matrix):
    """
    Exact TSP solution using backtracking with branch-and-bound.
    dist_matrix[i][j] = distance from city i to city j.
    """
    n = len(dist_matrix)
    min_cost = [float('inf')]
    best_path = []
    
    def bound(path, visited):
        """Lower bound on remaining cost"""
        # Minimum cost to complete tour from current state
        cost = 0
        
        # Add cost of path so far
        for i in range(len(path) - 1):
            cost += dist_matrix[path[i]][path[i+1]]
        
        # Add minimum outgoing edge from last city
        if path:
            last = path[-1]
            min_out = min(dist_matrix[last][j] for j in range(n) 
                         if j not in visited)
            cost += min_out
        
        # Add minimum edges for unvisited cities
        for city in range(n):
            if city not in visited:
                min_edge = min(dist_matrix[city])
                cost += min_edge
        
        return cost
    
    def backtrack(path, visited, cost):
        if len(path) == n:
            # Complete tour, return to start
            total = cost + dist_matrix[path[-1]][path[0]]
            if total < min_cost[0]:
                min_cost[0] = total
                best_path[:] = path + [path[0]]
            return
        
        # Branch and bound: prune if lower bound exceeds best
        if bound(path, visited) >= min_cost[0]:
            return
        
        last = path[-1] if path else 0
        
        for next_city in range(n):
            if next_city not in visited:
                new_cost = cost + dist_matrix[last][next_city]
                
                path.append(next_city)
                visited.add(next_city)
                
                backtrack(path, visited, new_cost)
                
                path.pop()
                visited.remove(next_city)
    
    backtrack([0], {0}, 0)
    return min_cost[0], best_path

# PRUNING STRATEGIES

class BacktrackingOptimizer:
    """Collection of pruning and optimization techniques"""
    
    @staticmethod
    def forward_checking(domain, constraints):
        """
        Forward checking: remove inconsistent values from future variables.
        """
        pass
    
    @staticmethod
    def arc_consistency(domain, constraints):
        """
        Maintain arc consistency during search.
        """
        pass
    
    @staticmethod
    def minimum_remaining_values(variables, domains):
        """
        MRV heuristic: choose variable with fewest remaining values.
        """
        return min(variables, key=lambda v: len(domains[v]))
    
    @staticmethod
    def degree_heuristic(variables, graph):
        """
        Degree heuristic: choose variable with most constraints.
        """
        return max(variables, key=lambda v: len(graph[v]))
    
    @staticmethod
    def least_constraining_value(variable, values, constraints):
        """
        LCV: choose value that rules out fewest choices for neighbors.
        """
        def count_conflicts(value):
            count = 0
            for neighbor in constraints[variable]:
                # Count how many values this eliminates for neighbor
                pass
            return count
        
        return sorted(values, key=count_conflicts)
```

### Test Cases
- N-Queens for various n (compare with/without optimization)
- Sudoku puzzles of varying difficulty
- Graph coloring instances
- Subset sum with different targets
- Small TSP instances (exact solution)

### Deliverables
- `n_queens.{py,go,zig}` - Both basic and optimized versions
- `sudoku_solver.{py,go,zig}` - With constraint propagation
- `graph_coloring.{py,go,zig}` - Backtracking coloring
- `tsp_exact.{py,go,zig}` - Branch-and-bound TSP
- `constraint_propagation.{py,go,zig}` - CSP utilities
- `pruning_analysis.md` - Effectiveness of pruning
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All backtracking algorithms find correct solutions
- Understanding of pruning effectiveness
- Branch-and-bound reduces search space significantly
- Can apply backtracking to new problems

---

---

## Back to Module

[← Back to Module 10: Advanced Algorithm Techniques](../module_10/README.md)

## Navigation

- [Previous Lab](./lab_10_1.md) (if exists)
- [Next Lab](./lab_10_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 10, Lab 2 of 5*
