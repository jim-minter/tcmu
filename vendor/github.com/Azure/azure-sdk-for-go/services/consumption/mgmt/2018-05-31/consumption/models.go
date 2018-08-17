package consumption

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
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// Bound enumerates the values for bound.
type Bound string

const (
	// Lower ...
	Lower Bound = "Lower"
	// Upper ...
	Upper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{Lower, Upper}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// Daily ...
	Daily Grain = "Daily"
	// Monthly ...
	Monthly Grain = "Monthly"
	// Yearly ...
	Yearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{Daily, Monthly, Yearly}
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The reason
// is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Forecast a forecast resource.
type Forecast struct {
	*ForecastProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Forecast.
func (f Forecast) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if f.ForecastProperties != nil {
		objectMap["properties"] = f.ForecastProperties
	}
	if f.ID != nil {
		objectMap["id"] = f.ID
	}
	if f.Name != nil {
		objectMap["name"] = f.Name
	}
	if f.Type != nil {
		objectMap["type"] = f.Type
	}
	if f.Tags != nil {
		objectMap["tags"] = f.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Forecast struct.
func (f *Forecast) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var forecastProperties ForecastProperties
				err = json.Unmarshal(*v, &forecastProperties)
				if err != nil {
					return err
				}
				f.ForecastProperties = &forecastProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				f.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				f.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				f.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				f.Tags = tags
			}
		}
	}

	return nil
}

// ForecastProperties the properties of the forecast charge.
type ForecastProperties struct {
	// UsageDate - The usage date of the forecast.
	UsageDate *string `json:"usageDate,omitempty"`
	// Grain - The granularity of forecast. Possible values include: 'Daily', 'Monthly', 'Yearly'
	Grain Grain `json:"grain,omitempty"`
	// Charge - The amount of charge
	Charge *decimal.Decimal `json:"charge,omitempty"`
	// Currency - The ISO currency in which the meter is charged, for example, USD.
	Currency *string `json:"currency,omitempty"`
	// ChargeType - The type of the charge. Could be actual or forecast. Possible values include: 'ChargeTypeActual', 'ChargeTypeForecast'
	ChargeType ChargeType `json:"chargeType,omitempty"`
	// ConfidenceLevels - The details about the forecast confidence levels. This is populated only when chargeType is Forecast.
	ConfidenceLevels *[]ForecastPropertiesConfidenceLevelsItem `json:"confidenceLevels,omitempty"`
}

// ForecastPropertiesConfidenceLevelsItem ...
type ForecastPropertiesConfidenceLevelsItem struct {
	// Percentage - The percentage level of the confidence
	Percentage *decimal.Decimal `json:"percentage,omitempty"`
	// Bound - The boundary of the percentage, values could be 'Upper' or 'Lower'. Possible values include: 'Upper', 'Lower'
	Bound Bound `json:"bound,omitempty"`
	// Value - The amount of forecast within the percentage level
	Value *decimal.Decimal `json:"value,omitempty"`
}

// ForecastsListResult result of listing forecasts. It contains a list of available forecasts.
type ForecastsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of forecasts.
	Value *[]Forecast `json:"value,omitempty"`
}

// MeterDetails the properties of the meter detail.
type MeterDetails struct {
	// MeterName - The name of the meter, within the given meter category
	MeterName *string `json:"meterName,omitempty"`
	// MeterCategory - The category of the meter, for example, 'Cloud services', 'Networking', etc..
	MeterCategory *string `json:"meterCategory,omitempty"`
	// MeterSubCategory - The subcategory of the meter, for example, 'A6 Cloud services', 'ExpressRoute (IXP)', etc..
	MeterSubCategory *string `json:"meterSubCategory,omitempty"`
	// Unit - The unit in which the meter consumption is charged, for example, 'Hours', 'GB', etc.
	Unit *string `json:"unit,omitempty"`
	// MeterLocation - The location in which the Azure service is available.
	MeterLocation *string `json:"meterLocation,omitempty"`
	// TotalIncludedQuantity - The total included quantity associated with the offer.
	TotalIncludedQuantity *decimal.Decimal `json:"totalIncludedQuantity,omitempty"`
	// PretaxStandardRate - The pretax listing price.
	PretaxStandardRate *decimal.Decimal `json:"pretaxStandardRate,omitempty"`
}

// Operation a Consumption REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Consumption.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: UsageDetail, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing consumption operations. It contains a list of operations and a URL link to
// get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of consumption operations supported by the Microsoft.Consumption resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// PriceSheetModel price sheet result. It contains the pricesheet associated with billing period
type PriceSheetModel struct {
	// Pricesheets - Price sheet
	Pricesheets *[]PriceSheetProperties `json:"pricesheets,omitempty"`
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// PriceSheetProperties the properties of the price sheet.
type PriceSheetProperties struct {
	// BillingPeriodID - The id of the billing period resource that the usage belongs to.
	BillingPeriodID *string `json:"billingPeriodId,omitempty"`
	// MeterID - The meter id (GUID)
	MeterID *uuid.UUID `json:"meterId,omitempty"`
	// MeterDetails - The details about the meter. By default this is not populated, unless it's specified in $expand.
	MeterDetails *MeterDetails `json:"meterDetails,omitempty"`
	// UnitOfMeasure - Unit of measure
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// IncludedQuantity - Included quality for an offer
	IncludedQuantity *decimal.Decimal `json:"includedQuantity,omitempty"`
	// PartNumber - Part Number
	PartNumber *string `json:"partNumber,omitempty"`
	// UnitPrice - Unit Price
	UnitPrice *decimal.Decimal `json:"unitPrice,omitempty"`
	// CurrencyCode - Currency Code
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// OfferID - Offer Id
	OfferID *string `json:"offerId,omitempty"`
}

// PriceSheetResult an pricesheet resource.
type PriceSheetResult struct {
	autorest.Response `json:"-"`
	*PriceSheetModel  `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PriceSheetResult.
func (psr PriceSheetResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if psr.PriceSheetModel != nil {
		objectMap["properties"] = psr.PriceSheetModel
	}
	if psr.ID != nil {
		objectMap["id"] = psr.ID
	}
	if psr.Name != nil {
		objectMap["name"] = psr.Name
	}
	if psr.Type != nil {
		objectMap["type"] = psr.Type
	}
	if psr.Tags != nil {
		objectMap["tags"] = psr.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PriceSheetResult struct.
func (psr *PriceSheetResult) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var priceSheetModel PriceSheetModel
				err = json.Unmarshal(*v, &priceSheetModel)
				if err != nil {
					return err
				}
				psr.PriceSheetModel = &priceSheetModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				psr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				psr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				psr.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				psr.Tags = tags
			}
		}
	}

	return nil
}

// Resource the Resource model definition.
type Resource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// UsageDetail an usage detail resource.
type UsageDetail struct {
	*UsageDetailProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for UsageDetail.
func (ud UsageDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ud.UsageDetailProperties != nil {
		objectMap["properties"] = ud.UsageDetailProperties
	}
	if ud.ID != nil {
		objectMap["id"] = ud.ID
	}
	if ud.Name != nil {
		objectMap["name"] = ud.Name
	}
	if ud.Type != nil {
		objectMap["type"] = ud.Type
	}
	if ud.Tags != nil {
		objectMap["tags"] = ud.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for UsageDetail struct.
func (ud *UsageDetail) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var usageDetailProperties UsageDetailProperties
				err = json.Unmarshal(*v, &usageDetailProperties)
				if err != nil {
					return err
				}
				ud.UsageDetailProperties = &usageDetailProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ud.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ud.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ud.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ud.Tags = tags
			}
		}
	}

	return nil
}

// UsageDetailProperties the properties of the usage detail.
type UsageDetailProperties struct {
	// BillingPeriodID - The id of the billing period resource that the usage belongs to.
	BillingPeriodID *string `json:"billingPeriodId,omitempty"`
	// InvoiceID - The id of the invoice resource that the usage belongs to.
	InvoiceID *string `json:"invoiceId,omitempty"`
	// UsageStart - The start of the date time range covered by the usage detail.
	UsageStart *date.Time `json:"usageStart,omitempty"`
	// UsageEnd - The end of the date time range covered by the usage detail.
	UsageEnd *date.Time `json:"usageEnd,omitempty"`
	// InstanceName - The name of the resource instance that the usage is about.
	InstanceName *string `json:"instanceName,omitempty"`
	// InstanceID - The uri of the resource instance that the usage is about.
	InstanceID *string `json:"instanceId,omitempty"`
	// InstanceLocation - The location of the resource instance that the usage is about.
	InstanceLocation *string `json:"instanceLocation,omitempty"`
	// Currency - The ISO currency in which the meter is charged, for example, USD.
	Currency *string `json:"currency,omitempty"`
	// UsageQuantity - The quantity of usage.
	UsageQuantity *decimal.Decimal `json:"usageQuantity,omitempty"`
	// BillableQuantity - The billable usage quantity.
	BillableQuantity *decimal.Decimal `json:"billableQuantity,omitempty"`
	// PretaxCost - The amount of cost before tax.
	PretaxCost *decimal.Decimal `json:"pretaxCost,omitempty"`
	// IsEstimated - The estimated usage is subject to change.
	IsEstimated *bool `json:"isEstimated,omitempty"`
	// MeterID - The meter id (GUID).
	MeterID *uuid.UUID `json:"meterId,omitempty"`
	// MeterDetails - The details about the meter. By default this is not populated, unless it's specified in $expand.
	MeterDetails *MeterDetails `json:"meterDetails,omitempty"`
	// SubscriptionGUID - Subscription guid.
	SubscriptionGUID *uuid.UUID `json:"subscriptionGuid,omitempty"`
	// SubscriptionName - Subscription name.
	SubscriptionName *string `json:"subscriptionName,omitempty"`
	// AccountName - Account name.
	AccountName *string `json:"accountName,omitempty"`
	// DepartmentName - Department name.
	DepartmentName *string `json:"departmentName,omitempty"`
	// Product - Product name.
	Product *string `json:"product,omitempty"`
	// ConsumedService - Consumed service name.
	ConsumedService *string `json:"consumedService,omitempty"`
	// CostCenter - The cost center of this department if it is a department and a costcenter exists
	CostCenter *string `json:"costCenter,omitempty"`
	// PartNumber - Part Number
	PartNumber *string `json:"partNumber,omitempty"`
	// ResourceGUID - Resource Guid
	ResourceGUID *string `json:"resourceGuid,omitempty"`
	// OfferID - Offer Id
	OfferID *string `json:"offerId,omitempty"`
	// ChargesBilledSeparately - Charges billed separately
	ChargesBilledSeparately *bool `json:"chargesBilledSeparately,omitempty"`
	// AdditionalProperties - Additional details of this usage item. By default this is not populated, unless it's specified in $expand.
	AdditionalProperties *string `json:"additionalProperties,omitempty"`
}

// UsageDetailsListResult result of listing usage details. It contains a list of available usage details in reverse
// chronological order by billing period.
type UsageDetailsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of usage details.
	Value *[]UsageDetail `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// UsageDetailsListResultIterator provides access to a complete listing of UsageDetail values.
type UsageDetailsListResultIterator struct {
	i    int
	page UsageDetailsListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UsageDetailsListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UsageDetailsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UsageDetailsListResultIterator) Response() UsageDetailsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UsageDetailsListResultIterator) Value() UsageDetail {
	if !iter.page.NotDone() {
		return UsageDetail{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (udlr UsageDetailsListResult) IsEmpty() bool {
	return udlr.Value == nil || len(*udlr.Value) == 0
}

// usageDetailsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (udlr UsageDetailsListResult) usageDetailsListResultPreparer() (*http.Request, error) {
	if udlr.NextLink == nil || len(to.String(udlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(udlr.NextLink)))
}

// UsageDetailsListResultPage contains a page of UsageDetail values.
type UsageDetailsListResultPage struct {
	fn   func(UsageDetailsListResult) (UsageDetailsListResult, error)
	udlr UsageDetailsListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UsageDetailsListResultPage) Next() error {
	next, err := page.fn(page.udlr)
	if err != nil {
		return err
	}
	page.udlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UsageDetailsListResultPage) NotDone() bool {
	return !page.udlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UsageDetailsListResultPage) Response() UsageDetailsListResult {
	return page.udlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UsageDetailsListResultPage) Values() []UsageDetail {
	if page.udlr.IsEmpty() {
		return nil
	}
	return *page.udlr.Value
}
