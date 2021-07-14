package appgenerate

import (
	"context"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/otiai10/copy"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// GenerateProject generates a code base for the project and returns the file path in which
// it was generated in.
func GenerateCodeBaseForProject(ctx context.Context, pages []*schema.PageSnapshot) (CodeBasePath, error) {
	// Create a temporary directory on which a new project is to be generated.
	projectPath, err := newCodeBasePath()
	if err != nil {
		return "", err
	}

	// Remove the node modules, we may have downloaded it locally during development.
	err = removeNodeModulesIfExists(config.TemplateNextAppPath)
	if err != nil {
		return "", err
	}

	// Copy the template next app in the temp location. We are going to build on top of it.
	err = copy.Copy(config.TemplateNextAppPath, string(projectPath))
	if err != nil {
		return "", err
	}

	pagesPath := path.Join(string(projectPath), "pages")

	for _, page := range pages {
		if err := writePageInPath(page, pagesPath); err != nil {
			return projectPath, err
		}
	}

	return projectPath, nil
}

// writePageInPath writes a page component based on the given page information.
func writePageInPath(p *schema.PageSnapshot, pagesPath string) error {
	pageFilePath := path.Join(pagesPath, resolvePagePath(p.Route))

	// Create directories leading to the page file.
	pageDirPath := path.Dir(pageFilePath)
	if err := os.MkdirAll(pageDirPath, os.ModePerm); err != nil {
		return err
	}

	page, err := buildPage(p)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pageFilePath, []byte(page), fs.ModePerm)
}

// CodeBasePath is the path within which a project is generated.
// This path is a temporary directory which needs to be manually cleaned up
// after use. On Cloud Run, the filesystem actually resides on the RAM, so
// need to be careful to clean things up otherwise it can clog up memory
// between requests (as the container may still be kept alive).
// https://cloud.google.com/run/docs/reference/container-contract#filesystem
type CodeBasePath string

// newCodeBasePath creates a new code base path for a project to be generated in.
func newCodeBasePath() (CodeBasePath, error) {
	err := os.MkdirAll("deployApps", os.ModePerm)
	if err != nil {
		return "", err
	}

	path, err := ioutil.TempDir("deployApps", "app")
	if err != nil {
		return "", err
	}

	return CodeBasePath(path), nil
}

// Cleanup removes all the files within the code base path.
func (c CodeBasePath) Cleanup() error {
	return os.RemoveAll(string(c))
}

// resolvePagePath gets the file path for the route in the pattern of the NextJS
// page directory structure.
func resolvePagePath(route string) string {
	// There can be the following kinds of path definitions:
	// 1. /
	// 2. /name
	// 3. /name/help
	//
	// And the one at 1 can be represented as `/index.js`.
	// While 2 can be represented as `/name.js` or `/name/index.js`.
	// And 3 can be represented as `/name/help.js` or `/name/help/index.js`.
	//
	// If 2 and 3 were to happen at once, it can be much simpler to organize
	// and define the routes by appending index.js.
	return path.Join(route, "index.js")
}

// removeNodeModulesIfExists removes node modules if it exists.
// This is only useful during development. During production, node_modules
// does not exist.
func removeNodeModulesIfExists(templatePath string) error {
	nodeModules := path.Join(templatePath, "node_modules")
	return os.RemoveAll(nodeModules)
}
