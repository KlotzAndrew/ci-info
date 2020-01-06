package checker_test

import (
	"os"
	"testing"

	c "github.com/klotzandrew/ci-info/checker"
	"github.com/stretchr/testify/assert"
)

var defaultVendors = c.Vendors

func TestNotCI(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, c.IsCI())
	assert.Equal(t, false, c.IsPR())
	assert.Equal(t, "", c.CIName())
}

func TestChecker(t *testing.T) {
	vendors := []c.Vendor{
		c.Vendor{Name: "has", Constant: "c", Env: c.Has{"has-e"}, PR: c.Has{"has-pr"}},
		c.Vendor{Name: "any", Constant: "c", Env: c.Any{[]string{"any-e"}}, PR: c.Any{[]string{"any-pr"}}},
	}

	c.Vendors = vendors
	defer func() { c.Vendors = defaultVendors }()

	for _, vendor := range vendors {
		os.Clearenv()

		if has, ok := vendor.Env.(c.Has); ok {
			assert.NoError(t, os.Setenv(has.Env, "zzz"))
			assert.Equal(t, true, c.IsCI(), vendor)
			assert.Equal(t, vendor.Name, c.CIName())
			os.Clearenv()
		}

		if has, ok := vendor.PR.(c.Has); ok {
			assert.NoError(t, os.Setenv(has.Env, "zzz"))
			assert.Equal(t, true, c.IsPR(), vendor)
			assert.Equal(t, "", c.CIName())
			os.Clearenv()
		}

		if any, ok := vendor.Env.(c.Any); ok {
			assert.NoError(t, os.Setenv(any.Envs[0], "zzz"))
			assert.Equal(t, true, c.IsCI(), vendor)
			assert.Equal(t, vendor.Name, c.CIName())
			os.Clearenv()
		}

		if any, ok := vendor.PR.(c.Any); ok {
			assert.NoError(t, os.Setenv(any.Envs[0], "zzz"))
			assert.Equal(t, true, c.IsPR(), vendor)
			assert.Equal(t, "", c.CIName())
			os.Clearenv()
		}
	}
}
