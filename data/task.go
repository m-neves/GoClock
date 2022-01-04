package data

import (
	"fmt"
	"sort"
	"time"
)

type Task struct {
	Id     uint   `json:"id"`
	Name   string `json:"task_name"`
	UserId uint   `json:"-"`
}

type TaskCompletion struct {
	Id          uint      `json:"id"`
	TaskId      uint      `json:"task_id"`
	CompletedAt time.Time `json:"completed_at"`
}

type TaskCompletions []*TaskCompletion

func (tc *TaskCompletions) GroupByTaskId() {
	sort.Slice(*tc, func(i, j int) bool {
		return (*tc)[i].CompletedAt.Before((*tc)[j].CompletedAt)
	})

}

func (tc *TaskCompletions) Streaks() {

	byTask := map[int][]*TaskCompletion{}

	tc.GroupByTaskId()

	for _, v := range *tc {
		byTask[int(v.TaskId)] = append(byTask[int(v.TaskId)], v)
	}

	var maxStreak int = 0
	var curStreak int = 0
	var prev *TaskCompletion = nil

	for _, v := range byTask {
		prev = nil
		maxStreak = 0
		curStreak = 0

		for _, tc := range v {
			if prev != nil {

				if tc.CompletedAt.Sub(prev.CompletedAt).Hours() <= 24 {
					curStreak++
				} else {
					curStreak = 1
				}

				if curStreak >= maxStreak {
					maxStreak = curStreak
				}
			} else {
				curStreak++
			}
			prev = tc
		}
	}

	fmt.Print("Streak: ", maxStreak)
}
