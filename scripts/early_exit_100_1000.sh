#!/bin/bash

# Run the generate command for the early_exit
for i in {100..1000..100}; do
    opa-benchmak generate early_exit -n $i
done

# Run the test command for all cases in parallel
for i in {100..1000..100}; do
    opa-benchmak test early_exit -n $i -c 6 &
done
wait

# Run the compare command for all cases
opa-benchmak compare ./results/early_exit/100.txt ./results/early_exit/200.txt \
    ./results/early_exit/300.txt ./results/early_exit/400.txt ./results/early_exit/500.txt \
    ./results/early_exit/600.txt ./results/early_exit/700.txt ./results/early_exit/800.txt \
    ./results/early_exit/900.txt ./results/early_exit/1000.txt \
    >./results/early_exit/100_1000.txt
echo "Results of compare saved to results/early_exit/100_1000.txt"
