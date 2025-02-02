package common

import (
	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/apis/v1alpha1"
)

var (
	rhmi2ProductVersions = map[integreatlyv1alpha1.StageName]map[integreatlyv1alpha1.ProductName]integreatlyv1alpha1.ProductVersion{
		integreatlyv1alpha1.AuthenticationStage: {
			integreatlyv1alpha1.ProductRHSSO: integreatlyv1alpha1.VersionRHSSO,
		},
		integreatlyv1alpha1.ObservabilityStage: {
			integreatlyv1alpha1.ProductObservability: integreatlyv1alpha1.VersionObservability,
		},
		integreatlyv1alpha1.CloudResourcesStage: {
			integreatlyv1alpha1.ProductCloudResources: integreatlyv1alpha1.VersionCloudResources,
		},
		integreatlyv1alpha1.ProductsStage: {
			integreatlyv1alpha1.Product3Scale:              integreatlyv1alpha1.Version3Scale,
			integreatlyv1alpha1.ProductFuse:                integreatlyv1alpha1.VersionFuseOnOpenshift,
			integreatlyv1alpha1.ProductRHSSOUser:           integreatlyv1alpha1.VersionRHSSOUser,
			integreatlyv1alpha1.ProductCodeReadyWorkspaces: integreatlyv1alpha1.VersionCodeReadyWorkspaces,
			integreatlyv1alpha1.ProductAMQOnline:           integreatlyv1alpha1.VersionAMQOnline,
			integreatlyv1alpha1.ProductUps:                 integreatlyv1alpha1.VersionUps,
			integreatlyv1alpha1.ProductApicurito:           integreatlyv1alpha1.VersionApicurito,
		},
		integreatlyv1alpha1.SolutionExplorerStage: {
			integreatlyv1alpha1.ProductSolutionExplorer: integreatlyv1alpha1.VersionSolutionExplorer,
		},
	}

	managedApiProductVersions = map[integreatlyv1alpha1.StageName]map[integreatlyv1alpha1.ProductName]integreatlyv1alpha1.ProductVersion{
		integreatlyv1alpha1.AuthenticationStage: {
			integreatlyv1alpha1.ProductRHSSO: integreatlyv1alpha1.VersionRHSSO,
		},
		integreatlyv1alpha1.ObservabilityStage: {
			integreatlyv1alpha1.ProductObservability: integreatlyv1alpha1.VersionObservability,
		},
		integreatlyv1alpha1.CloudResourcesStage: {
			integreatlyv1alpha1.ProductCloudResources: integreatlyv1alpha1.VersionCloudResources,
		},
		integreatlyv1alpha1.ProductsStage: {
			integreatlyv1alpha1.Product3Scale:    integreatlyv1alpha1.Version3Scale,
			integreatlyv1alpha1.ProductRHSSOUser: integreatlyv1alpha1.VersionRHSSOUser,
		},
	}
)

func TestProductVersions(t TestingTB, ctx *TestingContext) {

	rhmi, err := GetRHMI(ctx.Client, true)

	if err != nil {
		t.Fatalf("failed to get the RHMI: %s", err)
	}

	productVersions := getProductVersions(rhmi.Spec.Type)

	for stage := range productVersions {
		for productName, productVersion := range productVersions[stage] {
			productStatus := rhmi.Status.Stages[stage].Products[productName]
			clusterVersion := productStatus.Version
			if clusterVersion != productVersion {
				t.Errorf("Error with version of %s deployed on cluster. Expected %s. Got %s\nProduct status: %v", productName, productVersion, clusterVersion, productStatus)
			}
		}

	}
}

func getProductVersions(installType string) map[integreatlyv1alpha1.StageName]map[integreatlyv1alpha1.ProductName]integreatlyv1alpha1.ProductVersion {
	if integreatlyv1alpha1.IsRHOAM(integreatlyv1alpha1.InstallationType(installType)) {
		return managedApiProductVersions
	} else {
		return rhmi2ProductVersions
	}
}
