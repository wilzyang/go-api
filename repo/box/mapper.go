package box

import "github.com/wilzyang/go-api/domain/box"

func toDomainFileId(rr ResearchReport) box.ResearchReport {
	return box.ResearchReport{
		FileId: rr.FileId,
	}
}
