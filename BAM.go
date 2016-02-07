// package for reading modifying .bam files 
package main

import (
       "bufio"
       "fmt"
       "os"
)

func main() {
     input := bufio.NewScanner(os.Stdin)
