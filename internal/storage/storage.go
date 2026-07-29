package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/ctrl-shift-markus/yaat/internal/activity"
)

func Path() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "yaat", "activities.json"), nil
}

func LogActivity(path string, description string, duration int) error {
	activities, err := Load(path)
	if err != nil {
		return err
	}

	ID := 1
	if len(activities) > 0 {

		maxID := 0
		for _, activity := range activities {
			if activity.ID > maxID {
				maxID = activity.ID
			}
		}
		ID = maxID + 1
	}

	newItem := activity.Item{
		Description: description,
		Duration:    duration,
		ID:          ID,
	}

	activities = append(activities, newItem)

	err = save(path, activities)
	if err != nil {
		return err
	}

	return nil
}

func Load(path string) ([]activity.Item, error) {
	cleanPath := filepath.Clean(path)

	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return []activity.Item{}, nil
	}

	jsonActivities, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}

	var activities []activity.Item
	err = json.Unmarshal(jsonActivities, &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func save(path string, activities []activity.Item) error {
	jsonActivities, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		return err
	}

	cleanPath := filepath.Clean(path)

	dir := filepath.Dir(cleanPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return err
	}

	err = os.WriteFile(cleanPath, jsonActivities, 0600)
	if err != nil {
		return err
	}

	return nil
}
