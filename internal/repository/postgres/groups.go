package postgres

import (
	"database/sql"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
)

func (r *Repository) GetGroups(languageCode string) map[string]string {
	return r.getGroups(languageCode, nil)
}

func (r *Repository) getGroups(languageCode string, groupIds *[]string) map[string]string {
	groups := make(map[string]string)

	query := fmt.Sprintf(`
		SELECT group_id, %[2]v
		FROM %[1]v
	`, groupsTable, r.getNameColumn(languageCode))

	var rows *sql.Rows
	var err error

	if groupIds != nil {
		query = query + " WHERE group_id=ANY($1)"
		rows, err = r.db.Query(query, *groupIds)
	} else {
		rows, err = r.db.Query(query)
	}

	if err != nil {
		log.Errorf("unable to get groups: %s", err)
		return map[string]string{}
	}

	for rows.Next() {
		var id string
		var name *string
		if err = rows.Scan(&id, &name); err != nil {
			log.Errorf("unable to parse group: %s", err)
			continue
		}
		if name == nil {
			name = r.getFallbackGroupName(id, languageCode)
		}
		if name != nil {
			groups[id] = *name
		}
	}

	return groups
}

func (r *Repository) getFallbackGroupName(groupId string, languageCode string) *string {
	var name *string

	query := fmt.Sprintf(`
		SELECT %[2]v
		FROM %[1]v
		WHERE group_id=$1
	`, groupsTable, r.getFallbackNameColumn(languageCode))

	row := r.db.QueryRow(query, groupId)
	_ = row.Scan(&name)

	return name
}
