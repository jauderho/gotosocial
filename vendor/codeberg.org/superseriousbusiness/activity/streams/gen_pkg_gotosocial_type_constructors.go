// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeannounceapproval "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_announceapproval"
	typecanannounce "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_canannounce"
	typecanlike "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_canlike"
	typecanreply "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_canreply"
	typeinteractionpolicy "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_interactionpolicy"
	typelikeapproval "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_likeapproval"
	typereplyapproval "codeberg.org/superseriousbusiness/activity/streams/impl/gotosocial/type_replyapproval"
	vocab "codeberg.org/superseriousbusiness/activity/streams/vocab"
)

// NewGoToSocialAnnounceApproval creates a new GoToSocialAnnounceApproval
func NewGoToSocialAnnounceApproval() vocab.GoToSocialAnnounceApproval {
	return typeannounceapproval.NewGoToSocialAnnounceApproval()
}

// NewGoToSocialCanAnnounce creates a new GoToSocialCanAnnounce
func NewGoToSocialCanAnnounce() vocab.GoToSocialCanAnnounce {
	return typecanannounce.NewGoToSocialCanAnnounce()
}

// NewGoToSocialCanLike creates a new GoToSocialCanLike
func NewGoToSocialCanLike() vocab.GoToSocialCanLike {
	return typecanlike.NewGoToSocialCanLike()
}

// NewGoToSocialCanReply creates a new GoToSocialCanReply
func NewGoToSocialCanReply() vocab.GoToSocialCanReply {
	return typecanreply.NewGoToSocialCanReply()
}

// NewGoToSocialInteractionPolicy creates a new GoToSocialInteractionPolicy
func NewGoToSocialInteractionPolicy() vocab.GoToSocialInteractionPolicy {
	return typeinteractionpolicy.NewGoToSocialInteractionPolicy()
}

// NewGoToSocialLikeApproval creates a new GoToSocialLikeApproval
func NewGoToSocialLikeApproval() vocab.GoToSocialLikeApproval {
	return typelikeapproval.NewGoToSocialLikeApproval()
}

// NewGoToSocialReplyApproval creates a new GoToSocialReplyApproval
func NewGoToSocialReplyApproval() vocab.GoToSocialReplyApproval {
	return typereplyapproval.NewGoToSocialReplyApproval()
}
