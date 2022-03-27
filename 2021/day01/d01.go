//https://adventofcode.com/2021/day/1

package day01

//NumberOfIncreased returns how many interval is actually increased
//obsolete. replaced by NumberOfIncreasedNInterval(depths, 1)
func NumberOfIncreased(depths []int) int {
	prev := depths[0]
	ret := 0
	for i := 1; i < len(depths); i++ {
		if prev < depths[i] {
			ret++
		}
		prev = depths[i]
	}
	return ret
}

//ThreeSumArray makes [a,b,c,d,e] to [a+b+c,b+c+d,c+d+e,...]
// in order to caliculate second part of the problem
//obsolete. replaced by NumberOfIncreasedNInterval(depths, 1)
func ThreeSumArray(depths []int) []int {
	sum := depths[0] + depths[1]
	var ret []int
	for i := 2; i < len(depths); i++ {
		sum += depths[i]
		ret = append(ret, sum)
		sum -= depths[i-2]
	}
	return ret
}

//NumberOfIncreasedNInterval returns how many
//is actually increased between depths[i] and depth[i+interval]
func NumberOfIncreasedNInterval(depths []int, interval int) int {
	if len(depths) < interval+1 {
		return 0
	}
	ret := 0
	for i := 0; i < len(depths)-interval; i++ {
		if depths[i] < depths[i+interval] {
			ret++
		}
	}
	return ret
}
