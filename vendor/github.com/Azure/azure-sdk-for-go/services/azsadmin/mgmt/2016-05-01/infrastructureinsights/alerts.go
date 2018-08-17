package infrastructureinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// AlertsClient is the infrastructureInsights Admin Client
type AlertsClient struct {
	BaseClient
}

// NewAlertsClient creates an instance of the AlertsClient client.
func NewAlertsClient(subscriptionID string) AlertsClient {
	return NewAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertsClientWithBaseURI creates an instance of the AlertsClient client.
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return AlertsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Close close an alert.
//
// location is location name. alertName is name of the alert. userParameter is the username used to perform the
// operation. alert is updated Alert Parameter.
func (client AlertsClient) Close(ctx context.Context, location string, alertName string, userParameter string, alert Alert) (result Alert, err error) {
	req, err := client.ClosePreparer(ctx, location, alertName, userParameter, alert)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Close", nil, "Failure preparing request")
		return
	}

	resp, err := client.CloseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Close", resp, "Failure sending request")
		return
	}

	result, err = client.CloseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Close", resp, "Failure responding to request")
	}

	return
}

// ClosePreparer prepares the Close request.
func (client AlertsClient) ClosePreparer(ctx context.Context, location string, alertName string, userParameter string, alert Alert) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertName":      autorest.Encode("path", alertName),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"user":        autorest.Encode("query", userParameter),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.InfrastructureInsights.Admin/regionHealths/{location}/alerts/{alertName}", pathParameters),
		autorest.WithJSON(alert),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CloseSender sends the Close request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) CloseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CloseResponder handles the response to the Close request. The method always
// closes the http.Response Body.
func (client AlertsClient) CloseResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get an alert.
//
// location is location name. alertName is name of the alert.
func (client AlertsClient) Get(ctx context.Context, location string, alertName string) (result Alert, err error) {
	req, err := client.GetPreparer(ctx, location, alertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AlertsClient) GetPreparer(ctx context.Context, location string, alertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertName":      autorest.Encode("path", alertName),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.InfrastructureInsights.Admin/regionHealths/{location}/alerts/{alertName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns the list of all alerts in a given location.
//
// location is location name. filter is oData filter parameter.
func (client AlertsClient) List(ctx context.Context, location string, filter string) (result AlertListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "List", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AlertsClient) ListPreparer(ctx context.Context, location string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.InfrastructureInsights.Admin/regionHealths/{location}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AlertsClient) ListResponder(resp *http.Response) (result AlertList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AlertsClient) listNextResults(lastResults AlertList) (result AlertList, err error) {
	req, err := lastResults.alertListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "infrastructureinsights.AlertsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AlertsClient) ListComplete(ctx context.Context, location string, filter string) (result AlertListIterator, err error) {
	result.page, err = client.List(ctx, location, filter)
	return
}
