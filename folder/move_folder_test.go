package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	org1Id := uuid.Must(uuid.NewV4())
	org2Id := uuid.Must(uuid.NewV4())

	// Initial folder structure
	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1Id},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1Id},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1Id},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1Id},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1Id},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2Id},
		{Name: "golf", Paths: "golf", OrgId: org1Id},
	}

	// Each test case runs independently, but changes to the folder structure within one test may affect subsequent tests 
	// since they change the memory addess in place.
	tests := []struct {
		name          string
		source        string
		destination   string
		wantLen       int
		expectError   bool
		errorMessage  string
		expectedPaths map[string]string
	}{
		{
			name:        "Move bravo to delta",
			source:      "bravo",
			destination: "delta",
			wantLen:     7,
			expectError: false,
			expectedPaths: map[string]string{
				"bravo":   "alpha.delta.bravo",
				"charlie": "alpha.delta.bravo.charlie",
			},
		},
		{
			name:        "Move charlie to delta",
			source:      "charlie",
			destination: "delta",
			wantLen:     7,
			expectError: false,
			expectedPaths :map[string]string{
				"charlie": "alpha.delta.charlie",
			},
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
			name:        "Move folders at same level",
			source:      "bravo",
			destination: "charlie",
			wantLen:     7,
			expectError: false,
			expectedPaths: map[string]string {
				"bravo":"alpha.delta.charlie.bravo",
			},
		},
		{
			name:        "Move folder to its child",
			source:      "delta",
			destination: "charlie",
			wantLen:     0,
			expectError: true,
			errorMessage: "Cannot move a folder to a child of itself",
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

			// Verify that the folder paths have been updated correctly
			
			for _, folder := range res {
				expectedPath, exists := tt.expectedPaths[folder.Name]
				if exists {
					assert.Equal(t, expectedPath, folder.Paths, "Folder path mismatch")
				}
			}
		})
	}
}
