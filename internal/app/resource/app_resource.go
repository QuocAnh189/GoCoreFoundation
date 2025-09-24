package resource

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/lingos"
)

type AppResource struct {
	Env      *configs.Env
	Db       db.IDatabase
	LingoSvc *lingos.Service
}
