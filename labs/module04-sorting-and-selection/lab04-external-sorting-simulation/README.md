# Lab 4.4: External Sorting Simulation

**Module**: Module 4 - Sorting & Selection  
**Duration**: 5-6 hours  
**Difficulty**: Hard

---

Lab 4.4: External Sorting Simulation

**Duration**: 5-6 hours  
**Difficulty**: Hard

### Objectives
- Sort data larger than available memory
- Understand multi-pass external sorting
- Minimize disk I/O operations

### Requirements

1. **Simulate memory constraints**:
   - Limit usable memory to M elements
   - Track disk reads and writes
   - Implement external mergesort

2. **External sorting algorithm**:
   - Phase 1: Create sorted runs of size M
   - Phase 2: k-way merge with k = M/B (B = block size)
   - Minimize number of passes over data

3. **Optimizations**:
   - Replacement selection for longer initial runs
   - Polyphase merge
   - Overlapping I/O with computation

4. **Measure**:
   - Total disk I/O operations
   - Number of passes
   - Wall-clock time with simulated disk latency

### Implementation Details

```python
class DiskSimulator:
    def __init__(self, read_latency_ms=10, write_latency_ms=10):
        self.read_latency = read_latency_ms / 1000.0
        self.write_latency = write_latency_ms / 1000.0
        self.reads = 0
        self.writes = 0
        self.disk_storage = {}
    
    def read_block(self, filename, block_id):
        time.sleep(self.read_latency)
        self.reads += 1
        return self.disk_storage.get((filename, block_id), [])
    
    def write_block(self, filename, block_id, data):
        time.sleep(self.write_latency)
        self.writes += 1
        self.disk_storage[(filename, block_id)] = data[:]

def external_sort(input_file, output_file, memory_size, block_size, disk):
    # Phase 1: Create sorted runs
    runs = []
    chunk_id = 0
    
    while has_more_data(input_file):
        chunk = read_chunk(input_file, memory_size, disk)
        chunk.sort()  # In-memory sort
        
        run_file = f"run_{chunk_id}.tmp"
        write_run_to_disk(chunk, run_file, block_size, disk)
        runs.append(run_file)
        chunk_id += 1
    
    # Phase 2: k-way merge
    k = memory_size // block_size
    
    while len(runs) > 1:
        new_runs = []
        
        for i in range(0, len(runs), k):
            batch = runs[i:i+k]
            merged_file = f"merged_{len(new_runs)}.tmp"
            k_way_merge_external(batch, merged_file, memory_size, block_size, disk)
            new_runs.append(merged_file)
        
        # Clean up old runs
        for run in runs:
            delete_file(run, disk)
        
        runs = new_runs
    
    # Final run is the sorted file
    rename_file(runs[0], output_file, disk)
    
    return disk.reads, disk.writes

def k_way_merge_external(run_files, output_file, memory_size, block_size, disk):
    import heapq
    
    # Allocate buffers
    k = len(run_files)
    buffer_size = memory_size // (k + 1)  # k input buffers + 1 output buffer
    
    # Initialize input buffers
    input_buffers = []
    heap = []
    
    for i, run_file in enumerate(run_files):
        buffer = read_block_from_run(run_file, 0, buffer_size, disk)
        if buffer:
            input_buffers.append({
                'run_file': run_file,
                'buffer': buffer,
                'block_id': 0,
                'position': 0
            })
            heapq.heappush(heap, (buffer[0], i, 0))
    
    # Output buffer
    output_buffer = []
    output_block_id = 0
    
    # Merge
    while heap:
        value, run_idx, pos = heapq.heappop(heap)
        output_buffer.append(value)
        
        # Flush output buffer if full
        if len(output_buffer) >= buffer_size:
            write_block_to_file(output_file, output_block_id, output_buffer, disk)
            output_buffer = []
            output_block_id += 1
        
        # Refill input buffer if needed
        buf = input_buffers[run_idx]
        buf['position'] += 1
        
        if buf['position'] >= len(buf['buffer']):
            # Read next block
            buf['block_id'] += 1
            new_buffer = read_block_from_run(
                buf['run_file'], 
                buf['block_id'], 
                buffer_size, 
                disk
            )
            
            if new_buffer:
                buf['buffer'] = new_buffer
                buf['position'] = 0
                heapq.heappush(heap, (new_buffer[0], run_idx, 0))
        else:
            heapq.heappush(heap, (buf['buffer'][buf['position']], run_idx, buf['position']))
    
    # Flush remaining output
    if output_buffer:
        write_block_to_file(output_file, output_block_id, output_buffer, disk)

def replacement_selection(input_file, memory_size, disk):
    # Generate longer initial runs using a heap
    import heapq
    
    heap = []
    runs = []
    current_run = []
    run_id = 0
    
    # Fill initial heap
    for _ in range(memory_size):
        value = read_next_value(input_file, disk)
        if value is not None:
            heapq.heappush(heap, (value, run_id))
    
    while heap:
        value, run_tag = heapq.heappop(heap)
        
        if run_tag == run_id:
            current_run.append(value)
            
            # Read next value
            next_value = read_next_value(input_file, disk)
            if next_value is not None:
                if next_value >= value:
                    heapq.heappush(heap, (next_value, run_id))
                else:
                    heapq.heappush(heap, (next_value, run_id + 1))
        else:
            # Start new run
            write_run_to_disk(current_run, f"run_{run_id}.tmp", disk)
            runs.append(f"run_{run_id}.tmp")
            run_id += 1
            current_run = [value]
    
    if current_run:
        write_run_to_disk(current_run, f"run_{run_id}.tmp", disk)
        runs.append(f"run_{run_id}.tmp")
    
    return runs
```

### Test Cases
- Data size 10x memory size
- Data size 100x memory size
- Different block sizes
- Varying number of runs
- Compare replacement selection vs. simple run creation

### Deliverables
- `external_sort.{py,go,zig}` - Complete external sort
- `disk_simulator.{py,go,zig}` - Simulated disk I/O
- `run_generator.{py,go,zig}` - Create sorted runs
- `replacement_selection.{py,go,zig}` - Advanced run creation
- `analysis.md` - I/O analysis and optimizations
- Graphs showing I/O vs. data size

### Success Criteria
- Correctly sorts data larger than memory
- I/O count matches theoretical predictions
- Understanding of external sorting tradeoffs
- Replacement selection produces longer runs

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_3.md) (if exists)
- [Next Lab](./lab_4_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 4 of 7*
