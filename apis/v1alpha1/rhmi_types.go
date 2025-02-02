/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatusPhase string
type InstallationType string
type ProductName string
type ProductVersion string
type OperatorVersion string
type PreflightStatus string
type StageName string

var (
	PhaseNone                   StatusPhase = ""
	PhaseAccepted               StatusPhase = "accepted"
	PhaseCreatingSubscription   StatusPhase = "creating subscription"
	PhaseAwaitingOperator       StatusPhase = "awaiting operator"
	PhaseAwaitingCloudResources StatusPhase = "awaiting cloud resources"
	PhaseCreatingComponents     StatusPhase = "creating components"
	PhaseAwaitingComponents     StatusPhase = "awaiting components"

	PhaseInProgress StatusPhase = "in progress"
	PhaseCompleted  StatusPhase = "completed"
	PhaseFailed     StatusPhase = "failed"

	InstallationTypeWorkshop              InstallationType = "workshop"
	InstallationTypeManaged               InstallationType = "managed"
	InstallationTypeManagedApi            InstallationType = "managed-api"
	InstallationTypeSelfManaged           InstallationType = "self-managed"
	InstallationTypeMultitenantManagedApi InstallationType = "multitenant-managed-api"

	BootstrapStage               StageName = "bootstrap"
	CloudResourcesStage          StageName = "cloud-resources"
	MonitoringStage              StageName = "monitoring"
	ObservabilityStage           StageName = "observability"
	AuthenticationStage          StageName = "authentication"
	ProductsStage                StageName = "products"
	CompleteStage                StageName = "complete"
	SolutionExplorerStage        StageName = "solution-explorer"
	UninstallProductsStage       StageName = "uninstall - products"
	UninstallMonitoringStage     StageName = "uninstall - monitoring"
	UninstallObservabilityStage  StageName = "uninstall - observability"
	UninstallCloudResourcesStage StageName = "uninstall - cloud-resources"
	UninstallBootstrap           StageName = "uninstall - bootstrap"

	ProductAMQStreams          ProductName = "amqstreams"
	ProductAMQOnline           ProductName = "amqonline"
	ProductSolutionExplorer    ProductName = "solution-explorer"
	ProductRHSSO               ProductName = "rhsso"
	ProductRHSSOUser           ProductName = "rhssouser"
	ProductCodeReadyWorkspaces ProductName = "codeready-workspaces"
	ProductFuse                ProductName = "fuse"
	ProductFuseOnOpenshift     ProductName = "fuse-on-openshift"
	Product3Scale              ProductName = "3scale"
	ProductUps                 ProductName = "ups"
	ProductApicurioRegistry    ProductName = "apicurio-registry"
	ProductApicurito           ProductName = "apicurito"
	ProductMonitoring          ProductName = "middleware-monitoring"
	ProductObservability       ProductName = "observability"
	ProductCloudResources      ProductName = "cloud-resources"
	ProductDataSync            ProductName = "datasync"
	ProductMonitoringSpec      ProductName = "monitoring-spec"
	ProductMarin3r             ProductName = "marin3r"
	ProductGrafana             ProductName = "grafana"

	// Could not find a way to determine these versions dynamically, so they are hard-coded
	// It is preferable to determine the version of a product dynamically (from a CR, or configmap, etc)
	// Follow up Jira: https://issues.redhat.com/browse/INTLY-5946
	VersionAMQOnline           ProductVersion = "1.4"
	VersionApicurioRegistry    ProductVersion = "1.2.3.final"
	VersionApicurito           ProductVersion = "7.6"
	VersionAMQStreams          ProductVersion = "1.1.0"
	VersionCodeReadyWorkspaces ProductVersion = "2.1.1"
	VersionFuseOnOpenshift     ProductVersion = "7.6"
	VersionMonitoring          ProductVersion = "1.8.0"
	Version3Scale              ProductVersion = "2.11.0"
	VersionUps                 ProductVersion = "2.3.2"
	VersionCloudResources      ProductVersion = "0.32.0"
	VersionFuseOnline          ProductVersion = "7.6"
	VersionDataSync            ProductVersion = "0.9.4"
	VersionRHSSO               ProductVersion = "7.5"
	VersionRHSSOUser           ProductVersion = "7.5"
	VersionMonitoringSpec      ProductVersion = "1.0"
	VersionSolutionExplorer    ProductVersion = "2.28.0"
	VersionMarin3r             ProductVersion = "0.8.0"
	VersionGrafana             ProductVersion = "3.10.4"
	VersionObservability       ProductVersion = "3.0.7"

	// Versioning for Fuse on OpenShift does not follow a similar pattern to other products.
	// It is currently implicitly tied to version 7.6 of Fuse, hence the 7.6 value for VersionFuseOnOpenshift above
	// The need for the below tag references is due to the nature of installing which is documented here: (for 7.6)
	// https://access.redhat.com/documentation/en-us/red_hat_fuse/7.6/html-single/fuse_on_openshift_guide/index#install-fuse-on-openshift4
	TagFuseOnOpenShiftCore        string = "application-templates-2.1.0.fuse-760043-redhat-00001"
	TagFuseOnOpenShiftSpringBoot2 string = "application-templates-2.1.0.fuse-sb2-760039-redhat-00001"

	PreflightInProgress PreflightStatus = ""
	PreflightSuccess    PreflightStatus = "successful"
	PreflightFail       PreflightStatus = "failed"

	// Operator image tags
	OperatorVersionAMQStreams       OperatorVersion = "1.1.0"
	OperatorVersionAMQOnline        OperatorVersion = "1.4"
	OperatorVersionMonitoring       OperatorVersion = "1.8.0"
	OperatorVersionSolutionExplorer OperatorVersion = "0.0.62"
	OperatorVersionRHSSO            OperatorVersion = "15.0.2"
	OperatorVersionRHSSOUser        OperatorVersion = "15.0.2"

	OperatorVersionCodeReadyWorkspaces OperatorVersion = "2.1.1"
	OperatorVersion3Scale              OperatorVersion = "0.8.0"
	OperatorVersionFuse                OperatorVersion = "1.6.0"
	OperatorVersionCloudResources      OperatorVersion = "0.32.0"
	OperatorVersionUPS                 OperatorVersion = "0.5.0"
	OperatorVersionApicurioRegistry    OperatorVersion = "0.0.3"
	OperatorVersionApicurito           OperatorVersion = "1.6.0"
	OperatorVersionMonitoringSpec      OperatorVersion = "1.0"
	OperatorVersionMarin3r             OperatorVersion = "0.8.0"
	OperatorVersionGrafana             OperatorVersion = "3.10.4"
	OperatorVersionObservability       OperatorVersion = "3.0.7"

	// Event reasons to be used when emitting events
	EventProcessingError       string = "ProcessingError"
	EventInstallationCompleted string = "InstallationCompleted"
	EventPreflightCheckPassed  string = "PreflightCheckPassed"
	EventUpgradeApproved       string = "UpgradeApproved"

	DefaultOriginPullSecretName      = "pull-secret"
	DefaultOriginPullSecretNamespace = "openshift-config"

	EnvKeyAlertSMTPFrom = "ALERT_SMTP_FROM"
	EnvKeyQuota         = "QUOTA"
)

// RHMISpec defines the desired state of RHMI
type RHMISpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Type                   string                 `json:"type"`
	RoutingSubdomain       string                 `json:"routingSubdomain,omitempty"`
	MasterURL              string                 `json:"masterURL,omitempty"`
	NamespacePrefix        string                 `json:"namespacePrefix"`
	RebalancePods          bool                   `json:"rebalancePods,omitempty"`
	SelfSignedCerts        bool                   `json:"selfSignedCerts,omitempty"`
	PullSecret             PullSecretSpec         `json:"pullSecret,omitempty"`
	UseClusterStorage      string                 `json:"useClusterStorage,omitempty"`
	AlertingEmailAddress   string                 `json:"alertingEmailAddress,omitempty"`
	PriorityClassName      string                 `json:"priorityClassName,omitempty"`
	AlertingEmailAddresses AlertingEmailAddresses `json:"alertingEmailAddresses,omitempty"`
	AlertFromAddress       string                 `json:"alertFromAddress,omitempty"`

	// OperatorsInProductNamespace is a flag that decides if
	// the product operators should be installed in the product
	// namespace (when set to true) or in standalone namespace
	// (when set to false, default). Standalone namespace will
	// be used only for those operators that support it.
	OperatorsInProductNamespace bool `json:"operatorsInProductNamespace,omitempty"`

	// SMTPSecret is the name of a secret in the installation
	// namespace containing SMTP connection details. The secret
	// must contain the following fields:
	//
	// host
	// port
	// tls
	// username
	// password
	SMTPSecret string `json:"smtpSecret,omitempty"`

	// PagerDutySecret is the name of a secret in the
	// installation namespace containing PagerDuty account
	// details. The secret must contain the following fields:
	//
	// serviceKey
	PagerDutySecret string `json:"pagerDutySecret,omitempty"`

	// DeadMansSnitchSecret is the name of a secret in the
	// installation namespace containing connection details
	// for Dead Mans Snitch. The secret must contain the
	// following fields:
	//
	// url
	DeadMansSnitchSecret string `json:"deadMansSnitchSecret,omitempty"`
}

type PullSecretSpec struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type AlertingEmailAddresses struct {
	BusinessUnit string `json:"businessUnit"`
	CSSRE        string `json:"cssre"`
}

// RHMIStatus defines the observed state of RHMI
type RHMIStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Stages             map[StageName]RHMIStageStatus `json:"stages"`
	Stage              StageName                     `json:"stage"`
	PreflightStatus    PreflightStatus               `json:"preflightStatus,omitempty"`
	PreflightMessage   string                        `json:"preflightMessage,omitempty"`
	LastError          string                        `json:"lastError"`
	GitHubOAuthEnabled bool                          `json:"gitHubOAuthEnabled,omitempty"`
	SMTPEnabled        bool                          `json:"smtpEnabled,omitempty"`
	Version            string                        `json:"version,omitempty"`
	ToVersion          string                        `json:"toVersion,omitempty"`
	Quota              string                        `json:"quota,omitempty"`
	ToQuota            string                        `json:"toQuota,omitempty"`
}

type RHMIStageStatus struct {
	Name     StageName                         `json:"name"`
	Phase    StatusPhase                       `json:"phase"`
	Products map[ProductName]RHMIProductStatus `json:"products,omitempty"`
}

type RHMIProductStatus struct {
	Name            ProductName     `json:"name"`
	OperatorVersion OperatorVersion `json:"operator,omitempty"`
	Version         ProductVersion  `json:"version"`
	Host            string          `json:"host"`
	Type            string          `json:"type,omitempty"`
	Mobile          bool            `json:"mobile,omitempty"`
	Phase           StatusPhase     `json:"status"`
	Uninstall       bool            `json:"uninstall,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RHMI is the Schema for the rhmis API
type RHMI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RHMISpec   `json:"spec,omitempty"`
	Status RHMIStatus `json:"status,omitempty"`
}

func (i *RHMI) GetProductStatusObject(product ProductName) *RHMIProductStatus {
	for _, stage := range i.Status.Stages {
		if product, ok := stage.Products[product]; ok {
			return &product
		}
	}
	return &RHMIProductStatus{
		Name: product,
	}
}

func (i *RHMI) GetPullSecretSpec() *PullSecretSpec {
	if i.Spec.PullSecret.Name != "" && i.Spec.PullSecret.Namespace != "" {
		return &(i.Spec.PullSecret)
	} else {
		return &PullSecretSpec{Name: DefaultOriginPullSecretName, Namespace: DefaultOriginPullSecretNamespace}
	}
}

// +kubebuilder:object:root=true

// RHMIList contains a list of RHMI
type RHMIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RHMI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RHMI{}, &RHMIList{})
}
