/*
Copyright 2020 The Kubernetes Authors.

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

package rbac

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/model/file"
)

var _ file.Template = &AuthProxyRoleBinding{}

// AuthProxyRoleBinding scaffolds a file that defines the role binding for the auth proxy
type AuthProxyRoleBinding struct {
	file.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *AuthProxyRoleBinding) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("config", "rbac", "auth_proxy_role_binding.yaml")
	}

	f.TemplateBody = proxyRoleBindinggTemplate

	return nil
}

const proxyRoleBindinggTemplate = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: proxy-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: proxy-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: system
`
