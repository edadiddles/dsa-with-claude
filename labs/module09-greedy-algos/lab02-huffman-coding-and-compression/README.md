# Lab 9.2: Huffman Coding & Compression

**Module**: Module 9 - Greedy Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 9.2: Huffman Coding & Compression

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

---

## Back to Module

[← Back to Module 9: Greedy Algorithms](../module_09/README.md)

## Navigation

- [Previous Lab](./lab_9_1.md) (if exists)
- [Next Lab](./lab_9_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 9, Lab 2 of 5*
