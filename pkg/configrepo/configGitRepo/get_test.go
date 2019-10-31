package configGitRepo

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"testing"
)

func TestConfigRepo_GetNearestBranch_FullMatch(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	gitRepo := cfgRepo.(*GitConfigRepo)
	branch, err := gitRepo.GetNearestBranch("app1", "v1.0.0")
	assert.NoError(t, err)
	assert.NotNil(t, branch)
	assert.Contains(t, branch.Name().String(), "app1/v1.0.0")
}

func TestConfigRepo_GetNearestBranch_Between(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	gitRepo := cfgRepo.(*GitConfigRepo)
	branch, err := gitRepo.GetNearestBranch("app1", "v5.0.0")
	assert.NoError(t, err)
	assert.NotNil(t, branch)
	assert.Contains(t, branch.Name().String(), "refs/tags/app1/v1.0.1")
}

func TestConfigRepo_GetNearestBranch_Over(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	gitRepo := cfgRepo.(*GitConfigRepo)
	branch, err := gitRepo.GetNearestBranch("app1", "v10.0.0")
	assert.NoError(t, err)
	assert.NotNil(t, branch)
	assert.Contains(t, branch.Name().String(), "app1/v6.0.0")
}

func TestConfigRepo_GetFile(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	tests := []struct {
		name            string
		version         string
		expectedVersion string
	}{
		{"version 1.0.0", "1.0.0", "1.0.0"},
		{"version 1.0.1", "1.0.1", "1.0.1"},
		{"version 5.0.0", "5.0.0", "1.0.1"},
		{"version 10.0.0", "10.0.0", "6.0.0"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			configContent := getConfigYml(t, cfgRepo, "app1", test.version)
			assert.Equal(t, "dev", configContent["environment"])
			assert.Equal(t, test.expectedVersion, configContent["ver"])
		})
	}
}

func TestConfigRepo_GetNearestBranch_AppNotFound(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	gitRepo := cfgRepo.(*GitConfigRepo)
	_, err = gitRepo.GetNearestBranch("not-exist-app", "v1.0.0")
	assert.Error(t, err)
}

func TestConfigRepo_GetFile_AppNotFound(t *testing.T) {
	localRepo, remoteRepo := InitRepos(t)
	cfgRepo, err := NewConfigRepo(localRepo, &git.CloneOptions{URL: remoteRepo})
	assert.NoError(t, err)
	assert.NotNil(t, cfgRepo)
	assert.NoError(t, cfgRepo.Init())
	_, err = cfgRepo.GetFile("not-exist-app", "v1.0.0", "config.yml")
	assert.Error(t, err)
}