package lib

import (
	"encoding/json"
	"os/exec"

	"github.com/cgardner/taskhelper/types"
)

func GetNote(id string) (types.TaskExport, error) {
	var task []types.TaskExport
	taskJson, err := exec.Command("task", "export", id).Output()
	if err != nil {
		return task[0], err
	}

	err = json.Unmarshal(taskJson, &task)
	if err != nil {
		return task[0], err
	}

	return task[0], nil

}
