//
// Copyright 2022 IBM Corporation
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

package service

import (
	operatorv1alpha1 "github.com/ibm/ibm-licensing-operator/api/v1alpha1"
	"github.com/ibm/ibm-licensing-operator/controllers/resources"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const APIUploadTokenName = "ibm-licensing-upload-token"
const APISecretTokenKeyName = "token"
const APIUploadTokenKeyName = "token-upload"

const ReporterSecretTokenKeyName = "token"

const URLConfigMapKey = "url"
const CrtConfigMapKey = "crt"

func GetAPISecretToken(instance *operatorv1alpha1.IBMLicensing) (*corev1.Secret, error) {
	return resources.GetSecretToken(instance.Spec.APISecretToken, instance.Spec.InstanceNamespace, APISecretTokenKeyName, LabelsForMeta(instance))
}

func GetUploadToken(instance *operatorv1alpha1.IBMLicensing) (*corev1.Secret, error) {
	return resources.GetSecretToken(APIUploadTokenName, instance.Spec.InstanceNamespace, APIUploadTokenKeyName, LabelsForMeta(instance))
}

func GetUploadConfigMap(instance *operatorv1alpha1.IBMLicensing) *corev1.ConfigMap {
	metaLabels := LabelsForMeta(instance)
	expectedCM := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ibm-licensing-upload-config",
			Namespace: instance.Spec.InstanceNamespace,
			Labels:    metaLabels,
		},
		Data: map[string]string{URLConfigMapKey: GetServiceURL(instance)},
	}
	return expectedCM
}

func GetInfoConfigMap(instance *operatorv1alpha1.IBMLicensing, internalCertData string) *corev1.ConfigMap {
	metaLabels := LabelsForMeta(instance)
	expectedCM := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ibm-licensing-info",
			Namespace: instance.Spec.InstanceNamespace,
			Labels:    metaLabels,
		},
		Data: map[string]string{
			URLConfigMapKey: GetServiceURL(instance),
			CrtConfigMapKey: internalCertData,
		},
	}
	return expectedCM
}
