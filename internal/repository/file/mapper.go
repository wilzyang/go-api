package file

import "github.com/wilzyang/go-api/internal/core/domain/file"

func mapFileListToDomain(fl FileList) file.FileList {
	return file.FileList{
		Filename:     fl.Filename,
		Size:         fl.Size,
		MediaLink:    fl.MediaLink,
		DownloadLink: fl.DownloadLink,
		FileType:     fl.FileType,
		CreatedAt:    fl.CreatedAt,
		UpdatedAt:    fl.UpdatedAt,
	}
}
