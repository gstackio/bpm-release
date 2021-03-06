// Copyright (C) 2018-Present CloudFoundry.org Foundation, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package config

import (
	"io/ioutil"
	"path/filepath"
)

const DefaultBoshRoot = "/var/vcap"

type Bosh struct {
	root string
}

func NewBosh(root string) *Bosh {
	if root == "" {
		root = DefaultBoshRoot
	}

	return &Bosh{
		root: root,
	}
}

func (b *Bosh) Root() string {
	return b.root
}

func (b *Bosh) JobPaths() []string {
	var jobDirs []string

	dirInfos, err := ioutil.ReadDir(b.root)
	if err != nil {
		return jobDirs
	}

	for _, dirInfo := range dirInfos {
		if dirInfo.IsDir() {
			jobDirs = append(jobDirs, filepath.Join(b.root, dirInfo.Name()))
		}
	}

	return jobDirs
}
