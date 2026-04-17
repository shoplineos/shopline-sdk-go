package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/customer"
	"github.com/stretchr/testify/assert"
)

func customerURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Customer
// ══════════════════════════════════════════════════════════════════════════════

func TestGetCustomers(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customers":[{"id":"4201825054","first_name":"Bob","last_name":"Norman","email":"bob@example.com","currency":"USD"}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, "v2/customers.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.GetCustomersAPIReq{}
	apiResp := &customer.GetCustomersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Customers, 1)
	assert.Equal(t, "4201825054", apiResp.Customers[0].Id)
	assert.Equal(t, "Bob", apiResp.Customers[0].FirstName)
}

func TestGetACustomer(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "4201825054"
	mockResp := `{"customer":{"id":"4201825054","first_name":"Bob","last_name":"Norman","email":"bob@example.com","currency":"USD","state":3}}`
	httpmock.RegisterResponder("GET", customerURL(cli, fmt.Sprintf("customers/v2/%s.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.GetACustomerAPIReq{Id: customerId}
	apiResp := &customer.GetACustomerAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "4201825054", apiResp.Customer.Id)
	assert.Equal(t, "Bob", apiResp.Customer.FirstName)
}

func TestGetACustomer_MissingId(t *testing.T) {
	err := (&customer.GetACustomerAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCreateACustomer(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customer":{"id":"4201825054","first_name":"Bob","last_name":"Norman","email":"bob@example.com","state":2,"currency":"USD"}}`
	httpmock.RegisterResponder("POST", customerURL(cli, "customers.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.CreateACustomerAPIReq{
		Customer: customer.CreateACustomerAPIReqCustomer{
			FirstName: "Bob",
			LastName:  "Norman",
			Email:     "bob@example.com",
		},
	}
	apiResp := &customer.CreateACustomerAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "4201825054", apiResp.Customer.Id)
	assert.Equal(t, "Bob", apiResp.Customer.FirstName)
}

func TestUpdateACustomer(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "4201825054"
	mockResp := `{"customer":{"id":"4201825054","first_name":"Bob","last_name":"Smith","email":"bob@example.com"}}`
	httpmock.RegisterResponder("PUT", customerURL(cli, fmt.Sprintf("customers/%s.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.UpdateACustomerAPIReq{
		Id: customerId,
		Customer: customer.UpdateACustomerAPIReqCustomer{
			LastName: "Smith",
		},
	}
	apiResp := &customer.UpdateACustomerAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "4201825054", apiResp.Customer.Id)
	assert.Equal(t, "Smith", apiResp.Customer.LastName)
}

func TestUpdateACustomer_MissingId(t *testing.T) {
	err := (&customer.UpdateACustomerAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteCustomerInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "4211465524"
	httpmock.RegisterResponder("DELETE", customerURL(cli, fmt.Sprintf("customers/%s.json", customerId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.DeleteCustomerInformationAPIReq{Id: customerId}
	apiResp := &customer.DeleteCustomerInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteCustomerInformation_MissingId(t *testing.T) {
	err := (&customer.DeleteCustomerInformationAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSearchForCustomers(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customers":[{"id":"4201825054","first_name":"Bob","last_name":"Norman","email":"bob@example.com"}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, "customers/v2/search.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.SearchForCustomersAPIReq{QueryParam: "bob"}
	apiResp := &customer.SearchForCustomersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Customers, 1)
	assert.Equal(t, "Bob", apiResp.Customers[0].FirstName)
}

func TestCustomerNumberQuery(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", customerURL(cli, "customers/v2/count.json"),
		httpmock.NewStringResponder(200, `{"count":100}`))

	apiReq := &customer.CustomerNumberQueryAPIReq{}
	apiResp := &customer.CustomerNumberQueryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 100, apiResp.Count)
}

func TestActivateTheCustomer(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "4201825054"
	mockResp := `{"data":{"account_activation_url":"https://example.myshopline.com/account/activate"}}`
	httpmock.RegisterResponder("POST", customerURL(cli, fmt.Sprintf("customers/%s/account_activation_url.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.ActivateTheCustomerAPIReq{Id: customerId}
	apiResp := &customer.ActivateTheCustomerAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "https://example.myshopline.com/account/activate", apiResp.Data.AccountActivationUrl)
}

func TestActivateTheCustomer_MissingId(t *testing.T) {
	err := (&customer.ActivateTheCustomerAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSendActivationMail(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "4201801159"
	mockResp := `{"customer_invite":{"from":"noreply@example.com","to":"customer@example.com"}}`
	httpmock.RegisterResponder("POST", customerURL(cli, fmt.Sprintf("customers/%s/send_invite.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.SendActivationMailAPIReq{Id: customerId}
	apiResp := &customer.SendActivationMailAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "noreply@example.com", apiResp.CustomerInvite.From)
	assert.Equal(t, "customer@example.com", apiResp.CustomerInvite.To)
}

func TestSendActivationMail_MissingId(t *testing.T) {
	err := (&customer.SendActivationMailAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCheckCustomerInformationViaEmail(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customer":{"id":"2277270272","first_name":"Bob","last_name":"Norman","email":"bob@example.com","state":3}}`
	httpmock.RegisterResponder("POST", customerURL(cli, "customers/query_user_by_email.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.CheckCustomerInformationViaEmailAPIReq{Email: "bob@example.com"}
	apiResp := &customer.CheckCustomerInformationViaEmailAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "2277270272", apiResp.Customer.Id)
	assert.Equal(t, "Bob", apiResp.Customer.FirstName)
}

func TestCheckCustomerInformationViaEmail_MissingEmail(t *testing.T) {
	err := (&customer.CheckCustomerInformationViaEmailAPIReq{}).Verify()
	assert.EqualError(t, err, "Email is required")
}

func TestQuerySpecifyCustomerOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "2002697754"
	mockResp := `{"list":[{"id":"21050224312121887324667162","financial_status":"paid","email":"hok@shoplineapp.com"}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, fmt.Sprintf("customers/%s/orders.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.QuerySpecifyCustomerOrderAPIReq{
		Id:     customerId,
		Status: "any",
	}
	apiResp := &customer.QuerySpecifyCustomerOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.List, 1)
	assert.Equal(t, "21050224312121887324667162", apiResp.List[0].Id)
	assert.Equal(t, "paid", apiResp.List[0].FinancialStatus)
}

func TestQuerySpecifyCustomerOrder_MissingId(t *testing.T) {
	err := (&customer.QuerySpecifyCustomerOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQuerySpecifyCustomerOrder_MissingStatus(t *testing.T) {
	err := (&customer.QuerySpecifyCustomerOrderAPIReq{Id: "2002697754"}).Verify()
	assert.EqualError(t, err, "Status is required")
}

func TestBatchQueryUserSubscriptionInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"user_subscribe_info_list":[{"customer_id":"4201799145","status":1,"subscribe_account":"user@example.com","subscribe_account_type":0}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, "customers/subscribe.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.BatchQueryUserSubscriptionInformationAPIReq{CustomerIds: "4201799145"}
	apiResp := &customer.BatchQueryUserSubscriptionInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.UserSubscribeInfoList, 1)
	assert.Equal(t, "4201799145", apiResp.UserSubscribeInfoList[0].CustomerId)
	assert.Equal(t, 1, apiResp.UserSubscribeInfoList[0].Status)
}

func TestAddACustomerToTheBlacklist(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", customerURL(cli, "customers/add_blacklist.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.AddACustomerToTheBlacklistAPIReq{CustomerId: "46032231994"}
	apiResp := &customer.AddACustomerToTheBlacklistAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddACustomerToTheBlacklist_MissingCustomerId(t *testing.T) {
	err := (&customer.AddACustomerToTheBlacklistAPIReq{}).Verify()
	assert.EqualError(t, err, "CustomerId is required")
}

func TestRemoveACustomerFromTheBlacklist(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", customerURL(cli, "customers/remove_blacklist.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.RemoveACustomerFromTheBlacklistAPIReq{CustomerId: "46032231994"}
	apiResp := &customer.RemoveACustomerFromTheBlacklistAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteSpecifyCustomerTag(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", customerURL(cli, "customer_tags_remove.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.DeleteSpecifyCustomerTagAPIReq{
		UserId: "2277270272",
		TagIds: []string{"SL201UT5491137867706559244"},
	}
	apiResp := &customer.DeleteSpecifyCustomerTagAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteSpecifyCustomerTag_MissingUserId(t *testing.T) {
	err := (&customer.DeleteSpecifyCustomerTagAPIReq{}).Verify()
	assert.EqualError(t, err, "UserId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Customer Address
// ══════════════════════════════════════════════════════════════════════════════

func TestAddAddress20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "2277270272"
	mockResp := `{"customer_address":{"id":"SL201UA5006511321220969539","first_name":"Bob","last_name":"Norman","address1":"7720 Cherokee Road","city":"Hagerman","country":"United States","country_code":"US","customer_id":"2277270272"}}`
	httpmock.RegisterResponder("POST", customerURL(cli, fmt.Sprintf("customers/%s/addresses.json", customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.AddAddress20APIReq{
		Id: customerId,
		CustomerAddress: customer.AddAddress20APIReqCustomerAddress{
			FirstName: "Bob",
			LastName:  "Norman",
			Address1:  "7720 Cherokee Road",
			City:      "Hagerman",
			Country:   "United States",
		},
	}
	apiResp := &customer.AddAddress20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "SL201UA5006511321220969539", apiResp.CustomerAddress.Id)
	assert.Equal(t, "Bob", apiResp.CustomerAddress.FirstName)
}

func TestAddAddress20_MissingId(t *testing.T) {
	err := (&customer.AddAddress20APIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestUpdateAddress20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	addressId := "SL201UA5006511321220969539"
	customerId := "2277270272"
	mockResp := `{"customer_address":{"id":"SL201UA5006511321220969539","first_name":"maybe","last_name":"joho","address1":"7720 Cherokee Road","city":"Hagerman","country":"United States","country_code":"US","customer_id":"2277270272"}}`
	httpmock.RegisterResponder("PUT", customerURL(cli, fmt.Sprintf("customers/%s/addresses/%s.json", addressId, customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.UpdateAddress20APIReq{
		AddressId: addressId,
		Id:        customerId,
		CustomerAddress: customer.UpdateAddress20APIReqCustomerAddress{
			FirstName: "maybe",
			LastName:  "joho",
		},
	}
	apiResp := &customer.UpdateAddress20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "SL201UA5006511321220969539", apiResp.CustomerAddress.Id)
	assert.Equal(t, "maybe", apiResp.CustomerAddress.FirstName)
}

func TestUpdateAddress20_MissingAddressId(t *testing.T) {
	err := (&customer.UpdateAddress20APIReq{}).Verify()
	assert.EqualError(t, err, "AddressId is required")
}

func TestUpdateAddress20_MissingId(t *testing.T) {
	err := (&customer.UpdateAddress20APIReq{AddressId: "SL201UA5006511321220969539"}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteCustomerAddress20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	addressId := "SL201UA6021464182578751091"
	customerId := "2277270272"
	httpmock.RegisterResponder("DELETE", customerURL(cli, fmt.Sprintf("customers/%s/addresses/%s.json", addressId, customerId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.DeleteCustomerAddress20APIReq{
		AddressId: addressId,
		Id:        customerId,
	}
	apiResp := &customer.DeleteCustomerAddress20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteCustomerAddress20_MissingAddressId(t *testing.T) {
	err := (&customer.DeleteCustomerAddress20APIReq{}).Verify()
	assert.EqualError(t, err, "AddressId is required")
}

func TestDeleteCustomerAddress20_MissingId(t *testing.T) {
	err := (&customer.DeleteCustomerAddress20APIReq{AddressId: "SL201UA6021464182578751091"}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryCustomerAddressDetails20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	addressId := "SL201UA5006511321220969539"
	customerId := "2277270272"
	mockResp := `{"customer_address":{"id":"SL201UA5006511321220969539","first_name":"Bob","last_name":"Norman","address1":"7720 Cherokee Road","city":"Hagerman","country":"United States","country_code":"US","customer_id":"2277270272"}}`
	httpmock.RegisterResponder("GET", customerURL(cli, fmt.Sprintf("customers/%s/addresses/%s.json", addressId, customerId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.QueryCustomerAddressDetails20APIReq{
		AddressId: addressId,
		Id:        customerId,
	}
	apiResp := &customer.QueryCustomerAddressDetails20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "SL201UA5006511321220969539", apiResp.CustomerAddress.Id)
	assert.Equal(t, "Bob", apiResp.CustomerAddress.FirstName)
}

func TestQueryCustomerAddressDetails20_MissingAddressId(t *testing.T) {
	err := (&customer.QueryCustomerAddressDetails20APIReq{}).Verify()
	assert.EqualError(t, err, "AddressId is required")
}

func TestQueryCustomerAddressDetails20_MissingId(t *testing.T) {
	err := (&customer.QueryCustomerAddressDetails20APIReq{AddressId: "SL201UA5006511321220969539"}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSetDefaultAddress20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	addressId := "SL201UA6021464182578751091"
	customerId := "2277270272"
	httpmock.RegisterResponder("PUT", customerURL(cli, fmt.Sprintf("customers/%s/addresses/%s/default.json", addressId, customerId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.SetDefaultAddress20APIReq{
		AddressId: addressId,
		Id:        customerId,
	}
	apiResp := &customer.SetDefaultAddress20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSetDefaultAddress20_MissingAddressId(t *testing.T) {
	err := (&customer.SetDefaultAddress20APIReq{}).Verify()
	assert.EqualError(t, err, "AddressId is required")
}

func TestSetDefaultAddress20_MissingId(t *testing.T) {
	err := (&customer.SetDefaultAddress20APIReq{AddressId: "SL201UA6021464182578751091"}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestBatchAddress(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	customerId := "3300004100"
	httpmock.RegisterResponder("PUT", customerURL(cli, fmt.Sprintf("customers/%s/addresses/set.json", customerId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.BatchAddressAPIReq{
		Id:         customerId,
		AddressIds: "SL201UA4581724836447062019",
		Operation:  "destroy",
	}
	apiResp := &customer.BatchAddressAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestBatchAddress_MissingId(t *testing.T) {
	err := (&customer.BatchAddressAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestBatchAddress_MissingAddressIds(t *testing.T) {
	err := (&customer.BatchAddressAPIReq{Id: "3300004100"}).Verify()
	assert.EqualError(t, err, "AddressIds is required")
}

func TestBatchAddress_MissingOperation(t *testing.T) {
	err := (&customer.BatchAddressAPIReq{Id: "3300004100", AddressIds: "SL201UA4581724836447062019"}).Verify()
	assert.EqualError(t, err, "Operation is required")
}

func TestBatchQueryCustomerAddress20(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"user_address":[{"uid":"2277270272","addresses":[{"id":"SL201UA5006511321220969539","address1":"7720 Cherokee Road","city":"Hagerman","country":"United States","country_code":"US"}]}]}`
	httpmock.RegisterResponder("POST", customerURL(cli, "customers/addresses/list.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.BatchQueryCustomerAddress20APIReq{
		Uids: []string{"2277270272"},
	}
	apiResp := &customer.BatchQueryCustomerAddress20APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.UserAddress, 1)
	assert.Equal(t, "2277270272", apiResp.UserAddress[0].Uid)
	assert.Len(t, apiResp.UserAddress[0].Addresses, 1)
}

// ══════════════════════════════════════════════════════════════════════════════
// Customer Grouping
// ══════════════════════════════════════════════════════════════════════════════

func TestGetCustomerGroupingInBulk(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customer_saved_searches":[{"id":"GROUP5483719394486658833","name":"Returning","query":"orders_count:>1","created_at":"2022-07-07T08:40:55+00:00","updated_at":"2022-07-07T08:40:55+00:00"}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, "customer_saved_searches.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.GetCustomerGroupingInBulkAPIReq{}
	apiResp := &customer.GetCustomerGroupingInBulkAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.CustomerSavedSearches, 1)
	assert.Equal(t, "GROUP5483719394486658833", apiResp.CustomerSavedSearches[0].Id)
	assert.Equal(t, "Returning", apiResp.CustomerSavedSearches[0].Name)
}

func TestSpecifyCustomerGroupingQuery(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	groupId := "GROUP5429808178858062838"
	mockResp := `{"customer_saved_search":{"id":"GROUP5429808178858062838","name":"Returning","query":"orders_count:>1","created_at":"2022-07-07T08:40:55+00:00","updated_at":"2022-07-07T08:40:55+00:00"}}`
	httpmock.RegisterResponder("GET", customerURL(cli, fmt.Sprintf("customer_saved_searches/%s.json", groupId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.SpecifyCustomerGroupingQueryAPIReq{Id: groupId}
	apiResp := &customer.SpecifyCustomerGroupingQueryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "GROUP5429808178858062838", apiResp.CustomerSavedSearch.Id)
	assert.Equal(t, "Returning", apiResp.CustomerSavedSearch.Name)
}

func TestSpecifyCustomerGroupingQuery_MissingId(t *testing.T) {
	err := (&customer.SpecifyCustomerGroupingQueryAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCreateCustomerGrouping(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"customer_saved_search":{"id":"GROUP5483719394486658833","name":"orders count greater than one","query":"orders_count:>1","created_at":"2022-07-07T08:40:55+00:00","updated_at":"2022-07-07T08:40:55+00:00"}}`
	httpmock.RegisterResponder("POST", customerURL(cli, "customer_saved_searches.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.CreateCustomerGroupingAPIReq{
		CustomerSavedSearch: customer.CreateCustomerGroupingAPIReqCustomerSavedSearch{
			Name:  "orders count greater than one",
			Query: "orders_count:>1",
		},
	}
	apiResp := &customer.CreateCustomerGroupingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "GROUP5483719394486658833", apiResp.CustomerSavedSearch.Id)
	assert.Equal(t, "orders count greater than one", apiResp.CustomerSavedSearch.Name)
}

func TestUpdateCustomerGrouping(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	groupId := "GROUP5573201978623266205"
	mockResp := `{"customer_saved_search":{"id":"GROUP5573201978623266205","name":"modify group name","query":"orders_count:<1","created_at":"2022-07-07T08:40:55+00:00","updated_at":"2022-07-07T08:40:55+00:00"}}`
	httpmock.RegisterResponder("PUT", customerURL(cli, fmt.Sprintf("customer_saved_searches/%s.json", groupId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.UpdateCustomerGroupingAPIReq{
		Id: groupId,
		CustomerSavedSearch: customer.UpdateCustomerGroupingAPIReqCustomerSavedSearch{
			Name: "modify group name",
		},
	}
	apiResp := &customer.UpdateCustomerGroupingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "GROUP5573201978623266205", apiResp.CustomerSavedSearch.Id)
	assert.Equal(t, "modify group name", apiResp.CustomerSavedSearch.Name)
}

func TestUpdateCustomerGrouping_MissingId(t *testing.T) {
	err := (&customer.UpdateCustomerGroupingAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteCustomerGrouping(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	groupId := "GROUP5702593537218454508"
	httpmock.RegisterResponder("DELETE", customerURL(cli, fmt.Sprintf("customer_saved_searches/%s.json", groupId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.DeleteCustomerGroupingAPIReq{Id: groupId}
	apiResp := &customer.DeleteCustomerGroupingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteCustomerGrouping_MissingId(t *testing.T) {
	err := (&customer.DeleteCustomerGroupingAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryStoreGrouping(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", customerURL(cli, "customer_saved_searches/count.json"),
		httpmock.NewStringResponder(200, `{"count":5}`))

	apiReq := &customer.QueryStoreGroupingAPIReq{}
	apiResp := &customer.QueryStoreGroupingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 5, apiResp.Count)
}

func TestQueryAllCustomersInGroups(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	groupId := "GROUP5483719394486658833"
	mockResp := `{"customers":[{"id":"4201825054","first_name":"Bob","last_name":"Norman","email":"zhangsan@gmail.com","currency":"USD","state":2}]}`
	httpmock.RegisterResponder("GET", customerURL(cli, fmt.Sprintf("customer_saved_searches/%s/customers.json", groupId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.QueryAllCustomersInGroupsAPIReq{Id: groupId}
	apiResp := &customer.QueryAllCustomersInGroupsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Customers, 1)
	assert.Equal(t, "4201825054", apiResp.Customers[0].Id)
	assert.Equal(t, "Bob", apiResp.Customers[0].FirstName)
}

func TestQueryAllCustomersInGroups_MissingId(t *testing.T) {
	err := (&customer.QueryAllCustomersInGroupsAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Third Party Login
// ══════════════════════════════════════════════════════════════════════════════

func TestRetrieveAListOfThirdPartyLoginConfigurationsOfStore(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"third_party_login_configurations":[{"login_channel":"google","login_platform_type":"web","third_app_id":"demo app id","third_key":"demo key"}]}}`
	httpmock.RegisterResponder("GET", customerURL(cli, "customers_social_login.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &customer.RetrieveAListOfThirdPartyLoginConfigurationsOfStoreAPIReq{LoginChannel: "google"}
	apiResp := &customer.RetrieveAListOfThirdPartyLoginConfigurationsOfStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data.ThirdPartyLoginConfigurations, 1)
	assert.Equal(t, "google", apiResp.Data.ThirdPartyLoginConfigurations[0].LoginChannel)
	assert.Equal(t, "web", apiResp.Data.ThirdPartyLoginConfigurations[0].LoginPlatformType)
}

func TestUpdateTheThirdPartyLoginConfigurationOfStore(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", customerURL(cli, "customers_social_login.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.UpdateTheThirdPartyLoginConfigurationOfStoreAPIReq{
		LoginChannel:      "google",
		LoginPlatformType: "app",
		ThirdAppId:        "demo third app id",
		ThirdKey:          "demo third key",
	}
	apiResp := &customer.UpdateTheThirdPartyLoginConfigurationOfStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateTheThirdPartyLoginConfigurationOfStore_MissingLoginChannel(t *testing.T) {
	err := (&customer.UpdateTheThirdPartyLoginConfigurationOfStoreAPIReq{}).Verify()
	assert.EqualError(t, err, "LoginChannel is required")
}

func TestDeleteTheThirdPartyLoginConfigurationOfStore(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE", customerURL(cli, "customers_social_login.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &customer.DeleteTheThirdPartyLoginConfigurationOfStoreAPIReq{
		LoginChannel:      "google",
		LoginPlatformType: "app",
	}
	apiResp := &customer.DeleteTheThirdPartyLoginConfigurationOfStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteTheThirdPartyLoginConfigurationOfStore_MissingLoginChannel(t *testing.T) {
	err := (&customer.DeleteTheThirdPartyLoginConfigurationOfStoreAPIReq{}).Verify()
	assert.EqualError(t, err, "LoginChannel is required")
}

func TestDeleteTheThirdPartyLoginConfigurationOfStore_MissingLoginPlatformType(t *testing.T) {
	err := (&customer.DeleteTheThirdPartyLoginConfigurationOfStoreAPIReq{LoginChannel: "google"}).Verify()
	assert.EqualError(t, err, "LoginPlatformType is required")
}
