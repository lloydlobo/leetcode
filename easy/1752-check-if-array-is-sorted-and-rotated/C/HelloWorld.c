// #include "HelloWorld.h"
#include <stdio.h>

// main
//
// $ gcc HelloWorld.c   // Compile the C program.
// $ ./a.out            // Run the output binary.
int main() {
  // printf() displays the string inside quotation
  printf("Hello, World!");
  return 0;
}

/*
# GCC: $ gcc <file_name.c>

  Preprocess and compile C and C++ source files, then assemble and link them
together. More information: <https://gcc.gnu.org>.

  Compile multiple source files into executable:
      $ gcc path/to/source1.c path/to/source2.c ... -o path/to/output_executable

  Allow warnings, debug symbols in output:
      $ gcc path/to/source.c -Wall -Og -o path/to/output_executable

  Include libraries from a different path:
      $ gcc path/to/source.c -o path/to/output_executable -Ipath/to/header
-Lpath/to/library -llibrary_name

  Compile source code into Assembler instructions:
      $ gcc -S path/to/source.c

  Compile source code into an object file without linking:
      $ gcc -c path/to/source.c
*/
