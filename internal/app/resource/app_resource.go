package resource

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

type AppResource struct {
	Env        *configs.Env
	HostConfig root.HostConfig
	Db         db.IDatabase
}
