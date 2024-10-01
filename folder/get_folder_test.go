package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	org1ID := uuid.Must(uuid.NewV4())
	org2ID := uuid.Must(uuid.NewV4())
	org3ID := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "returns folders for org1",
			orgID: org1ID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1ID},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: org2ID},
				{Name: "charlie", Paths: "alpha.charlie", OrgId: org3ID},
			},
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1ID},
			},
		},
		{
			name:  "returns no folders for org3",
			orgID: org3ID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1ID},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: org2ID},
			},
			want: []folder.Folder{}, // No folders should match
		},
		{
			name:  "returns empty for no folders",
			orgID: org1ID,
			folders: []folder.Folder{},
			want:    []folder.Folder{}, // Expecting an empty result
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			assert.ElementsMatch(t, tt.want, got, "Folders returned do not match expected output")
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	org1ID := uuid.Must(uuid.NewV4())
	org2ID := uuid.Must(uuid.NewV4())

	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1ID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1ID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1ID},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1ID},
		{Name: "echo", Paths: "echo", OrgId: org1ID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2ID},
	}

	tests := []struct {
		name         string
		folderName   string
		orgID        uuid.UUID
		wantLen      int
		expectError  bool
		errorMessage string
	}{
		{	// Normal Case
			name:       "Returns child folders for alpha",
			orgID:      org1ID,
			wantLen:    3,
			expectError: false,
		},
		{	 // Normal case
			name:       "Returns no child folders for bravo",
			folderName: "bravo",
			orgID:      org1ID,
			wantLen:    1,
			expectError: false,
		},
		{	// Edge Case
			name:       "returns error for not existing folder name",
			folderName: "invalid",
			orgID:      org1ID,
			wantLen:    0,
			expectError: true,
			errorMessage: "Folder does not exist",
		},
		{	// Edge Case
			name:       "returns error for not existing folder name in specific organization",
			folderName: "foxtrot",
			orgID:      org1ID,
			wantLen:    0,
			expectError: true,
			errorMessage: "Folder does not exist in the specified organization",
		},
		{	// Edge Case
			name:       "returns error for not existing folder name in specific organization",
			folderName: "bravo",
			orgID:      org2ID,
			wantLen:    0,
			expectError: true,
			errorMessage: "Folder does not exist in the specified organization",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(folders)
			res, err := folderDriver.GetAllChildFolders(tt.orgID, tt.folderName)

			if tt.expectError {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, res, tt.wantLen, "Expected number of child folders does not match")
		})
	}
}
