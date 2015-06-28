package schedule

import (
	"testing"
)

func TestRandomSchedule(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{1, 1}, {2, 8}, {10, 100},
	}
	for _, test := range tests {
		git := &MockGit{}
		git.Init()
		filegen := MockFileGenerator{}
		RandomSchedule(test.min, test.max, git, filegen)
		if git.numInitCalls != 1 {
			t.Errorf("Expected one init call, but got %d", git.numInitCalls)
		}
		if git.numAddCalls != git.numCommitCalls {
			t.Error("Add calls should happen as often as commit calls.")
		}
		if test.min*366 > git.numCommitCalls || test.max*366 < git.numCommitCalls {
			fmt := "Total commits should be between %d and %d, but was %d"
			t.Errorf(fmt, test.min*366, test.max*366, git.numCommitCalls)
		}
	}
}

type MockFileGenerator struct{}

func (m MockFileGenerator) CreateFile() string {
	return ""
}

type MockGit struct {
	numInitCalls   int
	numAddCalls    int
	numCommitCalls int
}

func (m *MockGit) Init() {
	m.numInitCalls++
}

func (m *MockGit) Add(filename string) {
	m.numAddCalls++
}

func (m *MockGit) Commit(message, date string) {
	m.numCommitCalls++
}
