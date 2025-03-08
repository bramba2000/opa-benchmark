#!/bin/bash

# Run the generate command for the conditions
for i in {100..1000..100}; do
    opa-benchmak generate conditions -n $i # changed command
done

# Run the test command for all cases in parallel
for i in {100..1000..100}; do
    opa-benchmak test conditions -n $i -c 6 & # changed command
done
wait

# Run the compare command for all cases
opa-benchmak compare ./results/conditions/100.txt ./results/conditions/200.txt \
    ./results/conditions/300.txt ./results/conditions/400.txt ./results/conditions/500.txt \
    ./results/conditions/600.txt ./results/conditions/700.txt ./results/conditions/800.txt \
    ./results/conditions/900.txt ./results/conditions/1000.txt \
    >./results/conditions/100_1000.txt # changed command
echo "Results of compare saved to results/conditions/100_1000.txt"
