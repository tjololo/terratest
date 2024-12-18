package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
	"github.com/stretchr/testify/require"
	"testing"
)

func ContainerAppsEnvironmentExists(t *testing.T, environmentName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := ContainerAppsEnvironmentExistsE(environmentName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

func ContainerAppsEnvironmentExistsE(environmentName string, resourceGroupName string, subscriptionID string) (bool, error) {
	client, err := CreateManagedEnvironmentsClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	_, err = client.Get(context.Background(), resourceGroupName, environmentName, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetContainerAppsEnvironmentE(environmentName string, resourceGroupName string, subscriptionID string) (*armappcontainers.ManagedEnvironment, error) {
	client, err := CreateManagedEnvironmentsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	env, err := client.Get(context.Background(), resourceGroupName, environmentName, nil)
	if err != nil {
		return nil, err
	}
	return &env.ManagedEnvironment, nil
}
