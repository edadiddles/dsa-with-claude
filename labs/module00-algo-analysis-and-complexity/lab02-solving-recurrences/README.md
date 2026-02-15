# Lab 0.2: Solving Recurrences

**Module**: Module 0 - Foundations & Analysis  
**Duration**: 2-3 hours  
**Difficulty**: Medium

---

Lab 0.2: Solving Recurrences

**Duration**: 2-3 hours  
**Difficulty**: Medium

### Objectives
- Solve recurrences using substitution, recursion tree, and Master theorem
- Verify solutions empirically
- Understand the gap between theory and practice

### Requirements

1. **Implement 5 recursive algorithms** with different recurrence relations:
   - T(n) = T(n-1) + O(1) - Linear recursion
   - T(n) = 2T(n/2) + O(n) - Divide and conquer
   - T(n) = T(n/2) + O(1) - Binary search pattern
   - T(n) = 2T(n/2) + O(1) - Tree recursion
   - T(n) = T(n-1) + O(n) - Quadratic recursion

2. **For each algorithm**:
   - Derive the closed-form solution mathematically
   - Implement the recursive function
   - Measure actual runtime empirically
   - Compare theoretical vs. empirical complexity

3. **Create a recurrence solver tool** that:
   - Takes a recurrence relation as input
   - Applies Master theorem when applicable
   - Generates recursion tree visualization
   - Estimates the solution

### Test Cases
- Verify solutions match known complexity classes
- Test edge cases (n=0, n=1, large n)
- Compare recursive vs. iterative implementations

### Deliverables
- `recurrence_examples.{py,go,zig}` - Implementations of all 5 patterns
- `recurrence_solver.{py,go,zig}` - Tool for solving recurrences
- `analysis.md` - Mathematical derivations and empirical results
- Charts showing theoretical vs. actual performance

### Success Criteria
- All recurrences solved correctly
- Empirical data matches theoretical predictions within margin of error
- Can explain when/why empirical results diverge from theory

---

---

## Back to Module

[← Back to Module 0: Foundations & Analysis](..)

## Navigation

- [Previous Lab](../lab01-build-analysis-framework/README.md)
- [Next Lab](../lab03-cache-effects-real-hardware/README.md)
- [All Labs](../..)

---

*Part of the comprehensive DSA Curriculum*  
*Module 0, Lab 2 of 4*
