package checker

import "os"

import "strings"

func IsCI() bool {
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

type Checker interface {
	Check() bool
}

type Vendor struct {
	Name, Constant string
	Env, PR        Checker
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
