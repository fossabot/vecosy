package configGitRepo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vecosy/vecosy/v2/pkg/configrepo"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/storage"
	"time"
)

func (cr *GitConfigRepo) StartPullingEvery(period time.Duration) error {
	t := time.NewTicker(period)
	go func() {
		for {
			select {
			case t := <-t.C:
				logrus.Debugf("Auto pull :%+s", t)
				cr.pushError(cr.Pull())
			case <-cr.pullCh:
				t.Stop()
				return
			}
		}
	}()
	return nil
}

func (cr *GitConfigRepo) StopPulling() {
	cr.pullCh <- true
}

func (cr *GitConfigRepo) Pull() error {
	logrus.Info("Pull")
	if cr.cloneOpts != nil {
		refs := []config.RefSpec{"*:*"}
		fetchOpts := &git.FetchOptions{Auth: cr.cloneOpts.Auth, Force: true, Tags: git.AllTags, RefSpecs: refs}
		err := cr.repo.Fetch(fetchOpts)
		if err != nil && err != storage.ErrReferenceHasChanged {
			if err != git.NoErrAlreadyUpToDate {
				logrus.Errorf("Error pulling :%s", err)
				return err
			} else {
				logrus.Info("already up to date")
			}
		} else {
			newApps, err := cr.loadApps()
			if err != nil {
				return err
			}
			changes := detectChanges(cr.Apps, newApps)
			if len(changes) > 0 {
				cr.Apps = newApps
				cr.callChangeHandlers(changes)
			} else {
				logrus.Debugf("no changes detected")
			}
		}
	} else {
		return fmt.Errorf("cannot pull:no remote information found")
	}
	return nil
}

func detectChanges(oldApps, newApps map[string]*app) []configrepo.ApplicationVersion {
	changes := make([]configrepo.ApplicationVersion, 0)
	for newAppName, newApp := range newApps {
		if oldApp, alreadyExist := oldApps[newAppName]; alreadyExist {
			// checking references
			for verName, verRef := range newApp.Branches {
				if oldVer, oldVerExist := oldApp.Branches[verName]; !oldVerExist || oldVer.Hash().String() != verRef.Hash().String() {
					logrus.Debugf("new version or different branch hash detected oldVersionExist version:%s", verName)
					changes = append(changes, configrepo.ApplicationVersion{AppName: newAppName, AppVersion: verName})
				}
			}
		} else {
			logrus.Infof("new application detected:%s", newAppName)
			for _, newAppVersion := range newApp.Versions {
				changes = append(changes, configrepo.ApplicationVersion{AppName: newAppName, AppVersion: newAppVersion.String()})
			}
		}
	}
	return changes
}

func (cr *GitConfigRepo) callChangeHandlers(changes []configrepo.ApplicationVersion) {
	logrus.Infof("callChangeHandlers: %+v", changes)
	for _, chHandler := range cr.changesHandlers {
		for _, change := range changes {
			chHandler(change)
		}
	}
}
