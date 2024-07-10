## Disclaimer
I didn't have experience coding in Go but since Tomas mentioned GO is what you mostly used, I decided to go with it.

I had a lot of help from Copilot when solving this exercise. I do understand everything that goes on in the code but it would have been a lot harder it would have taken me a lot longer without the help of AI. I left the code very commented as a guide to myself.

It also helped me in figuring out the time and space complexity of the program.


## To run:
```bash
docker run --rm -it -v $(pwd)/input.txt:/app/input.txt -v $(pwd)/output:/app/output top-numbers-app -n 5 --input-file input.txt --output-file /app/output/output.txt
```

### ERROR: Number of top results must be bigger than 0.
```bash
docker run --rm -it -v $(pwd)/input.txt:/app/input.txt -v $(pwd)/output:/app/output top-numbers-app -n 0 --input-file input.txt --output-file /app/output/output.txt
```

#### ERROR: The maximum number of top results must be less or equal than 30000000.
```bash
docker run --rm -it -v $(pwd)/input.txt:/app/input.txt -v $(pwd)/output:/app/output top-numbers-app -n 30000001 --input-file input.txt --output-file /app/output/output.txt
```

#### ERROR: input file does not exist.
```bash
docker run --rm -it  -v $(pwd)/output:/app/output top-numbers-app -n 5 --input-file input.txt --output-file /app/output/output.txt
```

### ERROR: input file is not readable.
```bash
chmod 000 input.txt
go build -o top_numbers .
./top_numbers -n 5 --input-file input.txt --output-file output.txt
```


## What would be the time and space complexity of your program? Why?

### Time Complexity
1. Reading the input file:
   - The program reads the file line by line using a scanner. This operation is O(n), where n is the number of lines in the file.

2. Parsing and storing numbers:
   - For each line, the program parses the string to a uint64. This is generally considered O(1) for each number.
   - Appending each number to the slice is amortized O(1).
   - Overall, this step is O(n).

3. Sorting the numbers:
   - The program uses Go's sort.Slice function, which implements a quicksort algorithm.
   - The average time complexity of quicksort is O(n log n), where n is the number of elements.

4. Selecting top N numbers:
   - This is a simple slice operation, which is O(1).

5. Writing output:
   - Writing N numbers to the output file is O(N), where N is the number of top results requested.

The dominant factor in this algorithm is the sorting step. So the overall time complexity of this program is:

- Average case: O(n log n)

Where n is the number of valid numbers in the input file.

It's worth noting that while the program allows for up to 30,000,000 top results, the actual number of results (N) doesn't significantly affect the time complexity because the entire dataset is sorted regardless of N. The N value only affects the final output step, which is linear.


### Space Complexity
1. Input storage:
   - The program stores all valid numbers from the input file in a slice called `numbers`.
   - This requires O(n) space, where n is the number of valid numbers in the input file.

2. Sorting:
   - Go's sort.Slice function typically uses quicksort, which requires O(log n) additional space for the call stack due to its recursive nature.
   - However, the sorting is done in-place on the existing slice, so it doesn't require additional space proportional to the input size.

3. Output selection:
   - The program potentially truncates the sorted slice to the top N elements, but this doesn't increase the space complexity.

4. Other variables:
   - The program uses a few other variables (like `n`, `inputFile`, `outputFile`, etc.) but these require constant space.

5. File I/O buffers:
   - The program uses buffered I/O for reading and writing, which uses a fixed amount of memory regardless of input size.

The dominant factor in space usage is the slice that stores all the numbers from the input file. Therefore, the overall space complexity of this program is O(n), where n is the number of valid numbers in the input file.


## Do you think there is still room for improvement in your solution? If so, please elaborate.

Efficiency in Sorting:
Currently, the program sorts all valid numbers to find the top N largest. For very large input files with tens of millions of lines, this sorting operation could be optimized further. One approach could involve using a min-heap (priority queue) of size N to efficiently maintain the top N largest numbers as we iterate through the input.

Parallelism:
Given the potential size of input files, parallelizing the input reading and possibly the sorting process (using Go's goroutines) could provide performance benefits, especially on multi-core systems.

Optimized Input Parsing:
Enhancing the parsing logic to skip over invalid lines more efficiently could reduce unnecessary processing time, especially in cases where there are many invalid lines.

Memory Efficiency:
Although Go's standard libraries are used, exploring more memory-efficient data structures or algorithms for sorting and maintaining the top N elements could be beneficial, especially if memory constraints are a concern.
In conclusion, while the current implementation meets the basic requirements, there are indeed areas where further optimization and improvement could enhance its performance, especially when dealing with very large input files and strict performance requirements.