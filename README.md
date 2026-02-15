# Data Structures & Algorithms: A Lab-Based Curriculum

## Course Philosophy

This curriculum is designed for experienced engineers who want to build deep, foundational understanding of data structures and algorithms through implementation and experimentation. Rather than memorizing patterns, you'll develop intuition by building these structures from scratch, analyzing their behavior, and understanding the engineering tradeoffs that make certain approaches better than others in specific contexts.

**Core Principle**: Understanding comes from doing. Every concept includes hands-on labs where you'll implement, test, benchmark, and break things to see how they really work.

## Course Structure

The curriculum is organized into modules that build upon each other. Each module contains:

- **Theory Foundation**: Core concepts with emphasis on *why* these approaches exist
- **Implementation Labs**: Build it yourself in your choice of language (Python, Go, Zig, Rust, or Elixir)
- **Analysis Exercises**: Prove complexity bounds, analyze tradeoffs
- **Practical Applications**: Real-world scenarios where this structure/algorithm shines
- **Challenge Problems**: Push your understanding with edge cases and optimizations

## Prerequisites

- Comfort with at least one programming language
- Basic understanding of time/space complexity notation (O, Ω, Θ)
- Willingness to struggle productively with implementation details

## Module Breakdown

### Module 0: Foundations & Analysis (2 weeks)
**Why this matters**: You can't evaluate solutions without being able to analyze them rigorously.

**Topics**:
- Mathematical foundations (summations, recurrences, proof techniques)
- Asymptotic notation deep dive
- Empirical performance analysis and profiling
- RAM model vs. real hardware

**Labs**:
1. Implement timing harnesses and visualization tools for algorithm analysis
2. Solve recurrences both mathematically and empirically
3. Cache effects lab: observe real performance deviations from theoretical models
4. Build a micro-benchmarking framework you'll use throughout the course

**CLRS Reference**: Chapters 1-4

---

### Module 1: Elementary Data Structures (3 weeks)
**Why this matters**: These are the building blocks. Understanding their implementation details and tradeoffs is crucial for everything that follows.

**Topics**:
- Arrays and dynamic arrays (amortized analysis)
- Linked lists (singly, doubly, circular)
- Stacks and queues
- Hash tables (collision resolution strategies, load factors, hash function design)

**Labs**:
1. Implement a dynamic array from scratch with detailed amortization analysis
2. Build multiple hash table implementations (chaining, open addressing with different probing strategies)
3. Design and test custom hash functions for different data types
4. Memory layout exploration: measure actual memory usage and access patterns
5. Performance comparison: when does a linked list actually beat an array?

**CLRS Reference**: Chapters 10-11

---

### Module 2: Trees & Tree Algorithms (4 weeks)
**Why this matters**: Tree structures model hierarchical relationships and enable efficient search, insertion, and deletion operations.

**Topics**:
- Binary search trees (basic operations, traversals)
- Balanced trees: AVL, Red-Black trees, B-trees
- Tree traversals and their applications
- Augmented data structures

**Labs**:
1. Implement a basic BST with all standard operations
2. Build a self-balancing tree (AVL or Red-Black, your choice)
3. Augment trees with additional information (size, rank, interval data)
4. Implement B-tree for disk-based storage simulation
5. Visualizer: create tools to visualize tree operations and rotations
6. Comparative analysis: measure practical performance differences between tree types

**CLRS Reference**: Chapters 12-14, 18

---

### Module 3: Heaps & Priority Queues (2 weeks)
**Why this matters**: Priority queues are fundamental to many algorithms (graph algorithms, scheduling, compression).

**Topics**:
- Binary heaps (min and max)
- Heap operations and heap sort
- Advanced heaps: Fibonacci heaps, binomial heaps
- Priority queue applications

**Labs**:
1. Implement binary heap from scratch with heapify operations
2. Build heap sort and analyze in-place vs. non-in-place variants
3. Create a priority queue supporting various operations
4. Implement d-ary heaps and find optimal d for different workloads
5. Application lab: build a task scheduler using priority queues
6. Performance comparison: heap vs. BST for priority queue operations

**CLRS Reference**: Chapters 6, 19

---

### Module 4: Sorting & Selection (3 weeks)
**Why this matters**: Sorting is everywhere, and different algorithms shine in different contexts. Understanding when to use which is key.

**Topics**:
- Comparison sorts: quicksort, mergesort, heapsort
- Lower bounds for comparison sorting
- Linear-time sorts: counting, radix, bucket sort
- Selection algorithms (finding kth element)

**Labs**:
1. Implement 5+ sorting algorithms from scratch
2. Partition strategies lab: test different pivot selection methods for quicksort
3. Build adaptive sorting algorithms that detect pre-sorted data
4. External sorting simulation: sort data larger than memory
5. Implement randomized and deterministic selection algorithms
6. Comprehensive performance comparison across different input distributions
7. Custom sort: design a hybrid algorithm optimized for a specific data type

**CLRS Reference**: Chapters 7-9

---

### Module 5: Advanced Data Structures (3 weeks)
**Why this matters**: Specialized structures for specific problem domains can give massive performance wins.

**Topics**:
- Disjoint-set (Union-Find) with optimizations
- Tries and suffix trees
- Segment trees and Fenwick trees (Binary Indexed Trees)
- Skip lists
- Bloom filters and probabilistic structures

**Labs**:
1. Implement Union-Find with path compression and union by rank
2. Build a trie for autocomplete functionality
3. Create segment tree for range query operations
4. Implement skip list and compare with balanced BST
5. Build a Bloom filter and measure false positive rates
6. Application lab: use suffix tree for string pattern matching
7. Design challenge: create a data structure for a specific real-world problem

**CLRS Reference**: Chapters 21, some topics from advanced sources

---

### Module 6: Graph Algorithms I - Fundamentals (3 weeks)
**Why this matters**: Graphs model relationships and networks. These algorithms are used everywhere from maps to social networks to compilers.

**Topics**:
- Graph representations (adjacency matrix, adjacency list, edge list)
- BFS and DFS with applications
- Topological sort
- Strongly connected components
- Minimum spanning trees (Kruskal's, Prim's)

**Labs**:
1. Implement graph representations and conversion between them
2. Build BFS and DFS with various applications (connectivity, cycle detection, etc.)
3. Implement topological sort and detect cycles in directed graphs
4. Find strongly connected components (Kosaraju's or Tarjan's algorithm)
5. Build Kruskal's and Prim's MST algorithms
6. Graph visualization: create tools to visualize graph algorithms
7. Application lab: build a dependency resolver using topological sort

**CLRS Reference**: Chapters 22-23

---

### Module 7: Graph Algorithms II - Shortest Paths (2 weeks)
**Why this matters**: Finding optimal paths is a fundamental problem with countless applications.

**Topics**:
- Single-source: Dijkstra's, Bellman-Ford
- All-pairs: Floyd-Warshall, Johnson's algorithm
- Shortest paths in DAGs
- Negative cycles

**Labs**:
1. Implement Dijkstra's algorithm with different priority queue implementations
2. Build Bellman-Ford and detect negative cycles
3. Implement Floyd-Warshall for all-pairs shortest paths
4. Optimize shortest paths in DAGs
5. A* search implementation with heuristic design
6. Comparative analysis: when does each algorithm shine?
7. Application lab: build a route planner with realistic constraints

**CLRS Reference**: Chapters 24-25

---

### Module 8: Dynamic Programming (4 weeks)
**Why this matters**: DP is a problem-solving paradigm that appears everywhere. Understanding how to recognize and solve DP problems is invaluable.

**Topics**:
- DP fundamentals: optimal substructure, overlapping subproblems
- Memoization vs. tabulation
- Classic problems: LCS, edit distance, knapsack, matrix chain multiplication
- Advanced DP: bitmask DP, DP on trees, DP optimization techniques

**Labs**:
1. Solve 20+ classic DP problems with both top-down and bottom-up approaches
2. Implement space-optimized DP solutions
3. Build a framework to automatically memoize recursive functions
4. Reconstruct solutions, not just optimal values
5. Advanced DP challenge problems
6. Create visualization tools for DP state spaces
7. Application lab: build a text editor with efficient diff/patch using LCS

**CLRS Reference**: Chapter 15

---

### Module 9: Greedy Algorithms (2 weeks)
**Why this matters**: Greedy algorithms are intuitive but proving correctness is subtle. Learning to recognize when greedy works is a valuable skill.

**Topics**:
- Greedy strategy and proof techniques
- Activity selection, fractional knapsack
- Huffman coding
- Greedy graph algorithms (MST, Dijkstra's revisited)

**Labs**:
1. Implement classic greedy algorithms
2. Build Huffman encoder/decoder for file compression
3. Greedy scheduling algorithms
4. Counterexamples: create problems where greedy fails
5. Proof practice: formally prove greedy correctness
6. Application lab: resource allocation system using greedy strategies

**CLRS Reference**: Chapter 16

---

### Module 10: Advanced Algorithm Techniques (3 weeks)
**Why this matters**: These techniques expand your problem-solving toolkit beyond basic paradigms.

**Topics**:
- Divide and conquer (Strassen's, master theorem)
- Backtracking and branch-and-bound
- Randomized algorithms (QuickSort revisited, hashing, Monte Carlo vs. Las Vegas)
- Amortized analysis techniques

**Labs**:
1. Implement Strassen's matrix multiplication and measure crossover points
2. Build a sudoku solver using backtracking
3. Create a randomized algorithm and analyze expected performance
4. Implement and analyze skip list with probabilistic guarantees
5. Advanced amortization: dynamic tables, splay trees
6. Challenge: implement a SAT solver using various techniques

**CLRS Reference**: Chapters 4, 5, 17, various chapters

---

### Module 11: String Algorithms (2 weeks)
**Why this matters**: String processing is ubiquitous in real applications, and specialized algorithms can be orders of magnitude faster than naive approaches.

**Topics**:
- String matching: naive, Rabin-Karp, KMP, Boyer-Moore
- Suffix arrays and suffix trees
- String compression (LZ77, LZ78)
- Edit distance and sequence alignment

**Labs**:
1. Implement 4+ string matching algorithms
2. Build suffix array with efficient construction
3. Create a simple compression/decompression tool
4. Implement edit distance with various costs
5. DNA sequence alignment using string algorithms
6. Performance comparison on real-world text data
7. Application lab: build a plagiarism detector

**CLRS Reference**: Chapter 32

---

### Module 12: Computational Geometry (Optional, 2 weeks)
**Why this matters**: Geometric algorithms have applications in graphics, GIS, robotics, and CAD.

**Topics**:
- Convex hull (Graham scan, Jarvis march)
- Line segment intersection
- Closest pair of points
- Range searching

**Labs**:
1. Implement convex hull algorithms with visualization
2. Build a line segment intersection detector
3. Solve closest pair problem with divide-and-conquer
4. Create a simple collision detection system
5. Application lab: computational geometry in graphics or mapping

**CLRS Reference**: Chapter 33

---

### Module 13: NP-Completeness & Approximation (2 weeks)
**Why this matters**: Understanding computational limits helps you recognize when to search for optimal solutions vs. good-enough approximations.

**Topics**:
- P vs. NP, NP-completeness
- Classic NP-complete problems (SAT, Clique, Vertex Cover, TSP)
- Reductions
- Approximation algorithms and ratios

**Labs**:
1. Implement brute force solutions for small NP-complete problems
2. Build approximation algorithms with provable bounds
3. Construct reductions between NP-complete problems
4. Heuristics lab: test various heuristics for TSP
5. Compare exact vs. approximate solutions on real instances
6. Application lab: practical problem solving with approximations

**CLRS Reference**: Chapters 34-35

---

## Language-Specific Tracks

While you can use any language for implementation, here are recommended approaches for each:

**Python**: Great for rapid prototyping and algorithm exploration. Focus on clean, readable implementations. Use profiling tools to understand performance.

**Go**: Excellent for concurrent algorithm implementations. Explore parallel versions of algorithms, focus on practical systems programming approaches.

**Zig**: Perfect for understanding low-level details and manual memory management. Build high-performance implementations with explicit control.

**Rust**: Ideal for exploring ownership and safety in data structure design. Challenge: how do you build complex structures (like graphs) in Rust's ownership model?

**Elixir**: Great for functional approaches and concurrent/distributed algorithms. Explore immutable data structures and persistent data structures.

## Projects & Capstones

After completing core modules, choose 2-3 capstone projects:

1. **Build a Database Index**: B-trees, hash indexes, LSM trees
2. **Compiler/Interpreter Component**: Parser, AST, optimization passes
3. **Network Router Simulator**: Graph algorithms, priority queues, optimization
4. **Distributed System Primitive**: Consensus algorithm, distributed hash table
5. **Game AI**: Pathfinding, decision trees, optimization
6. **Custom Data Structure Library**: Production-quality implementation with tests and benchmarks

## Learning Resources

**Primary**:
- CLRS 4th Edition (your reference)
- Your implementations and experiments

**Supplementary**:
- Algorithm Design Manual (Skiena) - great for practical insights
- Competitive Programming 3 (Halim & Halim) - problem sets
- Papers and original sources for specific algorithms

**Online Resources**:
- Visualization tools (VisuAlgo, Algorithm Visualizer)
- Competitive programming platforms (LeetCode, Codeforces) for practice problems
- Performance profiling tools for your chosen languages

## Assessment & Progress

Track your progress through:
- **Implementation checkpoints**: Did you build it? Does it work?
- **Analysis exercises**: Can you prove the complexity? Explain the tradeoffs?
- **Performance benchmarks**: How does your implementation perform?
- **Problem solving**: Can you apply these techniques to new problems?

## Time Estimates

- **Core curriculum** (Modules 0-11): 35-40 weeks at 15-20 hours/week
- **Optional modules**: +2 weeks each
- **Capstone projects**: 4-8 weeks each

This is flexible - go faster if you have more time, slower if you need it. The key is depth of understanding, not speed of completion.

## Getting Started

1. Set up your development environment with tooling for your chosen language(s)
2. Create a repository for your implementations
3. Start with Module 0 - build your analysis tools first
4. For each topic: read theory → implement → experiment → reflect
5. Keep a learning journal documenting insights and "aha moments"

## Philosophy on "Getting Stuck"

Getting stuck is part of the process. When you hit a wall:
1. Try to implement it anyway, even if inefficiently
2. Analyze what makes it hard - that's where the learning is
3. Look at hints, not solutions
4. If you must look at a solution, re-implement from scratch afterwards
5. Understand *why* the solution works, not just *that* it works

The goal is not to accumulate code, but to build intuition and problem-solving ability that will serve you for your entire career.

---

## Next Steps

This README provides the overall structure. For each module, we can create:
- Detailed topic breakdowns
- Specific lab assignments with starter code
- Problem sets and challenges
- Testing frameworks
- Recommended reading sections from CLRS
