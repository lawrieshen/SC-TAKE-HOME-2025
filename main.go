package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
    res := folders.GenerateData()

    folders.PrettyPrint(res)

    folders.WriteSampleData(res)
}
