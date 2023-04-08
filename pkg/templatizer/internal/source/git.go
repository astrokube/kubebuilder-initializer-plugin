/*
Package source provides us with the implementations that can be used to fetch the source code in the templates.
*/
package source

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

var repoRegex = regexp.MustCompile(`(?P<protocol>(https|ssh))://((?P<auth>(.*))(@))?(?P<path>([a-zA-Z0-9./\-_]+))(#(?P<branch>(.*)))?`)

// A GitSource struct is used to fetch the template from a Git repository.
type GitSource struct {
	url       string
	branch    string
	auth      http.AuthMethod
	fs        billy.Filesystem
	refOrigin string
}

// NewGitSource returns a pointer of a new instance of GitSource.
//
// The function receives a parameter, named conn,  that contains the details of the Git connection. The function
// uses the below regular expression to extract the  details in the parameter.
//
// (?P<protocol>(https|ssh))://((?P<auth>(.*))(@))?(?P<path>([a-zA-Z0-9./\-_]+))(#(?P<branch>(.*)))?
//
// The connection url is composed by the following parts
// protocol: It's required a must be https or ssh depending on the chosen mechanism to establish the connection with the remote repository.
// auth: It's an optional param (not required for public repositories) and we can provide both user credentials or a token. If we pass the credentials the format must be username:password
// path: The repository url. For instance, github.com/astrokube/kubebuilder-initializer-plugin if we use the https protocol, or github.com/astrokube/kubebuilder-initializer-plugin.git if we use the ssh protocol
// branch: Optional param, we can use a specific branch by passing the name of the desired branch after symbol `#`
func NewGitSource(conn string) *GitSource {
	match := repoRegex.FindStringSubmatch(conn)

	paramsMap := make(map[string]string)
	for i, name := range repoRegex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
			fmt.Printf("%s -> %s", name, match[i])
		}
	}
	s := &GitSource{}
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
	s.url = paramsMap["path"]
	return s
}

// GetTemplateContent returns a map with an entry per each file in the repository. The key of this map
// is the absolute path to the file from the root of the repository and the value is the content of the file.
func (s *GitSource) GetTemplateContent() (map[string]string, error) {
	if err := s.loadFileSystem(); err != nil {
		return nil, err
	}
	out := make(map[string]string, 0)
	err := s.iterateOverTheFiles("/", out)
	return out, err
}

func (s *GitSource) loadFileSystem() error {
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL:          s.url,
		Auth:         s.auth,
		SingleBranch: true,
	})
	if err != nil {
		return fmt.Errorf("error cloning the repository '%w'", err)
	}
	w, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("error getting the working tree of the repository '%w'", err)
	}
	if s.branch != "" {
		if err := w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewRemoteReferenceName(s.refOrigin, s.branch),
		}); err != nil {
			return fmt.Errorf("error checking out branch '%s' from the repository '%w'", s.branch, err)
		}
	}
	s.fs = w.Filesystem
	return nil
}

func (s *GitSource) iterateOverTheFiles(path string, files map[string]string) error {
	resources, err := s.fs.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading folder '%s' from repository: '%w", path, err)
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
			return fmt.Errorf("error opening file '%s' from repository: '%w", filename, err)
		}
		if resource.Size() == 0 {
			files[filename] = ""
			continue
		}
		bytes := make([]byte, resource.Size())
		if _, err = src.Read(bytes); err != nil {
			return fmt.Errorf("error reading  '%v' file '%s' size '%v' from repository: '%w", resource, filename,
				resource.Size(), err)
		}
		files[filename] = string(bytes)
	}
	return nil
}
