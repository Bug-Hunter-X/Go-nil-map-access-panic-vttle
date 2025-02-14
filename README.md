# Go nil map access panic

This repository demonstrates a common runtime panic in Go programs: accessing a nil map.

## The Bug

The `main` function attempts to assign a value to a map that is not yet initialized.  This directly results in a runtime panic.

## The Solution

The solution initializes the map before any access is made.  This is a common way to avoid this type of runtime error.
