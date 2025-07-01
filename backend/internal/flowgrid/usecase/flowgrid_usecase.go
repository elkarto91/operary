package usecase

import "github.com/elkarto91/operary/internal/flowgrid/model"

// MatchSkills iterates over tasks and assigns them to available workers
// with matching skills. A worker can only be assigned once.
func MatchSkills(workers []model.Worker, tasks []model.Task) []model.Assignment {
	used := make(map[string]bool)
	var assignments []model.Assignment

	for _, task := range tasks {
		for _, w := range workers {
			if !w.Available || used[w.ID] {
				continue
			}
			if hasSkills(w.Skills, task.RequiredSkills) {
				assignments = append(assignments, model.Assignment{
					TaskID:   task.ID,
					WorkerID: w.ID,
				})
				used[w.ID] = true
				break
			}
		}
	}

	return assignments
}

func hasSkills(workerSkills, required []string) bool {
	skillMap := make(map[string]bool, len(workerSkills))
	for _, s := range workerSkills {
		skillMap[s] = true
	}
	for _, s := range required {
		if !skillMap[s] {
			return false
		}
	}
	return true
}
