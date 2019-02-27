package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type IntervalsStart []Interval

func (in IntervalsStart) Len() int {
	return len(in)
}

func (in IntervalsStart) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.Start < y.Start
}

func (in IntervalsStart) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

//对 [l1,r1], [l2,r2]，如果 r1 > l2，则 r1 = max(r1, r2)
func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Sort(IntervalsStart(intervals))
	fmt.Println("sort by start:", intervals)
	var r []Interval
	start, end := intervals[0].Start, intervals[0].End

	for _, v := range intervals {
		if v.Start <= end {
			if v.End > end {
				end = v.End
			}
		} else {
			r = append(r, Interval{Start: start, End: end})
			start = v.Start
			end = v.End
		}
	}
	r = append(r, Interval{Start: start, End: end})
	return r
}

func main() {
	intervals := []Interval{
		{Start: 1, End: 3},
		{Start: 2, End: 6},
		{Start: 15, End: 18},
		{Start: 8, End: 10},
	}
	fmt.Println(merge(intervals))
}
