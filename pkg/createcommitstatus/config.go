// Copyright 2024 Mike Nguyen <hey(at)mike.ee>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package createcommitstatus

import githubactions "github.com/sethvargo/go-githubactions"

type Config struct {
	GitHubActionName       string
	GitHubActionPath       string
	GitHubActionRepository string // for a step executing an action, this is the owner and repository name of the action.
	GitHubActorName        string
	GitHubAPIUrl           string // e.g. https://api.github.com
	GitHubEventName        string
	GitHubEventPath        string
	GitHubBaseRef          string // populated on a pull_request event
	GitHubHeadRef          string
	GitHubJobID            string
	GitHubRef              string // fully formed ref of the branch tag that triggered the workflow run
	GitHubRepository       string
	GitHubRepositoryOwner  string
	RunnerDebug            bool // the RUNNER_DEBUG int is set to 1 if debug logging is enabled
}

func NewConfig(action *githubactions.Action) (*Config, error) {
	runner_debug := false
	if action.Getenv("RUNNER_DEBUG") == "1" {
		runner_debug = true
	}

	config := &Config{
		GitHubActionName:       action.Getenv("GITHUB_ACTION"),
		GitHubActionPath:       action.Getenv("GITHUB_ACTION_PATH"),
		GitHubActionRepository: action.Getenv("GITHUB_ACTION_REPOSITORY"),
		GitHubActorName:        action.Getenv("GITHUB_ACTOR"),
		GitHubAPIUrl:           action.Getenv("GITHUB_API_URL"),
		GitHubEventName:        action.Getenv("GITHUB_EVENT_NAME"),
		GitHubEventPath:        action.Getenv("GITHUB_EVENT_PATH"),
		GitHubBaseRef:          action.Getenv("GITHUB_BASE_REF"),
		GitHubHeadRef:          action.Getenv("GITHUB_HEAD_REF"),
		GitHubJobID:            action.Getenv("GITHUB_JOB"),
		GitHubRef:              action.Getenv("GITHUB_REF"),
		GitHubRepository:       action.Getenv("GITHUB_REPOSITORY"),
		GitHubRepositoryOwner:  action.Getenv("GITHUB_REPOSITORY_OWNER"),
		RunnerDebug:            runner_debug,
	}

	return config, nil
}
