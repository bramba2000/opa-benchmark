# OPA Benchmark Tool

This tool allows you to benchmark the Open Policy Agent (OPA) engine under various policy setups. It includes:

- **Policy Generation:** Create different cases (e.g., conditions, array, early_exit) using the `generate` command.
- **Testing:** Run benchmark tests with the `test` command.
- **Results Comparison:** Compare test results using the provided scripts and tools.

## Directory Structure

- `policies/` – Contains generated policy cases organized by case name and number of policies.
- `results/` – Stores test results for each benchmark case.
- `cmd/` – Source code for the CLI commands (generate, test, etc.).
- `scripts/` – Utility scripts to run multiple tests (e.g., for early_exit cases).

## Commands

### Generate Policies

Generate a test case:

```
opa-benchmark generate <case> --num <number_of_policies>
```

For example, to generate 10 policies for the `early_exit` case:

```
opa-benchmark generate early_exit --num 10
```

### Run Benchmark Test

Run the benchmark test on a generated policy case:

```
opa-benchmark test <case> --num <number_of_policies> --count <test_iterations>
```

For example, to run a test on the `early_exit` case with 10 policies repeated 6 times:

```
opa-benchmark test early_exit --num 10 --count 6
```

### Compare Results

Use provided scripts (like in the `scripts/` folder) to generate aggregate comparisons:

```
./scripts/early_exit_100_1000.sh
```

## Build & Run

- **Build:**  
  Use the Makefile to build the binary:
  ```
  make build
  ```
- **Run:**  
  Once built, execute the commands as shown in the examples above.

Happy benchmarking!
