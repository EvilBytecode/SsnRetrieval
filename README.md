# SsnRetrieval
Loads NTDLL, parses the PE file, extracts "Zw" functions, retrieves their System Service Numbers (SSNs), and prints each function’s name, SSN, and address.

## Execution Process:
- 1st > Load the NTDLL Libary.
- 2nd > Parse the PE file to get the structure and find important directories like the export directory.
- 3rd > Extract function names and addresses, look for functions that start with "Zw", and find their System Service Numbers (SSNs).
- 4th > Collect and print the SSN, function name, and address for each "Zw" function.

# Enjoy 
