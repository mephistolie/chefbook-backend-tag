# ChefBook Backend Tag Service

The tag service owns tag taxonomy and localized lookup used by recipe search, filtering, and recipe metadata.

## Responsibilities

- Store tag groups.
- Store localized tags and emoji metadata.
- Provide tag lists, tag maps, single tag lookup, and tag group lookup.
- Keep fallback localization behavior close to tag persistence.

## Main RPC Families

- `GetTags`
- `GetTagsMap`
- `GetTag`
- `GetTagGroups`

## Dependencies

- Owns its PostgreSQL schema and migrations.
- Does not call other ChefBook services for core tag lookup.

## Database Ownership

Owns:

- `groups` - localized tag group names.
- `tags` - localized tag names, emoji, and group membership.

```mermaid
erDiagram
    TAG_GROUPS {
        varchar group_id PK
        varchar name_en
        varchar name_ru
    }

    TAGS {
        varchar tag_id PK
        varchar name_en
        varchar name_ru
        varchar emoji
        varchar group_id FK
    }

    TAG_GROUPS ||--o{ TAGS : contains
```

Important constraints:

- Tags belong to local tag groups.
- Localization fallback behavior is owned by this service.
