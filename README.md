[![Build Status](https://travis-ci.com/KlotzAndrew/ci-info.svg?branch=master)](https://travis-ci.com/KlotzAndrew/ci-info)
[![codecov](https://codecov.io/gh/KlotzAndrew/ci-info/branch/master/graph/badge.svg)](https://codecov.io/gh/KlotzAndrew/ci-info)
[![Coverage Status](https://coveralls.io/repos/github/KlotzAndrew/ci-info/badge.svg?branch=master)](https://coveralls.io/github/KlotzAndrew/ci-info?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/1e6e6435a35df9bed74f/maintainability)](https://codeclimate.com/github/KlotzAndrew/ci-info/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1e6e6435a35df9bed74f/test_coverage)](https://codeclimate.com/github/KlotzAndrew/ci-info/test_coverage)

# ci-info

Get details about the current Continuous Integration environment. Download as a
binary and get info regardless of language

## Installation


```bash
curl -L https://github.com//KlotzAndrew/ci-info/releases/latest/download/ci-info.linux-amd64 > ./ci-info
chmod +x ./ci-info

./ci-info help
```

You can find the latest release here: https://github.com/KlotzAndrew/ci-info/releases/latest

## Installation from source

```bash
go get -u https://github.com/KlotzAndrew/ci-info

ci-info help
```


## Usage in CLI

```bash
ci-info isci
# true

ci-info ispr
# flase

ci-info ciname
# Travis CI
```

## Usage in a GO project

```go
package main

import (
	"fmt"

	"github.com/klotzandrew/ci-info/checker"
)

func main() {
	fmt.Printf(
		"ci: %v, pr %v, name: %v",
		checker.IsCI(),
		checker.IsPR(),
		checker.CIName(),
	)
}
```

## Updating Supported CI tools

Data here: checker/vendors.go

## Supported CI tools

| Name | Constant | isPR |
|------|----------|------|
| [AWS CodeBuild](https://aws.amazon.com/codebuild/) | `ci.CODEBUILD` | 🚫 |
| [AppVeyor](http://www.appveyor.com) | `ci.APPVEYOR` | ✅ |
| [Azure Pipelines](https://azure.microsoft.com/en-us/services/devops/pipelines/) | `ci.AZURE_PIPELINES` | ✅ |
| [Bamboo](https://www.atlassian.com/software/bamboo) by Atlassian | `ci.BAMBOO` | 🚫 |
| [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines) | `ci.BITBUCKET` | ✅ |
| [Bitrise](https://www.bitrise.io/) | `ci.BITRISE` | ✅ |
| [Buddy](https://buddy.works/) | `ci.BUDDY` | ✅ |
| [Buildkite](https://buildkite.com) | `ci.BUILDKITE` | ✅ |
| [CircleCI](http://circleci.com) | `ci.CIRCLE` | ✅ |
| [Cirrus CI](https://cirrus-ci.org) | `ci.CIRRUS` | ✅ |
| [Codeship](https://codeship.com) | `ci.CODESHIP` | 🚫 |
| [Drone](https://drone.io) | `ci.DRONE` | ✅ |
| [dsari](https://github.com/rfinnie/dsari) | `ci.DSARI` | 🚫 |
| [GitLab CI](https://about.gitlab.com/gitlab-ci/) | `ci.GITLAB` | 🚫 |
| [GoCD](https://www.go.cd/) | `ci.GOCD` | 🚫 |
| [Heroku](https://www.heroku.com) | `ci.HEROKU` | 🚫 |
| [Hudson](http://hudson-ci.org) | `ci.HUDSON` | 🚫 |
| [Jenkins CI](https://jenkins-ci.org) | `ci.JENKINS` | ✅ |
| [Magnum CI](https://magnum-ci.com) | `ci.MAGNUM` | 🚫 |
| [Netlify CI](https://www.netlify.com/) | `ci.NETLIFY` | ✅ |
| [Nevercode](http://nevercode.io/) | `ci.NEVERCODE` | ✅ |
| [Sail CI](https://sail.ci/) | `ci.SAIL` | ✅ |
| [Semaphore](https://semaphoreci.com) | `ci.SEMAPHORE` | ✅ |
| [Shippable](https://www.shippable.com/) | `ci.SHIPPABLE` | ✅ |
| [Solano CI](https://www.solanolabs.com/) | `ci.SOLANO` | ✅ |
| [Strider CD](https://strider-cd.github.io/) | `ci.STRIDER` | 🚫 |
| [TaskCluster](http://docs.taskcluster.net) | `ci.TASKCLUSTER` | 🚫 |
| [TeamCity](https://www.jetbrains.com/teamcity/) by JetBrains | `ci.TEAMCITY` | 🚫 |
| [Travis CI](http://travis-ci.org) | `ci.TRAVIS` | ✅ |
