package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
    // Write your code here
    var i int32 
    var j int32 
    for i=0 ; i < n ; i++{
        for j = 0 ; j < n ;j++{
            if (n - j - 1 <= i){
                fmt.Print("#")
            }else{
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
