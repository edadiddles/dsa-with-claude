# Lab 9.4: Greedy Correctness Proofs

**Module**: Module 9 - Greedy Algorithms  
**Duration**: 3-4 hours  
**Difficulty**: Hard

---

Lab 9.4: Greedy Correctness Proofs

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

---

## Back to Module

[← Back to Module 9: Greedy Algorithms](../module_09/README.md)

## Navigation

- [Previous Lab](./lab_9_3.md) (if exists)
- [Next Lab](./lab_9_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 9, Lab 4 of 5*
