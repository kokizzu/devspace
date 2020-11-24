package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetBranch(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	gitRepo, err := NewGitCLIRepository(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	err = gitRepo.Clone(CloneOptions{
		URL:    testRepo,
		Branch: testBranch,
	})
	if err != nil {
		t.Fatal(err)
	}

	resultBranch, err := GetBranch(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	if resultBranch != testBranch {
		t.Fatalf("Wrong branch, got %s, expected %s", resultBranch, testBranch)
	}
}
