package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Solver struct {
	N       int
	nums    []int
	target  int
	answers []string
}

func (S *Solver) solve() {
	perms := permutate(S.nums)
	// opss := getOperators(S.N)
	opss := getOperators2()
	for _, perm := range perms {
		for _, ops := range opss {
			S.check(perm, ops)
		}
	}
}

// Nが4限定の書き方
// TODO: 一般化
func (S *Solver) check(perm []int, ops []byte) {
	target := float32(S.target)
	a := float32(perm[0])
	b := float32(perm[1])
	c := float32(perm[2])
	d := float32(perm[3])

	op0 := ops[0]
	op1 := ops[1]
	op2 := ops[2]

	// ((a op0 b) op1 c) op2 d
	res0 := op(op(op(a, b, op0), c, op1), d, op2)
	if res0 == target {
		ans := fmt.Sprintf("((%d %c %d) %c %d) %c %d", int(a), op0, int(b), op1, int(c), op2, int(d))
		S.answers = append(S.answers, ans)
	}

	// (a op0 (b op1 c)) op2 d
	res1 := op(op(a, op(b, c, op1), op0), d, op2)
	if res1 == target {
		ans := fmt.Sprintf("(%d %c (%d %c %d)) %c %d", int(a), op0, int(b), op1, int(c), op2, int(d))
		S.answers = append(S.answers, ans)
	}

	// a op0 ((b op1 c) op2 d)
	res2 := op(a, op(op(b, c, op1), d, op2), op0)
	if res2 == target {
		ans := fmt.Sprintf("%d %c ((%d %c %d) %c %d)", int(a), op0, int(b), op1, int(c), op2, int(d))
		S.answers = append(S.answers, ans)
	}

	// a op0 (b op1 (c op2 d))
	res3 := op(a, op(b, op(c, d, op2), op1), op0)
	if res3 == target {
		ans := fmt.Sprintf("%d %c (%d %c (%d %c %d))", int(a), op0, int(b), op1, int(c), op2, int(d))
		S.answers = append(S.answers, ans)
	}

	// (a op0 b) op1 (c op2 d)
	res4 := op(op(a, b, op0), op(c, d, op2), op1)
	if res4 == target {
		ans := fmt.Sprintf("(%d %c %d) %c (%d %c %d)", int(a), op0, int(b), op1, int(c), op2, int(d))
		S.answers = append(S.answers, ans)
	}
}

func op(a, b float32, o byte) float32 {
	switch o {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return 0
	}
}

// getOperators()がうまくいっていないので暫定的な書き方
// TODO: 一般化
func getOperators2() [][]byte {
	var opss [][]byte
	var arith_ops []byte = []byte{
		'+', '-', '*', '/',
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				opss = append(opss, []byte{
					arith_ops[i],
					arith_ops[j],
					arith_ops[k],
				})
			}
		}
	}

	return opss
}

func getOperators(operands_num int) [][]byte {
	var result = [][]byte{
		[]byte{},
	}
	return getOperatorsRec(0, operands_num, result)
}

func getOperatorsRec(cur_i, operands_num int, result [][]byte) [][]byte {
	if cur_i >= operands_num {
		return result
	}

	var tmp [][]byte
	for _, ops := range result {
		tmp = append(tmp, getOperatorsAppend(ops)...)
	}
	result = tmp
	return getOperatorsRec(cur_i+1, operands_num, result)
}

func getOperatorsAppend(operator []byte) [][]byte {
	o0 := append(operator, '+')
	o1 := append(operator, '-')
	o2 := append(operator, '*')
	o3 := append(operator, '/')
	return [][]byte{
		o0, o1, o2, o3,
	}
}

func (S Solver) printAnswers() {
	fmt.Println("\nResult")
	for i, ans := range S.answers {
		fmt.Printf("%dth: %s\n", i, ans)
	}
}

func (S *Solver) storeNum() {
	var sc = bufio.NewScanner(os.Stdin)
	for i := 0; i < S.N; i++ {
		if sc.Scan() {
			t := sc.Text()
			n, _ := strconv.Atoi(t)
			S.nums = append(S.nums, n)
		}
	}
}

func permutate(nums []int) [][]int {
	var perms [][]int = [][]int{
		[]int{nums[0]},
	}

	for i := 1; i < len(nums); i++ {
		var tmp [][]int
		for _, perm := range perms {
			tmp = append(tmp, permutateInsert(perm, nums[i])...)
		}
		perms = tmp
	}
	return perms
}

func permutateInsert(nums []int, n int) [][]int {
	var result [][]int
	extends := getExtendSlices(nums)
	for i, extend := range extends {
		extend[i] = n
		result = append(result, extend)
	}
	return result
}

func getExtendSlices(s []int) [][]int {
	var result_slices [][]int
	for pos := 0; pos < len(s)+1; pos++ {
		result_slice := make([]int, len(s)+1)
		// posは 0 が入る位置
		for f_half := 0; f_half < pos; f_half++ {
			result_slice[f_half] = s[f_half]
		}
		for s_half := pos + 1; s_half < len(s)+1; s_half++ {
			result_slice[s_half] = s[s_half-1]
		}
		result_slices = append(result_slices, result_slice)
	}
	return result_slices
}

func main() {
	S := Solver{
		N:      4,
		target: 10,
	}
	S.storeNum()
	S.solve()
	S.printAnswers()
}
