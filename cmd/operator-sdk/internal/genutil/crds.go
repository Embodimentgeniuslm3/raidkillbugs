// Copyright 2018 The Operator-SDK Authors
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

package genutil

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/operator-framework/operator-sdk/internal/scaffold"
	"github.com/operator-framework/operator-sdk/internal/scaffold/input"
	"github.com/operator-framework/operator-sdk/internal/util/k8sutil"
	"github.com/operator-framework/operator-sdk/internal/util/projutil"

	log "github.com/sirupsen/logrus"
)

// CRDGen generates CRDs for all APIs in pkg/apis.
func CRDGen() error {
	projutil.MustInProjectRoot()

	absProjectPath := projutil.MustGetwd()
	repoPkg := projutil.GetGoPkg()

	gvMap, err := k8sutil.ParseGroupSubpackages(scaffold.ApisDir)
	if err != nil {
		return fmt.Errorf("failed to parse group versions: (%v)", err)
	}
	gvb := &strings.Builder{}
	for g, vs := range gvMap {
		gvb.WriteString(fmt.Sprintf("%s:%v, ", g, vs))
	}

	log.Infof("Running CRD generation for Custom Resource group versions: [%v]\n", gvb.String())

	s := &scaffold.Scaffold{}
	cfg := &input.Config{
		Repo:           repoPkg,
		AbsProjectPath: absProjectPath,
		ProjectName:    filepath.Base(absProjectPath),
	}
	crds, err := k8sutil.GetCRDs(scaffold.CRDsDir)
	if err != nil {
		return err
	}
	for _, crd := range crds {
		g, v, k := crd.Spec.Group, crd.Spec.Version, crd.Spec.Names.Kind
		if v == "" {
			if len(crd.Spec.Versions) != 0 {
				v = crd.Spec.Versions[0].Name
			} else {
				return fmt.Errorf("crd of group %s kind %s has no version", g, k)
			}
		}
		r, err := scaffold.NewResource(g+"/"+v, k)
		if err != nil {
			return err
		}
		err = s.Execute(cfg,
			&scaffold.CRD{Resource: r, IsOperatorGo: projutil.IsOperatorGo()},
		)
		if err != nil {
			return err
		}
	}

	log.Info("CRD generation complete.")
	return nil
}
