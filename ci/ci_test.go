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
	assert.Equal(t, "", ci.Name())
}

func TestUnknownCI(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("CI", "true"))
	assert.Equal(t, true, ci.IsCI())
	assert.Equal(t, false, ci.IsPR())
	assert.Equal(t, "", ci.Name())
}

func TestChecker(t *testing.T) {
	for _, tt := range []struct {
		vendor ci.Vendor
		want   ci.Info
		envs   map[string]string
	}{
		{
			vendor: ci.Vendor{Name: "d", Constant: "a", Env: ci.Has{"b"}, PR: ci.Has{"c"}},
			want:   ci.Info{Name: "d", IsCI: true, CanCheckPR: true, IsPR: true},
			envs:   map[string]string{"b": "t", "c": "t"},
		},
		{
			vendor: ci.Vendor{Name: "anon", Constant: "a", Env: ci.Has{"b"}},
			want:   ci.Info{Name: "", IsCI: true, CanCheckPR: false, IsPR: false},
			envs:   map[string]string{"CI": "t"},
		},
		{
			vendor: ci.Vendor{Name: "none", Constant: "a", Env: ci.Has{"b"}},
			want:   ci.Info{Name: "", IsCI: false, CanCheckPR: false, IsPR: false},
			envs:   map[string]string{},
		},
	} {
		os.Clearenv()
		ci.Vendors = []ci.Vendor{tt.vendor}
		defer func() { ci.Vendors = defaultVendors }()
		for e, v := range tt.envs {
			assert.NoError(t, os.Setenv(e, v))
		}

		assert.Equal(t, tt.want.Name, ci.Name())
		assert.Equal(t, tt.want.IsCI, ci.IsCI())
		assert.Equal(t, tt.want.IsPR, ci.IsPR())
		assert.Equal(t, tt.want.CanCheckPR, ci.CanCheckPR())
		assert.Equal(t, tt.want, ci.GetInfo())
	}
}
