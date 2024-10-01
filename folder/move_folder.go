package folder

import (
	"strings"
	"errors"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	folders := f.folders
	var source *Folder
	var destination *Folder

	// Input check
	if name == dst {
		return nil, errors.New("Cannot move a folder to itself")
	}

	// Locate source and destination
	for _, folder := range folders {

		if folder.Name == name {
			source = &folder
		} 
		if folder.Name == dst {
			destination = &folder
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
    for _, folder := range folders {
        if folder.Name == name && strings.Contains(folder.Paths, source.Paths) {
            relativePaths := strings.TrimPrefix(folder.Paths, source.Paths)
            folder.Paths = destination.Paths + "." + source.Name + relativePaths
        }
    }

	return folders, nil
}
