package easy

/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers
directly above it as shown:

      1
	  1 1
	 1 2 1
   1 3 3 1
  1 4 6 4 1


Example 1:
   Input: numRows = 5
   Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:
   Input: numRows = 1
   Output: [[1]]

Constraints:
   1 <= numRows <= 30

*/
func Generate(numRows int) [][]int {
	tri := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		tri[i] = make([]int, i+1)
		tri[i][0] = 1
		tri[i][i] = 1

		for j := 1; j < i; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}
	return tri
}
