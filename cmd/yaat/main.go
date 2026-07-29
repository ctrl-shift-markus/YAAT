package main

import (
	"fmt"

	"github.com/ctrl-shift-markus/yaat/internal/activity"

	"github.com/alecthomas/kong"
	"github.com/ctrl-shift-markus/yaat/internal/storage"
)

type LogCmd struct {
	Description string `arg:"" help:"The description of the activity to log."`
	Duration    int    `arg:"" help:"The duration of the activity to log (in minutes with no suffix, eg. 41)."`
}

func (l *LogCmd) Run() error {
	path, err := storage.Path()
	if err != nil {
		return err
	}

	err = storage.LogActivity(path, l.Description, l.Duration)
	if err != nil {
		return fmt.Errorf("could not log activity: %w", err)
	}

	fmt.Printf("Logged activity: %s - %s\n", l.Description, activity.FormatDuration(l.Duration))
	return nil
}

type ListCmd struct{}

func (c *ListCmd) Run() error {
	path, err := storage.Path()
	if err != nil {
		return err
	}

	activities, err := storage.Load(path)
	if err != nil {
		return fmt.Errorf("could not list activities: %w", err)
	}

	if len(activities) == 0 {
		fmt.Println("No activities found.")
		return nil
	}

	for _, act := range activities {
		fmt.Printf("[%d] %s - %s\n", act.ID, act.Description, activity.FormatDuration(act.Duration))
	}

	return nil
}

var cli struct {
	Log  LogCmd  `cmd:"" help:"Log a new activity."`
	List ListCmd `cmd:"" help:"List all logged activities."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
