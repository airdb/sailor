package enum

type Env uint

const (
	EnvDev Env = iota + 1
	EnvTest
	EnvStable
	EnvStaging
	EnvUAT
	EnvLivesh
	EnvLive
)

func FromEnv(env Env) string {
	switch env {
	case EnvTest:
		return "TEST"
	case EnvStable:
		return "STABLE"
	case EnvStaging:
		return "STAGING"
	case EnvUAT:
		return "UAT"
	case EnvLivesh:
		return "LIVESH"
	case EnvLive:
		return "LIVE"
	default:
		return "DEV"
	}
}

func ToEnv(sEnv string) Env {
	switch sEnv {
	case "TEST":
		return EnvTest
	case "STABLE":
		return EnvStable
	case "STAGING":
		return EnvStaging
	case "UAT":
		return EnvUAT
	case "LIVESH":
		return EnvLivesh
	case "LIVE":
		return EnvLive
	default:
		return EnvDev
	}
}

func IsLiveEnv(sEnv string) bool {
	return sEnv == FromEnv(EnvLive) || sEnv == FromEnv(EnvLivesh)
}

func GetEnvList() (envList []string) {
	envList = append(envList,
		FromEnv(EnvDev),
		FromEnv(EnvTest),
		FromEnv(EnvStable),
		FromEnv(EnvStaging),
		FromEnv(EnvUAT),
		FromEnv(EnvLivesh),
		FromEnv(EnvLive),
	)
	return
}

func GetNonLiveEnvList() (envList []string) {
	envList = append(envList,
		FromEnv(EnvDev),
		FromEnv(EnvTest),
		FromEnv(EnvStable),
		FromEnv(EnvStaging),
		FromEnv(EnvUAT),
	)
	return
}
