//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1beta1.DaemonSet{}, func(obj interface{}) { SetObjectDefaults_DaemonSet(obj.(*v1beta1.DaemonSet)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.DaemonSetList{}, func(obj interface{}) { SetObjectDefaults_DaemonSetList(obj.(*v1beta1.DaemonSetList)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.Deployment{}, func(obj interface{}) { SetObjectDefaults_Deployment(obj.(*v1beta1.Deployment)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.DeploymentList{}, func(obj interface{}) { SetObjectDefaults_DeploymentList(obj.(*v1beta1.DeploymentList)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.Ingress{}, func(obj interface{}) { SetObjectDefaults_Ingress(obj.(*v1beta1.Ingress)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.IngressList{}, func(obj interface{}) { SetObjectDefaults_IngressList(obj.(*v1beta1.IngressList)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.NetworkPolicy{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicy(obj.(*v1beta1.NetworkPolicy)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.NetworkPolicyList{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicyList(obj.(*v1beta1.NetworkPolicyList)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.ReplicaSet{}, func(obj interface{}) { SetObjectDefaults_ReplicaSet(obj.(*v1beta1.ReplicaSet)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.ReplicaSetList{}, func(obj interface{}) { SetObjectDefaults_ReplicaSetList(obj.(*v1beta1.ReplicaSetList)) })
	return nil
}

func SetObjectDefaults_DaemonSet(in *v1beta1.DaemonSet) {
	SetDefaults_DaemonSet(in)
	v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		v1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			if a.VolumeSource.ISCSI.ISCSIInterface == "" {
				a.VolumeSource.ISCSI.ISCSIInterface = "default"
			}
		}
		if a.VolumeSource.RBD != nil {
			if a.VolumeSource.RBD.RBDPool == "" {
				a.VolumeSource.RBD.RBDPool = "rbd"
			}
			if a.VolumeSource.RBD.RadosUser == "" {
				a.VolumeSource.RBD.RadosUser = "admin"
			}
			if a.VolumeSource.RBD.Keyring == "" {
				a.VolumeSource.RBD.Keyring = "/etc/ceph/keyring"
			}
		}
		if a.VolumeSource.DownwardAPI != nil {
			v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					v1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			v1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
		if a.VolumeSource.Ephemeral != nil {
			if a.VolumeSource.Ephemeral.VolumeClaimTemplate != nil {
				v1.SetDefaults_PersistentVolumeClaimSpec(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests)
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.EphemeralContainers {
		a := &in.Spec.Template.Spec.EphemeralContainers[i]
		v1.SetDefaults_EphemeralContainer(a)
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	v1.SetDefaults_ResourceList(&in.Spec.Template.Spec.Overhead)
}

func SetObjectDefaults_DaemonSetList(in *v1beta1.DaemonSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_DaemonSet(a)
	}
}

func SetObjectDefaults_Deployment(in *v1beta1.Deployment) {
	SetDefaults_Deployment(in)
	v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		v1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			if a.VolumeSource.ISCSI.ISCSIInterface == "" {
				a.VolumeSource.ISCSI.ISCSIInterface = "default"
			}
		}
		if a.VolumeSource.RBD != nil {
			if a.VolumeSource.RBD.RBDPool == "" {
				a.VolumeSource.RBD.RBDPool = "rbd"
			}
			if a.VolumeSource.RBD.RadosUser == "" {
				a.VolumeSource.RBD.RadosUser = "admin"
			}
			if a.VolumeSource.RBD.Keyring == "" {
				a.VolumeSource.RBD.Keyring = "/etc/ceph/keyring"
			}
		}
		if a.VolumeSource.DownwardAPI != nil {
			v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					v1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			v1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
		if a.VolumeSource.Ephemeral != nil {
			if a.VolumeSource.Ephemeral.VolumeClaimTemplate != nil {
				v1.SetDefaults_PersistentVolumeClaimSpec(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests)
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.EphemeralContainers {
		a := &in.Spec.Template.Spec.EphemeralContainers[i]
		v1.SetDefaults_EphemeralContainer(a)
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	v1.SetDefaults_ResourceList(&in.Spec.Template.Spec.Overhead)
}

func SetObjectDefaults_DeploymentList(in *v1beta1.DeploymentList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Deployment(a)
	}
}

func SetObjectDefaults_Ingress(in *v1beta1.Ingress) {
	for i := range in.Spec.Rules {
		a := &in.Spec.Rules[i]
		if a.IngressRuleValue.HTTP != nil {
			for j := range a.IngressRuleValue.HTTP.Paths {
				b := &a.IngressRuleValue.HTTP.Paths[j]
				SetDefaults_HTTPIngressPath(b)
			}
		}
	}
}

func SetObjectDefaults_IngressList(in *v1beta1.IngressList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Ingress(a)
	}
}

func SetObjectDefaults_NetworkPolicy(in *v1beta1.NetworkPolicy) {
	SetDefaults_NetworkPolicy(in)
}

func SetObjectDefaults_NetworkPolicyList(in *v1beta1.NetworkPolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NetworkPolicy(a)
	}
}

func SetObjectDefaults_ReplicaSet(in *v1beta1.ReplicaSet) {
	SetDefaults_ReplicaSet(in)
	v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		v1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			if a.VolumeSource.ISCSI.ISCSIInterface == "" {
				a.VolumeSource.ISCSI.ISCSIInterface = "default"
			}
		}
		if a.VolumeSource.RBD != nil {
			if a.VolumeSource.RBD.RBDPool == "" {
				a.VolumeSource.RBD.RBDPool = "rbd"
			}
			if a.VolumeSource.RBD.RadosUser == "" {
				a.VolumeSource.RBD.RadosUser = "admin"
			}
			if a.VolumeSource.RBD.Keyring == "" {
				a.VolumeSource.RBD.Keyring = "/etc/ceph/keyring"
			}
		}
		if a.VolumeSource.DownwardAPI != nil {
			v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					v1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			v1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
		if a.VolumeSource.Ephemeral != nil {
			if a.VolumeSource.Ephemeral.VolumeClaimTemplate != nil {
				v1.SetDefaults_PersistentVolumeClaimSpec(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits)
				v1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests)
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			v1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.EphemeralContainers {
		a := &in.Spec.Template.Spec.EphemeralContainers[i]
		v1.SetDefaults_EphemeralContainer(a)
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		v1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			v1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet != nil {
				v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					v1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	v1.SetDefaults_ResourceList(&in.Spec.Template.Spec.Overhead)
}

func SetObjectDefaults_ReplicaSetList(in *v1beta1.ReplicaSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ReplicaSet(a)
	}
}
