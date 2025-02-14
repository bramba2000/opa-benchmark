# Generate tag gets a num and a file name and run the corresponding go cmd script	
gen:
	go run cmd/generate/$(case)/main.go $(num)

run:
	opa test policies/$(case)/$(num)_policies -v --bench --format gobench --count 6 | tee ./results/$(case)/$(num)_result.txt

cmp: 
	~/go/bin/benchstat $(list)

.PHONY: gen run cmp