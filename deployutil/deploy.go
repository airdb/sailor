package deployutil

import (
	"os"
	"strings"
)

const (
	CloudPlatformAWS    = "aws"
	CloudPlatformAliyun = "aliyun"

	// The name from github, XD.
	CloudPlatformTencent = "tencentyun"
)

type DeployPolicy uint

const (
	DeployPolicyBlue DeployPolicy = iota + 1
	DeployPolicyGreen
	DeployPolicyRed
	DeployPolicyBlack
)

const (
	DeployStageDev     = "dev"
	DeployStageTest    = "test"
	DeployStagePre     = "pre"
	DeployStageRelease = "release"
)

var DeployPolicyMap = map[DeployPolicy]string{
	// Blue strategy serves about at 0-20% traffic.
	DeployPolicyBlue: "blue",

	// Green strategy serves about at 20-50% traffic.
	DeployPolicyGreen: "green",

	// Red strategy serves about at 50-100% traffic.
	DeployPolicyRed: "red",

	// Black strategy is for hot back, can serves all(100%) traffic.
	DeployPolicyBlack: "black",
}

func FromDeployPolicy(policy DeployPolicy) string {
	if result, ok := DeployPolicyMap[policy]; ok {
		return result
	}

	return DeployPolicyMap[policy]
}

func ToEnv(sPolicy string) DeployPolicy {
	for k, v := range DeployPolicyMap {
		if v == strings.ToLower(sPolicy) {
			return k
		}
	}

	// Fallback to default if nothing else works.
	return DeployPolicyBlue
}

func GetDeployStage() string {
	stage := os.Getenv("env")

	if stage == "" {
		stage = DeployStageDev
	}

	return stage
}

func IsStageDev() bool {
	return GetDeployStage() == DeployStageDev
}

func IsStageTest() bool {
	return GetDeployStage() == DeployStageTest
}
