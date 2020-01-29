package enum

import "strings"

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

var EnvMap = map[Env]string{
	EnvDev:     "dev",
	EnvTest:    "test",
	EnvStable:  "stable",
	EnvStaging: "staging",
	EnvUAT:     "uat",
	EnvLivesh:  "livesh",
	EnvLive:    "live",
}

func FromEnv(env Env) string {
	if result, ok := EnvMap[env]; ok {
		return result
	}

	return EnvMap[EnvDev]
}

func ToEnv(sEnv string) Env {
	for k, v := range EnvMap {
		if v == strings.ToLower(sEnv) {
			return k
		}
	}

	return EnvDev
}

func IsLiveEnv(sEnv string) bool {
	return sEnv == EnvMap[EnvLive] || sEnv == EnvMap[EnvLivesh]
}

func GetEnvList() (envList []string) {
	for _, v := range EnvMap {
		envList = append(envList, v)
	}

	return
}

func GetNonLiveEnvList() (envList []string) {
	for k, v := range EnvMap {
		if k == EnvLivesh || k == EnvLive {
			envList = append(envList, v)
		}
	}

	return
}
