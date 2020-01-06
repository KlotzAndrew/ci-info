package ci_test

import (
	"os"
	"testing"

	"github.com/klotzandrew/ci-info/ci"
	"github.com/stretchr/testify/assert"
)

func TestTravis(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("TRAVIS", "true"))
	assert.Equal(t, true, ci.IsCI())
	assert.Equal(t, "Travis CI", ci.CIName())

	assert.NoError(t, os.Setenv("TRAVIS_PULL_REQUEST", "true"))
	assert.Equal(t, true, ci.IsPR())
	assert.Equal(t, "Travis CI", ci.CIName())
}

func TestBitBucket(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("BITBUCKET_PR_ID", "42"))
	assert.Equal(t, true, ci.IsPR())
}

func TestNevercode(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("NEVERCODE_PULL_REQUEST", "42"))
	assert.Equal(t, true, ci.IsPR())

	os.Clearenv()
	assert.Equal(t, false, ci.IsPR())
	assert.NoError(t, os.Setenv("CHANGE_ID", ""))
	assert.Equal(t, true, ci.IsPR())
}

func TestJenkins(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("JENKINS_URL", "42"))
	assert.Equal(t, true, ci.IsCI())
	assert.Equal(t, "Jenkins", ci.CIName())

	os.Clearenv()
	assert.NoError(t, os.Setenv("BUILD_ID", "42"))
	assert.Equal(t, true, ci.IsCI())

	assert.NoError(t, os.Setenv("ghprbPullId", "42"))
	assert.Equal(t, true, ci.IsPR())

	assert.NoError(t, os.Setenv("DRONE_BUILD_EVENT", "pull_request"))
	assert.Equal(t, true, ci.IsPR())
}

func TestBuildkite(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, false, ci.IsCI())

	assert.NoError(t, os.Setenv("BUILDKITE", "true"))
	assert.Equal(t, true, ci.IsCI())

	assert.NoError(t, os.Setenv("BUILDKITE_PULL_REQUEST", "42"))
	assert.Equal(t, true, ci.IsPR())
	assert.Equal(t, "Buildkite", ci.CIName())
}

func TestHeroku(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("NODE", "/app/.heroku/node/bin/node"))
	assert.Equal(t, true, ci.IsCI())
	assert.Equal(t, false, ci.IsPR())
	assert.Equal(t, "Heroku", ci.CIName())
}

func TestGiHubActions(t *testing.T) {
	os.Clearenv()
	assert.NoError(t, os.Setenv("GITHUB_ACTIONS", "true"))
	assert.NoError(t, os.Setenv("GITHUB_EVENT_NAME", "push"))

	assert.Equal(t, true, ci.IsCI())
	assert.Equal(t, false, ci.IsPR())
	assert.Equal(t, "GitHub Actions", ci.CIName())

	assert.NoError(t, os.Setenv("GITHUB_EVENT_NAME", "pull_request"))
	assert.Equal(t, true, ci.IsPR())
}
