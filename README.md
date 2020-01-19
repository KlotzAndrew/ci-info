[![Build Status](https://travis-ci.com/KlotzAndrew/ci-info.svg?branch=master)](https://travis-ci.com/KlotzAndrew/ci-info)
[![codecov](https://codecov.io/gh/KlotzAndrew/ci-info/branch/master/graph/badge.svg)](https://codecov.io/gh/KlotzAndrew/ci-info)
[![Coverage Status](https://coveralls.io/repos/github/KlotzAndrew/ci-info/badge.svg?branch=master)](https://coveralls.io/github/KlotzAndrew/ci-info?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/1e6e6435a35df9bed74f/maintainability)](https://codeclimate.com/github/KlotzAndrew/ci-info/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1e6e6435a35df9bed74f/test_coverage)](https://codeclimate.com/github/KlotzAndrew/ci-info/test_coverage)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5d251f0736d64ea29cfa4334c9aaeabb)](https://www.codacy.com/manual/KlotzAndrew/ci-info?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=KlotzAndrew/ci-info&amp;utm_campaign=Badge_Grade)


![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/KlotzAndrew/ci-info)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/KlotzAndrew/ci-info)

# ci-info

Get details about the current Continuous Integration environment.

Download as a binary and get info regardless of language

## Contributing
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github/KlotzAndrew/ci-info/issues)


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

ci-info cancheckpr
# false

ci-info ciname
# Travis CI
```

## Usage in a GO project

```go
package main

import (
	"fmt"

	"github.com/klotzandrew/ci-info/ci"
)

func main() {
	fmt.Printf(
		"ci: %v, can_check_pr: %v, pr %v, name: %v",
    ci.IsCI(),
    ci.CanCheckPR(),
		ci.IsPR(),
		ci.Name(),
  )

  info := ci.Info()
  fmt.Printf(
		"ci: %v, can_check_pr: %v, pr %v, name: %v",
    info.IsCI,
    info.CanCheckPR,
		info.IsPR,
		info.Name,
  )
}
```

## Updating Supported CI tools

Data here: ci/vendors.go

## Supported CI tools
| Name                                                                            | CanCheckPR |
|---------------------------------------------------------------------------------|------------|
| [AWS CodeBuild](https://aws.amazon.com/codebuild/)                              | 🚫         |
| [AppVeyor](http://www.appveyor.com)                                             | ✅         |
| [Azure Pipelines](https://azure.microsoft.com/en-us/services/devops/pipelines/) | ✅         |
| [Bamboo](https://www.atlassian.com/software/bamboo) by Atlassian                | 🚫         |
| [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines)         | ✅         |
| [Bitrise](https://www.bitrise.io/)                                              | ✅         |
| [Buddy](https://buddy.works/)                                                   | ✅         |
| [Buildkite](https://buildkite.com)                                              | ✅         |
| [CircleCI](http://circleci.com)                                                 | ✅         |
| [Cirrus CI](https://cirrus-ci.org)                                              | ✅         |
| [Codeship](https://codeship.com)                                                | 🚫         |
| [Drone](https://drone.io)                                                       | ✅         |
| [dsari](https://github.com/rfinnie/dsari)                                       | 🚫         |
| [GitHub Actions](https://github.com/features/actions/)                          | ✅         |
| [GitLab CI](https://about.gitlab.com/gitlab-ci/)                                | 🚫         |
| [GoCD](https://www.go.cd/)                                                      | 🚫         |
| [Heroku](https://www.heroku.com)                                                | 🚫         |
| [Hudson](http://hudson-ci.org)                                                  | 🚫         |
| [Jenkins CI](https://jenkins-ci.org)                                            | ✅         |
| [Magnum CI](https://magnum-ci.com)                                              | 🚫         |
| [ZEIT Now](https://zeit.co/)                                                    | 🚫         |
| [Netlify CI](https://www.netlify.com/)                                          | ✅         |
| [Nevercode](http://nevercode.io/)                                               | ✅         |
| [Render](https://render.com/)                                                   | ✅         |
| [Sail CI](https://sail.ci/)                                                     | ✅         |
| [Semaphore](https://semaphoreci.com)                                            | ✅         |
| [Shippable](https://www.shippable.com/)                                         | ✅         |
| [Solano CI](https://www.solanolabs.com/)                                        | ✅         |
| [Strider CD](https://strider-cd.github.io/)                                     | 🚫         |
| [TaskCluster](http://docs.taskcluster.net)                                      | 🚫         |
| [TeamCity](https://www.jetbrains.com/teamcity/) by JetBrains                    | 🚫         |
| [Travis CI](http://travis-ci.org)                                               | ✅         |
