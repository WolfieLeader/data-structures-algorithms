# Data Structures and Algorithms (DSA) 🏛️

## Arrays

### Array Types:

- **Static Arrays**: Fixed size, allocated at compile time.
- **Dynamic Arrays**: Resizable, allocated at runtime (e.g., slices in Go).
- **Multidimensional Arrays**: Arrays of arrays, useful for matrices.

![Array Visualization](./assets/01_array/arrays/array.png)

![2D Array Visualization](./assets/01_array/arrays/2d_array.png)

## Searching Algorithms

1. **Linear Search** – Checks each element sequentially.

   ![Linear Search](./assets/01_array/searching/linear.gif)

2. **Binary Search** – Efficient search on sorted arrays, divides the array in half each time. Works only on sorted arrays.

   ![Binary Search](./assets/01_array/searching/binary.gif)

### Sorting Algorithms

1. **Bubble Sort** – Repeatedly swaps adjacent elements if they are in the wrong order.

   ![Bubble Sort](./assets/01_array/sorting/bubble.gif)

2. **Selection Sort** – Selects the minimum element and places it at the beginning.

   ![Selection Sort](./assets/01_array/sorting/selection.gif)

3. **Insertion Sort** – Builds a sorted array one element at a time (good for small datasets).

   ![Insertion Sort](./assets/01_array/sorting/insertion.gif)

4. **Gnome Sort** – Swaps like bubble sort but steps back when out of order.

   ![Gnome Sort](./assets/01_array/sorting/gnome.gif)

5. **Cocktail Sort** – A bi-directional version of bubble sort.

   ![Cocktail Sort](./assets/01_array/sorting/cocktail.gif)

6. **Comb Sort** – Bubble sort with gap shrinkage to eliminate "turtles"(small values at the end).

   ![Comb Sort](./assets/01_array/sorting/comb.gif)

7. **Shell Sort** – Generalized insertion sort using decreasing gaps.

   ![Shell Sort](./assets/01_array/sorting/shell.gif)

8. **Merge Sort** – Recursively divides and merges sorted halves (stable, O(n log n)).

   ![Merge Sort](./assets/01_array/sorting/merge.gif)

9. **Quick Sort** – Divide-and-conquer using a pivot for partitioning.

   - Lomuto partition scheme (classic)
   - Hoare partition scheme (optimized)
   - Dutch national flag (Three-Way partitioning for duplicates)

   ![Quick Sort](./assets/01_array/sorting/quick.gif)

10. **Heap Sort** – Converts array into a heap, then extracts elements in order.

    ![Heap Sort](./assets/01_array/sorting/heap.gif)

11. **Counting Sort** – Counts occurrences (great for small integer ranges).

    ![Counting Sort](./assets/01_array/sorting/counting.gif)

12. **Radix Sort** – Sorts by individual digits using stable sort per digit.

    ![Radix Sort](./assets/01_array/sorting/radix.gif)

13. **Bucket Sort** – Distributes elements into buckets, sorts each.

    ![Bucket Sort](./assets/01_array/sorting/bucket.png)

## Linked Lists

### Linked List Types:

- **Singly Linked List**: Each node points to the next node.
- **Doubly Linked List**: Each node points to both the next and previous nodes
- **Circular Linked List**: Last node points back to the first node, forming a circle.

## Stacks

- Array-based Stack
- Linked List-based Stack

## Queues

- Array-based Queue
- Linked List-based Queue

## Hash Tables

## Trees

## Graphs

#### NOTE❗

All visuals are sourced from Wikipedia, GeeksForGeeks and FreeCodeCamp and are used for educational purposes only.
