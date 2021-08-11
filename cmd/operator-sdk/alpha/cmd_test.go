// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alpha

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Running an alpha command", func() {
	Describe("NewCmd", func() {
		It("builds and returns a cobra command with the correct subcommand", func() {
			cmd := NewCmd()
			Expect(cmd).NotTo(BeNil())

			subcommands := cmd.Commands()
			Expect(len(subcommands)).To(Equal(1))
			Expect(subcommands[0].Use).To(Equal("scorecard"))
		})
	})
})
