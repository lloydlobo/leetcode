package algo

type Solution interface{ CountQuadruplets() int }
type Example struct{ nums []int }

func ExecuteAlgo1(nums []int) int {
	var s Solution
	s = Example{nums: nums}
	return s.CountQuadruplets()
}

// Easy solution by using simple formula
// a+b+c=d to a+b=d-c which reduces
// O(n^3) to O(N^2) as we equalized both the sides
func (e Example) CountQuadruplets() (count int) {
	MAP := make(map[int]int)
	n := len(e.nums)
	// a value from 0 to b-1; b values b
	// c values b+1; d values b+2
	for b := 1; b < n; b++ {
		// a + b
		for a := 0; a < b; a++ {
			MAP[e.nums[a]+e.nums[b]]++
		}
		//d - c
		c := (b + 1)
		for d := b + 2; d < n; d++ {
			count += MAP[e.nums[d]-e.nums[c]]
		}
	}
	return count
}

/*
	Easy solution by using simple formula a+b+c=d to a+b=d-c which reduces O(n^3) to O(N^2) as we equalized both the sides
class Solution {
public:
	int countQuadruplets(vector<int>& nums) {
	    unordered_map<int,int> map;
	    int n(nums.size());
	    int count(0);
	    for(int b=1;b<n;b++){
//             a+b
	for(int a=0;a<b;a++){
	    map[nums[a]+nums[b]]++;
	}

//             d-c
//             a value from 0 to b-1
//             b values b
//             c values b+1
//             d values b+2

	            int c(b+1);
	            for(int d=b+2;d<n;d++){
	                count+=map[nums[d]-nums[c]];
	            }
	        }

	        return count;
	    }
	};
*/
