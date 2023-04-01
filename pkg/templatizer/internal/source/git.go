package source

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

var repoRegex = regexp.MustCompile("((?P<auth>(.*))(@))?(?P<path>([a-zA-Z0-9./\\-_]+))(#(?P<branch>(.*)))?")

type GitSource struct {
	url       string
	branch    string
	auth      http.AuthMethod
	fs        billy.Filesystem
	refOrigin string
}

func NewGitSource(conn string) (s *GitSource) {
	match := repoRegex.FindStringSubmatch(conn)

	paramsMap := make(map[string]string)
	for i, name := range repoRegex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	if authVal, ok := paramsMap["auth"]; ok && authVal != "" {
		if strings.Contains(authVal, ":") {
			authValParts := strings.Split(authVal, ":")
			s.auth = &http.BasicAuth{
				Username: authValParts[0],
				Password: authValParts[1],
			}
		} else {
			s.auth = &http.TokenAuth{
				Token: authVal,
			}
		}
	}
	s.refOrigin = "origin"
	s.branch = paramsMap["branch"]
	s.url = fmt.Sprintf("https://%s", paramsMap["path"])
	return
}

func (s *GitSource) GetTemplateContent() (out map[string]string, err error) {
	if err = s.loadFileSystem(); err != nil {
		return
	}
	err = s.iterateOverTheFiles("/", out)
	return
}

func (s *GitSource) loadFileSystem() error {
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL:             s.url,
		Auth:            s.auth,
		Progress:        os.Stdout,
		InsecureSkipTLS: true,
	})
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	if s.branch != "" {
		if err := w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewRemoteReferenceName(s.refOrigin, s.branch),
		}); err != nil {
			return err
		}
	}
	s.fs = w.Filesystem
	return nil
}

func (s *GitSource) iterateOverTheFiles(path string, files map[string]string) error {
	resources, err := s.fs.ReadDir(path)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		filename := filepath.Join(path, resource.Name())
		if resource.IsDir() {
			if err := s.iterateOverTheFiles(filename, files); err != nil {
				return err
			}
			continue
		}
		src, err := s.fs.Open(filename)
		if err != nil {
			return err
		}
		bytes := make([]byte, resource.Size())
		if _, err = src.Read(bytes); err != nil {
			return err
		}
		files[filename] = string(bytes)
	}
	return nil
}
