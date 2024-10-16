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
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v66/github"
	githubactions "github.com/sethvargo/go-githubactions"
)

func Run(action *githubactions.Action) error {
	config, err := NewConfig(action)
	if err != nil {
		return err
	}

	inputs, err := NewInputs(action, config)
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(inputs.GitHubToken)

	status, resp, err := client.Repositories.CreateStatus(
		context.Background(),
		inputs.RepositoryOwner,
		inputs.RepositoryName,
		inputs.SHA,
		&github.RepoStatus{
			State:       (*string)(&inputs.State),
			TargetURL:   inputs.TargetURL,
			Description: inputs.Description,
			Context:     &inputs.Context,
		})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error creating status - expected code: %v, received code: %v", http.StatusCreated, resp.StatusCode)
	}

	if status.State != (*string)(&inputs.State) {
		return errors.New("returned status does not match input")
	}
	return nil
}
