package version

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/airdb/sailor/sliceutil"
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	Env       string
	Repo      string
	Version   string
	Build     string
	BuildTime string
	CreatedAt time.Time
}

var (
	Repo      string
	Version   string
	Build     string
	BuildTime string
	CreatedAt time.Time
)

func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		GoVersion: runtime.Version(),
		Env:       os.Getenv("ENV"),
		Repo:      Repo,
		Version:   Version,
		Build:     Build,
		BuildTime: BuildTime,
		CreatedAt: CreatedAt,
	}
}

func (info *BuildInfo) ToString() string {
	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (info *BuildInfo) ToProject() string {
	return sliceutil.LastStringWithSplit(info.Repo, "/")
}
