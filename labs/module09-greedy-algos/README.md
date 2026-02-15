# Module 9: Greedy Algorithms

**Duration**: 2 weeks  
**Difficulty**: Medium

## Overview

Greedy algorithms build solutions incrementally by making locally optimal choices at each step, hoping to find a global optimum. Unlike dynamic programming, greedy algorithms don't reconsider choices once made. This module teaches you when greedy works, how to prove correctness, and when to recognize that DP is needed instead.

**Why this matters**: Many real-world optimization problems have greedy solutions - task scheduling, resource allocation, compression, network routing. Knowing when greedy works saves you from unnecessary complexity. More importantly, understanding when greedy fails teaches you to recognize problems requiring DP or other approaches.

## Learning Objectives

- Recognize greedy algorithm opportunities
- Prove greedy correctness using exchange arguments
- Understand the greedy choice property and optimal substructure
- Distinguish between problems solvable by greedy vs. DP
- Identify and construct counterexamples for incorrect greedy approaches
- Master classic greedy algorithms (Huffman, activity selection, etc.)

## Topics Covered

- Greedy choice property and optimal substructure
- Activity selection and interval scheduling
- Huffman coding and compression
- Fractional knapsack vs. 0/1 knapsack
- Job sequencing and scheduling
- Minimum spanning trees (Kruskal's and Prim's - from Module 6)
- Shortest paths (Dijkstra's - from Module 7)
- Exchange arguments and correctness proofs
- Greedy vs. DP decision making

---

## Lab 9.1: Classic Greedy Problems

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement fundamental greedy algorithms
- Understand greedy choice property through examples
- Build intuition for when greedy works

### Requirements

1. **Implement classic greedy algorithms**:
   - **Activity Selection** (interval scheduling)
   - **Fractional Knapsack**
   - **Coin Change** (greedy version)
   - **Job Sequencing with Deadlines**
   - **Minimum Platforms** (train scheduling)
   - **Gas Station** (circular tour)

2. **For each problem**:
   - Prove greedy choice property
   - Show optimal substructure
   - Analyze time complexity
   - Compare with non-greedy alternatives

3. **Demonstrate greedy fails**:
   - Show when greedy doesn't work
   - Provide counterexamples
   - Compare with DP solutions

### Implementation Details

```python
# ACTIVITY SELECTION - Classic Greedy

def activity_selection(activities):
    """
    Select maximum number of non-overlapping activities.
    Activities: list of (start, finish) tuples.
    Greedy: always pick activity that finishes earliest.
    """
    if not activities:
        return []
    
    # Sort by finish time
    activities = sorted(activities, key=lambda x: x[1])
    
    selected = [activities[0]]
    last_finish = activities[0][1]
    
    for start, finish in activities[1:]:
        if start >= last_finish:
            selected.append((start, finish))
            last_finish = finish
    
    return selected

def activity_selection_proof():
    """
    Proof of correctness using exchange argument:
    
    Let A = optimal solution
    Let G = greedy solution
    
    Claim: |G| = |A|
    
    Proof:
    1. Let a1 be the first activity in G (finishes earliest)
    2. Let b1 be the first activity in A
    3. If a1 = b1, done
    4. If a1 != b1, then finish(a1) <= finish(b1)
    5. Replace b1 with a1 in A → still valid solution
    6. Repeat for remaining activities
    7. Therefore G is optimal
    """
    pass

# FRACTIONAL KNAPSACK

def fractional_knapsack(weights, values, capacity):
    """
    Fractional knapsack: can take fractions of items.
    Greedy: sort by value-to-weight ratio, take greedily.
    """
    n = len(weights)
    
    # Compute value-to-weight ratios
    items = [(values[i] / weights[i], weights[i], values[i]) 
             for i in range(n)]
    
    # Sort by ratio (descending)
    items.sort(reverse=True)
    
    total_value = 0
    remaining_capacity = capacity
    taken = []
    
    for ratio, weight, value in items:
        if remaining_capacity == 0:
            break
        
        if weight <= remaining_capacity:
            # Take whole item
            taken.append((weight, value, 1.0))
            total_value += value
            remaining_capacity -= weight
        else:
            # Take fraction
            fraction = remaining_capacity / weight
            taken.append((weight, value, fraction))
            total_value += value * fraction
            remaining_capacity = 0
    
    return total_value, taken

def fractional_vs_01_knapsack():
    """
    Demonstration that greedy works for fractional but not 0/1:
    
    Example:
    weights = [10, 20, 30]
    values = [60, 100, 120]
    capacity = 50
    
    Greedy (by ratio):
    - Item 0: ratio = 6.0, take all → 60
    - Item 1: ratio = 5.0, take all → 100
    - Total: 160
    
    Optimal for 0/1:
    - Items 1, 2: 100 + 120 = 220 (better!)
    
    Greedy fails for 0/1 knapsack!
    """
    pass

# COIN CHANGE - Greedy Version

def coin_change_greedy(coins, amount):
    """
    Coin change using greedy approach.
    WARNING: Only works for certain coin systems!
    
    Works for: US coins (1, 5, 10, 25)
    Fails for: (1, 3, 4) with amount=6
    """
    coins = sorted(coins, reverse=True)
    
    count = 0
    used = []
    
    for coin in coins:
        while amount >= coin:
            amount -= coin
            count += 1
            used.append(coin)
    
    if amount == 0:
        return count, used
    else:
        return -1, []  # Impossible

def coin_change_greedy_counterexample():
    """
    Greedy fails for coins = [1, 3, 4], amount = 6:
    
    Greedy: 4 + 1 + 1 = 3 coins
    Optimal: 3 + 3 = 2 coins
    
    Must use DP for general coin change!
    """
    coins = [1, 3, 4]
    amount = 6
    
    greedy_count, _ = coin_change_greedy(coins, amount)
    # greedy_count = 3
    
    # DP gives optimal = 2
    print("Greedy fails! Need DP for coin change.")

# JOB SEQUENCING WITH DEADLINES

def job_sequencing(jobs):
    """
    Jobs with deadlines and profits.
    Each job takes 1 unit time.
    Maximize profit.
    
    jobs: list of (id, deadline, profit)
    Greedy: sort by profit (descending), schedule as late as possible.
    """
    # Sort by profit (descending)
    jobs = sorted(jobs, key=lambda x: x[2], reverse=True)
    
    max_deadline = max(job[1] for job in jobs)
    
    # Track which time slots are filled
    slots = [-1] * max_deadline
    total_profit = 0
    scheduled = []
    
    for job_id, deadline, profit in jobs:
        # Try to schedule in latest possible slot before deadline
        for t in range(min(deadline, max_deadline) - 1, -1, -1):
            if slots[t] == -1:
                slots[t] = job_id
                total_profit += profit
                scheduled.append((job_id, t + 1, profit))
                break
    
    return total_profit, scheduled

# MINIMUM PLATFORMS

def min_platforms(arrivals, departures):
    """
    Minimum platforms needed for trains.
    arrivals[i] = arrival time of train i
    departures[i] = departure time of train i
    
    Greedy: track overlapping intervals.
    """
    events = []
    
    for arr in arrivals:
        events.append((arr, 'arrival'))
    for dep in departures:
        events.append((dep, 'departure'))
    
    # Sort events by time
    events.sort()
    
    platforms_needed = 0
    max_platforms = 0
    
    for time, event_type in events:
        if event_type == 'arrival':
            platforms_needed += 1
            max_platforms = max(max_platforms, platforms_needed)
        else:
            platforms_needed -= 1
    
    return max_platforms

# GAS STATION (Circular Tour)

def can_complete_circuit(gas, cost):
    """
    Circular route of gas stations.
    gas[i] = gas at station i
    cost[i] = gas needed to go from i to i+1
    
    Return starting station index, or -1 if impossible.
    
    Greedy insight: if total gas >= total cost, solution exists.
    Start from first station where we can begin accumulating surplus.
    """
    total_gas = sum(gas)
    total_cost = sum(cost)
    
    if total_gas < total_cost:
        return -1
    
    tank = 0
    start = 0
    
    for i in range(len(gas)):
        tank += gas[i] - cost[i]
        
        if tank < 0:
            # Can't reach next station, try starting from next
            start = i + 1
            tank = 0
    
    return start

def gas_station_proof():
    """
    Proof that greedy works:
    
    1. If total_gas < total_cost, impossible
    2. If total_gas >= total_cost, solution exists
    3. If we fail starting from station i, all stations 
       from i to current are invalid starts
    4. Start from first station after failure point
    5. If we complete loop, this is the answer
    
    Key insight: We only need to find ONE valid start.
    """
    pass

# GREEDY INTERVAL SCHEDULING VARIANTS

def weighted_interval_scheduling(intervals):
    """
    Activity selection with weights.
    This is NOT solvable by greedy - need DP!
    
    Counterexample:
    intervals = [(0, 10, 100), (1, 2, 10), (3, 4, 10)]
    
    Greedy (by finish time): (1, 2, 10), (3, 4, 10) → 20
    Optimal: (0, 10, 100) → 100
    
    Must use DP!
    """
    # This is a counterexample showing greedy fails
    pass

def interval_partitioning(intervals):
    """
    Partition intervals into minimum number of groups
    such that no intervals in same group overlap.
    
    Greedy: sort by start time, use earliest finishing group.
    """
    if not intervals:
        return []
    
    intervals = sorted(intervals)
    
    groups = []
    
    for interval in intervals:
        start, end = interval
        
        # Find group where this interval fits
        placed = False
        for group in groups:
            # Check if compatible with this group
            if start >= group[-1][1]:  # No overlap
                group.append(interval)
                placed = True
                break
        
        if not placed:
            # Need new group
            groups.append([interval])
    
    return groups

# COMPARISON: Greedy vs DP

def compare_greedy_dp():
    """
    Problems where greedy works:
    1. Activity selection
    2. Fractional knapsack
    3. Huffman coding
    4. MST (Kruskal, Prim)
    5. Dijkstra's shortest path
    
    Problems where greedy fails (need DP):
    1. 0/1 knapsack
    2. Weighted interval scheduling
    3. Longest common subsequence
    4. Edit distance
    5. General coin change
    
    Key difference: Greedy has NO reconsideration.
    DP explores all possibilities.
    """
    pass
```

### Test Cases
- Activity selection with various overlaps
- Fractional knapsack vs. 0/1 comparison
- Coin change with canonical and non-canonical systems
- Job sequencing with tight deadlines
- Platform problems with complex schedules
- Gas station with various configurations

### Deliverables
- `activity_selection.{py,go,zig}` - Complete implementation with proof
- `fractional_knapsack.{py,go,zig}` - With comparison to 0/1
- `coin_change_greedy.{py,go,zig}` - With counterexamples
- `job_sequencing.{py,go,zig}` - Job scheduling implementation
- `platform_scheduling.{py,go,zig}` - Train platform problem
- `gas_station.{py,go,zig}` - Circular tour problem
- `greedy_proofs.md` - Correctness proofs for each algorithm
- `counterexamples.md` - When greedy fails
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All greedy algorithms work correctly
- Can prove correctness using exchange arguments
- Understanding of greedy choice property
- Can identify when greedy fails
- Know difference between greedy and DP problems

---

## Lab 9.2: Huffman Coding & Compression

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement Huffman coding algorithm
- Understand optimal prefix-free codes
- Build practical compression tool

### Requirements

1. **Implement Huffman coding**:
   - Build Huffman tree from frequencies
   - Generate optimal prefix codes
   - Encode and decode messages
   - Calculate compression ratio

2. **Prove optimality**:
   - Show greedy choice property
   - Prove using exchange argument
   - Compare with fixed-length encoding

3. **Extensions**:
   - Adaptive Huffman coding
   - Canonical Huffman codes
   - Comparison with other compression

### Implementation Details

```python
import heapq
from collections import Counter, defaultdict

class HuffmanNode:
    def __init__(self, char=None, freq=0, left=None, right=None):
        self.char = char
        self.freq = freq
        self.left = left
        self.right = right
    
    def __lt__(self, other):
        return self.freq < other.freq

def build_huffman_tree(text):
    """
    Build Huffman tree from text.
    Greedy: always combine two lowest frequency nodes.
    """
    # Count frequencies
    freq = Counter(text)
    
    # Create leaf nodes
    heap = [HuffmanNode(char, f) for char, f in freq.items()]
    heapq.heapify(heap)
    
    # Build tree bottom-up
    while len(heap) > 1:
        left = heapq.heappop(heap)
        right = heapq.heappop(heap)
        
        merged = HuffmanNode(
            freq=left.freq + right.freq,
            left=left,
            right=right
        )
        
        heapq.heappush(heap, merged)
    
    return heap[0] if heap else None

def generate_codes(root):
    """Generate Huffman codes from tree"""
    if root is None:
        return {}
    
    codes = {}
    
    def dfs(node, code):
        if node.char is not None:
            # Leaf node
            codes[node.char] = code if code else '0'
        else:
            if node.left:
                dfs(node.left, code + '0')
            if node.right:
                dfs(node.right, code + '1')
    
    dfs(root, '')
    return codes

def huffman_encode(text):
    """Encode text using Huffman coding"""
    if not text:
        return '', {}, None
    
    # Build tree and generate codes
    tree = build_huffman_tree(text)
    codes = generate_codes(tree)
    
    # Encode text
    encoded = ''.join(codes[char] for char in text)
    
    return encoded, codes, tree

def huffman_decode(encoded, tree):
    """Decode Huffman encoded text"""
    if not encoded or tree is None:
        return ''
    
    decoded = []
    current = tree
    
    for bit in encoded:
        if bit == '0':
            current = current.left
        else:
            current = current.right
        
        if current.char is not None:
            # Reached leaf
            decoded.append(current.char)
            current = tree
    
    return ''.join(decoded)

def compression_ratio(original, encoded):
    """Calculate compression ratio"""
    original_bits = len(original) * 8  # Assuming ASCII
    encoded_bits = len(encoded)
    
    ratio = (1 - encoded_bits / original_bits) * 100
    return ratio

def huffman_proof_of_optimality():
    """
    Proof that Huffman coding is optimal:
    
    Lemma 1: There exists an optimal code where two lowest-frequency
             characters are siblings at maximum depth.
    
    Proof by exchange argument:
    - Let x, y be lowest frequency characters
    - In any optimal tree T, let a, b be deepest siblings
    - Swap x with a, y with b
    - This doesn't increase cost (x,y have lower freq)
    - Therefore optimal tree has x,y as siblings
    
    Lemma 2: Greedy choice property
    - Combine x, y into z with freq f(x) + f(y)
    - Solve remaining problem recursively
    - This gives optimal solution
    
    Therefore Huffman is optimal!
    """
    pass

# ADAPTIVE HUFFMAN CODING

class AdaptiveHuffman:
    """
    Adaptive Huffman coding: update tree as we encode.
    Useful for streaming data.
    """
    def __init__(self):
        self.frequencies = defaultdict(int)
        self.tree = None
        self.codes = {}
    
    def update(self, char):
        """Update tree with new character"""
        self.frequencies[char] += 1
        
        # Rebuild tree (in practice, use FGK algorithm for efficiency)
        chars = list(self.frequencies.items())
        
        heap = [HuffmanNode(c, f) for c, f in chars]
        heapq.heapify(heap)
        
        while len(heap) > 1:
            left = heapq.heappop(heap)
            right = heapq.heappop(heap)
            merged = HuffmanNode(freq=left.freq + right.freq, left=left, right=right)
            heapq.heappush(heap, merged)
        
        self.tree = heap[0] if heap else None
        self.codes = generate_codes(self.tree)
    
    def encode_stream(self, chars):
        """Encode stream of characters, updating tree as we go"""
        encoded = []
        
        for char in chars:
            if char in self.codes:
                encoded.append(self.codes[char])
            else:
                # First occurrence - use escape code
                encoded.append('ESC')  # Simplified
            
            self.update(char)
        
        return ''.join(encoded)

# CANONICAL HUFFMAN CODES

def canonical_huffman_codes(codes):
    """
    Convert Huffman codes to canonical form.
    Benefits: can transmit code using just lengths.
    """
    # Sort by code length, then lexicographically
    items = sorted(codes.items(), key=lambda x: (len(x[1]), x[0]))
    
    canonical = {}
    code = 0
    prev_len = 0
    
    for char, original_code in items:
        code_len = len(original_code)
        
        if code_len > prev_len:
            code <<= (code_len - prev_len)
        
        canonical[char] = format(code, f'0{code_len}b')
        code += 1
        prev_len = code_len
    
    return canonical

# COMPARISON WITH OTHER COMPRESSION

def compare_compression_methods(text):
    """
    Compare different compression approaches:
    
    1. Fixed-length encoding (ASCII): 8 bits per char
    2. Huffman coding: Variable length based on frequency
    3. Run-length encoding: Good for repetitive data
    4. LZW: Dictionary-based compression
    """
    # Huffman
    encoded_huffman, codes, tree = huffman_encode(text)
    huffman_size = len(encoded_huffman)
    
    # Fixed-length
    fixed_size = len(text) * 8
    
    # Run-length (simplified)
    rle_size = len(run_length_encode(text)) * 8
    
    print(f"Original size: {fixed_size} bits")
    print(f"Huffman: {huffman_size} bits ({compression_ratio(text, encoded_huffman):.1f}% reduction)")
    print(f"Run-length: {rle_size} bits")

def run_length_encode(text):
    """Simple run-length encoding"""
    if not text:
        return ''
    
    encoded = []
    current_char = text[0]
    count = 1
    
    for char in text[1:]:
        if char == current_char:
            count += 1
        else:
            encoded.append(f"{count}{current_char}")
            current_char = char
            count = 1
    
    encoded.append(f"{count}{current_char}")
    return ''.join(encoded)

# PRACTICAL FILE COMPRESSION

class HuffmanCompressor:
    """Practical file compression using Huffman coding"""
    
    def compress_file(self, input_path, output_path):
        """Compress a file"""
        with open(input_path, 'r') as f:
            text = f.read()
        
        if not text:
            return
        
        # Encode
        encoded, codes, tree = huffman_encode(text)
        
        # Save compressed data with header
        with open(output_path, 'wb') as f:
            # Write header: code table
            self._write_header(f, codes)
            
            # Write encoded data as bytes
            self._write_bits(f, encoded)
        
        original_size = len(text) * 8
        compressed_size = len(encoded)
        
        print(f"Compressed: {original_size} → {compressed_size} bits")
        print(f"Compression ratio: {compression_ratio(text, encoded):.1f}%")
    
    def decompress_file(self, input_path, output_path):
        """Decompress a file"""
        with open(input_path, 'rb') as f:
            # Read header
            codes = self._read_header(f)
            
            # Rebuild tree from codes
            tree = self._rebuild_tree(codes)
            
            # Read encoded bits
            encoded = self._read_bits(f)
        
        # Decode
        decoded = huffman_decode(encoded, tree)
        
        with open(output_path, 'w') as f:
            f.write(decoded)
    
    def _write_header(self, f, codes):
        """Write code table to file"""
        # Simplified - in practice use more efficient format
        import pickle
        pickle.dump(codes, f)
    
    def _read_header(self, f):
        """Read code table from file"""
        import pickle
        return pickle.load(f)
    
    def _rebuild_tree(self, codes):
        """Rebuild Huffman tree from codes"""
        root = HuffmanNode()
        
        for char, code in codes.items():
            current = root
            for bit in code:
                if bit == '0':
                    if current.left is None:
                        current.left = HuffmanNode()
                    current = current.left
                else:
                    if current.right is None:
                        current.right = HuffmanNode()
                    current = current.right
            current.char = char
        
        return root
    
    def _write_bits(self, f, bits):
        """Write bit string to file as bytes"""
        # Pad to byte boundary
        padding = (8 - len(bits) % 8) % 8
        bits += '0' * padding
        
        # Write padding length
        f.write(bytes([padding]))
        
        # Convert bits to bytes
        for i in range(0, len(bits), 8):
            byte = int(bits[i:i+8], 2)
            f.write(bytes([byte]))
    
    def _read_bits(self, f):
        """Read bits from file"""
        padding = f.read(1)[0]
        
        bits = ''
        while True:
            byte = f.read(1)
            if not byte:
                break
            bits += format(byte[0], '08b')
        
        # Remove padding
        if padding > 0:
            bits = bits[:-padding]
        
        return bits
```

### Test Cases
- Text with uniform frequency (all chars equal)
- Text with skewed frequency (some chars very common)
- Binary data
- Long texts (measure actual compression)
- Edge cases (single character, empty string)

### Deliverables
- `huffman_coding.{py,go,zig}` - Complete implementation
- `adaptive_huffman.{py,go,zig}` - Adaptive version
- `canonical_codes.{py,go,zig}` - Canonical Huffman
- `file_compressor.{py,go,zig}` - Practical file compression
- `huffman_visualizer.{py,go,zig}` - Visualize Huffman tree
- `optimality_proof.md` - Formal proof of optimality
- `compression_comparison.md` - Compare with other methods
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Correct Huffman tree construction
- Successful encode/decode round trip
- Understanding of optimality proof
- Practical compression works on files
- Can explain why Huffman is greedy

---

## Lab 9.3: Greedy vs. DP Decision Making

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Distinguish greedy-solvable from DP-required problems
- Build decision framework
- Practice on ambiguous problems

### Requirements

1. **Analyze problem characteristics**:
   - Greedy choice property presence
   - Optimal substructure verification
   - Local vs. global optimization

2. **Implement both approaches** for comparison:
   - Coin change (greedy fails, DP works)
   - Weighted job scheduling (greedy fails, DP works)
   - Activity selection (greedy works)
   - Minimum spanning tree (greedy works)

3. **Decision framework**:
   - Checklist for identifying greedy problems
   - Red flags indicating DP needed
   - Testing approach validity

### Implementation Details

```python
# DECISION FRAMEWORK

def is_greedy_applicable(problem_description):
    """
    Checklist for greedy algorithms:
    
    ✓ Greedy Choice Property:
      - Can we make a locally optimal choice?
      - Does this choice lead to global optimum?
      - Can we prove it with exchange argument?
    
    ✓ Optimal Substructure:
      - After greedy choice, is remaining problem similar?
      - Can we solve recursively?
    
    ✓ No Reconsideration:
      - Once choice is made, never revisit?
      - No need to explore alternatives?
    
    ✗ Red flags (use DP instead):
      - Multiple interdependent choices
      - Need to consider all possibilities
      - Greedy choice depends on future choices
      - Can construct counterexample
    """
    pass

# EXAMPLE 1: Coin Change

def coin_change_comparison(coins, amount):
    """
    Compare greedy vs. DP for coin change.
    
    Greedy approach: Always take largest coin.
    DP approach: Try all possibilities.
    """
    # Greedy
    def greedy(coins, amount):
        coins = sorted(coins, reverse=True)
        count = 0
        for coin in coins:
            while amount >= coin:
                amount -= coin
                count += 1
        return count if amount == 0 else -1
    
    # DP
    def dp(coins, amount):
        dp_table = [float('inf')] * (amount + 1)
        dp_table[0] = 0
        
        for i in range(1, amount + 1):
            for coin in coins:
                if coin <= i:
                    dp_table[i] = min(dp_table[i], 1 + dp_table[i - coin])
        
        return dp_table[amount] if dp_table[amount] != float('inf') else -1
    
    greedy_result = greedy(coins, amount)
    dp_result = dp(coins, amount)
    
    return {
        'greedy': greedy_result,
        'dp': dp_result,
        'greedy_optimal': greedy_result == dp_result
    }

def when_greedy_coin_change_works():
    """
    Greedy works for canonical coin systems:
    - US coins: [1, 5, 10, 25]
    - Euro coins: [1, 2, 5, 10, 20, 50]
    
    Canonical system: greedy always gives optimal solution.
    
    Non-canonical example: [1, 3, 4]
    - For amount=6: greedy gives 3 coins (4+1+1)
    - DP gives 2 coins (3+3)
    """
    # US coins - greedy works
    result1 = coin_change_comparison([1, 5, 10, 25], 41)
    print(f"US coins, amount=41: {result1}")
    
    # Non-canonical - greedy fails
    result2 = coin_change_comparison([1, 3, 4], 6)
    print(f"[1,3,4], amount=6: {result2}")

# EXAMPLE 2: Weighted Interval Scheduling

def weighted_interval_comparison(intervals):
    """
    intervals: [(start, end, weight), ...]
    
    Greedy approaches (all fail):
    1. Earliest finish time
    2. Shortest interval
    3. Highest weight
    
    DP approach (works):
    Sort by finish time, DP on whether to include each interval.
    """
    # Greedy by weight (fails)
    def greedy_by_weight(intervals):
        sorted_intervals = sorted(intervals, key=lambda x: x[2], reverse=True)
        
        selected = []
        for interval in sorted_intervals:
            start, end, weight = interval
            
            # Check if compatible with selected
            compatible = True
            for s, e, w in selected:
                if not (end <= s or start >= e):
                    compatible = False
                    break
            
            if compatible:
                selected.append(interval)
        
        return sum(w for s, e, w in selected)
    
    # DP (correct)
    def dp_weighted_intervals(intervals):
        # Sort by finish time
        intervals = sorted(intervals, key=lambda x: x[1])
        n = len(intervals)
        
        # Find latest non-overlapping interval for each
        prev = [-1] * n
        for i in range(n):
            for j in range(i - 1, -1, -1):
                if intervals[j][1] <= intervals[i][0]:
                    prev[i] = j
                    break
        
        # DP
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            # Don't include interval i-1
            without = dp[i-1]
            
            # Include interval i-1
            weight = intervals[i-1][2]
            with_interval = weight + (dp[prev[i-1] + 1] if prev[i-1] != -1 else 0)
            
            dp[i] = max(without, with_interval)
        
        return dp[n]
    
    greedy_result = greedy_by_weight(intervals)
    dp_result = dp_weighted_intervals(intervals)
    
    return {
        'greedy': greedy_result,
        'dp': dp_result,
        'greedy_optimal': greedy_result == dp_result
    }

# TESTING FRAMEWORK

class GreedyValidator:
    """Framework for testing greedy algorithms"""
    
    def __init__(self, greedy_func, brute_force_func):
        self.greedy = greedy_func
        self.brute_force = brute_force_func
    
    def validate(self, test_cases):
        """Test greedy against brute force"""
        results = []
        
        for test_input in test_cases:
            greedy_output = self.greedy(test_input)
            optimal_output = self.brute_force(test_input)
            
            is_optimal = (greedy_output == optimal_output)
            
            results.append({
                'input': test_input,
                'greedy': greedy_output,
                'optimal': optimal_output,
                'correct': is_optimal
            })
        
        return results
    
    def find_counterexample(self, input_generator, max_tries=1000):
        """Try to find counterexample where greedy fails"""
        for _ in range(max_tries):
            test_input = input_generator()
            
            greedy_output = self.greedy(test_input)
            optimal_output = self.brute_force(test_input)
            
            if greedy_output != optimal_output:
                return {
                    'found': True,
                    'input': test_input,
                    'greedy': greedy_output,
                    'optimal': optimal_output
                }
        
        return {'found': False}

# PROBLEM CLASSIFICATION

def classify_problem(problem_characteristics):
    """
    Classify whether problem is greedy or DP based on characteristics.
    
    Greedy indicators:
    - Sorting helps
    - Local optimal leads to global optimal
    - Proof by exchange argument exists
    - Never need to reconsider choices
    
    DP indicators:
    - Need to explore all possibilities
    - Optimal depends on solving subproblems optimally
    - Can't determine best choice without looking ahead
    - Overlapping subproblems exist
    """
    score_greedy = 0
    score_dp = 0
    
    if problem_characteristics.get('sorting_helps'):
        score_greedy += 2
    
    if problem_characteristics.get('local_optimal_works'):
        score_greedy += 3
    
    if problem_characteristics.get('exchange_argument_exists'):
        score_greedy += 3
    
    if problem_characteristics.get('overlapping_subproblems'):
        score_dp += 3
    
    if problem_characteristics.get('need_all_possibilities'):
        score_dp += 3
    
    if problem_characteristics.get('future_dependent'):
        score_dp += 2
    
    if score_greedy > score_dp:
        return 'greedy'
    elif score_dp > score_greedy:
        return 'dp'
    else:
        return 'unclear'

# EXAMPLES OF CLASSIFICATION

def problem_examples():
    """
    Clear greedy problems:
    1. Activity selection: Sort by finish, take earliest
    2. Huffman coding: Always merge two smallest
    3. MST (Kruskal): Always add cheapest edge
    4. Fractional knapsack: Take by value/weight ratio
    
    Clear DP problems:
    1. 0/1 knapsack: Can't decide without seeing future
    2. Longest common subsequence: Need all possibilities
    3. Edit distance: Multiple ways to transform
    4. Matrix chain multiplication: Order matters
    
    Ambiguous (need analysis):
    1. Coin change: Depends on coin system
    2. Interval scheduling: Depends on whether weighted
    3. Subset sum: Generally DP, but special cases greedy
    """
    pass
```

### Test Cases
- Problems with known greedy solutions
- Problems with known DP solutions
- Ambiguous problems requiring analysis
- Generate counterexamples systematically

### Deliverables
- `greedy_vs_dp_framework.{py,go,zig}` - Decision framework
- `coin_change_analysis.{py,go,zig}` - Greedy vs. DP comparison
- `interval_scheduling_analysis.{py,go,zig}` - Weighted vs. unweighted
- `greedy_validator.{py,go,zig}` - Testing framework
- `problem_classifier.{py,go,zig}` - Automatic classification
- `decision_guide.md` - How to choose approach
- `examples_catalog.md` - Categorized problem examples

### Success Criteria
- Can distinguish greedy from DP problems
- Understanding of why certain problems need DP
- Can construct counterexamples
- Know how to validate greedy algorithms

---

## Lab 9.4: Greedy Correctness Proofs

**Duration**: 3-4 hours  
**Difficulty**: Hard

### Objectives
- Master exchange argument technique
- Write formal correctness proofs
- Understand proof patterns

### Requirements

1. **Learn proof techniques**:
   - Exchange argument
   - Greedy stays ahead
   - Structural induction

2. **Prove correctness for**:
   - Activity selection
   - Huffman coding
   - Fractional knapsack
   - Job sequencing
   - Dijkstra's algorithm (review from Module 7)

3. **Practice proof writing**:
   - State theorem clearly
   - Provide rigorous proof
   - Handle edge cases

### Implementation Details

```python
# EXCHANGE ARGUMENT TEMPLATE

def exchange_argument_template():
    """
    Template for exchange argument proofs:
    
    1. Let G = greedy solution
    2. Let O = optimal solution
    3. Assume G ≠ O
    4. Find first difference between G and O
    5. Show we can exchange elements
    6. Prove exchange doesn't worsen solution
    7. Repeat to show G is optimal
    
    Key insight: Transform O to look like G without worsening cost.
    """
    pass

# PROOF 1: Activity Selection

def activity_selection_proof():
    """
    Theorem: Activity selection by earliest finish time is optimal.
    
    Proof (by exchange argument):
    
    Let G = greedy solution (sorted by finish time)
    Let O = optimal solution
    
    Claim: |G| = |O| (both have same number of activities)
    
    Proof:
    1. Let g1 be first activity in G (earliest finish)
    2. Let o1 be first activity in O
    
    Case 1: g1 = o1
       Recurse on remaining problem. Done.
    
    Case 2: g1 ≠ o1
       - finish(g1) ≤ finish(o1) (by greedy choice)
       - Replace o1 with g1 in O
       - Still valid solution (g1 finishes no later than o1)
       - Same number of activities
    
    3. Repeat for remaining activities
    
    Therefore G is optimal. QED.
    
    Time complexity: O(n log n) for sorting
    """
    pass

def activity_selection_greedy_stays_ahead():
    """
    Alternative proof: "Greedy Stays Ahead"
    
    Invariant: After selecting k activities, greedy finishes 
               no later than any optimal solution.
    
    Proof by induction:
    
    Base case (k=0): True (both start at time 0)
    
    Inductive step:
    Assume after k activities, greedy finishes at time t
    and optimal finishes at time t' where t ≤ t'.
    
    Greedy chooses activity finishing earliest ≥ t
    Optimal must choose activity finishing ≥ t' ≥ t
    
    Therefore greedy still finishes no later. QED.
    """
    pass

# PROOF 2: Huffman Coding

def huffman_proof():
    """
    Theorem: Huffman coding produces optimal prefix-free code.
    
    Proof (by structural induction on tree):
    
    Lemma 1: Optimal tree has two lowest-frequency chars as siblings.
    
    Proof of Lemma 1:
    1. Let x, y be lowest frequency chars
    2. In any optimal tree T, let a, b be deepest siblings
    3. Since x, y have lowest freq: freq(x) ≤ freq(a), freq(y) ≤ freq(b)
    4. Exchange x with a, y with b
    5. Cost change = (freq(a) - freq(x)) * depth(a) 
                    + (freq(b) - freq(y)) * depth(b)
                    ≤ 0
    6. Therefore exchange doesn't increase cost
    7. So optimal tree exists with x, y as deepest siblings. QED Lemma 1.
    
    Main theorem:
    1. Let x, y be lowest frequency
    2. Create new char z with freq(z) = freq(x) + freq(y)
    3. Optimal solution for {z, rest} gives optimal for {x, y, rest}
       (by Lemma 1 and induction)
    4. Greedy combines x, y → this is optimal choice
    5. Recurse on smaller problem
    
    Therefore Huffman is optimal. QED.
    """
    pass

# PROOF 3: Fractional Knapsack

def fractional_knapsack_proof():
    """
    Theorem: Taking items by value/weight ratio is optimal.
    
    Proof (by exchange argument):
    
    Let G = greedy solution (sorted by ratio)
    Let O = optimal solution
    
    Assume G ≠ O
    
    1. Let i be first item where G and O differ in amount taken
    2. Let g_i = amount of i in G, o_i = amount of i in O
    3. Assume g_i > o_i (G takes more of i than O)
    
    4. O must take more of some item j where ratio(j) < ratio(i)
    5. Exchange: take less of j, more of i in O
       - Value change = Δ * (ratio(i) - ratio(j)) > 0
       - Weight stays same
    6. This improves O, contradiction!
    
    Therefore G is optimal. QED.
    
    Key: Items with higher ratio always preferable.
    """
    pass

# PROOF 4: Job Sequencing

def job_sequencing_proof():
    """
    Theorem: Scheduling jobs by profit (descending) is optimal.
    
    Proof (by exchange argument):
    
    Schedule: array of time slots
    
    Let G = greedy schedule (by profit)
    Let O = optimal schedule
    
    Claim: profit(G) = profit(O)
    
    Proof:
    1. Let j be highest profit job in O but not in G
    2. j must have been scheduled in G (or rejected for valid reason)
    
    Case 1: j scheduled at different time in G
       - Can swap without changing feasibility
       - Same profit
    
    Case 2: j not scheduled in G
       - Some lower profit job k scheduled in j's slot
       - Swap j and k
       - Increases profit, contradiction!
    
    Therefore G is optimal. QED.
    """
    pass

# PROOF 5: MST (Kruskal's Algorithm)

def kruskal_proof():
    """
    Theorem: Kruskal's algorithm produces MST.
    
    Proof (by cut property):
    
    Cut property: For any cut, minimum weight edge crossing 
                  cut is in some MST.
    
    Kruskal's algorithm:
    1. Sort edges by weight
    2. Add edge if doesn't create cycle
    
    Proof:
    At each step, edge e = (u,v) is added.
    - Consider cut S vs V-S where u ∈ S, v ∈ V-S
    - e is minimum weight edge crossing cut (processed in order)
    - By cut property, e is in some MST
    - Greedy choice is safe
    
    By induction, all greedy choices are safe.
    Therefore Kruskal produces MST. QED.
    """
    pass

# COMMON PROOF PATTERNS

class ProofPatterns:
    """Common patterns in greedy proofs"""
    
    @staticmethod
    def exchange_argument_pattern():
        """
        1. Assume greedy ≠ optimal
        2. Find first difference
        3. Exchange elements
        4. Show exchange is safe (doesn't worsen)
        5. Repeat to transform optimal → greedy
        6. Conclude greedy is optimal
        """
        pass
    
    @staticmethod
    def greedy_stays_ahead_pattern():
        """
        1. Define "ahead" metric
        2. Prove greedy starts ahead
        3. Prove if greedy ahead, stays ahead after next choice
        4. By induction, greedy always ahead
        5. Conclude greedy is optimal
        """
        pass
    
    @staticmethod
    def structural_induction_pattern():
        """
        1. Prove base case (smallest problem)
        2. Assume greedy works for smaller problems
        3. Show greedy choice for current problem
        4. Reduce to smaller problem
        5. By induction, greedy works for all sizes
        """
        pass

# FORMAL PROOF TEMPLATE

def formal_proof_template(problem, greedy_algorithm):
    """
    Template for writing formal greedy proofs:
    
    Theorem: [State what you're proving]
    
    Proof:
    
    1. Setup:
       - Define notation
       - Let G = greedy solution
       - Let O = optimal solution
    
    2. Assumption:
       - Assume G ≠ O (for contradiction or exchange)
    
    3. Analysis:
       - Identify first difference
       - Describe exchange
       - Prove exchange is safe
    
    4. Conclusion:
       - Show G is optimal
       - State time complexity
    
    QED.
    """
    pass
```

### Test Cases
- Verify proofs with small examples
- Check edge cases in proofs
- Test counterexamples to failed "greedy" algorithms

### Deliverables
- `proof_templates.md` - Templates for each proof type
- `activity_selection_proof.md` - Complete formal proof
- `huffman_proof.md` - Complete formal proof
- `fractional_knapsack_proof.md` - Complete formal proof
- `job_sequencing_proof.md` - Complete formal proof
- `kruskal_proof.md` - Complete formal proof
- `proof_patterns.md` - Common patterns guide
- `proof_checker.{py,go,zig}` - Validate proofs with examples

### Success Criteria
- Can write exchange argument proofs
- Understanding of greedy stays ahead technique
- Can identify proof pattern for new problems
- Proofs are rigorous and complete

---

## Lab 9.5: Counterexamples & When Greedy Fails

**Duration**: 2-3 hours  
**Difficulty**: Medium

### Objectives
- Systematically find counterexamples
- Understand why greedy fails
- Build intuition for problem difficulty

### Requirements

1. **Catalog of failures**:
   - 0/1 knapsack
   - Weighted interval scheduling
   - Longest path in graph
   - Traveling salesman
   - General coin change

2. **Counterexample construction**:
   - Systematic search
   - Minimal counterexamples
   - Explain why greedy fails

3. **Analysis of failure modes**:
   - Irreversible choices
   - Short-term vs. long-term optimization
   - Dependencies between choices

### Implementation Details

```python
# COUNTEREXAMPLE CATALOG

class GreedyCounterexamples:
    """Collection of problems where greedy fails"""
    
    @staticmethod
    def knapsack_01():
        """
        0/1 Knapsack: Greedy by ratio fails
        
        Items: [(weight, value)]
        1. (10, 60) → ratio = 6.0
        2. (20, 100) → ratio = 5.0
        3. (30, 120) → ratio = 4.0
        
        Capacity = 50
        
        Greedy (by ratio): Take 1, 2 → value = 160
        Optimal: Take 2, 3 → value = 220
        
        Why greedy fails:
        - Committing to item 1 (high ratio) prevents taking 2+3
        - Can't undo choice of item 1
        - Need to explore all possibilities (DP)
        """
        return {
            'items': [(10, 60), (20, 100), (30, 120)],
            'capacity': 50,
            'greedy_value': 160,
            'optimal_value': 220,
            'reason': 'Irreversible early choice'
        }
    
    @staticmethod
    def weighted_interval_scheduling():
        """
        Weighted Interval Scheduling: All greedy heuristics fail
        
        Intervals: [(start, end, weight)]
        1. (0, 100, 50)
        2. (1, 2, 40)
        3. (3, 4, 40)
        
        Greedy by finish: Take 2, 3 → 80
        Greedy by weight: Take 1 → 50
        Optimal: Take 2, 3 → 80? Or take 1 → 50?
        
        Actually, this example: 2+3 = 80 is optimal.
        
        Better example:
        1. (0, 10, 100)
        2. (1, 2, 10)
        3. (3, 4, 10)
        
        Greedy by finish: 2, 3 → 20
        Greedy by weight: 1 → 100
        Optimal: 1 → 100
        
        But this doesn't show greedy by weight fails...
        
        Actual counterexample:
        1. (0, 5, 6)
        2. (1, 3, 4)
        3. (4, 8, 4)
        4. (5, 7, 3)
        5. (6, 9, 2)
        
        Greedy by weight: 1, 3, 5 → 12
        Optimal: 2, 4, 5 → 9? No...
        
        Need: 1, 4 vs 2, 3
        """
        return {
            'intervals': [(0, 10, 100), (1, 2, 10), (3, 4, 10)],
            'greedy_by_finish': 20,
            'greedy_by_weight': 100,
            'optimal': 100,
            'reason': 'Multiple greedy strategies, none guaranteed'
        }
    
    @staticmethod
    def coin_change():
        """
        Coin Change: Greedy fails for non-canonical systems
        
        Coins: [1, 3, 4]
        Amount: 6
        
        Greedy: 4, 1, 1 → 3 coins
        Optimal: 3, 3 → 2 coins
        
        Why greedy fails:
        - Taking largest coin (4) prevents optimal solution
        - Need to try not taking 4
        - Requires exploring alternatives (DP)
        """
        return {
            'coins': [1, 3, 4],
            'amount': 6,
            'greedy': 3,
            'optimal': 2,
            'reason': 'Largest coin blocks better solution'
        }
    
    @staticmethod
    def longest_path():
        """
        Longest Path in Graph: All greedy strategies fail
        
        Graph with weighted edges.
        Greedy: always take heaviest edge
        
        Problem: May get stuck in local optimum
        
        Example:
            A --10--> B --1--> C
            |                  |
            +-------100------->+
        
        Path A → C:
        Greedy: A → B → C = 11
        Optimal: A → C = 100
        
        Why greedy fails:
        - First choice (A→B) looks good locally
        - Blocks globally optimal path
        - NP-hard problem, no greedy solution
        """
        pass
    
    @staticmethod
    def tsp():
        """
        Traveling Salesman Problem: Greedy fails
        
        Nearest neighbor heuristic:
        - Start at city
        - Go to nearest unvisited city
        - Repeat
        
        This can give arbitrarily bad solutions!
        
        Example where nearest neighbor fails:
        
        Cities arranged in cross pattern with center:
        - Center has distance 1 to all outer cities
        - Outer cities have distance 2 to each other
        
        Greedy from outer city: visits center last
        Optimal: visit center in middle of tour
        
        TSP is NP-hard, no polynomial greedy solution.
        """
        pass

# COUNTEREXAMPLE GENERATOR

class CounterexampleFinder:
    """Systematically search for counterexamples"""
    
    def find_knapsack_counterexample(self, num_items=3, max_weight=10, max_value=20):
        """
        Search for 0/1 knapsack instances where greedy fails.
        """
        import itertools
        import random
        
        for _ in range(1000):
            # Generate random items
            items = [(random.randint(1, max_weight), 
                     random.randint(1, max_value)) 
                    for _ in range(num_items)]
            
            capacity = random.randint(max_weight, num_items * max_weight)
            
            # Try greedy
            greedy_value = self._knapsack_greedy(items, capacity)
            
            # Try optimal (brute force for small instances)
            optimal_value = self._knapsack_optimal(items, capacity)
            
            if greedy_value < optimal_value:
                return {
                    'items': items,
                    'capacity': capacity,
                    'greedy': greedy_value,
                    'optimal': optimal_value
                }
        
        return None
    
    def _knapsack_greedy(self, items, capacity):
        """Greedy by value/weight ratio"""
        items_with_ratio = [(v/w, w, v) for w, v in items]
        items_with_ratio.sort(reverse=True)
        
        total_value = 0
        remaining = capacity
        
        for ratio, weight, value in items_with_ratio:
            if weight <= remaining:
                total_value += value
                remaining -= weight
        
        return total_value
    
    def _knapsack_optimal(self, items, capacity):
        """Brute force optimal"""
        n = len(items)
        max_value = 0
        
        for mask in range(1 << n):
            total_weight = 0
            total_value = 0
            
            for i in range(n):
                if mask & (1 << i):
                    total_weight += items[i][0]
                    total_value += items[i][1]
            
            if total_weight <= capacity:
                max_value = max(max_value, total_value)
        
        return max_value

# FAILURE ANALYSIS

def analyze_greedy_failure(problem, greedy_solution, optimal_solution):
    """
    Analyze why greedy failed:
    
    1. Irreversible choice:
       - Early decision blocks better later options
       - Example: 0/1 knapsack
    
    2. Missing global view:
       - Greedy sees only local information
       - Example: Longest path
    
    3. Interdependent choices:
       - Earlier choices affect later constraints
       - Example: TSP
    
    4. Multiple optimal criteria:
       - No clear "greedy choice"
       - Example: Weighted interval scheduling
    """
    failure_type = None
    
    if 'blocked_by_early_choice' in analyze_choices(greedy_solution, optimal_solution):
        failure_type = 'irreversible_choice'
    
    if 'requires_global_view' in analyze_structure(problem):
        failure_type = 'missing_global_view'
    
    if 'choices_interdependent' in analyze_dependencies(problem):
        failure_type = 'interdependent_choices'
    
    return failure_type

# MINIMAL COUNTEREXAMPLES

def find_minimal_counterexample(problem_class, greedy_algorithm):
    """
    Find smallest counterexample where greedy fails.
    
    Start with small instances, increase size until greedy fails.
    """
    size = 2
    
    while size < 20:
        # Generate random instance of given size
        instance = problem_class.generate_random(size)
        
        greedy_result = greedy_algorithm(instance)
        optimal_result = solve_optimal(instance)
        
        if greedy_result != optimal_result:
            # Found counterexample, try to minimize
            return minimize_counterexample(instance, greedy_algorithm)
        
        size += 1
    
    return None

def minimize_counterexample(instance, greedy_algorithm):
    """
    Given counterexample, try to make it smaller while preserving failure.
    """
    # Try removing elements one by one
    # If still counterexample, keep smaller version
    # Repeat until minimal
    
    current = instance
    
    improved = True
    while improved:
        improved = False
        
        for i in range(len(current)):
            smaller = current[:i] + current[i+1:]
            
            if is_counterexample(smaller, greedy_algorithm):
                current = smaller
                improved = True
                break
    
    return current

def is_counterexample(instance, greedy_algorithm):
    """Check if instance is counterexample"""
    greedy = greedy_algorithm(instance)
    optimal = solve_optimal(instance)
    return greedy != optimal
```

### Test Cases
- Verify all counterexamples are correct
- Search for new counterexamples
- Analyze why each greedy approach fails

### Deliverables
- `counterexample_catalog.{py,go,zig}` - Collection of failures
- `counterexample_finder.{py,go,zig}` - Systematic search
- `minimal_counterexamples.{py,go,zig}` - Minimization tool
- `failure_analysis.{py,go,zig}` - Analyze why greedy fails
- `greedy_failures.md` - Complete catalog with explanations
- `when_to_avoid_greedy.md` - Guide for recognizing failures

### Success Criteria
- Understanding of common greedy failure modes
- Can construct counterexamples for incorrect greedy approaches
- Know when to immediately reject greedy and use DP
- Can explain why each counterexample breaks greedy

---

## Module Resources

### CLRS References
- Chapter 16: Greedy Algorithms

### Additional Reading
- "Algorithm Design" by Kleinberg & Tardos (greedy chapter)
- "Introduction to Algorithms" greedy algorithms section
- Exchange argument proofs

### Key Takeaways

By the end of this module, you should:
1. Recognize greedy algorithm opportunities
2. Prove correctness using exchange arguments
3. Distinguish greedy from DP problems confidently
4. Construct counterexamples for invalid greedy approaches
5. Master classic greedy algorithms
6. Understand when greedy fails and why

### Next Module

**Module 10: Advanced Algorithm Techniques** - You'll explore divide-and-conquer, backtracking, randomized algorithms, and other advanced techniques.

---

## Tips for Success

1. **Always try exchange argument first** - Most greedy proofs use this
2. **Look for sorting** - Greedy often involves sorting
3. **Test with counterexamples** - Try to break your greedy algorithm
4. **Compare with DP** - Understand the difference
5. **Study the classics** - Activity selection, Huffman, MST

## Common Pitfalls

- **Assuming greedy works** - Always verify!
- **Incomplete proofs** - Exchange argument must be rigorous
- **Missing edge cases** - Test thoroughly
- **Confusing local vs. global optimum** - Greedy makes local choices
- **Not recognizing when DP needed** - Some problems just need DP
- **Weak counterexamples** - Make them minimal and clear
