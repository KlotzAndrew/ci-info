package checker_test

import (
	"os"
	"testing"

	"github.com/klotzandrew/ci-info/checker"
	"github.com/stretchr/testify/assert"
)

func TestNotCI(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, checker.IsCI())
	assert.Equal(t, "", checker.CIName())
}

func TestIsCI(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("TRAVIS", "true"))
	assert.Equal(t, true, checker.IsCI())
	assert.Equal(t, "Travis CI", checker.CIName())
}

func TestNotPR(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, checker.IsPR())
}

func TestPRString(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("BITBUCKET_PR_ID", "42"))
	assert.Equal(t, true, checker.IsPR())
}

func TestPREnv(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("NEVERCODE_PULL_REQUEST", "42"))
	assert.Equal(t, true, checker.IsPR())

	os.Clearenv()
	assert.Equal(t, false, checker.IsPR())
	assert.NoError(t, os.Setenv("CHANGE_ID", ""))
	assert.Equal(t, true, checker.IsPR())
}

func TestPRAny(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("ghprbPullId", "42"))
	assert.Equal(t, true, checker.IsPR())

	os.Clearenv()
	assert.Equal(t, false, checker.IsPR())
	assert.NoError(t, os.Setenv("DRONE_BUILD_EVENT", "pull_request"))
	assert.Equal(t, true, checker.IsPR())
}

func TestBuildkite(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, checker.IsCI())

	assert.NoError(t, os.Setenv("BUILDKITE_PULL_REQUEST", "42"))
	assert.NoError(t, os.Setenv("BUILDKITE", "true"))
	assert.Equal(t, true, checker.IsCI())
	assert.Equal(t, true, checker.IsPR())
	assert.Equal(t, "Buildkite", checker.CIName())

	os.Clearenv()
	assert.NoError(t, os.Setenv("BUILDKITE_PULL_REQUEST", "true"))
	assert.Equal(t, true, checker.IsPR())

	os.Clearenv()
	assert.NoError(t, os.Setenv("BUILDKITE_PULL_REQUEST", "false"))
	assert.Equal(t, false, checker.IsCI())
}
