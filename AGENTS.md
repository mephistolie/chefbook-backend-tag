# Tag Service Agents Guide

This service owns tag and category-like metadata used by the platform.

## Scope

- tag entities and storage
- tag lookup and transport contracts
- metadata relationships exposed to other services or clients

## Working Rules

- Keep taxonomy logic here instead of duplicating it in recipe or client modules.
- Be deliberate about backward compatibility for tag identifiers and API shapes.
