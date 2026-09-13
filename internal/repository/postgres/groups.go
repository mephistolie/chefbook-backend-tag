package postgres

import (
	"context"
	"database/sql"
	"fmt"

	eventlog "github.com/mephistolie/chefbook-backend-tag/internal/logging"
)

var events = eventlog.NewEvents()

func (r *Repository) GetGroups(ctx context.Context, languageCode string) map[string]string {
	return r.getGroups(ctx, languageCode, nil)
}

func (r *Repository) getGroups(ctx context.Context, languageCode string, groupIds *[]string) map[string]string {
	groups := make(map[string]string)

	query := fmt.Sprintf(`
		SELECT group_id, %[2]v
		FROM %[1]v
	`, groupsTable, r.getNameColumn(languageCode))

	var rows *sql.Rows
	var err error

	if groupIds != nil {
		query = query + " WHERE group_id=ANY($1)"
		rows, err = r.db.QueryContext(ctx, query, *groupIds)
	} else {
		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		events.PostgresQueryFailed(ctx, eventlog.PostgresOperation{
			Operation: "get_groups",
			Entity:    "group",
		}, err)
		return map[string]string{}
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name *string
		if err = rows.Scan(&id, &name); err != nil {
			events.PostgresRowScanFailed(ctx, eventlog.PostgresOperation{
				Operation: "get_groups",
				Entity:    "group",
			}, err)
			continue
		}
		if name == nil {
			name = r.getFallbackGroupName(ctx, id, languageCode)
		}
		if name != nil {
			groups[id] = *name
		}
	}
	if err = rows.Err(); err != nil {
		events.PostgresRowsIterationFailed(ctx, eventlog.PostgresOperation{
			Operation: "get_groups",
			Entity:    "group",
		}, err)
		return map[string]string{}
	}

	return groups
}

func (r *Repository) getFallbackGroupName(ctx context.Context, groupId string, languageCode string) *string {
	var name *string

	query := fmt.Sprintf(`
		SELECT %[2]v
		FROM %[1]v
		WHERE group_id=$1
	`, groupsTable, r.getFallbackNameColumn(languageCode))

	row := r.db.QueryRowContext(ctx, query, groupId)
	_ = row.Scan(&name)

	return name
}
