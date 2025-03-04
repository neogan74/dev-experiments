package main

func main() {
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}



func findFrequentTreeSum(root *TreeNode) []int {
    sumFreq := make(map[int]int)
    maxFreq := 0

    var subtreeSum func(node *TreeNode) int
    subtreeSum = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        leftSum := subtreeSum(node.Left)
        rightSum := subtreeSum(node.Right)
        currSum := node.Val + leftSum + rightSum
        sumFreq[currSum]++
        if sumFreq[currSum] > maxFreq {
            maxFreq = sumFreq[currSum]
        }
        return currSum
    }

    subtreeSum(root)
    var result []int
    for sum, freq := range sumFreq {
        if freq == maxFreq {
            result = append(result, sum)
        }
    }
    return result
}
}
