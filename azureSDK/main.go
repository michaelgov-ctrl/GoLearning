package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

/*
$env:AZURE_SUBSCRIPTION_ID=""
$env:AZURE_TENANT_ID=""
$env:AZURE_CLIENT_ID=""
$env:AZURE_CLIENT_SECRET=""
*/

const (
	resourceGropLocation = "eastus"
	deploymentName       = "deployARM-how-to"
	templateFile         = "PhishingAzVmDeploy.json"
)

var (
	ctx     = context.Background()
	vmCount = 1
)

type deployParams struct {
	VmCount string `json:"vmCount"`
}

func readJSON(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	contents := make(map[string]interface{})
	_ = json.Unmarshal(data, &contents)
	return contents, nil
}

func getEnvVarOrExit(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		log.Fatalf("Missing environment variable %s\n", varName)
	}
	return value
}

func checkCurrentVirtualMachines(vmClient *armcompute.VirtualMachinesClient, resourceGroupName string, numberOfVMsToCreate int) (int, []*armcompute.VirtualMachine, error) {

	vms := []*armcompute.VirtualMachine{}
	pager := vmClient.NewListPager(resourceGroupName, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to page through response: %v", err)
		}
		for _, v := range page.Value {
			vms = append(vms, v)
		}
	}

	numberOfExistingVms := len(vms)
	if numberOfExistingVms == 3 {
		return 0, nil, fmt.Errorf("3 VMs already exist, please delete unused resources and try again")
	}

	TotalRequestedVMsCheck := numberOfExistingVms + numberOfVMsToCreate
	if TotalRequestedVMsCheck >= 3 {
		return 3, vms, nil
	}

	return TotalRequestedVMsCheck, vms, nil

}

func main() {
	resourceGroupName := getEnvVarOrExit("USERNAME") + "-RG"
	subscriptionId := getEnvVarOrExit("AZURE_SUBSCRIPTION_ID")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain credential: %v", err)
	}
	client, err := armresources.NewResourceGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		log.Fatalf("failed to create a client: %v", err)
	}

	//
	existenceCheck, err := client.CheckExistence(ctx, resourceGroupName, nil)
	if err != nil {
		log.Fatalf("failed to detect if resource group exists: %v", err)
	}
	log.Printf("rg existence check type(%T) output(%v)", existenceCheck, existenceCheck)
	//

	if existenceCheck.Success != true {
		// If resource group does not already exist create it
		resp, err := client.CreateOrUpdate(ctx, resourceGroupName, armresources.ResourceGroup{
			Location: to.Ptr(resourceGropLocation),
		}, nil)
		if err != nil {
			log.Fatalf("failed to obtain a response: %v", err)
		}
		log.Printf("resource group ID: %s\n", *resp.ResourceGroup.ID)
	} else {
		// If resource group does exist enumerate existing VMs and calculate how many VMs should be in the resource group
		vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionId, cred, nil)
		if err != nil {
			log.Fatalf("failed to create armcompute.NewVirtualMachinesClient: %v", err)
		}

		numberOfVms, VMs, err := checkCurrentVirtualMachines(vmClient, resourceGroupName, 5)
		if err != nil {
			log.Fatalf("failed to enumerate virtual machines: %v", err)
		}

		fmt.Printf("vmcount: %d\n", numberOfVms)
		for _, vm := range VMs {
			fmt.Println(*vm.Name)
		}

		vmCount = numberOfVms
	}

	fmt.Println(vmCount)
	template, err := readJSON(templateFile)
	if err != nil {
		return
	}

	deployParam := map[string]interface{}{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#",
		"contentVersion": "1.0.0.0",
		"parameters": map[string]interface{}{
			"vmCount": 3,
		},
	}
	jsonData, err := json.Marshal(deployParam)
	if err != nil {
		log.Fatalf("shit")
	}

	//test := `{"$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#","contentVersion": "1.0.0.0", "parameters": {"vmCount": "3"}}`
	deploymentsClient, err := armresources.NewDeploymentsClient(subscriptionId, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	//deployParam := map[string]interface{}{"vmCount": vmCount}
	deploy, err := deploymentsClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		armresources.Deployment{
			Properties: &armresources.DeploymentProperties{
				Template:   template,
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: jsonData,
			},
		},
		nil,
	)
	if err != nil {
		log.Fatalf("failed to deploy template: %v", err)
	}

	//look into polling
	//deploy.PollUntilDone(ctx, nil)

	fmt.Println(deploy)

}
