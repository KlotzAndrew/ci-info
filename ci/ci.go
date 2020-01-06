package ci

import (
	"os"
	"strings"
)

func IsCI() bool {
	if knownVendor() {
		return true
	}

	if unknownVendor() {
		return true
	}

	return false
}

var anonymousEnvs = []string{
	"CI",                     // Travis CI, CircleCI, Cirrus CI, Gitlab CI, Appveyor, CodeShip, dsari
	"CONTINUOUS_INTEGRATION", // Travis CI, Cirrus CI
	"BUILD_NUMBER",           // Jenkins, TeamCity
	"RUN_ID",                 // TaskCluster, dsari
}

func unknownVendor() bool {
	for _, env := range anonymousEnvs {
		_, found := os.LookupEnv(env)
		return found
	}
	return false
}

func knownVendor() bool {
	for _, vendor := range Vendors {
		if vendor.Env.Check() {
			return true
		}
	}
	return false
}

func IsPR() bool {
	for _, vendor := range Vendors {
		if vendor.PR != nil && vendor.PR.Check() {
			return true
		}
	}
	return false
}

func CIName() string {
	for _, vendor := range Vendors {
		if vendor.Env.Check() {
			return vendor.Name
		}
	}
	return ""
}

type Vendor struct {
	Name, Constant string
	Env, PR        Checker
}

type Checker interface {
	Check() bool
}

type Has struct{ Env string }

func (n Has) Check() bool {
	_, found := os.LookupEnv(n.Env)
	return found
}

type Match struct{ Env, Val string }

func (c Match) Check() bool {
	val, found := os.LookupEnv(c.Env)
	return found && val == c.Val
}

type NotEqual struct{ Env, Ne string }

func (n NotEqual) Check() bool {
	val, found := os.LookupEnv(n.Env)
	return found && val != n.Ne
}

type Any struct{ Envs []string }

func (a Any) Check() bool {
	for _, val := range a.Envs {
		if _, found := os.LookupEnv(val); found {
			return true
		}
	}
	return false
}

type Includes struct{ Env, Sub string }

func (c Includes) Check() bool {
	val, found := os.LookupEnv(c.Env)
	return found && strings.Contains(val, c.Sub)
}
