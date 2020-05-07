//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	operatorv1alpha1 "github.com/ibm/ibm-licensing-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const APIUploadTokenName = "ibm-licensing-upload-token"
const APISecretTokenKeyName = "token"
const APIUploadTokenKeyName = "token-upload"

func GetAPISecretToken(instance *operatorv1alpha1.IBMLicensing) *corev1.Secret {
	metaLabels := LabelsForLicensingMeta(instance)
	expectedSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Spec.APISecretToken,
			Namespace: instance.Spec.InstanceNamespace,
			Labels:    metaLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{APISecretTokenKeyName: RandString(24)},
	}
	return expectedSecret
}

func GetUploadToken(instance *operatorv1alpha1.IBMLicensing) *corev1.Secret {
	metaLabels := LabelsForLicensingMeta(instance)
	expectedSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      APIUploadTokenName,
			Namespace: instance.Spec.InstanceNamespace,
			Labels:    metaLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{APIUploadTokenKeyName: RandString(24)},
	}
	return expectedSecret
}

func GetUploadConfigMap(instance *operatorv1alpha1.IBMLicensing) *corev1.ConfigMap {
	metaLabels := LabelsForLicensingMeta(instance)
	expectedCM := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ibm-licensing-upload-config",
			Namespace: instance.Spec.InstanceNamespace,
			Labels:    metaLabels,
		},
		Data: map[string]string{"url": GetUploadURL(instance)},
	}
	return expectedCM
}
