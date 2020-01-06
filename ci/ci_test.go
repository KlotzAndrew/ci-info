package ci_test

import (
	"os"
	"testing"

	"github.com/klotzandrew/ci-info/ci"
	"github.com/stretchr/testify/assert"
)

var defaultVendors = ci.Vendors

func TestNotCI(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, ci.IsCI())
	assert.Equal(t, false, ci.IsPR())
	assert.Equal(t, "", ci.CIName())
}

func TestChecker(t *testing.T) {
	vendors := []ci.Vendor{
		ci.Vendor{Name: "has", Constant: "c", Env: ci.Has{"has-e"}, PR: ci.Has{"has-pr"}},
		ci.Vendor{Name: "any", Constant: "c", Env: ci.Any{[]string{"any-e"}}, PR: ci.Any{[]string{"any-pr"}}},
	}

	ci.Vendors = vendors
	defer func() { ci.Vendors = defaultVendors }()

	for _, vendor := range vendors {
		os.Clearenv()

		if has, ok := vendor.Env.(ci.Has); ok {
			assert.NoError(t, os.Setenv(has.Env, "zzz"))
			assert.Equal(t, true, ci.IsCI(), vendor)
			assert.Equal(t, vendor.Name, ci.CIName())
			os.Clearenv()
		}

		if has, ok := vendor.PR.(ci.Has); ok {
			assert.NoError(t, os.Setenv(has.Env, "zzz"))
			assert.Equal(t, true, ci.IsPR(), vendor)
			assert.Equal(t, "", ci.CIName())
			os.Clearenv()
		}

		if any, ok := vendor.Env.(ci.Any); ok {
			assert.NoError(t, os.Setenv(any.Envs[0], "zzz"))
			assert.Equal(t, true, ci.IsCI(), vendor)
			assert.Equal(t, vendor.Name, ci.CIName())
			os.Clearenv()
		}

		if any, ok := vendor.PR.(ci.Any); ok {
			assert.NoError(t, os.Setenv(any.Envs[0], "zzz"))
			assert.Equal(t, true, ci.IsPR(), vendor)
			assert.Equal(t, "", ci.CIName())
			os.Clearenv()
		}
	}
}
