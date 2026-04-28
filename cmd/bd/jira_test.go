package main

import (
	"testing"

	"github.com/steveyegge/beads/internal/tracker"
)

func TestJiraSyncStats(t *testing.T) {
	stats := tracker.SyncStats{}

	if stats.Pulled != 0 {
		t.Errorf("expected Pulled to be 0, got %d", stats.Pulled)
	}
	if stats.Pushed != 0 {
		t.Errorf("expected Pushed to be 0, got %d", stats.Pushed)
	}
	if stats.Created != 0 {
		t.Errorf("expected Created to be 0, got %d", stats.Created)
	}
	if stats.Updated != 0 {
		t.Errorf("expected Updated to be 0, got %d", stats.Updated)
	}
	if stats.Skipped != 0 {
		t.Errorf("expected Skipped to be 0, got %d", stats.Skipped)
	}
	if stats.Errors != 0 {
		t.Errorf("expected Errors to be 0, got %d", stats.Errors)
	}
	if stats.Conflicts != 0 {
		t.Errorf("expected Conflicts to be 0, got %d", stats.Conflicts)
	}
}

func TestJiraSyncResult(t *testing.T) {
	result := tracker.SyncResult{
		Success: true,
		Stats: tracker.SyncStats{
			Created: 5,
			Updated: 3,
		},
		LastSync: "2025-01-15T10:30:00Z",
	}

	if !result.Success {
		t.Error("expected Success to be true")
	}
	if result.Stats.Created != 5 {
		t.Errorf("expected Created to be 5, got %d", result.Stats.Created)
	}
	if result.Stats.Updated != 3 {
		t.Errorf("expected Updated to be 3, got %d", result.Stats.Updated)
	}
	if result.LastSync != "2025-01-15T10:30:00Z" {
		t.Errorf("unexpected LastSync value: %s", result.LastSync)
	}
	if result.Error != "" {
		t.Errorf("expected Error to be empty, got %s", result.Error)
	}
	if len(result.Warnings) != 0 {
		t.Errorf("expected Warnings to be empty, got %v", result.Warnings)
	}
}

func TestJiraConfigValueWithEnv(t *testing.T) {
	t.Setenv("JIRA_PROJECTS", "ENV")

	if got := configValueWithEnv("CONFIG", "JIRA_PROJECTS"); got != "CONFIG" {
		t.Errorf("configValueWithEnv should prefer config value, got %q", got)
	}
	if got := configValueWithEnv("", "JIRA_PROJECTS"); got != "ENV" {
		t.Errorf("configValueWithEnv should fall back to env value, got %q", got)
	}
	if got := configValueWithEnv("", ""); got != "" {
		t.Errorf("configValueWithEnv with no env var = %q, want empty", got)
	}
}

func TestJiraProjectEnvResolution(t *testing.T) {
	t.Setenv("JIRA_PROJECTS", "PS,ENG")

	projects := tracker.ResolveProjectIDs(
		nil,
		configValueWithEnv("", "JIRA_PROJECTS"),
		configValueWithEnv("", "JIRA_PROJECT"),
	)

	if len(projects) != 2 || projects[0] != "PS" || projects[1] != "ENG" {
		t.Fatalf("resolved projects = %v, want [PS ENG]", projects)
	}
}
