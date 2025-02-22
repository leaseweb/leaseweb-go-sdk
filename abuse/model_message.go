/*
Abuse Reports

This API provides ways to manage the abuse reports you might receive from Leaseweb. To use this API, please request access via your account manager and/or compliance officer. **LIMITED ACCESS** 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package abuse

import (
	"encoding/json"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message struct for Message
type Message struct {
	// A string to indicating if the message was posted by the customer or the abuse agent.
	PostedBy *string `json:"postedBy,omitempty"`
	// The timestamp when the message was posted.
	PostedAt *string `json:"postedAt,omitempty"`
	// The posted message.
	Body *string `json:"body,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetPostedBy returns the PostedBy field value if set, zero value otherwise.
func (o *Message) GetPostedBy() string {
	if o == nil || IsNil(o.PostedBy) {
		var ret string
		return ret
	}
	return *o.PostedBy
}

// GetPostedByOk returns a tuple with the PostedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPostedByOk() (*string, bool) {
	if o == nil || IsNil(o.PostedBy) {
		return nil, false
	}
	return o.PostedBy, true
}

// HasPostedBy returns a boolean if a field has been set.
func (o *Message) HasPostedBy() bool {
	if o != nil && !IsNil(o.PostedBy) {
		return true
	}

	return false
}

// SetPostedBy gets a reference to the given string and assigns it to the PostedBy field.
func (o *Message) SetPostedBy(v string) {
	o.PostedBy = &v
}

// GetPostedAt returns the PostedAt field value if set, zero value otherwise.
func (o *Message) GetPostedAt() string {
	if o == nil || IsNil(o.PostedAt) {
		var ret string
		return ret
	}
	return *o.PostedAt
}

// GetPostedAtOk returns a tuple with the PostedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetPostedAtOk() (*string, bool) {
	if o == nil || IsNil(o.PostedAt) {
		return nil, false
	}
	return o.PostedAt, true
}

// HasPostedAt returns a boolean if a field has been set.
func (o *Message) HasPostedAt() bool {
	if o != nil && !IsNil(o.PostedAt) {
		return true
	}

	return false
}

// SetPostedAt gets a reference to the given string and assigns it to the PostedAt field.
func (o *Message) SetPostedAt(v string) {
	o.PostedAt = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Message) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Message) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *Message) SetBody(v string) {
	o.Body = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *Message) GetAttachment() Attachment {
	if o == nil || IsNil(o.Attachment) {
		var ret Attachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *Message) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given Attachment and assigns it to the Attachment field.
func (o *Message) SetAttachment(v Attachment) {
	o.Attachment = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostedBy) {
		toSerialize["postedBy"] = o.PostedBy
	}
	if !IsNil(o.PostedAt) {
		toSerialize["postedAt"] = o.PostedAt
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	varMessage := _Message{}

	err = json.Unmarshal(data, &varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "postedBy")
		delete(additionalProperties, "postedAt")
		delete(additionalProperties, "body")
		delete(additionalProperties, "attachment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


