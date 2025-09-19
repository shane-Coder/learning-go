# Go Sorting: A Deep Dive

## Objective
This example provides a comprehensive guide to the two primary methods for sorting slices of custom structs in Go: the flexible `sort.Slice` function and the powerful `sort.Interface` pattern. Understanding both is crucial for writing clean, efficient, and idiomatic Go code.

---

## Method 1: `sort.Slice` (The Quick & Flexible Way)

This is a convenient and common way to handle one-off or simple sorting needs.

### How It Works
You provide the slice you want to sort and a custom **"less" function**. This function tells Go how to compare any two elements in your slice. It takes two indices, `i` and `j`, and must return `true` if the element at `i` should come before the element at `j`.

### Use Case
Perfect for when you need to sort a slice in a specific way for a particular function, but you don't want to define a global sorting order for your data type.

---

## Method 2: `sort.Interface` (The Reusable & Professional Way)

This is the more traditional and powerful pattern for making a custom data type "natively" sortable.

### How It Works
You teach your custom slice type how to sort itself by implementing the three methods required by Go's `sort.Interface`:

1.  **`Len() int`**: Returns the number of elements in the slice.
2.  **`Swap(i, j int)`**: Swaps the elements at indices `i` and `j`.
3.  **`Less(i, j int) bool`**: This is where your core sorting logic goes. It returns `true` if the element at `i` should come before `j`.

Once your type has these three methods, you can sort it anywhere in your code by just calling `sort.Sort()`.

### Use Case
Ideal for when your data type has a "natural" or primary sort order that will be used frequently throughout your application.

---

## Running the Example
To see both methods in action, navigate to this directory and run:
```bash
go run main.go

--- Sorting with sort.Slice ---
Profiles sorted by name: [{Name:Alice Age:30} {Name:Bob Age:22} {Name:Shivam Age:25}]

--- Sorting with sort.Interface (by Age) ---
Profiles sorted by age: [{Name:Bob Age:22} {Name:Shivam Age:25} {Name:Alice Age:30}]