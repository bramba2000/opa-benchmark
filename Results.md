## Evaluation Environment

All evaluation was executed on an **Ubuntu 24.04.02 Linux x86** machine, equipped with an **i7-11700KF processor**, **32GB RAM**, and **1TB SSD**.  
All tests were performed using the [OPA-provided CLI](https://github.com/open-policy-agent/opa) as the test runner and Golang scripts as the generator of the test cases.  
The test repository ([opa-benchmark](https://github.com/bramba2000/opa-benchmark)) provides guides on how to generate the environment and the `opa-benchmark` tool required to run the tests. **Go version 1.24+ is required**.

### Test Case Generation

The main command to generate tests is:

```
opa-benchmark generate <test_case_name>
```
where `<test_case_name>` can be one of the following:
- `conditions`
- `array`
- `early-exit`

You can configure the cardinality of the test case using the `-n <num>` flag (default: 100).

The generated test case is stored in `./policies/<test_case>/<num>`, containing one or more Rego files and a `test.rego` file. `test.rego` contains the input data for the test case, while the other files contain the policies loaded into OPA during evaluation.

### Execution & Data Collection

- Each test measures the time required for the OPA engine to execute the query in the corresponding `test.rego`.
- The `opa-benchmark` tool runs each test case continuously for 30 seconds to collect a large enough sample, then averages the metrics.
- The entire process is repeated **6 times** for robust data collection.
- Data for each run is collected in `results/<test_case>/`, as `<num>.txt` files in [gobench](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) format, which enables easy comparison with other test cases.

---

## Test Case: `conditions`

The `conditions` test case, generated with $n$ policies, consists of $n$ Rego files, each with a simple conditional rule.  

We evaluated two queries:
- One returning `allow=true`
- One returning `allow=false`

### Query Response Time

|                | 100   | 200   | 300   | 400   | 500   | 600   | 700   | 800   | 900   | 1000  |
|----------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| **Test Allow** |       |       |       |       |       |       |       |       |       |       |
| Average (μs)   | 40.68 | 41.05 | 40.33 | 41.33 | 41.10 | 41.27 | 41.34 | 40.94 | 40.90 | 41.73 |
| Rel. Error (%) | 4     | 4     | 2     | 4     | 4     | 2     | 4     | 3     | 3     | 4     |
| **Test Deny**  |       |       |       |       |       |       |       |       |       |       |
| Average (μs)   | 34.89 | 34.89 | 34.87 | 35.78 | 34.91 | 35.47 | 35.77 | 35.41 | 35.32 | 35.84 |
| Rel. Error (%) | 4     | 3     | 3     | 5     | 4     | 3     | 4     | 3     | 4     | 4     |

### Space Allocated per Query

|                | 100    | 200    | 300    | 400    | 500    | 600    | 700    | 800    | 900    | 1000   |
|----------------|--------|--------|--------|--------|--------|--------|--------|--------|--------|--------|
| **Test Allow** |        |        |        |        |        |        |        |        |        |        |
| Avg (KiB)      | 7.133  | 7.127  | 7.123  | 7.120  | 7.118  | 7.116  | 7.115  | 7.114  | 7.113  | 7.113  |
| Incr. Rate (%) | -      | -0.08  | -0.14  | -0.18  | -0.19  | -0.24  | -0.25  | -0.26  | -0.28  | -0.29  |
| Rel. Error (%) | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      |
| **Test Deny**  |        |        |        |        |        |        |        |        |        |        |
| Avg (KiB)      | 5.534  | 5.530  | 5.526  | 5.524  | 5.525  | 5.521  | 5.520  | 5.520  | 5.519  | 5.518  |
| Incr. Rate (%) | -      | -0.08  | -0.14  | -0.18  | -0.19  | -0.24  | -0.25  | -0.26  | -0.28  | -0.29  |
| Rel. Error (%) | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      | 1      |

**Conclusion:**  
With ideal conditions, evaluating a high number of policies occurs in approximately constant time.

---

## Test Case: `early-exit`

The `early_exit` test case, generated with $n$ policies, uses a single Rego file containing $n$ policies. Each policy is an allow rule with a random set of equality conditions.  
The rules cannot be optimized by the OPA engine due to overlapping conditions and incomplete coverage. 
### Query Response Time

|                | 100   | 200   | 300   | 400   | 500   | 600   | 700   | 800   | 900   | 1000  |
|----------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| **Test Allow** |       |       |       |       |       |       |       |       |       |       |
| Average (μs)   | 42.10 | 44.20 | 46.50 | 48.90 | 51.40 | 54.10 | 56.90 | 59.90 | 63.00 | 66.30 |
| Incr. Rate (%) | -     | 4.99  | 10.45 | 16.15 | 22.09 | 28.50 | 35.15 | 42.28 | 49.64 | 57.48 |
| Rel. Error (%) | 3     | 4     | 3     | 5     | 4     | 6     | 5     | 4     | 3     | 4     |

### Space Allocated per Query

|                | 100   | 200   | 300   | 400   | 500   | 600   | 700   | 800   | 900   | 1000  |
|----------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| **Test Allow** |       |       |       |       |       |       |       |       |       |       |
| Avg (KiB)      | 7.196 | 7.695 | 7.683 | 7.720 | 7.792 | 7.971 | 8.246 | 8.267 | 8.282 | 8.338 |
| Incr. Rate (%) | -     | 6.94  | 6.77  | 7.28  | 8.29  | 10.78 | 14.59 | 14.89 | 15.10 | 15.88 |
| Rel. Error (%) | 1     | 1     | 1     | 1     | 1     | 1     | 1     | 1     | 1     | 1     |

---

## Test Case: `array`

The `array` use case, with cardinality $n$, uses a single Rego file with $n$ object entries.  
Each object simulates a user database entry, populating the `data.arrays.roles` variable.  
The test measures the response time for a lookup of a user in `roles` where the user role matches a random value.

**Example Policy:**
```rego
test_allow if {
    some a in data.arrays[_]
    a.role == "role_50"
}
```

### Query Response Time

|                | 100   | 200    | 300    | 400    | 500    | 600    | 700    | 800    | 900    | 1000   |
|----------------|-------|--------|--------|--------|--------|--------|--------|--------|--------|--------|
| **Test Allow** |       |        |        |        |        |        |        |        |        |        |
| Avg (μs)       | 45.25 | 134.81 | 174.37 | 222.01 | 258.99 | 298.68 | 337.34 | 377.44 | 416.02 | 457.39 |
| Incr. Rate (%) | -     | 197.90 | 285.31 | 390.58 | 472.29 | 560.02 | 645.44 | 734.05 | 819.30 | 910.72 |
| Rel. Error (%) | 103   | 51     | 47     | 29     | 30     | 6      | 3      | 3      | 2      | 2      |

### Space Allocated per Query

|                | 100   | 200   | 300   | 400   | 500   | 600   | 700   | 800   | 900   | 1000  |
|----------------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| **Test Allow** |       |       |       |       |       |       |       |       |       |       |
| Avg (KiB)      | 327.0 | 527.0 | 727.0 | 927.0 | 1127.0| 1327.0| 1527.0| 1727.0| 1927.0| 2127.0|
| Incr. Rate (%) | -     | 61.16 |122.32 |183.49 |244.65 |305.81 |366.97 |428.13 |489.30 |550.46 |
| Rel. Error (%) | 0     | 0     | 0     | 0     | 0     | 0     | 0     | 0     | 0     | 0     |

---

## References

- [opa-benchmark repository](https://github.com/open-policy-agent/opa-benchmark)
- [gobench format and tools](https://pkg.go.dev/golang.org/x/perf)
