package checker

var Vendors = []Vendor{
	Vendor{
		Name:     "AppVeyor",
		Constant: "APPVEYOR",
		Env:      Has{"APPVEYOR"},
		PR:       Has{"APPVEYOR_PULL_REQUEST_NUMBER"},
	},
	Vendor{
		Name:     "Azure Pipelines",
		Constant: "AZURE_PIPELINES",
		Env:      Has{"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI"},
		PR:       Has{"SYSTEM_PULLREQUEST_PULLREQUESTID"},
	},
	Vendor{
		Name:     "Bamboo",
		Constant: "BAMBOO",
		Env:      Has{"bamboo_planKey"},
	},
	Vendor{
		Name:     "Bitbucket Pipelines",
		Constant: "BITBUCKET",
		Env:      Has{"BITBUCKET_COMMIT"},
		PR:       Has{"BITBUCKET_PR_ID"},
	},
	Vendor{
		Name:     "Bitrise",
		Constant: "BITRISE",
		Env:      Has{"BITRISE_IO"},
		PR:       Has{"BITRISE_PULL_REQUEST"},
	},
	Vendor{
		Name:     "Buddy",
		Constant: "BUDDY",
		Env:      Has{"BUDDY_WORKSPACE_ID"},
		PR:       Has{"BUDDY_EXECUTION_PULL_REQUEST_ID"},
	},
	Vendor{
		Name:     "Buildkite",
		Constant: "BUILDKITE",
		Env:      Has{"BUILDKITE"},
		PR:       NotEqual{"BUILDKITE_PULL_REQUEST", "false"},
	},
	Vendor{
		Name:     "CircleCI",
		Constant: "CIRCLE",
		Env:      Has{"CIRCLECI"},
		PR:       Has{"CIRCLE_PULL_REQUEST"},
	},
	Vendor{
		Name:     "Cirrus CI",
		Constant: "CIRRUS",
		Env:      Has{"CIRRUS_CI"},
		PR:       Has{"CIRRUS_PR"},
	},
	Vendor{
		Name:     "AWS CodeBuild",
		Constant: "CODEBUILD",
		Env:      Has{"CODEBUILD_BUILD_ARN"},
	},
	Vendor{
		Name:     "Codeship",
		Constant: "CODESHIP",
		Env:      Match{"CI_NAME", "codeship"},
	},
	Vendor{
		Name:     "Drone",
		Constant: "DRONE",
		Env:      Has{"DRONE"},
		PR:       Match{"DRONE_BUILD_EVENT", "pull_request"},
	},
	Vendor{
		Name:     "dsari",
		Constant: "DSARI",
		Env:      Has{"DSARI"},
	},
	Vendor{
		Name:     "GitHub Actions",
		Constant: "GITHUB_ACTIONS",
		Env:      Has{"GITHUB_ACTIONS"},
		PR:       Match{"GITHUB_EVENT_NAME", "pull_request"},
	},
	Vendor{
		Name:     "GitLab CI",
		Constant: "GITLAB",
		Env:      Has{"GITLAB_CI"},
	},
	Vendor{
		Name:     "GoCD",
		Constant: "GOCD",
		Env:      Has{"GO_PIPELINE_LABEL"},
	},
	Vendor{
		Name:     "Hudson",
		Constant: "HUDSON",
		Env:      Has{"HUDSON_URL"},
	},
	Vendor{
		Name:     "Jenkins",
		Constant: "JENKINS",
		Env:      Any{Envs: []string{"JENKINS_URL", "BUILD_ID"}},
		PR:       Any{Envs: []string{"ghprbPullId", "CHANGE_ID"}},
	},
	Vendor{
		Name:     "ZEIT Now",
		Constant: "ZEIT_NOW",
		Env:      Has{"NOW_BUILDER"},
	},
	Vendor{
		Name:     "Magnum CI",
		Constant: "MAGNUM",
		Env:      Has{"MAGNUM"},
	},
	Vendor{
		Name:     "Netlify CI",
		Constant: "NETLIFY",
		Env:      Has{"NETLIFY_BUILD_BASE"},
		PR:       NotEqual{"PULL_REQUEST", "false"},
	},
	Vendor{
		Name:     "Nevercode",
		Constant: "NEVERCODE",
		Env:      Has{"NEVERCODE"},
		PR:       NotEqual{"NEVERCODE_PULL_REQUEST", "false"},
	},
	Vendor{
		Name:     "Sail CI",
		Constant: "SAIL",
		Env:      Has{"SAILCI"},
		PR:       Has{"SAIL_PULL_REQUEST_NUMBER"},
	},
	Vendor{
		Name:     "Render",
		Constant: "RENDER",
		Env:      Has{"RENDER"},
		PR:       Match{"IS_PULL_REQUEST", "true"},
	},
	Vendor{
		Name:     "Semaphore",
		Constant: "SEMAPHORE",
		Env:      Has{"SEMAPHORE"},
		PR:       Has{"PULL_REQUEST_NUMBER"},
	},
	Vendor{
		Name:     "Shippable",
		Constant: "SHIPPABLE",
		Env:      Has{"SHIPPABLE"},
		PR:       Match{"IS_PULL_REQUEST", "true"},
	},
	Vendor{
		Name:     "Solano CI",
		Constant: "SOLANO",
		Env:      Has{"TDDIUM"},
		PR:       Has{"TDDIUM_PR_ID"},
	},
	Vendor{
		Name:     "Strider CD",
		Constant: "STRIDER",
		Env:      Has{"STRIDER"},
	},
	Vendor{
		Name:     "TaskCluster",
		Constant: "TASKCLUSTER",
		Env:      Any{[]string{"TASK_ID", "RUN_ID"}},
	},
	Vendor{
		Name:     "TeamCity",
		Constant: "TEAMCITY",
		Env:      Has{"TEAMCITY_VERSION"},
	},
	Vendor{
		Name:     "Travis CI",
		Constant: "TRAVIS",
		Env:      Has{"TRAVIS"},
		PR:       NotEqual{"TRAVIS_PULL_REQUEST", "false"},
	},
	Vendor{
		Name:     "Heroku",
		Constant: "HEROKU",
		Env:      Includes{"NODE", "heroku"},
	},
}
