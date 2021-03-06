package storsimple8000series

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// HardwareComponentGroupsClient is the client for the HardwareComponentGroups methods of the Storsimple8000series
// service.
type HardwareComponentGroupsClient struct {
	ManagementClient
}

// NewHardwareComponentGroupsClient creates an instance of the HardwareComponentGroupsClient client.
func NewHardwareComponentGroupsClient(subscriptionID string) HardwareComponentGroupsClient {
	return NewHardwareComponentGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHardwareComponentGroupsClientWithBaseURI creates an instance of the HardwareComponentGroupsClient client.
func NewHardwareComponentGroupsClientWithBaseURI(baseURI string, subscriptionID string) HardwareComponentGroupsClient {
	return HardwareComponentGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ChangeControllerPowerState changes the power state of the controller. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// deviceName is the device name hardwareComponentGroupName is the hardware component group name. parameters is the
// controller power state change request. resourceGroupName is the resource group name managerName is the manager name
func (client HardwareComponentGroupsClient) ChangeControllerPowerState(deviceName string, hardwareComponentGroupName string, parameters ControllerPowerStateChangeRequest, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ControllerPowerStateChangeRequestProperties", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.HardwareComponentGroupsClient", "ChangeControllerPowerState")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ChangeControllerPowerStatePreparer(deviceName, hardwareComponentGroupName, parameters, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ChangeControllerPowerState", nil, "Failure preparing request")
			return
		}

		resp, err := client.ChangeControllerPowerStateSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ChangeControllerPowerState", resp, "Failure sending request")
			return
		}

		result, err = client.ChangeControllerPowerStateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ChangeControllerPowerState", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ChangeControllerPowerStatePreparer prepares the ChangeControllerPowerState request.
func (client HardwareComponentGroupsClient) ChangeControllerPowerStatePreparer(deviceName string, hardwareComponentGroupName string, parameters ControllerPowerStateChangeRequest, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":                 deviceName,
		"hardwareComponentGroupName": hardwareComponentGroupName,
		"managerName":                managerName,
		"resourceGroupName":          resourceGroupName,
		"subscriptionId":             client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups/{hardwareComponentGroupName}/changeControllerPowerState", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ChangeControllerPowerStateSender sends the ChangeControllerPowerState request. The method will close the
// http.Response Body if it receives an error.
func (client HardwareComponentGroupsClient) ChangeControllerPowerStateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ChangeControllerPowerStateResponder handles the response to the ChangeControllerPowerState request. The method always
// closes the http.Response Body.
func (client HardwareComponentGroupsClient) ChangeControllerPowerStateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDevice lists the hardware component groups at device-level.
//
// deviceName is the device name resourceGroupName is the resource group name managerName is the manager name
func (client HardwareComponentGroupsClient) ListByDevice(deviceName string, resourceGroupName string, managerName string) (result HardwareComponentGroupList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.HardwareComponentGroupsClient", "ListByDevice")
	}

	req, err := client.ListByDevicePreparer(deviceName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.HardwareComponentGroupsClient", "ListByDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client HardwareComponentGroupsClient) ListByDevicePreparer(deviceName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client HardwareComponentGroupsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client HardwareComponentGroupsClient) ListByDeviceResponder(resp *http.Response) (result HardwareComponentGroupList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
