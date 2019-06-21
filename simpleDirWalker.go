package simpleDirWalker

import (
	"io"
	"os"
	"path/filepath"
)

// PathAndFileInfo datat structure for save full path file & os.FileInfo
type PathAndFileInfo struct {
	FullPath string
	FileInfo os.FileInfo
}

// SDW return PathAndFileInfo on recursive read directories function
func SDW(dir string) ([]PathAndFileInfo, error) {
	// temporally variable
	var tmp []PathAndFileInfo

	// open directory
	dh, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer dh.Close()

	// read files
	for {
		fis, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}

		for _, fi := range fis {
			tmp = append(tmp, PathAndFileInfo{filepath.Join(dir, fi.Name()), fi})

			// recursive read directories
			if fi.IsDir() {
				SDW(dir + "/" + fi.Name())
			}
		}
	}

	return tmp, nil
}
