// Code generated by astool. DO NOT EDIT.

package streams

import (
	typepublickey "code.superseriousbusiness.org/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "code.superseriousbusiness.org/activity/streams/vocab"
)

// IsOrExtendsW3IDSecurityV1PublicKey returns true if the other provided type is
// the PublicKey type or extends from the PublicKey type.
func IsOrExtendsW3IDSecurityV1PublicKey(other vocab.Type) bool {
	return typepublickey.IsOrExtendsPublicKey(other)
}
