# Lab 8.6: DP Visualization Tools

**Module**: Module 8 - Dynamic Programming  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 8.6: DP Visualization Tools

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Build tools to visualize DP computation
- Aid understanding and debugging
- Create educational animations

### Requirements

1. **Visualization features**:
   - Display DP table/state space
   - Highlight computation order
   - Show dependencies between states
   - Animate filling of DP table

2. **Support for**:
   - 1D DP problems
   - 2D DP problems
   - State space visualization
   - Recurrence visualization

3. **Output formats**:
   - Terminal-based (ASCII)
   - HTML/JavaScript interactive
   - Step-by-step animation

### Implementation Details

```python
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import numpy as np

class DPVisualizer:
    """Visualize DP computation"""
    
    def visualize_table_2d(self, dp, title="DP Table", labels=None):
        """Visualize 2D DP table as heatmap"""
        fig, ax = plt.subplots(figsize=(10, 8))
        
        # Create heatmap
        im = ax.imshow(dp, cmap='YlOrRd', aspect='auto')
        
        # Add colorbar
        plt.colorbar(im, ax=ax)
        
        # Add labels
        if labels:
            ax.set_xlabel(labels.get('x', 'Column'))
            ax.set_ylabel(labels.get('y', 'Row'))
        
        # Add grid
        ax.set_xticks(np.arange(len(dp[0])))
        ax.set_yticks(np.arange(len(dp)))
        ax.grid(which='both', color='black', linewidth=0.5)
        
        # Add values as text
        for i in range(len(dp)):
            for j in range(len(dp[0])):
                val = dp[i][j]
                if val != float('inf'):
                    text = ax.text(j, i, f'{val:.0f}',
                                 ha="center", va="center", color="black")
        
        ax.set_title(title)
        plt.tight_layout()
        return fig
    
    def animate_filling(self, problem_func, *args):
        """
        Animate the filling of DP table.
        problem_func should yield (i, j, value) for each update.
        """
        updates = list(problem_func(*args))
        
        # Determine table size
        max_i = max(u[0] for u in updates) + 1
        max_j = max(u[1] for u in updates) + 1
        
        dp = [[0] * max_j for _ in range(max_i)]
        
        fig, ax = plt.subplots()
        im = ax.imshow(dp, cmap='YlOrRd', vmin=0, 
                      vmax=max(u[2] for u in updates))
        
        def update(frame):
            i, j, val = updates[frame]
            dp[i][j] = val
            im.set_array(dp)
            ax.set_title(f'Step {frame + 1}/{len(updates)}')
            return [im]
        
        ani = animation.FuncAnimation(fig, update, frames=len(updates),
                                     interval=200, blit=True)
        return ani
    
    def print_table_ascii(self, dp, title="DP Table"):
        """Print DP table in ASCII format"""
        print(f"\n{title}")
        print("=" * 50)
        
        if isinstance(dp[0], list):
            # 2D table
            for row in dp:
                print(" ".join(f"{val:5.0f}" if val != float('inf') else "  inf" 
                             for val in row))
        else:
            # 1D table
            print(" ".join(f"{val:5.0f}" if val != float('inf') else "  inf" 
                         for val in dp))
        print()

def visualize_lcs_computation(s1, s2):
    """Visualize LCS computation step by step"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    steps = []
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            
            steps.append({
                'i': i,
                'j': j,
                'value': dp[i][j],
                'table': [row[:] for row in dp],
                'match': s1[i-1] == s2[j-1]
            })
    
    return steps

def generate_dp_html(problem_name, steps):
    """Generate interactive HTML visualization"""
    html = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>{problem_name} - DP Visualization</title>
        <style>
            .dp-table {{ border-collapse: collapse; margin: 20px; }}
            .dp-cell {{ 
                border: 1px solid #ccc;
                padding: 10px;
                width: 40px;
                height: 40px;
                text-align: center;
            }}
            .current {{ background-color: #ffeb3b; }}
            .computed {{ background-color: #c8e6c9; }}
            .controls {{ margin: 20px; }}
        </style>
    </head>
    <body>
        <h1>{problem_name}</h1>
        <div class="controls">
            <button onclick="prevStep()">Previous</button>
            <button onclick="nextStep()">Next</button>
            <span id="step-info"></span>
        </div>
        <div id="visualization"></div>
        
        <script>
            const steps = {steps};
            let currentStep = 0;
            
            function renderStep() {{
                // Render logic here
            }}
            
            function nextStep() {{
                if (currentStep < steps.length - 1) {{
                    currentStep++;
                    renderStep();
                }}
            }}
            
            function prevStep() {{
                if (currentStep > 0) {{
                    currentStep--;
                    renderStep();
                }}
            }}
            
            renderStep();
        </script>
    </body>
    </html>
    """
    return html
```

### Test Cases
- Visualize classic DP problems
- Test animation on different problems
- Export to various formats
- Interactive controls work

### Deliverables
- `dp_visualizer.{py,go,zig}` - Core visualization tools
- `ascii_viz.{py,go,zig}` - Terminal visualization
- `matplotlib_viz.{py,go,zig}` - Graphical visualization
- `html_export.{py,go,zig}` - Interactive HTML export
- `animation_examples/` - Sample animations
- `visualization_guide.md` - How to use tools

### Success Criteria
- Clear, understandable visualizations
- Helps with debugging DP solutions
- Useful for learning and teaching
- Works for various DP problems

---

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_5.md) (if exists)
- [Next Lab](./lab_8_7.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 6 of 7*
