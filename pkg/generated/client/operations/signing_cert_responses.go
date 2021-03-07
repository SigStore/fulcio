// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Fulcio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/fulcio/pkg/generated/models"
)

// SigningCertReader is a Reader for the SigningCert structure.
type SigningCertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SigningCertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSigningCertCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSigningCertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSigningCertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSigningCertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSigningCertCreated creates a SigningCertCreated with default headers values
func NewSigningCertCreated() *SigningCertCreated {
	return &SigningCertCreated{}
}

/*SigningCertCreated handles this case with default header values.

Successful CSR Submit
*/
type SigningCertCreated struct {
	Payload *models.SubmitSuccess
}

func (o *SigningCertCreated) Error() string {
	return fmt.Sprintf("[POST /signingCert][%d] signingCertCreated  %+v", 201, o.Payload)
}

func (o *SigningCertCreated) GetPayload() *models.SubmitSuccess {
	return o.Payload
}

func (o *SigningCertCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubmitSuccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSigningCertBadRequest creates a SigningCertBadRequest with default headers values
func NewSigningCertBadRequest() *SigningCertBadRequest {
	return &SigningCertBadRequest{}
}

/*SigningCertBadRequest handles this case with default header values.

Bad Request
*/
type SigningCertBadRequest struct {
}

func (o *SigningCertBadRequest) Error() string {
	return fmt.Sprintf("[POST /signingCert][%d] signingCertBadRequest ", 400)
}

func (o *SigningCertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSigningCertUnauthorized creates a SigningCertUnauthorized with default headers values
func NewSigningCertUnauthorized() *SigningCertUnauthorized {
	return &SigningCertUnauthorized{}
}

/*SigningCertUnauthorized handles this case with default header values.

Unauthorized
*/
type SigningCertUnauthorized struct {
	Payload string
}

func (o *SigningCertUnauthorized) Error() string {
	return fmt.Sprintf("[POST /signingCert][%d] signingCertUnauthorized  %+v", 401, o.Payload)
}

func (o *SigningCertUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *SigningCertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSigningCertInternalServerError creates a SigningCertInternalServerError with default headers values
func NewSigningCertInternalServerError() *SigningCertInternalServerError {
	return &SigningCertInternalServerError{}
}

/*SigningCertInternalServerError handles this case with default header values.

Server error
*/
type SigningCertInternalServerError struct {
	Payload string
}

func (o *SigningCertInternalServerError) Error() string {
	return fmt.Sprintf("[POST /signingCert][%d] signingCertInternalServerError  %+v", 500, o.Payload)
}

func (o *SigningCertInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *SigningCertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}