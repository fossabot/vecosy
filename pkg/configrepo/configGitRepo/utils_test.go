package configGitRepo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vecosy/vecosy/v2/pkg/configrepo"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
	"time"
)

func getConfigYml(t *testing.T, cfgRepo configrepo.Repo, appName, targetVersion string) map[string]interface{} {
	cfgFl, err := cfgRepo.GetFile(appName, targetVersion, "config.yml")
	assert.NoError(t, err)
	configContent := make(map[string]interface{})
	assert.NoError(t, yaml.Unmarshal(cfgFl.Content, configContent))
	return configContent
}

func editAndPush(t *testing.T, remoteRepo, app, srcVersion, dstVersion, fileName, commitMsg string, content []byte) {
	editorRepoPath := fmt.Sprintf("%s_wk", remoteRepo)
	logrus.Debugf("editorPath:%s", editorRepoPath)
	editorRepo, err := git.PlainClone(editorRepoPath, false, &git.CloneOptions{URL: remoteRepo, NoCheckout: true})
	assert.NoError(t, err)

	wk, err := editorRepo.Worktree()
	assert.NoError(t, err)
	remoteAppBranch := plumbing.ReferenceName(fmt.Sprintf("refs/remotes/origin/%s/%s", app, srcVersion))
	err = wk.Checkout(&git.CheckoutOptions{Branch: remoteAppBranch, Force: true})
	assert.NoError(t, err)

	logrus.Debugf("current branch:%s", remoteAppBranch.String())
	fl, err := wk.Filesystem.OpenFile(fileName, os.O_RDWR, 0644)
	assert.NoError(t, err)
	defer fl.Close()
	_, err = fl.Write(content)
	assert.NoError(t, err)
	assert.NoError(t, fl.Close())

	_, err = wk.Add(fileName)
	assert.NoError(t, err)

	commitHash, err := wk.Commit(commitMsg, &git.CommitOptions{All: true, Author: editorSignature,})
	assert.NoError(t, err)
	logrus.Debugf("commitHash:%s", commitHash.String())

	localBranch := plumbing.NewBranchReferenceName(fmt.Sprintf("%s/%s", app, dstVersion))
	branchRef := plumbing.NewHashReference(localBranch, commitHash)
	assert.NoError(t, editorRepo.Storer.CheckAndSetReference(branchRef, nil))
	time.Sleep(1 * time.Second)
	assert.NoError(t, editorRepo.Push(&git.PushOptions{
		RemoteName: "origin",
		RefSpecs:   []config.RefSpec{config.DefaultPushRefSpec},
	}))
}
