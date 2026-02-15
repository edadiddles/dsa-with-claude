# Lab 11.4: Sequence Alignment & Edit Distance

**Module**: Module 11 - String Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 11.4: Sequence Alignment & Edit Distance

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement sequence alignment algorithms
- Understand edit distance variants
- Apply to bioinformatics problems
- Optimize alignment algorithms

### Requirements

1. **Implement alignment algorithms**:
   - **Edit distance** (Levenshtein)
   - **Needleman-Wunsch** (global alignment)
   - **Smith-Waterman** (local alignment)
   - **Longest Common Subsequence**

2. **Variants**:
   - Different scoring schemes
   - Gap penalties
   - Affine gap costs

3. **Applications**:
   - DNA sequence alignment
   - Spell checking
   - Diff utilities

### Implementation Details

```python
# EDIT DISTANCE (Review from Module 8)

def edit_distance(s1, s2):
    """
    Levenshtein distance: minimum edit operations.
    Operations: insert, delete, substitute (each cost 1).
    """
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1]
            else:
                dp[i][j] = 1 + min(
                    dp[i-1][j],      # Delete
                    dp[i][j-1],      # Insert
                    dp[i-1][j-1]     # Substitute
                )
    
    return dp[m][n]

# NEEDLEMAN-WUNSCH (Global Alignment)

def needleman_wunsch(seq1, seq2, match=1, mismatch=-1, gap=-1):
    """
    Global sequence alignment.
    Used in bioinformatics for DNA/protein alignment.
    """
    m, n = len(seq1), len(seq2)
    
    # Initialize DP table
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize first row and column with gap penalties
    for i in range(1, m + 1):
        dp[i][0] = i * gap
    for j in range(1, n + 1):
        dp[0][j] = j * gap
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            
            dp[i][j] = max(
                dp[i-1][j-1] + score_match,  # Match/mismatch
                dp[i-1][j] + gap,             # Gap in seq2
                dp[i][j-1] + gap              # Gap in seq1
            )
    
    # Traceback to get alignment
    align1, align2 = [], []
    i, j = m, n
    
    while i > 0 or j > 0:
        if i > 0 and j > 0:
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            if dp[i][j] == dp[i-1][j-1] + score_match:
                align1.append(seq1[i-1])
                align2.append(seq2[j-1])
                i -= 1
                j -= 1
                continue
        
        if i > 0 and dp[i][j] == dp[i-1][j] + gap:
            align1.append(seq1[i-1])
            align2.append('-')
            i -= 1
        else:
            align1.append('-')
            align2.append(seq2[j-1])
            j -= 1
    
    return ''.join(reversed(align1)), ''.join(reversed(align2)), dp[m][n]

# SMITH-WATERMAN (Local Alignment)

def smith_waterman(seq1, seq2, match=2, mismatch=-1, gap=-1):
    """
    Local sequence alignment: find best matching substring.
    Used for finding similar regions in sequences.
    """
    m, n = len(seq1), len(seq2)
    
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    max_score = 0
    max_pos = (0, 0)
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score_match = match if seq1[i-1] == seq2[j-1] else mismatch
            
            dp[i][j] = max(
                0,                              # Start new alignment
                dp[i-1][j-1] + score_match,    # Match/mismatch
                dp[i-1][j] + gap,               # Gap in seq2
                dp[i][j-1] + gap                # Gap in seq1
            )
            
            if dp[i][j] > max_score:
                max_score = dp[i][j]
                max_pos = (i, j)
    
    # Traceback from max position
    align1, align2 = [], []
    i, j = max_pos
    
    while i > 0 and j > 0 and dp[i][j] > 0:
        if seq1[i-1] == seq2[j-1]:
            score_match = match
        else:
            score_match = mismatch
        
        if dp[i][j] == dp[i-1][j-1] + score_match:
            align1.append(seq1[i-1])
            align2.append(seq2[j-1])
            i -= 1
            j -= 1
        elif dp[i][j] == dp[i-1][j] + gap:
            align1.append(seq1[i-1])
            align2.append('-')
            i -= 1
        else:
            align1.append('-')
            align2.append(seq2[j-1])
            j -= 1
    
    return ''.join(reversed(align1)), ''.join(reversed(align2)), max_score

# AFFINE GAP PENALTIES

def affine_gap_alignment(seq1, seq2, match=1, mismatch=-1, gap_open=-2, gap_extend=-1):
    """
    Sequence alignment with affine gap penalties.
    Gap cost = gap_open + k * gap_extend for gap of length k.
    More biologically realistic.
    """
    m, n = len(seq1), len(seq2)
    
    # Three DP tables
    M = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Match/mismatch
    X = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Gap in seq1
    Y = [[float('-inf')] * (n + 1) for _ in range(m + 1)]  # Gap in seq2
    
    # Initialize
    M[0][0] = 0
    for i in range(1, m + 1):
        Y[i][0] = gap_open + i * gap_extend
    for j in range(1, n + 1):
        X[0][j] = gap_open + j * gap_extend
    
    # Fill tables
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            score = match if seq1[i-1] == seq2[j-1] else mismatch
            
            M[i][j] = score + max(M[i-1][j-1], X[i-1][j-1], Y[i-1][j-1])
            X[i][j] = max(
                M[i][j-1] + gap_open + gap_extend,
                X[i][j-1] + gap_extend
            )
            Y[i][j] = max(
                M[i-1][j] + gap_open + gap_extend,
                Y[i-1][j] + gap_extend
            )
    
    return max(M[m][n], X[m][n], Y[m][n])

# DIFF UTILITY

def diff_lines(file1_lines, file2_lines):
    """
    Compute diff between two files using LCS.
    Similar to Unix diff command.
    """
    m, n = len(file1_lines), len(file2_lines)
    
    # Compute LCS
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if file1_lines[i-1] == file2_lines[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    # Traceback to generate diff
    diff = []
    i, j = m, n
    
    while i > 0 or j > 0:
        if i > 0 and j > 0 and file1_lines[i-1] == file2_lines[j-1]:
            diff.append(('  ', file1_lines[i-1]))
            i -= 1
            j -= 1
        elif j > 0 and (i == 0 or dp[i][j-1] >= dp[i-1][j]):
            diff.append(('+', file2_lines[j-1]))
            j -= 1
        else:
            diff.append(('-', file1_lines[i-1]))
            i -= 1
    
    return list(reversed(diff))
```

### Test Cases
- DNA sequences
- Protein sequences
- Similar strings with typos
- File comparison
- Performance on long sequences

### Deliverables
- `edit_distance.{py,go,zig}` - Various edit distances
- `needleman_wunsch.{py,go,zig}` - Global alignment
- `smith_waterman.{py,go,zig}` - Local alignment
- `affine_gaps.{py,go,zig}` - Affine gap penalties
- `diff_utility.{py,go,zig}` - File diff tool
- `bioinformatics_applications.md` - DNA/protein analysis
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All alignment algorithms produce correct results
- Understanding of scoring schemes
- Can apply to real biological sequences
- Diff utility matches expected output

---

---

## Back to Module

[← Back to Module 11: String Algorithms](../module_11/README.md)

## Navigation

- [Previous Lab](./lab_11_3.md) (if exists)
- [Next Lab](./lab_11_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 11, Lab 4 of 5*
