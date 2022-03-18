package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataWarehouseUserActivitiesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type DataWarehouseUserActivitiesClient struct {
	BaseClient
}

// NewDataWarehouseUserActivitiesClient creates an instance of the DataWarehouseUserActivitiesClient client.
func NewDataWarehouseUserActivitiesClient(subscriptionID string) DataWarehouseUserActivitiesClient {
	return NewDataWarehouseUserActivitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataWarehouseUserActivitiesClientWithBaseURI creates an instance of the DataWarehouseUserActivitiesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewDataWarehouseUserActivitiesClientWithBaseURI(baseURI string, subscriptionID string) DataWarehouseUserActivitiesClient {
	return DataWarehouseUserActivitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the user activities of a data warehouse which includes running and suspended queries
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client DataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DataWarehouseUserActivities, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataWarehouseUserActivitiesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataWarehouseUserActivitiesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":                  autorest.Encode("path", databaseName),
		"dataWarehouseUserActivityName": autorest.Encode("path", "current"),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"serverName":                    autorest.Encode("path", serverName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataWarehouseUserActivitiesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataWarehouseUserActivitiesClient) GetResponder(resp *http.Response) (result DataWarehouseUserActivities, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}