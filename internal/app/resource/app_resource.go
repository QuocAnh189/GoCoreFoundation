package resource

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/lingos"
)

type AppResource struct {
	Db       db.IDatabase
	LingoSvc *lingos.Service
}
