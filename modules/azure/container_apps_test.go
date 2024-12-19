package azure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManagedEnvironmentExists(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := ManagedEnvironmentExistsE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}

func TestGetManagedEnvironmentE(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := GetManagedEnvironmentE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}

func TestContainerAppExists(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := ContainerAppExistsE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}

func TestGetContainerAppE(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := GetContainerAppE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}

func TestContainerAppJobExists(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := ContainerAppJobExistsE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}

func TestGetContainerJobAppE(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := GetContainerAppJobE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}
