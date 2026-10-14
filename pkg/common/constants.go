/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

const (
	ErrorCodePodAlreadyExists = "code=409"
)

const (
	SparkApplicationFinalizerName          = "sparkoperator.k8s.io/finalizer"
	ScheduledSparkApplicationFinalizerName = "sparkoperator.k8s.io/finalizer"
)

const (
	RSAKeySize = 2048
)

const (
	// Internal Cert
	CAKeyPem      = "ca-key.pem"
	CACertPem     = "ca-cert.pem"
	ServerKeyPem  = "server-key.pem"
	ServerCertPem = "server-cert.pem"

	// Cert Manager
	CACert  = "ca.crt"
	TLSCert = "tls.crt"
	TLSKey  = "tls.key"
)

// Kubernetes volume types.
const (
	VolumeTypeEmptyDir              = "emptyDir"
	VolumeTypeHostPath              = "hostPath"
	VolumeTypeNFS                   = "nfs"
	VolumeTypePersistentVolumeClaim = "persistentVolumeClaim"
)

const (
	// Epsilon is a small number used to compare 64 bit floating point numbers.
	Epsilon = 1e-9
)

// Ingress controller types supported for path-based routing of the Spark web UI.
// Host-based routing (an ingress URL format without a path) works with any
// controller and is unaffected by this setting.
const (
	// IngressControllerNginx is the ingress-nginx controller. It is the default,
	// matching upstream behaviour.
	IngressControllerNginx = "nginx"
	// IngressControllerKong is the Kong ingress controller.
	IngressControllerKong = "kong"
)

// Ingress annotations used to strip the per-application subpath before the
// request reaches the Spark driver. Each controller spells this differently.
const (
	// NginxRewriteTargetAnnotation rewrites the request to the capture group that
	// createDriverIngress appends to the ingress path.
	NginxRewriteTargetAnnotation = "nginx.ingress.kubernetes.io/rewrite-target"
	// KongStripPathAnnotation tells Kong to strip the matched path prefix before
	// proxying to the upstream service. Kong matches the path as a literal prefix,
	// so it needs no capture group.
	KongStripPathAnnotation = "konghq.com/strip-path"
)
