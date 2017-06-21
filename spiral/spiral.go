package main

import "fmt"
import "os"
import "strconv"
import "math"

func main() {
    if (len(os.Args[1:]) != 1){
        fmt.Printf("Invalid number of arguments\n")
        os.Exit(1)
    } else {
        num, err := strconv.Atoi(os.Args[1])
        fmt.Printf("Using number: %v\n", num)
        if(err != nil){
            fmt.Println("Error")
            os.Exit(1)
        }
        grid := make([][]int, num)
        j := 0
        for j < num {
            grid[j] = make([]int, num)
            j++
        }
        row, col := 0, 0
        dRow, dCol := 0, 1
        i := 1
        for i <= num*num {
            if (row + dRow == num) || (row + dRow < 0) || (col + dCol == num) || (col + dCol < 0) || (grid[row + dRow][col + dCol] != 0) {
                dRow, dCol = dCol, -dRow
            }
            grid[row][col] = i
            row = row + dRow
            col = col + dCol
            i++
        }

        // print the results
        i = 0
        digits := math.Floor(math.Log10(float64(num*num))) + 1
        for i < num {
            j = 0
            l := len(grid[i])
            for j < l {
                thisDigits := math.Floor(math.Log10(float64(grid[i][j]))) + 1
                padding := int (digits - thisDigits)
                k := 0
                for k < padding {
                    fmt.Printf(" ")
                    k++
                }
                fmt.Printf("%v ", grid[i][j])
                j++
            }
            fmt.Printf("\n")
            i++
        }
    }
}
