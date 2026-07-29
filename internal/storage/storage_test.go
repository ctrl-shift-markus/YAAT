package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ctrl-shift-markus/yaat/internal/activity"
)

func TestPath(t *testing.T) {
	path, err := Path()
	if err != nil {
		t.Fatalf("Expected Path() to work, got error '%v'", err)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("Expected os.UserConfigDir() to work, got error '%v'", err)
	}

	expectedPath := filepath.Join(configDir, "yaat", "activities.json")
	if path != expectedPath {
		t.Errorf("Expected '%s', got '%s'", expectedPath, path)
	}
}

func TestSave(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "test.json")

	testJSON := []activity.Item{
		{ID: 0, Description: "TestSave()", Duration: 0},
	}

	err := save(testPath, testJSON)
	if err != nil {
		t.Fatalf("Expected save(testPath, testJSON) to work, got error '%v'", err)
	}

	JSON, err := os.ReadFile(testPath)
	if err != nil {
		t.Fatalf("Expected os.ReadFile(testPath) to work, got error '%v'", err)
	}
expectedJSON := `[
  {
    "id": 0,
    "description": "TestSave()",
    "duration": 0
  }
]`
	if string(JSON) != expectedJSON {
		t.Errorf("Expected %q, got %q", expectedJSON, string(JSON))
	}
}

func TestLoad(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "test.json")

	JSON := []byte(`[{"id":0,"description":"TestLoad()","duration":0}]`)
	err := os.WriteFile(testPath, JSON, 0640)
	if err != nil {
		t.Fatalf("Expected os.WriteFile(testPath, JSON, 0640) to work, got error '%v'", err)
	}

	activities, err := Load(testPath)
	if err != nil {
		t.Fatalf("Expected Load() to work, got error '%v'", err)
	}

	// Fatal because if it can't get len(activities) right, the ID numbering system will break
	if len(activities) != 1 {
		t.Fatalf("Expected '1' activity, got '%d' activity(ies)", len(activities))
	}

	if activities[0].ID != 0 {
		t.Errorf("Expected ID '0', got '%d'", activities[0].ID)
	}

	if activities[0].Description != "TestLoad()" {
		t.Errorf("Expected description 'TestLoad()', got %q", activities[0].Description)
	}

	if activities[0].Duration != 0 {
		t.Errorf("Expected duration '0', got '%d'", activities[0].Duration)
	}
}

func TestLogActivity(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "test.json")

	err := LogActivity(testPath, "TestLogActivity() 1", 0)
	if err != nil {
		t.Fatalf("Expected LogActivity() to work (first call), got error '%v'", err)
	}

	// Call LogActivity() twice to test if the ID numbering system works
	err = LogActivity(testPath, "TestLogActivity() 2", 0)
	if err != nil {
		t.Fatalf("Expected LogActivity() to work (second call), got error '%v'", err)
	}

	JSON, err := os.ReadFile(testPath)
	if err != nil {
		t.Fatalf("Expected to read file, got error '%v'", err)
	}

	// Indented JSON required or else it won't match	
	expectedJSON := `[
  {
    "id": 1,
    "description": "TestLogActivity() 1",
    "duration": 0
  },
  {
    "id": 2,
    "description": "TestLogActivity() 2",
    "duration": 0
  }
]`

	if string(JSON) != expectedJSON {
		t.Errorf("Expected %q, got %q", expectedJSON, string(JSON))
	}
}
