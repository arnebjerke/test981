package tmp_manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/flant/werf/pkg/werf"
)

const (
	projectsServiceDir          = "projects"
	dockerConfigsServiceDir     = "docker_configs"
	werfConfigRendersServiceDir = "werf_config_renders"

	CommonPrefix           = "werf-"
	ProjectDirPrefix       = CommonPrefix + "project-data-"
	DockerConfigDirPrefix  = CommonPrefix + "docker-config-"
	WerfConfigRenderPrefix = CommonPrefix + "config-render-"
)

func GetServiceTmpDir() string {
	return filepath.Join(werf.GetServiceDir(), "tmp")
}

func GetCreatedTmpDirs() string {
	return filepath.Join(GetServiceTmpDir(), "created")
}

func GetReleasedTmpDirs() string {
	return filepath.Join(GetServiceTmpDir(), "released")
}

func registerCreatedPath(newPath, createdPathsDir string) error {
	if err := os.MkdirAll(createdPathsDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create dir %s: %s", createdPathsDir, err)
	}

	createdPath := filepath.Join(createdPathsDir, filepath.Base(newPath))
	if err := os.Symlink(newPath, createdPath); err != nil {
		return fmt.Errorf("unable to create symlink %s -> %s: %s", createdPath, newPath, err)
	}

	return nil
}

func releasePath(path, createdPathsDir, releasedPathsDir string) error {
	if err := os.MkdirAll(releasedPathsDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create dir %s: %s", releasedPathsDir, err)
	}

	releasedPath := filepath.Join(releasedPathsDir, filepath.Base(path))
	if err := os.Symlink(path, releasedPath); err != nil {
		return fmt.Errorf("unable to create symlink %s -> %s: %s", releasedPath, path, err)
	}

	createdPath := filepath.Join(createdPathsDir, filepath.Base(path))
	if err := os.Remove(createdPath); err != nil {
		return fmt.Errorf("unable to remove %s: %s", createdPath, err)
	}

	return nil
}

func newTmpDir(prefix string) (string, error) {
	newDir, err := ioutil.TempDir(werf.GetTmpDir(), prefix)
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "darwin" {
		dir, err := filepath.EvalSymlinks(newDir)
		if err != nil {
			return "", fmt.Errorf("eval symlinks of path %s failed: %s", newDir, err)
		}
		newDir = dir
	}

	return newDir, nil
}

func newTmpFile(prefix string) (string, error) {
	newFile, err := ioutil.TempFile(werf.GetTmpDir(), prefix)
	if err != nil {
		return "", err
	}

	path := newFile.Name()

	err = newFile.Close()
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "darwin" {
		dir, err := filepath.EvalSymlinks(path)
		if err != nil {
			return "", fmt.Errorf("eval symlinks of path %s failed: %s", path, err)
		}
		path = dir
	}

	return path, nil
}
