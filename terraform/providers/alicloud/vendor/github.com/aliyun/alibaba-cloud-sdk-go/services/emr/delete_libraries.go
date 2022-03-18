package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteLibraries invokes the emr.DeleteLibraries API synchronously
func (client *Client) DeleteLibraries(request *DeleteLibrariesRequest) (response *DeleteLibrariesResponse, err error) {
	response = CreateDeleteLibrariesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLibrariesWithChan invokes the emr.DeleteLibraries API asynchronously
func (client *Client) DeleteLibrariesWithChan(request *DeleteLibrariesRequest) (<-chan *DeleteLibrariesResponse, <-chan error) {
	responseChan := make(chan *DeleteLibrariesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLibraries(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteLibrariesWithCallback invokes the emr.DeleteLibraries API asynchronously
func (client *Client) DeleteLibrariesWithCallback(request *DeleteLibrariesRequest, callback func(response *DeleteLibrariesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLibrariesResponse
		var err error
		defer close(result)
		response, err = client.DeleteLibraries(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteLibrariesRequest is the request struct for api DeleteLibraries
type DeleteLibrariesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LibraryBizIdList *[]string        `position:"Query" name:"LibraryBizIdList"  type:"Repeated"`
}

// DeleteLibrariesResponse is the response struct for api DeleteLibraries
type DeleteLibrariesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteLibrariesRequest creates a request to invoke DeleteLibraries API
func CreateDeleteLibrariesRequest() (request *DeleteLibrariesRequest) {
	request = &DeleteLibrariesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteLibraries", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLibrariesResponse creates a response to parse from DeleteLibraries response
func CreateDeleteLibrariesResponse() (response *DeleteLibrariesResponse) {
	response = &DeleteLibrariesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}