// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package migrations

import (
	"context"

	old_gtsmodel "code.superseriousbusiness.org/gotosocial/internal/db/bundb/migrations/20241121121623_enum_strings_to_ints"
	new_gtsmodel "code.superseriousbusiness.org/gotosocial/internal/gtsmodel"
	"code.superseriousbusiness.org/gotosocial/internal/log"
	"code.superseriousbusiness.org/gotosocial/internal/util"

	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		// Status visibility type indices.
		var statusVisIndices = []struct {
			name  string
			cols  []string
			order string
		}{
			{
				name:  "statuses_visibility_idx",
				cols:  []string{"visibility"},
				order: "",
			},
			{
				name:  "statuses_profile_web_view_idx",
				cols:  []string{"account_id", "visibility"},
				order: "id DESC",
			},
			{
				name:  "statuses_public_timeline_idx",
				cols:  []string{"visibility"},
				order: "id DESC",
			},
		}

		// Tables with visibility types.
		var visTables = []struct {
			Table                string
			Column               string
			Default              *new_gtsmodel.Visibility
			IndexCleanupCallback func(ctx context.Context, tx bun.Tx) error
			BatchByColumn        string
		}{
			{
				Table:  "statuses",
				Column: "visibility",
				IndexCleanupCallback: func(ctx context.Context, tx bun.Tx) error {
					// After new column has been created and
					// populated, drop indices relying on old column.
					for _, index := range statusVisIndices {
						log.Infof(ctx, "dropping old index %s...", index.name)
						if _, err := tx.NewDropIndex().
							Index(index.name).
							Exec(ctx); err != nil {
							return err
						}
					}
					return nil
				},
				BatchByColumn: "id",
			},
			{
				Table:         "sin_bin_statuses",
				Column:        "visibility",
				BatchByColumn: "id",
			},
			{
				Table:         "account_settings",
				Column:        "privacy",
				Default:       util.Ptr(new_gtsmodel.VisibilityDefault),
				BatchByColumn: "account_id",
			},

			{
				Table:         "account_settings",
				Column:        "web_visibility",
				Default:       util.Ptr(new_gtsmodel.VisibilityDefault),
				BatchByColumn: "account_id",
			},
		}

		// Get the mapping of old enum string values to new integer values.
		visibilityMapping := visibilityEnumMapping[old_gtsmodel.Visibility]()

		// Convert all visibility tables.
		for _, table := range visTables {

			// Perform each enum table conversion within its own transaction.
			if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
				return convertEnums(ctx, tx, table.Table, table.Column,
					visibilityMapping, table.Default, table.IndexCleanupCallback, table.BatchByColumn)
			}); err != nil {
				return err
			}
		}

		// Recreate the visibility indices.
		log.Info(ctx, "creating new visibility indexes...")
		for _, index := range statusVisIndices {
			log.Infof(ctx, "creating new index %s...", index.name)
			q := db.NewCreateIndex().
				Table("statuses").
				Index(index.name).
				Column(index.cols...)
			if index.order != "" {
				q = q.ColumnExpr(index.order)
			}
			if _, err := q.Exec(ctx); err != nil {
				return err
			}
		}

		// Get the mapping of old enum string values to the new integer value types.
		notificationMapping := notificationEnumMapping[old_gtsmodel.NotificationType]()

		// Migrate over old notifications table column to new type in tx.
		if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return convertEnums(ctx, tx, "notifications", "notification_type", //nolint:revive
				notificationMapping, nil, nil, "id")
		}); err != nil {
			return err
		}

		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}

// visibilityEnumMapping maps old Visibility enum values to their newer integer type.
func visibilityEnumMapping[T ~string]() map[T]new_gtsmodel.Visibility {
	return map[T]new_gtsmodel.Visibility{
		T(old_gtsmodel.VisibilityNone):          new_gtsmodel.VisibilityNone,
		T(old_gtsmodel.VisibilityPublic):        new_gtsmodel.VisibilityPublic,
		T(old_gtsmodel.VisibilityUnlocked):      new_gtsmodel.VisibilityUnlocked,
		T(old_gtsmodel.VisibilityFollowersOnly): new_gtsmodel.VisibilityFollowersOnly,
		T(old_gtsmodel.VisibilityMutualsOnly):   new_gtsmodel.VisibilityMutualsOnly,
		T(old_gtsmodel.VisibilityDirect):        new_gtsmodel.VisibilityDirect,
	}
}

// notificationEnumMapping maps old NotificationType enum values to their newer integer type.
func notificationEnumMapping[T ~string]() map[T]new_gtsmodel.NotificationType {
	return map[T]new_gtsmodel.NotificationType{
		T(old_gtsmodel.NotificationFollow):        new_gtsmodel.NotificationFollow,
		T(old_gtsmodel.NotificationFollowRequest): new_gtsmodel.NotificationFollowRequest,
		T(old_gtsmodel.NotificationMention):       new_gtsmodel.NotificationMention,
		T(old_gtsmodel.NotificationReblog):        new_gtsmodel.NotificationReblog,
		T(old_gtsmodel.NotificationFave):          new_gtsmodel.NotificationFavourite,
		T(old_gtsmodel.NotificationPoll):          new_gtsmodel.NotificationPoll,
		T(old_gtsmodel.NotificationStatus):        new_gtsmodel.NotificationStatus,
		T(old_gtsmodel.NotificationSignup):        new_gtsmodel.NotificationAdminSignup,
		T(old_gtsmodel.NotificationPendingFave):   new_gtsmodel.NotificationPendingFave,
		T(old_gtsmodel.NotificationPendingReply):  new_gtsmodel.NotificationPendingReply,
		T(old_gtsmodel.NotificationPendingReblog): new_gtsmodel.NotificationPendingReblog,
	}
}
