package azure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainerAppsEnvironmentExists(t *testing.T) {
	t.Parallel()

	environmentName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := ContainerAppsEnvironmentExistsE(environmentName, resourceGroupName, subscriptionID)
	require.Error(t, err)
}
