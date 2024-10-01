package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// func Test_folder_MoveFolder(t *testing.T) {
// 	// TODO: your tests here

// }

func TestMoveFolder(t *testing.T) {
	t.Parallel()
	org1Id := uuid.Must(uuid.NewV4())
	org2Id := uuid.Must(uuid.NewV4())
	
	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1Id},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1Id},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1Id},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1Id},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1Id},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2Id},
		{Name: "golf", Paths: "golf", OrgId: org1Id},
	}

	tests := []struct {
		name          string
		source        string
		destination   string
		wantLen       int
		expectError   bool
		errorMessage  string
	}{
		{
			name:        "Move bravo to delta",
			source:      "bravo",
			destination: "delta",
			wantLen:     7,
			expectError: false,
		},
		{
			name:        "Move charlie to delta",
			source:      "charlie",
			destination: "delta",
			wantLen:     7,
			expectError: false,
		},
		{
			name:        "Move non-existing source folder",
			source:      "folder",
			destination: "delta",
			wantLen:     0,
			expectError: true,
			errorMessage: "Source folder does not exist",
		},
		{
			name:        "Move non-existing destination folder",
			source:      "bravo",
			destination: "folder",
			wantLen:     0,
			expectError: true,
			errorMessage: "Destination folder does not exist",
		},
		{
			name:        "Move folders in different organization",
			source:      "bravo",
			destination: "foxtrot",
			wantLen:     0,
			expectError: true,
			errorMessage: "Cannot move a folder to a different organization",
		},
		{
			name:        "Move identical folders",
			source:      "bravo",
			destination: "bravo",
			wantLen:     0,
			expectError: true,
			errorMessage: "Cannot move a folder to itself",
		},
		{
			name:        "Move folders to its child",
			source:      "bravo",
			destination: "charlie",
			wantLen:     0,
			expectError: true,
			errorMessage: "Cannot move a folder to a child of itself",
		},
		{
			name:        "Move folders at same level",
			source:      "bravo",
			destination: "delta",
			wantLen:     7,
			expectError: false,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(folders)
			// Attempt to move the folder
			res, err := folderDriver.MoveFolder(tt.source, tt.destination)

			// Check for errors
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMessage)
			} else {
				assert.NoError(t, err)
			}

			// Verify the folder structure
			assert.Len(t, res, tt.wantLen, "Unexpected number of folders after move")
		})
	}
}
