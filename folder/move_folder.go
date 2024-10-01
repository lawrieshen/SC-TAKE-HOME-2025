package folder

import (
	"strings"
	"errors"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders
	var source *Folder
	var destination *Folder

	// Input check
	if name == dst {
		return nil, errors.New("Cannot move a folder to itself")
	}

	// Locate source and destination
	for i := range folders {
		if folders[i].Name == name {
			source = &folders[i]
		} 
		if folders[i].Name == dst {
			destination = &folders[i]
		}
	}

	// Check if source or destination is nil
	if source == nil {
		return nil, errors.New("Source folder does not exist")
	}

	if destination == nil {
		return nil, errors.New("Destination folder does not exist")
	}

	// Check if source and destination are under the same organization or not
	if source.OrgId != destination.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	// Check if source and destination are on the same path or not
	if strings.Contains(destination.Paths, source.Paths) {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	// Update the path of the source folder and its children
	oldPath := source.Paths
	newPath := destination.Paths + "." + source.Name

	for i := range folders {
		if strings.HasPrefix(folders[i].Paths, oldPath) {
			relativePaths := strings.TrimPrefix(folders[i].Paths, oldPath)
			folders[i].Paths = newPath + relativePaths
		}
	}

	return folders, nil
}
