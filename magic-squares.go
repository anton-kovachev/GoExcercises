package main

/*
    Write a function that determines if a two-dimensional slice of distinct
    positive integers representing a magic square. In this context, a
    magic square is a square slice where all numbers aligned horizontally,
    vertically and diagonally sum to the same value.
  
  	Take this square (2-dimensional) slice as an example:
        -------------
      | 2 | 7 | 6 |  15
      -------------
      | 9 | 5 | 1 |  15
      -------------
      | 4 | 3 | 8 |  15
      ------------- 
    15  15  15  15  15 
  
  
    As you can see each horizontal line, vertical line, and diagonal line all sum
    to the same value (the extra two 15's in the bottom are for the
    two diagonals).
  
  
    You may assume that the given square slice will always contain distinct,
    positive integers (excluding zeros) and that the minimum size of the square
    will be three (i.e., a width of three and height of three). You may also
    assume the dimensions of the slice will always be the same (i.e., it will be
    square).
  
  
    Note: Make sure to consider the case where the size of the square is greater
    than three. In this case you do NOT need to check more than two diagonals.
  
	Sample Input #1
	square = [][]int{ {5, 6, 19, 68}, {69, 18, 3, 8}, {4, 7, 70, 17}, {20, 67, 6, 5} }

	Sample Output #1
	true // all lines sum to 98
*/

func DetectMagicSquare(square [][]int) bool {
    size := len(square)
    sums := map[int]uint{}

    for i := 0; i < size; i++ {
        horizontal_sum := 0
        vertical_sum := 0
        
        for j := 0; j < size; j++ {
            horizontal_sum += square[i][j]
            vertical_sum += square[i][j]
        }

        sums[horizontal_sum] += 1
        sums[vertical_sum] += 1
    }

    diagonal_sum := 0
    reverse_diagonal_sum := 0
    for i := 0; i < size; i++ {
        diagonal_sum += square[i][i]
        reverse_diagonal_sum += square[i][size - i - 1]
    }

    sums[diagonal_sum] += 1
    sums[reverse_diagonal_sum] += 1
    
	return len(sums) == 1
}
