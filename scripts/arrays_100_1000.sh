#!/bin/bash

# Run the generate command for the array
for i in {100..1000..100}; do
    go run main.go generate array -n $i
done

# Run the test command for all cases in parallel
for i in {100..1000..100}; do
    go run main.go test array -n $i -c 6 &
done
wait

# Run the compare command for all cases
go run main.go compare ./results/array/100.txt ./results/array/200.txt ./results/array/300.txt ./results/array/400.txt ./results/array/500.txt ./results/array/600.txt ./results/array/700.txt ./results/array/800.txt ./results/array/900.txt ./results/array/1000.txt
