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

import (
	"errors"
	"fmt"
	"strings"

	githubactions "github.com/sethvargo/go-githubactions"
)

type Inputs struct {
	GitHubToken     string
	RepositoryOwner string
	RepositoryName  string
	SHA             string
	State           string
	TargetURL       *string
	Description     *string
	Context         string
}

func NewInputs(action *githubactions.Action, config *Config) (*Inputs, error) {
	sha, err := getSHA(action, config)
	if err != nil {
		return nil, err
	}

	status, err := getStatus(action)
	if err != nil {
		return nil, err
	}

	targetUrl := action.GetInput("target-url")
	description := action.GetInput("description")

	inputs := Inputs{
		GitHubToken:     action.GetInput("token"),
		RepositoryOwner: action.GetInput("repository-owner"),
		RepositoryName:  action.GetInput("repository"),
		SHA:             sha,
		State:           status,
		TargetURL:       &targetUrl,
		Description:     &description,
		Context:         action.GetInput("context"),
	}
	return &inputs, nil
}

func getStatus(action *githubactions.Action) (string, error) {
	input := strings.ToLower(action.GetInput("state"))
	switch input {
	case "":
		return "failure", nil
	case "error":
		return "error", nil
	case "failure":
		return "failure", nil
	case "pending":
		return "pending", nil
	case "success":
		return "success", nil
	}
	return "", fmt.Errorf("invalid state provided: %v", input)
}

func getSHA(action *githubactions.Action, config *Config) (string, error) {
	sha := action.GetInput("sha")
	if sha == "" {
		// get head ref
		if config.GitHubEventName == "pull_request" {
			return config.GitHubHeadRef, nil
		} else {
			actionContext, err := action.Context()
			if err != nil {
				return "", err
			}
			sha = actionContext.SHA
			if sha == "" {
				return "", errors.New("empty SHA found in the context")
			}
		}
	}
	return sha, nil
}
