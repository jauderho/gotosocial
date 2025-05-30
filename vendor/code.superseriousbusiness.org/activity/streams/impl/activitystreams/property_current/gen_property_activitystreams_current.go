// Code generated by astool. DO NOT EDIT.

package propertycurrent

import (
	vocab "code.superseriousbusiness.org/activity/streams/vocab"
	"fmt"
	"net/url"
)

// ActivityStreamsCurrentProperty is the functional property "current". It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsCurrentProperty struct {
	activitystreamsCollectionPageMember        vocab.ActivityStreamsCollectionPage
	activitystreamsLinkMember                  vocab.ActivityStreamsLink
	tootHashtagMember                          vocab.TootHashtag
	activitystreamsMentionMember               vocab.ActivityStreamsMention
	activitystreamsOrderedCollectionPageMember vocab.ActivityStreamsOrderedCollectionPage
	unknown                                    interface{}
	iri                                        *url.URL
	alias                                      string
}

// DeserializeCurrentProperty creates a "current" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeCurrentProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsCurrentProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "current"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "current")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsCurrentProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsCurrentProperty{
					activitystreamsCollectionPageMember: v,
					alias:                               alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsCurrentProperty{
					activitystreamsLinkMember: v,
					alias:                     alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeHashtagToot()(m, aliasMap); err == nil {
				this := &ActivityStreamsCurrentProperty{
					alias:             alias,
					tootHashtagMember: v,
				}
				return this, nil
			} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsCurrentProperty{
					activitystreamsMentionMember: v,
					alias:                        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsCurrentProperty{
					activitystreamsOrderedCollectionPageMember: v,
					alias: alias,
				}
				return this, nil
			}
		}
		this := &ActivityStreamsCurrentProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsCurrentProperty creates a new current property.
func NewActivityStreamsCurrentProperty() *ActivityStreamsCurrentProperty {
	return &ActivityStreamsCurrentProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsCurrentProperty) Clear() {
	this.activitystreamsCollectionPageMember = nil
	this.activitystreamsLinkMember = nil
	this.tootHashtagMember = nil
	this.activitystreamsMentionMember = nil
	this.activitystreamsOrderedCollectionPageMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetActivityStreamsCollectionPage returns the value of this property. When
// IsActivityStreamsCollectionPage returns false,
// GetActivityStreamsCollectionPage will return an arbitrary value.
func (this ActivityStreamsCurrentProperty) GetActivityStreamsCollectionPage() vocab.ActivityStreamsCollectionPage {
	return this.activitystreamsCollectionPageMember
}

// GetActivityStreamsLink returns the value of this property. When
// IsActivityStreamsLink returns false, GetActivityStreamsLink will return an
// arbitrary value.
func (this ActivityStreamsCurrentProperty) GetActivityStreamsLink() vocab.ActivityStreamsLink {
	return this.activitystreamsLinkMember
}

// GetActivityStreamsMention returns the value of this property. When
// IsActivityStreamsMention returns false, GetActivityStreamsMention will
// return an arbitrary value.
func (this ActivityStreamsCurrentProperty) GetActivityStreamsMention() vocab.ActivityStreamsMention {
	return this.activitystreamsMentionMember
}

// GetActivityStreamsOrderedCollectionPage returns the value of this property.
// When IsActivityStreamsOrderedCollectionPage returns false,
// GetActivityStreamsOrderedCollectionPage will return an arbitrary value.
func (this ActivityStreamsCurrentProperty) GetActivityStreamsOrderedCollectionPage() vocab.ActivityStreamsOrderedCollectionPage {
	return this.activitystreamsOrderedCollectionPageMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsCurrentProperty) GetIRI() *url.URL {
	return this.iri
}

// GetTootHashtag returns the value of this property. When IsTootHashtag returns
// false, GetTootHashtag will return an arbitrary value.
func (this ActivityStreamsCurrentProperty) GetTootHashtag() vocab.TootHashtag {
	return this.tootHashtagMember
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this ActivityStreamsCurrentProperty) GetType() vocab.Type {
	if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage()
	}
	if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink()
	}
	if this.IsTootHashtag() {
		return this.GetTootHashtag()
	}
	if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention()
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage()
	}

	return nil
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsCurrentProperty) HasAny() bool {
	return this.IsActivityStreamsCollectionPage() ||
		this.IsActivityStreamsLink() ||
		this.IsTootHashtag() ||
		this.IsActivityStreamsMention() ||
		this.IsActivityStreamsOrderedCollectionPage() ||
		this.iri != nil
}

// IsActivityStreamsCollectionPage returns true if this property has a type of
// "CollectionPage". When true, use the GetActivityStreamsCollectionPage and
// SetActivityStreamsCollectionPage methods to access and set this property.
func (this ActivityStreamsCurrentProperty) IsActivityStreamsCollectionPage() bool {
	return this.activitystreamsCollectionPageMember != nil
}

// IsActivityStreamsLink returns true if this property has a type of "Link". When
// true, use the GetActivityStreamsLink and SetActivityStreamsLink methods to
// access and set this property.
func (this ActivityStreamsCurrentProperty) IsActivityStreamsLink() bool {
	return this.activitystreamsLinkMember != nil
}

// IsActivityStreamsMention returns true if this property has a type of "Mention".
// When true, use the GetActivityStreamsMention and SetActivityStreamsMention
// methods to access and set this property.
func (this ActivityStreamsCurrentProperty) IsActivityStreamsMention() bool {
	return this.activitystreamsMentionMember != nil
}

// IsActivityStreamsOrderedCollectionPage returns true if this property has a type
// of "OrderedCollectionPage". When true, use the
// GetActivityStreamsOrderedCollectionPage and
// SetActivityStreamsOrderedCollectionPage methods to access and set this
// property.
func (this ActivityStreamsCurrentProperty) IsActivityStreamsOrderedCollectionPage() bool {
	return this.activitystreamsOrderedCollectionPageMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsCurrentProperty) IsIRI() bool {
	return this.iri != nil
}

// IsTootHashtag returns true if this property has a type of "Hashtag". When true,
// use the GetTootHashtag and SetTootHashtag methods to access and set this
// property.
func (this ActivityStreamsCurrentProperty) IsTootHashtag() bool {
	return this.tootHashtagMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsCurrentProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string
	if this.IsActivityStreamsCollectionPage() {
		child = this.GetActivityStreamsCollectionPage().JSONLDContext()
	} else if this.IsActivityStreamsLink() {
		child = this.GetActivityStreamsLink().JSONLDContext()
	} else if this.IsTootHashtag() {
		child = this.GetTootHashtag().JSONLDContext()
	} else if this.IsActivityStreamsMention() {
		child = this.GetActivityStreamsMention().JSONLDContext()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		child = this.GetActivityStreamsOrderedCollectionPage().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsCurrentProperty) KindIndex() int {
	if this.IsActivityStreamsCollectionPage() {
		return 0
	}
	if this.IsActivityStreamsLink() {
		return 1
	}
	if this.IsTootHashtag() {
		return 2
	}
	if this.IsActivityStreamsMention() {
		return 3
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return 4
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsCurrentProperty) LessThan(o vocab.ActivityStreamsCurrentProperty) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().LessThan(o.GetActivityStreamsCollectionPage())
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().LessThan(o.GetActivityStreamsLink())
	} else if this.IsTootHashtag() {
		return this.GetTootHashtag().LessThan(o.GetTootHashtag())
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().LessThan(o.GetActivityStreamsMention())
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().LessThan(o.GetActivityStreamsOrderedCollectionPage())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "current".
func (this ActivityStreamsCurrentProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "current"
	} else {
		return "current"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsCurrentProperty) Serialize() (interface{}, error) {
	if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().Serialize()
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().Serialize()
	} else if this.IsTootHashtag() {
		return this.GetTootHashtag().Serialize()
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().Serialize()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetActivityStreamsCollectionPage sets the value of this property. Calling
// IsActivityStreamsCollectionPage afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetActivityStreamsCollectionPage(v vocab.ActivityStreamsCollectionPage) {
	this.Clear()
	this.activitystreamsCollectionPageMember = v
}

// SetActivityStreamsLink sets the value of this property. Calling
// IsActivityStreamsLink afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.Clear()
	this.activitystreamsLinkMember = v
}

// SetActivityStreamsMention sets the value of this property. Calling
// IsActivityStreamsMention afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.Clear()
	this.activitystreamsMentionMember = v
}

// SetActivityStreamsOrderedCollectionPage sets the value of this property.
// Calling IsActivityStreamsOrderedCollectionPage afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetActivityStreamsOrderedCollectionPage(v vocab.ActivityStreamsOrderedCollectionPage) {
	this.Clear()
	this.activitystreamsOrderedCollectionPageMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetTootHashtag sets the value of this property. Calling IsTootHashtag
// afterwards returns true.
func (this *ActivityStreamsCurrentProperty) SetTootHashtag(v vocab.TootHashtag) {
	this.Clear()
	this.tootHashtagMember = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *ActivityStreamsCurrentProperty) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.ActivityStreamsCollectionPage); ok {
		this.SetActivityStreamsCollectionPage(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsLink); ok {
		this.SetActivityStreamsLink(v)
		return nil
	}
	if v, ok := t.(vocab.TootHashtag); ok {
		this.SetTootHashtag(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsMention); ok {
		this.SetActivityStreamsMention(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsOrderedCollectionPage); ok {
		this.SetActivityStreamsOrderedCollectionPage(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on current property: %T", t)
}
