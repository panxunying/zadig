/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import "github.com/koderover/zadig/pkg/microservice/aslan/core/common/service/fs"

type Chart struct {
	Name       string `json:"name"`
	CodehostID int    `json:"codehostID"`
	Owner      string `json:"owner"`
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	Path       string `json:"path"`

	Files []*fs.FileInfo `json:"files,omitempty"`
}

type DockerfileTemplate struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type DockerfileListObject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DockerfileDetail struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Content   string      `json:"content"`
	Variables []*Variable `json:"variable"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BuildReference struct {
	BuildName   string `json:"build_name"`
	ProjectName string `json:"project_name"`
}
