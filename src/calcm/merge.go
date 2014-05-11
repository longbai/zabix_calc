package main

import (
	"sort"
)

type MergeHour struct {
	Utc   int64
	Time  string
	Total int64
	Avg   []int64
}

type hourArray []*MergeHour

func (p hourArray) Len() int           { return len(p) }
func (p hourArray) Less(i, j int) bool { return p[i].Total < p[j].Total }
func (p hourArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p hourArray) Sort()              { sort.Sort(p) }

func mergeMax(args [][]RecordHour) (ret []*MergeHour) {
	store := map[int64]*MergeHour{}
	for _, arg := range args {
		for _, v := range arg {
			r, ok := store[v.Utc]
			if ok {
				r.Total += v.ValueMax
				r.Avg = append(r.Avg, v.ValueMax)
			} else {
				var m MergeHour
				m.Utc = v.Utc
				m.Time = v.Time
				m.Total += v.ValueMax
				m.Avg = append(m.Avg, v.ValueMax)
				store[v.Utc] = &m
			}
		}
	}

	ret = make([]*MergeHour, 0, len(store))

	for _, v := range store {
		ret = append(ret, v)
	}
	hourArray(ret).Sort()
	return
}

var total int64

func calcThoughput(records []*MergeHour) (amount int64) {
	for _, v := range records {
		amount += v.Total * 3600 / 8
	}
	return
}

func merge(args [][]RecordHour) (ret []*MergeHour) {
	store := map[int64]*MergeHour{}
	for _, arg := range args {
		for _, v := range arg {
			r, ok := store[v.Utc]
			if ok {
				r.Total += v.ValueAvg
				r.Avg = append(r.Avg, v.ValueAvg)
			} else {
				var m MergeHour
				m.Utc = v.Utc
				m.Time = v.Time
				m.Total += v.ValueAvg
				m.Avg = append(m.Avg, v.ValueAvg)
				store[v.Utc] = &m
			}
		}
	}

	ret = make([]*MergeHour, 0, len(store))

	for _, v := range store {
		ret = append(ret, v)
	}
	hourArray(ret).Sort()
	return
}
