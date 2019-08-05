# go-metabase
Metabase API Client for Golang

## Installation
```
go get github.com/kucuny/go-metabase/metabase
```

## Usage
```go
import "github.com/kucuny/go-metabase/metabase"
```

## Features (v0.32.10)

- [ ] Activity
  - [ ] GET /api/activity/
  - [ ] GET /api/activity/recent_views

- [ ] Card
  - [ ] DELETE /api/card/:card-id/favorite
  - [ ] DELETE /api/card/:id
  - [ ] GET /api/card/
  - [ ] GET /api/card/:id
  - [ ] POST /api/card/
  - [ ] POST /api/card/:card-id/favorite
  - [ ] POST /api/card/:card-id/labels
  - [ ] POST /api/card/:card-id/query
  - [ ] POST /api/card/:card-id/query/csv
  - [ ] POST /api/card/:card-id/query/json
  - [ ] PUT /api/card/:id

- [ ] Dashboard
  - [ ] DELETE /api/dashboard/:id
  - [ ] DELETE /api/dashboard/:id/cards
  - [ ] GET /api/dashboard/
  - [ ] GET /api/dashboard/:id
  - [ ] GET /api/dashboard/:id/revisions
  - [ ] POST /api/dashboard/
  - [ ] POST /api/dashboard/:id/cards
  - [ ] POST /api/dashboard/:id/revert
  - [ ] PUT /api/dashboard/:id
  - [ ] PUT /api/dashboard/:id/cards

- [ ] Database
  - [ ] DELETE /api/database/:id
  - [ ] GET /api/database/
  - [ ] GET /api/database/:id
  - [ ] GET /api/database/:id/autocomplete_suggestions
  - [ ] GET /api/database/:id/fields
  - [ ] GET /api/database/:id/idfields
  - [ ] GET /api/database/:id/metadata
  - [ ] POST /api/database/
  - [ ] POST /api/database/:id/sync
  - [ ] POST /api/database/sample_dataset
  - [ ] PUT /api/database/:id

- [ ] Dataset
  - [ ] POST /api/dataset/
  - [ ] POST /api/dataset/csv
  - [ ] POST /api/dataset/duration
  - [ ] POST /api/dataset/json

- [ ] Email
  - [ ] POST /api/email/test
  - [ ] PUT /api/email/

- [ ] Field
  - [ ] GET /api/field/:id
  - [ ] GET /api/field/:id/summary
  - [ ] GET /api/field/:id/values
  - [ ] POST /api/field/:id/value_map_update
  - [ ] PUT /api/field/:id

- [ ] Geojson
  - [ ] GET /api/geojson/:key

- [ ] Getting started
  - [ ] GET /api/getting-started/

- [ ] Label
  - [ ] DELETE /api/label/:id
  - [ ] GET /api/label/
  - [ ] POST /api/label/
  - [ ] PUT /api/label/:id

- [ ] Metric
  - [ ] DELETE /api/metric/:id
  - [ ] GET /api/metric/
  - [ ] GET /api/metric/:id
  - [ ] GET /api/metric/:id/revisions
  - [ ] POST /api/metric/
  - [ ] POST /api/metric/:id/revert
  - [ ] PUT /api/metric/:id
  - [ ] PUT /api/metric/:id/important_fields

- [ ] Notify
  - [ ] POST /api/notify/db/:id

- [ ] Permissions
  - [ ] DELETE /api/permissions/group/:group-id
  - [ ] DELETE /api/permissions/membership/:id
  - [ ] GET /api/permissions/graph
  - [ ] GET /api/permissions/group/:id
  - [ ] GET /api/permissions/membership
  - [ ] POST /api/permissions/group
  - [ ] POST /api/permissions/membership
  - [ ] PUT /api/permissions/graph
  - [ ] PUT /api/permissions/group/:group-id

- [ ] Pulse
  - [ ] DELETE /api/pulse/:id
  - [ ] GET /api/pulse/
  - [ ] GET /api/pulse/:id
  - [ ] GET /api/pulse/form_input
  - [ ] GET /api/pulse/preview_card/:id
  - [ ] GET /api/pulse/preview_card_info/:id
  - [ ] GET /api/pulse/preview_card_png/:id
  - [ ] POST /api/pulse/
  - [ ] POST /api/pulse/test
  - [ ] PUT /api/pulse/:id

- [ ] Revision
  - [ ] GET /api/revision/
  - [ ] POST /api/revision/revert

- [ ] Segment
  - [ ] DELETE /api/segment/:id
  - [ ] GET /api/segment/
  - [ ] GET /api/segment/:id
  - [ ] GET /api/segment/:id/revisions
  - [ ] POST /api/segment/
  - [ ] POST /api/segment/:id/revert
  - [ ] PUT /api/segment/:id

- [ ] Session
  - [x] DELETE /api/session/
  - [ ] GET /api/session/password_reset_token_valid
  - [ ] GET /api/session/properties
  - [x] POST /api/session/
  - [ ] POST /api/session/forgot_password
  - [ ] POST /api/session/google_auth
  - [ ] POST /api/session/reset_password

- [ ] Setting
  - [ ] GET /api/setting/
  - [ ] GET /api/setting/:key
  - [ ] PUT /api/setting/:key

- [ ] Setup
  - [ ] GET /api/setup/admin_checklist
  - [ ] POST /api/setup/
  - [ ] POST /api/setup/validate

- [ ] Slack
  - [ ] PUT /api/slack/settings

- [ ] Table
  - [ ] GET /api/table/
  - [ ] GET /api/table/:id
  - [ ] GET /api/table/:id/fks
  - [ ] GET /api/table/:id/query_metadata
  - [ ] PUT /api/table/:id

- [ ] Tiles
  - [ ] GET /api/tiles/:zoom/:x/:y/:lat-field-id/:lon-field-id/:lat-col-idx/:lon-col-idx/

- [ ] User
  - [ ] DELETE /api/user/:id
  - [ ] GET /api/user/
  - [ ] GET /api/user/:id
  - [ ] GET /api/user/current
  - [ ] POST /api/user/
  - [ ] POST /api/user/:id/send_invite
  - [ ] PUT /api/user/:id
  - [ ] PUT /api/user/:id/password
  - [ ] PUT /api/user/:id/qbnewb

- [ ] Util
  - [ ] GET /api/util/logs
  - [ ] GET /api/util/stats
  - [ ] POST /api/util/password_check

## References
Metabase - https://github.com/metabase/metabase/blob/master/docs/api-documentation.md
