package postgres

import (
	"database/sql"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-tag/internal/entity"
)

const (
	codeEn = "en"
	codeRu = "ru"
	codeUk = "uk"
	codeBe = "be"

	nameColumnPrefix = "name"
)

var (
	supportedLanguages   = []string{codeEn, codeRu}
	ruConsonantLanguages = []string{codeUk, codeBe}
)

func (r *Repository) GetTagsAndGroups(languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string) {
	tags, usedGroupIds := r.getTagsWithGroupsIds(languageCode, nil, groupIds)
	return tags, r.getGroups(languageCode, &usedGroupIds)
}

func (r *Repository) GetTagsMapWithGroups(tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string) {
	tags, usedGroupIds := r.getTagsWithGroupsIds(languageCode, &tagIds, nil)
	tagsMap := make(map[string]entity.Tag)
	for _, tag := range tags {
		tagsMap[tag.Id] = tag
	}
	return tagsMap, r.getGroups(languageCode, &usedGroupIds)
}

func (r *Repository) GetTagWithGroup(tagId, languageCode string) (entity.Tag, *string, error) {
	var tag entity.Tag
	var groupName *string

	getTagQuery := fmt.Sprintf(`
		SELECT tag_id, %[2]v, emoji, group_id
		FROM %[1]v
		WHERE tag_id=$1
	`, tagsTable, r.getNameColumn(languageCode))

	row := r.db.QueryRow(getTagQuery, tagId)
	if err := row.Scan(&tag.Id, &tag.Name, &tag.Emoji, &tag.GroupId); err != nil {
		log.Infof("unable to get tag %s: %s", tagId, err)
		return entity.Tag{}, nil, fail.GrpcNotFound
	}

	if tag.Name == nil {
		tag.Name = r.getFallbackTagName(tagId, languageCode)
		if tag.Name == nil {
			log.Warnf("unable to get tag %s name for language %s", tagId, languageCode)
			return entity.Tag{}, nil, fail.GrpcNotFound
		}
	}

	if tag.GroupId != nil {
		getGroupQuery := fmt.Sprintf(`
			SELECT name
			FROM %[1]v
			WHERE group_id=$1
		`, tagsTable, r.getNameColumn(languageCode))

		row = r.db.QueryRow(getGroupQuery, tag.GroupId)
		_ = row.Scan(&groupName)
	}

	return tag, groupName, nil
}
func (r *Repository) getTagsWithGroupsIds(languageCode string, tagIds *[]string, groupIds *[]string) ([]entity.Tag, []string) {
	var tags []entity.Tag
	usedGroupIdsSet := make(map[string]bool)

	query := fmt.Sprintf(`
		SELECT tag_id, %[2]v, emoji, group_id
		FROM %[1]v
	`, tagsTable, r.getNameColumn(languageCode))

	var rows *sql.Rows
	var err error

	if tagIds != nil && groupIds != nil {
		query = query + " WHERE tag_id=ANY($1) AND group_id=ANY($2)"
		rows, err = r.db.Query(query, *tagIds, *groupIds)
	} else if tagIds != nil {
		query = query + " WHERE tag_id=ANY($1)"
		rows, err = r.db.Query(query, *tagIds)
	} else if groupIds != nil {
		query = query + " WHERE group_id=ANY($1)"
		rows, err = r.db.Query(query, *groupIds)
	} else {
		rows, err = r.db.Query(query)
	}

	if err != nil {
		log.Errorf("unable to get tags: %s", err)
		return []entity.Tag{}, []string{}
	}

	for rows.Next() {
		var tag entity.Tag
		if err = rows.Scan(&tag.Id, &tag.Name, &tag.Emoji, &tag.GroupId); err != nil {
			log.Errorf("unable to parse tag: %s", err)
			continue
		}
		if tag.Name == nil {
			tag.Name = r.getFallbackTagName(tag.Id, languageCode)
		}
		if tag.Name != nil {
			tags = append(tags, tag)

			if tag.GroupId != nil {
				usedGroupIdsSet[*tag.GroupId] = true
			}
		}
	}

	var usedGroupIds []string
	for id := range usedGroupIdsSet {
		usedGroupIds = append(usedGroupIds, id)
	}

	return tags, usedGroupIds
}

func (r *Repository) getFallbackTagName(tagId string, languageCode string) *string {
	var name *string

	query := fmt.Sprintf(`
		SELECT %[2]v
		FROM %[1]v
		WHERE tag_id=$1
	`, tagsTable, r.getFallbackNameColumn(languageCode))

	row := r.db.QueryRow(query, tagId)
	_ = row.Scan(&name)

	return name
}
