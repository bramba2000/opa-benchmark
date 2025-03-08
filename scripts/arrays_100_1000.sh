#!/bin/bash

# Run the generate command for the array
for i in {100..1000..100}; do
    opa-benchmak generate array -n $i # changed command
done

# Run the test command for all cases in parallel
for i in {100..1000..100}; do
    opa-benchmak test array -n $i -c 6 & # changed command
done
wait

# Run the compare command for all cases
opa-benchmak compare ./results/array/100.txt ./results/array/200.txt ./results/array/300.txt ./results/array/400.txt ./results/array/500.txt ./results/array/600.txt ./results/array/700.txt ./results/array/800.txt ./results/array/900.txt ./results/array/1000.txt
