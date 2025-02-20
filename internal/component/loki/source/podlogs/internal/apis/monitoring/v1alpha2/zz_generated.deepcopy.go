//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodLogs) DeepCopyInto(out *PodLogs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodLogs.
func (in *PodLogs) DeepCopy() *PodLogs {
	if in == nil {
		return nil
	}
	out := new(PodLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodLogs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodLogsList) DeepCopyInto(out *PodLogsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*PodLogs, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PodLogs)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodLogsList.
func (in *PodLogsList) DeepCopy() *PodLogsList {
	if in == nil {
		return nil
	}
	out := new(PodLogsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodLogsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodLogsSpec) DeepCopyInto(out *PodLogsSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	if in.RelabelConfigs != nil {
		in, out := &in.RelabelConfigs, &out.RelabelConfigs
		*out = make([]*v1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodLogsSpec.
func (in *PodLogsSpec) DeepCopy() *PodLogsSpec {
	if in == nil {
		return nil
	}
	out := new(PodLogsSpec)
	in.DeepCopyInto(out)
	return out
}
