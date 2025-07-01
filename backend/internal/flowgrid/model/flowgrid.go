package model

// Worker represents an operator available for task assignments.
type Worker struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Skills    []string `json:"skills"`
	Available bool     `json:"available"`
}

// Task defines work requiring a particular skill set.
type Task struct {
	ID             string   `json:"id"`
	Description    string   `json:"description"`
	RequiredSkills []string `json:"required_skills"`
}

// ScheduleRequest bundles workers and tasks for scheduling.
type ScheduleRequest struct {
	Workers []Worker `json:"workers"`
	Tasks   []Task   `json:"tasks"`
}

// Assignment ties a task to the worker selected to perform it.
type Assignment struct {
	TaskID   string `json:"task_id"`
	WorkerID string `json:"worker_id"`
}
