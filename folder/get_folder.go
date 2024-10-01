package folder

import (
	"fmt"

	"errors"
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.folders
	var parentFolder *Folder


	// Find parent folder
	for _, folder := range folders {
		if folder.OrgId == orgID && folder.Name == name {
			parentFolder = &folder
			break
		}
	}

	// Handle parent folder not found error
	if parentFolder == nil {

		for _, folder := range folders {
			if folder.Name == name {
				fmt.Printf("Found folder with name: %s but OrgID does not match\n", folder.Name) // Debug: Folder name exists, but OrgID mismatch
				return nil, errors.New("Folder does not exist in the specified organization")
			}
		}
		return nil, errors.New("Folder does not exist")
	}


	// Find the child folders by validating the path prefix
	res := []Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID && len(folder.Paths) > len(parentFolder.Paths) &&
			folder.Paths[:len(parentFolder.Paths)] == parentFolder.Paths {
			res = append(res, folder)
		}
	}

	return res, nil
}

