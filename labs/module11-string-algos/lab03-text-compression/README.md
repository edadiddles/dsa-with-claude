# Lab 11.3: Text Compression

**Module**: Module 11 - String Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 11.3: Text Compression

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement compression algorithms
- Understand compression techniques
- Measure compression ratios
- Apply to real files

### Requirements

1. **Implement compression algorithms**:
   - **Run-Length Encoding (RLE)**
   - **LZ77**
   - **LZW (Lempel-Ziv-Welch)**
   - **Burrows-Wheeler Transform**
   - **Huffman coding** (from Module 9, review)

2. **Analysis**:
   - Compression ratio
   - Compression/decompression speed
   - Dictionary size

3. **Applications**:
   - File compression utility
   - Compare algorithms on different data types

### Implementation Details

```python
# RUN-LENGTH ENCODING

def rle_encode(data):
    """
    Run-length encoding: compress consecutive repeated characters.
    Good for data with many runs.
    """
    if not data:
        return ""
    
    encoded = []
    count = 1
    prev = data[0]
    
    for i in range(1, len(data)):
        if data[i] == prev:
            count += 1
        else:
            encoded.append(f"{count}{prev}")
            prev = data[i]
            count = 1
    
    encoded.append(f"{count}{prev}")
    
    return ''.join(encoded)

def rle_decode(encoded):
    """Decode run-length encoded data"""
    decoded = []
    i = 0
    
    while i < len(encoded):
        # Read count
        count = 0
        while i < len(encoded) and encoded[i].isdigit():
            count = count * 10 + int(encoded[i])
            i += 1
        
        # Read character
        if i < len(encoded):
            char = encoded[i]
            decoded.append(char * count)
            i += 1
    
    return ''.join(decoded)

# LZ77

def lz77_encode(data, window_size=4096, lookahead_size=18):
    """
    LZ77 compression: find longest match in sliding window.
    Output: (offset, length, next_char) tuples.
    """
    encoded = []
    i = 0
    
    while i < len(data):
        # Find longest match in window
        best_offset = 0
        best_length = 0
        
        window_start = max(0, i - window_size)
        
        for j in range(window_start, i):
            length = 0
            while (i + length < len(data) and 
                   length < lookahead_size and
                   data[j + length] == data[i + length]):
                length += 1
            
            if length > best_length:
                best_length = length
                best_offset = i - j
        
        # Output match or literal
        if best_length > 0:
            next_char = data[i + best_length] if i + best_length < len(data) else ''
            encoded.append((best_offset, best_length, next_char))
            i += best_length + (1 if next_char else 0)
        else:
            encoded.append((0, 0, data[i]))
            i += 1
    
    return encoded

def lz77_decode(encoded):
    """Decode LZ77 compressed data"""
    decoded = []
    
    for offset, length, char in encoded:
        if length > 0:
            # Copy from window
            start = len(decoded) - offset
            for i in range(length):
                decoded.append(decoded[start + i])
        
        if char:
            decoded.append(char)
    
    return ''.join(decoded)

# LZW (Lempel-Ziv-Welch)

def lzw_encode(data):
    """
    LZW compression: build dictionary of sequences.
    Output: sequence of codes.
    """
    # Initialize dictionary with single characters
    dict_size = 256
    dictionary = {chr(i): i for i in range(dict_size)}
    
    result = []
    w = ""
    
    for c in data:
        wc = w + c
        if wc in dictionary:
            w = wc
        else:
            result.append(dictionary[w])
            dictionary[wc] = dict_size
            dict_size += 1
            w = c
    
    if w:
        result.append(dictionary[w])
    
    return result

def lzw_decode(encoded):
    """Decode LZW compressed data"""
    # Initialize dictionary
    dict_size = 256
    dictionary = {i: chr(i) for i in range(dict_size)}
    
    result = []
    w = chr(encoded[0])
    result.append(w)
    
    for code in encoded[1:]:
        if code in dictionary:
            entry = dictionary[code]
        elif code == dict_size:
            entry = w + w[0]
        else:
            raise ValueError("Invalid LZW code")
        
        result.append(entry)
        
        dictionary[dict_size] = w + entry[0]
        dict_size += 1
        
        w = entry
    
    return ''.join(result)

# BURROWS-WHEELER TRANSFORM

def burrows_wheeler_transform(text):
    """
    Burrows-Wheeler Transform: reversible permutation.
    Makes similar characters cluster together.
    """
    # Add end-of-string marker
    text += '$'
    
    # Generate all rotations
    rotations = [text[i:] + text[:i] for i in range(len(text))]
    
    # Sort rotations
    rotations.sort()
    
    # Take last column
    bwt = ''.join(rotation[-1] for rotation in rotations)
    
    # Find index of original string
    original_index = rotations.index(text)
    
    return bwt, original_index

def inverse_burrows_wheeler(bwt, original_index):
    """
    Inverse Burrows-Wheeler Transform.
    Reconstruct original string.
    """
    n = len(bwt)
    
    # Build first column (sorted BWT)
    first = sorted(bwt)
    
    # Build next array: next[i] = index in BWT of character
    # that follows first[i] in original text
    next_array = [0] * n
    count = {}
    
    for i, char in enumerate(bwt):
        count[char] = count.get(char, 0)
        
        # Find position of this occurrence in first
        j = 0
        occurrences = 0
        while occurrences < count[char] or first[j] != char:
            if first[j] == char:
                occurrences += 1
            j += 1
        
        next_array[j] = i
        count[char] += 1
    
    # Reconstruct original text
    result = []
    idx = original_index
    
    for _ in range(n):
        result.append(first[idx])
        idx = next_array[idx]
    
    return ''.join(result[1:])  # Remove '$'

# COMPRESSION UTILITY

class CompressionAnalyzer:
    """Analyze compression performance"""
    
    @staticmethod
    def compression_ratio(original_size, compressed_size):
        """Calculate compression ratio"""
        return (1 - compressed_size / original_size) * 100 if original_size > 0 else 0
    
    @staticmethod
    def compare_algorithms(data):
        """Compare different compression algorithms"""
        import sys
        
        algorithms = {
            'RLE': (rle_encode, rle_decode),
            'LZ77': (lz77_encode, lz77_decode),
            'LZW': (lzw_encode, lzw_decode),
        }
        
        results = {}
        original_size = len(data)
        
        for name, (encode, decode) in algorithms.items():
            try:
                compressed = encode(data)
                
                # Estimate compressed size
                if isinstance(compressed, str):
                    compressed_size = len(compressed)
                else:
                    compressed_size = sys.getsizeof(compressed)
                
                # Verify decompression
                decompressed = decode(compressed)
                correct = (decompressed == data)
                
                results[name] = {
                    'compressed_size': compressed_size,
                    'ratio': CompressionAnalyzer.compression_ratio(original_size, compressed_size),
                    'correct': correct
                }
            except Exception as e:
                results[name] = {'error': str(e)}
        
        return results
```

### Test Cases
- Highly repetitive data (good for RLE)
- Natural text (good for LZW)
- Binary data
- Already compressed data (negative compression)
- Various file types

### Deliverables
- `rle.{py,go,zig}` - Run-length encoding
- `lz77.{py,go,zig}` - LZ77 implementation
- `lzw.{py,go,zig}` - LZW implementation
- `burrows_wheeler.{py,go,zig}` - BWT
- `compression_utility.{py,go,zig}` - File compression tool
- `compression_analysis.md` - Algorithm comparison
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All algorithms compress and decompress correctly
- Compression ratios measured accurately
- Understanding of when each algorithm works best
- Practical compression tool works on files

---

---

## Back to Module

[← Back to Module 11: String Algorithms](../module_11/README.md)

## Navigation

- [Previous Lab](./lab_11_2.md) (if exists)
- [Next Lab](./lab_11_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 11, Lab 3 of 5*
