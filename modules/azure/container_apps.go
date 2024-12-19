package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
	"github.com/stretchr/testify/require"
	"testing"
)

func ManagedEnvironmentExists(t *testing.T, environmentName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := ManagedEnvironmentExistsE(environmentName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

func ManagedEnvironmentExistsE(environmentName string, resourceGroupName string, subscriptionID string) (bool, error) {
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

func GetManagedEnvironmentE(environmentName string, resourceGroupName string, subscriptionID string) (*armappcontainers.ManagedEnvironment, error) {
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

func ContainerAppExists(t *testing.T, containerAppName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := ContainerAppExistsE(containerAppName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

func ContainerAppExistsE(containerAppName string, resourceGroupName string, subscriptionID string) (bool, error) {
	client, err := CreateContainerAppsClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	_, err = client.Get(context.Background(), resourceGroupName, containerAppName, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetContainerAppE(environmentName string, resourceGroupName string, subscriptionID string) (*armappcontainers.ContainerApp, error) {
	client, err := CreateContainerAppsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	app, err := client.Get(context.Background(), resourceGroupName, environmentName, nil)
	if err != nil {
		return nil, err
	}
	return &app.ContainerApp, nil
}
