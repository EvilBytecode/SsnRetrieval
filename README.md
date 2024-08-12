# SsnRetrieval
Loads NTDLL, parses the PE file, extracts "Zw" functions, retrieves their System Service Numbers (SSNs), and prints each function’s name, SSN, and address.

## Execution Process:
- 1st > Load the NTDLL Libary.
- 2nd > Parse the PE file to get the structure and find important directories like the export directory.
- 3rd > Extract function names and addresses, look for functions that start with "Zw", and find their System Service Numbers (SSNs).
- 4th > Collect and print the SSN, function name, and address for each "Zw" function.

# Build Process
- 1st -> ```go build main.go```
- if you want to run and test ```go run main.go```

## Enjoy - Made by EByte :Happy

# PoC
![image](https://github.com/user-attachments/assets/295d3e89-573b-43f5-8125-3199306d9adb)
