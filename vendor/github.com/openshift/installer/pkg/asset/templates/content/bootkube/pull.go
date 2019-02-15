package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	pullFileName = "pull.yaml.template"
)

var _ asset.WritableAsset = (*Pull)(nil)

// Pull is the constant to represent contents of pull.yaml.template file
type Pull struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *Pull) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *Pull) Name() string {
	return "Pull"
}

// Generate generates the actual files by this asset
func (t *Pull) Generate(parents asset.Parents) error {
	fileName := pullFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *Pull) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *Pull) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, pullFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
