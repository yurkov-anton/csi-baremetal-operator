/*
Copyright © 2021 Dell Inc. or its subsidiaries. All Rights Reserved.

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

package scenarios

import (
	"github.com/onsi/ginkgo"

	"github.com/dell/csi-baremetal-operator/test/e2e/common"
)

var _ = ginkgo.Describe("CSI Operator", func() {
	operatorCleanup := func() {}

	ginkgo.BeforeSuite(func() {
		clientset, err := common.GetGlobalClientSet()
		if err != nil {
			ginkgo.Fail(err.Error())
		}

		operatorCleanup, err = common.DeployOperator(clientset)
		if err != nil {
			ginkgo.Fail(err.Error())
		}
	})

	ginkgo.AfterSuite(func() {
		operatorCleanup()
	})

	ginkgo.Context("CSI Operator test suites", func() {
		CSIDeployTestSuite()
	})
})