// Copyright (c) 2021 Takayuki NAGATOMI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package takolabel

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

func ParseCreate(bytes []byte) (CreateTarget, error) {
	targetConfig := CreateTargetConfig{}
	err := yaml.Unmarshal(bytes, &targetConfig)
	if err != nil {
		return CreateTarget{}, fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	target := CreateTarget{Labels: targetConfig.Labels}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return CreateTarget{}, fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	return target, nil
}

func ParseDelete(bytes []byte) (DeleteTarget, error) {
	targetConfig := DeleteTargetConfig{}
	err := yaml.Unmarshal(bytes, &targetConfig)
	if err != nil {
		return DeleteTarget{}, err
	}

	target := DeleteTarget{Labels: targetConfig.Labels}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return DeleteTarget{}, fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	return target, nil
}
