package features

type UserFeatures struct {
	VirtualMachine         VirtualMachineFeatures
	VirtualMachineScaleSet VirtualMachineScaleSetFeatures
	KeyVault               KeyVaultFeatures
	Network                NetworkFeatures
	TemplateDeployment     TemplateDeploymentFeatures
	LogAnalyticsWorkspace  LogAnalyticsWorkspaceFeatures
}

type VirtualMachineFeatures struct {
	DeleteOSDiskOnDeletion bool
	GracefulShutdown       bool
	ForceDelete            bool
}

type VirtualMachineScaleSetFeatures struct {
	RollInstancesWhenRequired bool
}

type KeyVaultFeatures struct {
	PurgeSoftDeleteOnDestroy    bool
	RecoverSoftDeletedKeyVaults bool
}

type NetworkFeatures struct {
	RelaxedLocking bool
}

type TemplateDeploymentFeatures struct {
	DeleteNestedItemsDuringDeletion bool
}

type LogAnalyticsWorkspaceFeatures struct {
	PermanentlyDeleteOnDestroy bool
}
