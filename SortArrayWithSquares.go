package main

/*
	https://www.interviewbit.com/problems/sort-array-with-squares/
 */
import "fmt"

func main() {
	//small positive input
	//a := []int{-5, -4, -1, 0, 2}

	//large positive input
	a := []int { -855, -847, -747, -708, -701, -667, -666, -618, -609, -536, -533, -509, -500, -396, -336, -297, -284, -229, -173, -173, -132, -38, -5, 35, 141, 169, 281, 322, 358, 421, 436, 447, 478, 538, 547, 644, 667, 673, 705, 711, 718, 724, 726, 811, 869, 894, 895, 902, 912, 942, 961 }
	fmt.Println(sortArrayWithSquares(a))
}

func sortArrayWithSquares(a []int )  []int {
	n:=len(a)
	p:=n
	for i:=0; i<n;i++ {
		if a[i] >= 0 {
			p = i
			break
		}
	}
	i:=p-1
	j:=p
	result := make([]int, n)
	k:=0
	for i>=0 && j!=n {
		iSquare := a[i]*a[i]
		jSquare := a[j]*a[j]
		if iSquare < jSquare {
			result[k] = iSquare
			k+=1
			i-=1
		} else {
			result[k] = jSquare
			k+=1
			j+=1
		}
	}

	if i<0 {
		for j<n {
			result[k] = a[j]*a[j]
			k+=1
			j+=1
		}
	}
	if j==n {
		for i>=0 {
			result[k] = a[i]*a[i]
			k+=1
			i-=1
		}
	}
	return result
}
