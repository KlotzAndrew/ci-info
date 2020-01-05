package checker

var Vendors = []map[string]interface{}{
	map[string]interface{}{
		"name":     "AppVeyor",
		"constant": "APPVEYOR",
		"env":      "APPVEYOR",
		"pr":       "APPVEYOR_PULL_REQUEST_NUMBER",
	},
	map[string]interface{}{
		"name":     "Azure Pipelines",
		"constant": "AZURE_PIPELINES",
		"env":      "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
		"pr":       "SYSTEM_PULLREQUEST_PULLREQUESTID",
	},
	map[string]interface{}{
		"name":     "Bamboo",
		"constant": "BAMBOO",
		"env":      "bamboo_planKey",
	},
	map[string]interface{}{
		"name":     "Bitbucket Pipelines",
		"constant": "BITBUCKET",
		"env":      "BITBUCKET_COMMIT",
		"pr":       "BITBUCKET_PR_ID",
	},
	map[string]interface{}{
		"name":     "Bitrise",
		"constant": "BITRISE",
		"env":      "BITRISE_IO",
		"pr":       "BITRISE_PULL_REQUEST",
	},
	map[string]interface{}{
		"name":     "Buddy",
		"constant": "BUDDY",
		"env":      "BUDDY_WORKSPACE_ID",
		"pr":       "BUDDY_EXECUTION_PULL_REQUEST_ID",
	},
	map[string]interface{}{
		"name":     "Buildkite",
		"constant": "BUILDKITE",
		"env":      "BUILDKITE",
		"pr": map[string]string{
			"env": "BUILDKITE_PULL_REQUEST",
			"ne":  "false",
		},
	},
	map[string]interface{}{
		"name":     "CircleCI",
		"constant": "CIRCLE",
		"env":      "CIRCLECI",
		"pr":       "CIRCLE_PULL_REQUEST",
	},
	map[string]interface{}{
		"name":     "Cirrus CI",
		"constant": "CIRRUS",
		"env":      "CIRRUS_CI",
		"pr":       "CIRRUS_PR",
	},
	map[string]interface{}{
		"name":     "AWS CodeBuild",
		"constant": "CODEBUILD",
		"env":      "CODEBUILD_BUILD_ARN",
	},
	map[string]interface{}{
		"name":     "Codeship",
		"constant": "CODESHIP",
		"env": map[string]interface{}{
			"CI_NAME": "codeship",
		},
	},
	map[string]interface{}{
		"name":     "Drone",
		"constant": "DRONE",
		"env":      "DRONE",
		"pr": map[string]string{
			"DRONE_BUILD_EVENT": "pull_request",
		},
	},
	map[string]interface{}{
		"name":     "dsari",
		"constant": "DSARI",
		"env":      "DSARI",
	},
	map[string]interface{}{
		"name":     "GitLab CI",
		"constant": "GITLAB",
		"env":      "GITLAB_CI",
	},
	map[string]interface{}{
		"name":     "GoCD",
		"constant": "GOCD",
		"env":      "GO_PIPELINE_LABEL",
	},
	map[string]interface{}{
		"name":     "Hudson",
		"constant": "HUDSON",
		"env":      "HUDSON_URL",
	},
	map[string]interface{}{
		"name":     "Jenkins",
		"constant": "JENKINS",
		"env": []string{
			"JENKINS_URL",
			"BUILD_ID",
		},
		"pr": map[string][]string{
			"any": []string{
				"ghprbPullId",
				"CHANGE_ID",
			},
		},
	},
	map[string]interface{}{
		"name":     "Magnum CI",
		"constant": "MAGNUM",
		"env":      "MAGNUM",
	},
	map[string]interface{}{
		"name":     "Netlify CI",
		"constant": "NETLIFY",
		"env":      "NETLIFY_BUILD_BASE",
		"pr": map[string]string{
			"env": "PULL_REQUEST",
			"ne":  "false",
		},
	},
	map[string]interface{}{
		"name":     "Nevercode",
		"constant": "NEVERCODE",
		"env":      "NEVERCODE",
		"pr": map[string]string{
			"env": "NEVERCODE_PULL_REQUEST",
			"ne":  "false",
		},
	},
	map[string]interface{}{
		"name":     "Sail CI",
		"constant": "SAIL",
		"env":      "SAILCI",
		"pr":       "SAIL_PULL_REQUEST_NUMBER",
	},
	map[string]interface{}{
		"name":     "Semaphore",
		"constant": "SEMAPHORE",
		"env":      "SEMAPHORE",
		"pr":       "PULL_REQUEST_NUMBER",
	},
	map[string]interface{}{
		"name":     "Shippable",
		"constant": "SHIPPABLE",
		"env":      "SHIPPABLE",
		"pr": map[string]string{
			"IS_PULL_REQUEST": "true",
		},
	},
	map[string]interface{}{
		"name":     "Solano CI",
		"constant": "SOLANO",
		"env":      "TDDIUM",
		"pr":       "TDDIUM_PR_ID",
	},
	map[string]interface{}{
		"name":     "Strider CD",
		"constant": "STRIDER",
		"env":      "STRIDER",
	},
	map[string]interface{}{
		"name":     "TaskCluster",
		"constant": "TASKCLUSTER",
		"env": []string{
			"TASK_ID",
			"RUN_ID",
		},
	},
	map[string]interface{}{
		"name":     "TeamCity",
		"constant": "TEAMCITY",
		"env":      "TEAMCITY_VERSION",
	},
	map[string]interface{}{
		"name":     "Travis CI",
		"constant": "TRAVIS",
		"env":      "TRAVIS",
		"pr": map[string]string{
			"env": "TRAVIS_PULL_REQUEST",
			"ne":  "false",
		},
	},
}
