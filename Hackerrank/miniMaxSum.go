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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
// BEGIN MyCODE
func miniMaxSum(arr []int32) {
    // Write your code here
    var minSum int64 = int64(uint64(1 << 63) -1)
    var maxSum int64 = 0
    
    for i:=0 ; i<5;i++{
        var temp int64 =0
        for j := 0 ; j<4; j++{
            temp += int64(arr[(i+j)%5])
        }
        if temp < minSum{
            minSum = temp
        }
        if temp > maxSum{
            maxSum = temp
        }
    }
    fmt.Println(minSum,maxSum)
}
// END
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
