package postgres

import (
	"context"
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
	supportedLanguages   = []string{codeEn, codeRu, codeUk, codeBe}
	ruConsonantLanguages = []string{codeUk, codeBe}
)

func (r *Repository) GetTagsAndGroups(ctx context.Context, languageCode string, groupIds *[]string) ([]entity.Tag, map[string]string) {
	tags, usedGroupIds := r.getTagsWithGroupsIds(ctx, languageCode, nil, groupIds)
	return tags, r.getGroups(ctx, languageCode, &usedGroupIds)
}

func (r *Repository) GetTagsMapWithGroups(ctx context.Context, tagIds []string, languageCode string) (map[string]entity.Tag, map[string]string) {
	tags, usedGroupIds := r.getTagsWithGroupsIds(ctx, languageCode, &tagIds, nil)
	tagsMap := make(map[string]entity.Tag)
	for _, tag := range tags {
		tagsMap[tag.Id] = tag
	}
	return tagsMap, r.getGroups(ctx, languageCode, &usedGroupIds)
}

func (r *Repository) GetTagWithGroup(ctx context.Context, tagId, languageCode string) (entity.Tag, *string, error) {
	var tag entity.Tag
	var groupName *string

	getTagQuery := fmt.Sprintf(`
		SELECT tag_id, %[2]v, emoji, group_id
		FROM %[1]v
		WHERE tag_id=$1
	`, tagsTable, r.getNameColumn(languageCode))

	row := r.db.QueryRowContext(ctx, getTagQuery, tagId)
	if err := row.Scan(&tag.Id, &tag.Name, &tag.Emoji, &tag.GroupId); err != nil {
		log.AutoInfof("unable to get tag %s: %s", tagId, err)
		return entity.Tag{}, nil, fail.GrpcNotFound
	}

	if tag.Name == nil {
		tag.Name = r.getFallbackTagName(ctx, tagId, languageCode)
		if tag.Name == nil {
			log.AutoWarnf("unable to get tag %s name for language %s", tagId, languageCode)
			return entity.Tag{}, nil, fail.GrpcNotFound
		}
	}

	if tag.GroupId != nil {
		getGroupQuery := fmt.Sprintf(`
			SELECT %[2]v
			FROM %[1]v
			WHERE group_id=$1
		`, groupsTable, r.getNameColumn(languageCode))

		row = r.db.QueryRowContext(ctx, getGroupQuery, *tag.GroupId)
		_ = row.Scan(&groupName)
	}

	return tag, groupName, nil
}
func (r *Repository) getTagsWithGroupsIds(ctx context.Context, languageCode string, tagIds *[]string, groupIds *[]string) ([]entity.Tag, []string) {
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
		rows, err = r.db.QueryContext(ctx, query, *tagIds, *groupIds)
	} else if tagIds != nil {
		query = query + " WHERE tag_id=ANY($1)"
		rows, err = r.db.QueryContext(ctx, query, *tagIds)
	} else if groupIds != nil {
		query = query + " WHERE group_id=ANY($1)"
		rows, err = r.db.QueryContext(ctx, query, *groupIds)
	} else {
		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		log.AutoErrorf("unable to get tags: %s", err)
		return []entity.Tag{}, []string{}
	}
	defer rows.Close()

	for rows.Next() {
		var tag entity.Tag
		if err = rows.Scan(&tag.Id, &tag.Name, &tag.Emoji, &tag.GroupId); err != nil {
			log.AutoErrorf("unable to parse tag: %s", err)
			continue
		}
		if tag.Name == nil {
			tag.Name = r.getFallbackTagName(ctx, tag.Id, languageCode)
		}
		if tag.Name != nil {
			tags = append(tags, tag)

			if tag.GroupId != nil {
				usedGroupIdsSet[*tag.GroupId] = true
			}
		}
	}
	if err = rows.Err(); err != nil {
		log.AutoErrorf("unable to iterate tags: %s", err)
		return []entity.Tag{}, []string{}
	}

	var usedGroupIds []string
	for id := range usedGroupIdsSet {
		usedGroupIds = append(usedGroupIds, id)
	}

	return tags, usedGroupIds
}

func (r *Repository) getFallbackTagName(ctx context.Context, tagId string, languageCode string) *string {
	var name *string

	query := fmt.Sprintf(`
		SELECT %[2]v
		FROM %[1]v
		WHERE tag_id=$1
	`, tagsTable, r.getFallbackNameColumn(languageCode))

	row := r.db.QueryRowContext(ctx, query, tagId)
	_ = row.Scan(&name)

	return name
}
