{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (schema .Table.TableName) -}}
{{- $idxFields := (flatidxfields .) -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} `json:"{{ .Col.ColumnName }}" db:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
{{- if .PrimaryKey }}

	// xo fields
	_exists, _deleted bool
{{ end }}
}

{{ if .PrimaryKey }}
// Exists determines if the {{ .Name }} exists in the database.
func ({{ $short }} *{{ .Name }}) Exists() bool {
	return {{ $short }}._exists
}

// Deleted provides information if the {{ .Name }} has been deleted from the database.
func ({{ $short }} *{{ .Name }}) Deleted() bool {
	return {{ $short }}._deleted
}

// Insert inserts the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if {{ $short }}._exists {
		return errors.New("insert failed: already exists")
	}

	{{ gqlinsert . }}

	// set existence
	{{ $short }}._exists = true

	return nil
}

{{ if ne (fieldnamesmulti .Fields $short .PrimaryKeyFields) "" }}
	// Update updates the {{ .Name }} in the database.
	func ({{ $short }} *{{ .Name }}) Update(db XODB) error {
		var err error

		// if doesn't exist, bail
		if !{{ $short }}._exists {
			return errors.New("update failed: does not exist")
		}

		// if deleted, bail
		if {{ $short }}._deleted {
			return errors.New("update failed: marked for deletion")
		}

		{{ if gt ( len .PrimaryKeyFields ) 1 }}
			// sql query with composite primary key
			const sqlstr = `UPDATE "{{ $table }}" SET (` +
				`{{ colnamesmulti .Fields .PrimaryKeyFields }}` +
				`) = ( ` +
				`{{ colvalsmulti .Fields .PrimaryKeyFields }}` +
				`) WHERE {{ colnamesquerymulti .PrimaryKeyFields " AND " (getstartcount .Fields .PrimaryKeyFields) nil }}`

			// run query
			XOLog(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
			_, err = db.Exec(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
		return err
		{{- else }}
			// sql query
			const sqlstr = `UPDATE "{{ $table }}" SET (` +
				`{{ colnames .Fields .PrimaryKey.Name }}` +
				`) = ( ` +
				`{{ colvals .Fields .PrimaryKey.Name }}` +
				`) WHERE {{ colname .PrimaryKey.Col }} = ${{ colcount .Fields .PrimaryKey.Name }}`

			// run query
			XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
			_, err = db.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
			return err
		{{- end }}
	}

	// Save saves the {{ .Name }} to the database.
	func ({{ $short }} *{{ .Name }}) Save(db XODB) error {
		if {{ $short }}.Exists() {
			return {{ $short }}.Update(db)
		}

		return {{ $short }}.Insert(db)
	}

	// Upsert performs an upsert for {{ .Name }}.
	//
	// NOTE: PostgreSQL 9.5+ only
	func ({{ $short }} *{{ .Name }}) Upsert(db XODB) error {
		var err error

		// if already exist, bail
		if {{ $short }}._exists {
			return errors.New("insert failed: already exists")
		}

		// sql query
		const sqlstr = `INSERT INTO "{{ $table }}" (` +
			`{{ colnames .Fields }}` +
			`) VALUES (` +
			`{{ colvals .Fields }}` +
			`) ON CONFLICT ({{ colnames .PrimaryKeyFields }}) DO UPDATE SET (` +
			`{{ colnames .Fields }}` +
			`) = (` +
			`{{ colprefixnames .Fields "EXCLUDED" }}` +
			`)`

		// run query
		XOLog(sqlstr, {{ fieldnames .Fields $short }})
		_, err = db.Exec(sqlstr, {{ fieldnames .Fields $short }})
		if err != nil {
			return err
		}

		// set existence
		{{ $short }}._exists = true

		return nil
}
{{ else }}
	// Update statements omitted due to lack of fields other than primary key
{{ end }}

// Delete deletes the {{ .Name }} from the database.
func ({{ $short }} *{{ .Name }}) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !{{ $short }}._exists {
		return nil
	}

	// if deleted, bail
	if {{ $short }}._deleted {
		return nil
	}

	{{ if gt ( len .PrimaryKeyFields ) 1 }}
		// sql query with composite primary key
		const sqlstr = `DELETE FROM "{{ $table }}"  WHERE {{ colnamesquery .PrimaryKeyFields " AND " }}`

		// run query
		XOLog(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		_, err = db.Exec(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		if err != nil {
			return err
		}
	{{- else }}
		// sql query
		const sqlstr = `DELETE FROM "{{ $table }}" WHERE {{ colname .PrimaryKey.Col }} = $1`

		// run query
		XOLog(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		_, err = db.Exec(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		if err != nil {
			return err
		}
	{{- end }}

	// set deleted
	{{ $short }}._deleted = true

	return nil
}
{{- end }}

{{ if (existsqlfilter $table $idxFields) }}
// {{ .Name }}Filter related to {{ .Name }}QueryArguments
// struct field name contain table column name in Camel style and logic operator(lt, gt etc)
// only indexed column and special column defined in sqlSpecColFilterCtlMap declared in file xo/internal/funcs.go
type {{ .Name }}Filter struct {
	Conjunction	*string  // enum in "AND", "OR", nil(consider as single condition)
{{- range .Fields -}}
	{{- $ftyp := (sqlfilter $table . $idxFields) -}}
	{{- if (and (ne .Name $.PrimaryKey.Name) (ne $ftyp "unsupported")) -}}
		{{- if (or (eq $ftyp "Number") (eq $ftyp "String")) }}
			{{ .Name }} {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}"` // equal to {{ .Name }}
		{{- end -}}
		{{- if (eq $ftyp "String") }}
			{{ .Name }}Like {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_like"` // LIKE
			{{ .Name }}ILike {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_ilike"` // ILIKE case-insensitive
			{{ .Name }}NLike {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_nlike"` // NOT LIKE
			{{ .Name }}NILike {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_nilike"` // NOT ILIKE case-insensitive
		{{- end -}}
		{{- if (or (eq $ftyp "Number") (eq $ftyp "Time")) }}
			{{ .Name }}Lt {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_lt"` // less than {{ .Name }}
			{{ .Name }}Lte {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_lte"` // less than and equal to {{ .Name }}
			{{ .Name }}Gt {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_gt"` // greater than {{ .Name }}
			{{ .Name }}Gte {{ sqltogopointertype .Type .Col.IsPrimaryKey }} `json:"{{ togqlname .Name }}_gte"` // greater than and equal to {{ .Name }}
		{{- end -}}
	{{- end -}}
{{- end }}
}

// {{ .Name }}QueryArguments composed by Cursor, {{ .Name }}Filter and sql filter string
type {{ .Name }}QueryArguments struct{
	Cursor
	Where *{{ .Name }}Filter

	// non-export field
	filterArgs *filterArguments
}

// get{{ .Name }}Filter return the sql filter
func get{{ .Name }}Filter(filter *{{ .Name }}Filter) (*filterArguments, error){
	if filter == nil{
		return nil, nil
	}
	conjunction := ""
	conjCnt := 0
	var filterPairs []*filterPair
	if filter.Conjunction != nil{
		conjunction = *filter.Conjunction
		if _, ok := sqlConjunctionMap[conjunction]; !ok{
			return nil, fmt.Errorf("unsupported conjunction:%v", filter.Conjunction)
		}
	}
{{- range .Fields -}}
	{{- $ftyp := (sqlfilter $table . $idxFields) -}}
	{{- if (and (ne .Name $.PrimaryKey.Name) (ne $ftyp "unsupported")) -}}
		{{- if (or (eq $ftyp "Number") (eq $ftyp "String")) }}
			if filter.{{ .Name }} != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "=", value: *filter.{{ .Name }}})
			}
		{{- end -}}
		{{- if (eq $ftyp "String") }}
			if filter.{{ .Name }}Like != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "LIKE", value: *filter.{{ .Name }}Like})
			}
			if filter.{{ .Name }}ILike != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "ILIKE", value: *filter.{{ .Name }}ILike})
			}
			if filter.{{ .Name }}NLike != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "NOT LIKE", value: *filter.{{ .Name }}NLike})
			}
			if filter.{{ .Name }}NILike != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "NOT ILIKE", value: *filter.{{ .Name }}NILike})
			}
		{{- end -}}
		{{- if (eq $ftyp "Number") }}
			if filter.{{ .Name}}Lt != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "<", value: *filter.{{ .Name }}Lt})
			}else if filter.{{ .Name}}Lte != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "<=", value: *filter.{{ .Name }}Lte})
			}
			if filter.{{ .Name}}Gt != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: ">", value: *filter.{{ .Name }}Gt})
			}else if filter.{{ .Name}}Gte != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: ">=", value: *filter.{{ .Name }}Gte})
			}
		{{- end -}}
		{{- if (eq $ftyp "Time") }}
			if filter.{{ .Name}}Lt != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "<", value: filter.{{ .Name }}Lt.Time})
			}else if filter.{{ .Name}}Lte != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: "<=", value: filter.{{ .Name }}Lte.Time})
			}
			if filter.{{ .Name}}Gt != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: ">", value: filter.{{ .Name }}Gt.Time})
			}else if filter.{{ .Name}}Gte != nil{
				conjCnt++
				filterPairs = append(filterPairs, &filterPair{fieldName: "{{ .Col.ColumnName }}", option: ">=", value: filter.{{ .Name }}Gte.Time})
			}
		{{- end -}}
	{{- end -}}
{{- end }}
	if conjCnt == 0{
		return nil, nil
	}
	if len(conjunction)>0 && conjCnt < 2{
		return nil, fmt.Errorf("invalid filter conjunction: %v need more than 2 parameter but have: %v", *filter.Conjunction, conjCnt)
	}
	if len(conjunction) == 0 && conjCnt != 1{
		return nil, fmt.Errorf("multi field:%v should be connected by conjunction AND or OR", conjCnt)
	}
	filterArgs := &filterArguments{filterPairs: filterPairs, conjunction: conjunction, conjCnt: conjCnt}
	return filterArgs, nil
}

{{ else }}
// {{ .Name }}QueryArguments composed by Cursor, {{ .Name }}Filter and sql filter string
type {{ .Name }}QueryArguments struct{
	Cursor
}
{{ end }}

// Apply{{ .Name }}QueryArgsDefaults assigns default cursor values to non-nil fields.
func Apply{{ .Name }}QueryArgsDefaults(queryArgs *{{ .Name }}QueryArguments) *{{ .Name }}QueryArguments {
	if queryArgs == nil {
		queryArgs = &{{ .Name }}QueryArguments{
			Cursor:DefaultCursor,
		}
		return queryArgs
	}
	if queryArgs.Offset == nil {
		queryArgs.Offset = DefaultCursor.Offset
	}
	if queryArgs.Limit == nil {
		queryArgs.Limit = DefaultCursor.Limit
	}
	if queryArgs.Index == nil {
		queryArgs.Index = DefaultCursor.Index
	}
	if queryArgs.Desc == nil {
		queryArgs.Desc = DefaultCursor.Desc
	}
	if queryArgs.Dead == nil {
		queryArgs.Dead = DefaultCursor.Dead
	}
	if queryArgs.After == nil {
		queryArgs.After = DefaultCursor.After
	}
	if queryArgs.First == nil {
		queryArgs.First = DefaultCursor.First
	}
	if queryArgs.Before == nil {
		queryArgs.Before = DefaultCursor.Before
	}
	if queryArgs.Last == nil {
		queryArgs.Last = DefaultCursor.Last
	}
	return queryArgs
}

// GetMostRecent{{ .Name }} returns n most recent rows from '{{ .Schema }}.{{ .Table.TableName }}',
// ordered by "created_date" in descending order.
func GetMostRecent{{ .Name }}(db XODB, n int) ([]*{{ .Name }}, error) {
    const sqlstr = `SELECT ` +
        `{{ colnames .Fields }} ` +
        `FROM "{{ $table }}" ` +
        `ORDER BY created_date DESC LIMIT $1`

    q, err := db.Query(sqlstr, n)
    if err != nil {
        return nil, err
    }
    defer q.Close()

    // load results
    var res []*{{ .Name }}
    for q.Next() {
        {{ $short }} := {{ .Name }}{}

        // scan
        err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
        if err != nil {
            return nil, err
        }

        res = append(res, &{{ $short }})
    }

    return res, nil
}


// GetMostRecentChanged{{ .Name }} returns n most recent rows from '{{ .Schema }}.{{ .Table.TableName }}',
// ordered by "changed_date" in descending order.
func GetMostRecentChanged{{ .Name }}(db XODB, n int) ([]*{{ .Name }}, error) {
    const sqlstr = `SELECT ` +
        `{{ colnames .Fields }} ` +
        `FROM "{{ $table }}" ` +
        `ORDER BY changed_date DESC LIMIT $1`

    q, err := db.Query(sqlstr, n)
    if err != nil {
        return nil, err
    }
    defer q.Close()

    // load results
    var res []*{{ .Name }}
    for q.Next() {
        {{ $short }} := {{ .Name }}{}

        // scan
        err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
        if err != nil {
            return nil, err
        }

        res = append(res, &{{ $short }})
    }

    return res, nil
}

// GetAll{{ .Name }} returns all rows from '{{ .Schema }}.{{ .Table.TableName }}', based on the {{ .Name }}QueryArguments.
// If the {{ .Name }}QueryArguments is nil, it will use the default {{ .Name }}QueryArguments instead.
func GetAll{{ .Name }}(db XODB, queryArgs *{{ .Name }}QueryArguments) ([]*{{ .Name }}, error) { // nolint: gocyclo
	queryArgs = Apply{{ .Name }}QueryArgsDefaults(queryArgs)

	desc := ""
	if *queryArgs.Desc {
		desc = "DESC"
	}

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	var params []interface{}
	placeHolders := ""
{{- if (existsqlfilter $table $idxFields) }}
	if queryArgs.filterArgs != nil{
		pls := make([]string, len(queryArgs.filterArgs.filterPairs))
		for i, pair := range queryArgs.filterArgs.filterPairs {
		   pls[i] = fmt.Sprintf("%s %s $%d", pair.fieldName, pair.option, i+1)
		   params = append(params, pair.value)
	   }
	   placeHolders = strings.Join(pls, " " + queryArgs.filterArgs.conjunction + " ")
	   placeHolders = fmt.Sprintf("(%s) AND", placeHolders)
	}
{{- end }}
	params = append(params, *queryArgs.Offset)
	offsetPos := len(params)

	params = append(params, *queryArgs.Limit)
	limitPos := len(params)
	var sqlstr = fmt.Sprintf(
		`SELECT %s FROM %s WHERE %s deleted_date IS %s ORDER BY %s %s OFFSET $%d LIMIT $%d`,
		`{{ colnames .Fields }} `,
		`"{{ $table }}"`,
		placeHolders,
		dead,
		"{{ .PrimaryKey.Col.ColumnName }}",
		desc,
		offsetPos,
		limitPos)

	XOLog(sqlstr, params...)

	q, err := db.Query(sqlstr, params...)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*{{ .Name }}
	for q.Next() {
		{{ $short }} := {{ .Name }}{}

		// scan
		err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
		if err != nil {
		    return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
}

// CountAll{{ .Name }} returns a count of all rows from '{{ .Schema }}.{{ .Table.TableName }}'
func CountAll{{ .Name }}(db XODB, queryArgs *{{ .Name }}QueryArguments) (int, error) {
	queryArgs = Apply{{ .Name }}QueryArgsDefaults(queryArgs)

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	var params []interface{}
	placeHolders := ""
{{- if (existsqlfilter $table $idxFields) }}
	if queryArgs.filterArgs != nil{
		pls := make([]string, len(queryArgs.filterArgs.filterPairs))
		for i, pair := range queryArgs.filterArgs.filterPairs {
		   pls[i] = fmt.Sprintf("%s %s $%d", pair.fieldName, pair.option, i+1)
		   params = append(params, pair.value)
	   }
	   placeHolders = strings.Join(pls, " " + queryArgs.filterArgs.conjunction + " ")
	   placeHolders = fmt.Sprintf("(%s) AND", placeHolders)
	}
{{- end }}

	var err error
	var sqlstr = fmt.Sprintf(`SELECT count(*) from "{{ $table }}" WHERE %s deleted_date IS %s`, placeHolders, dead)
	XOLog(sqlstr)

	var count int
	err = db.QueryRow(sqlstr, params...).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

{{ range .ForeignKeys }}
	{{- $fnname := (print (plural $.Name) "By" .Field.Name "FK") -}}
	{{- if not (isdup $fnname) }}
	// {{ $fnname }} retrieves rows from {{ $table }} by foreign key {{.Field.Name}}.
	// Generated from foreign key {{.Name}}.
	func {{ $fnname }}(db XODB, {{ togqlname .Field.Name }} {{ .RefField.Type }}, queryArgs *{{ $.Name }}QueryArguments) ([]*{{$.Name}}, error) {
		queryArgs = Apply{{ $.Name }}QueryArgsDefaults(queryArgs)

		desc := ""
		if *queryArgs.Desc {
			desc = "DESC"
		}

		dead := "NULL"
		if *queryArgs.Dead {
			dead = "NOT NULL"
		}

		var params []interface{}
		placeHolders := ""
{{- if (existsqlfilter $table $idxFields) }}
		if queryArgs.filterArgs != nil{
			pos := 0
			pls := make([]string, 0, len(queryArgs.filterArgs.filterPairs))
			for _, pair := range queryArgs.filterArgs.filterPairs {
				if pair.fieldName == "{{ .Field.Col.ColumnName }}"{
					return nil, fmt.Errorf("already have condition on field:{{ .Field.Name }}, because of foregin key {{ .Name }}")
				}
				pos++
				pls = append(pls, fmt.Sprintf("%s %s $%d", pair.fieldName, pair.option, pos))
				params = append(params, pair.value)
			}
			placeHolders = strings.Join(pls, " " + queryArgs.filterArgs.conjunction + " ")
			placeHolders = fmt.Sprintf("(%s) AND", placeHolders)
		}
{{- end }}
		params = append(params, {{ togqlname .Field.Name }})
		placeHolders = fmt.Sprintf("%s {{ .Field.Col.ColumnName }} = $%d AND ", placeHolders, len(params))

		params = append(params, *queryArgs.Offset)
		offsetPos := len(params)

		params = append(params, *queryArgs.Limit)
		limitPos := len(params)

		var sqlstr = fmt.Sprintf(
			`SELECT %s FROM %s WHERE %s deleted_date IS %s ORDER BY %s %s OFFSET $%d LIMIT $%d`,
			`{{ colnames $.Fields }} `,
			`"{{ $table }}"`,
			placeHolders,
			dead,
			"{{ $.PrimaryKey.Col.ColumnName }}",
			desc,
			offsetPos,
			limitPos)

		q, err := db.Query(sqlstr, params...)
		if err != nil {
			return nil, err
		}
		defer q.Close()

		// load results
		var res []*{{ $.Name }}
		for q.Next() {
			{{ $short }} := {{ $.Name }}{}

			// scan
			err = q.Scan({{ fieldnames $.Fields (print "&" $short) }})
			if err != nil {
				return nil, err
			}

			res = append(res, &{{ $short }})
		}

		return res, nil
	}

	func Count{{ $fnname }}(db XODB, {{ togqlname .Field.Name }} {{ .RefField.Type }}, queryArgs *{{ $.Name }}QueryArguments) (int, error) {
		queryArgs = Apply{{ $.Name }}QueryArgsDefaults(queryArgs)

		dead := "NULL"
		if *queryArgs.Dead {
			dead = "NOT NULL"
		}

		var params []interface{}
		placeHolders := ""
{{- if (existsqlfilter $table $idxFields) }}
		if queryArgs.filterArgs != nil{
			pos := 0
			pls := make([]string, 0, len(queryArgs.filterArgs.filterPairs))
			for _, pair := range queryArgs.filterArgs.filterPairs {
				if pair.fieldName == "{{ .Field.Col.ColumnName }}"{
					return -1, fmt.Errorf("already have condition on field:{{ .Field.Name }}, because of foregin key {{ .Name }}")
				}
				pos++
				pls = append(pls, fmt.Sprintf("%s %s $%d", pair.fieldName, pair.option, pos))
				params = append(params, pair.value)
			}
			placeHolders = strings.Join(pls, " " + queryArgs.filterArgs.conjunction + " ")
			placeHolders = fmt.Sprintf("(%s) AND", placeHolders)
		}
{{- end }}
		params = append(params, {{ togqlname .Field.Name }})
		placeHolders = fmt.Sprintf("%s {{ .Field.Col.ColumnName }} = $%d AND ", placeHolders, len(params))

		var err error
		var sqlstr = fmt.Sprintf(`SELECT count(*) from "{{ $table }}" WHERE %s deleted_date IS %s`, placeHolders, dead)
		XOLog(sqlstr)

		var count int
		err = db.QueryRow(sqlstr, params...).Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	{{end}}
{{ end}}

const graphQL{{ .Name }}Queries = `
{{- if (existsqlfilter $table $idxFields) }}
	all{{ plural .Name }}(where: {{ .Name }}Filter, offset: Int, limit: Int): {{ .Name }}Connection!
{{- else }}
	all{{ plural .Name }}(offset: Int, limit: Int): {{ .Name }}Connection!
{{- end -}}
{{- range $x, $index := .Indexes }}
	{{ togqlname .FuncName }}(
	{{- range $i, $field := .Fields }}
	  {{- togqlname .Name -}}:
	  {{- sqltogqltype .Type .Col.IsPrimaryKey -}}
	  {{- if not (islast $i (len $index.Fields)) -}}
	  ,
	  {{- end -}}
	{{- end -}}
	):
	{{- if not .Index.IsUnique }}[{{ end }}{{- $.Name }}{{ if not .Index.IsUnique }}!]{{ end }}
{{- end }}
`

const graphQL{{ .Name }}Mutations = `
	insert{{ plural .Name  }}(input: [Insert{{ .Name }}Input!]!): [{{ .Name }}!]!
	update{{ plural .Name  }}(input: [Update{{ .Name }}Input!]!): [{{ .Name }}!]!
	delete{{ plural .Name  }}(input: [Delete{{ .Name }}Input!]!): [ID!]!
`

func (RootResolver) All{{ plural .Name }}(ctx context.Context, args *{{ .Name }}QueryArguments) (*{{ .Name }}ConnectionResolver, error) {
	return All{{ plural .Name }}(ctx, args)
}
{{- range $x := .Indexes }}
// {{ .FuncName }} generated by {{ .Index.IndexName }}
func (RootResolver) {{ .FuncName }}(ctx context.Context, args struct{
	{{- range $i, $field := .Fields }}
	  {{ .Name }} {{ sqltogotype .Type .Col.IsPrimaryKey }}
	{{ end -}}
	}) ({{ if not .Index.IsUnique }}*[]{{ else }}*{{ end }}{{ .Type.Name }}Resolver, error) {
		return {{ .FuncName }}GraphQL(ctx, args)
	}
{{- end }}

func (RootResolver) Insert{{ plural .Name  }}(ctx context.Context, args struct{ Input []Insert{{ .Name }}Input }) ([]{{ .Name }}Resolver, error) {
	return Insert{{ .Name  }}GraphQL(ctx, args.Input)
}

func (RootResolver) Update{{ plural .Name  }}(ctx context.Context, args struct{ Input []Update{{ .Name }}Input }) ([]{{ .Name }}Resolver, error) {
	return Update{{ .Name  }}GraphQL(ctx, args.Input)
}

func (RootResolver) Delete{{ plural .Name  }}(ctx context.Context, args struct{ Input []Delete{{ .Name }}Input }) ([]graphql.ID, error) {
	return Delete{{ .Name  }}GraphQL(ctx, args.Input)
}

func Get{{ .Name }}Queries() string {
	return graphQL{{ .Name }}Queries
}

func Get{{ .Name }}Mutations() string {
	return graphQL{{ .Name }}Mutations
}

// GraphQL{{ .Name }}Types specifies the GraphQL types for {{ .Name }}
const GraphQL{{ .Name }}Types = `
	type {{ .Name }} {
{{- range .Fields }}
	{{- $field := . -}}
	{{- with (getforeignkey .Name $.ForeignKeys) }}
		{{ togqlname (fkname $field.Name) }}: {{ .RefType.Name }}
	{{- else }}
		{{ togqlname .Name }}: {{ sqltogqltype .Type .Col.IsPrimaryKey }}
	{{- end }}
{{- end }}

{{- range .RefFKs -}}
{{- if (existsqlfilter $table $idxFields) }}
		{{ togqlname .FkReverseField }}(where: {{ .Type.Name }}Filter, offset: Int, limit: Int): {{ .Type.Name }}Connection!
{{- else }}
		{{ togqlname .FkReverseField }}(offset: Int, limit: Int): {{ .Type.Name }}Connection!
{{- end -}}
{{- end -}}
{{- ""}}
	}

	type {{ .Name }}Connection {
		pageInfo: PageInfo!
		edges: [{{ .Name }}Edge]
		totalCount: Int
		{{ plural (togqlname .Name) }}: [{{ .Name }}]
	}

	type {{ .Name }}Edge {
		node: {{ .Name }}
		cursor: ID!
	}
{{ if (existsqlfilter $table $idxFields) }}
	input {{ .Name }}Filter {
		conjunction: FilterConjunction
{{- range .Fields -}}
	{{- $ftyp := (sqlfilter $table . $idxFields) -}}
	{{- if (and (ne .Name $.PrimaryKey.Name) (ne $ftyp "unsupported")) -}}
		{{- if (or (eq $ftyp "Number") (eq $ftyp "String")) }}
		{{ togqlname .Name  }}: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}
		{{- end -}}
		{{- if (eq $ftyp "String") }}
		{{ togqlname .Name }}_like: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }} // LIKE
		{{ togqlname .Name }}_ilike: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }} // LIKE case insensitive
		{{ togqlname .Name }}_nlike: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}	// NOT LIKE
		{{ togqlname .Name }}_nilike: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }} // NOT LIKE case insensitive
		{{- end -}}
		{{- if (or (eq $ftyp "Number") (eq $ftyp "Time")) }}
		{{ togqlname .Name  }}_lt: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}
		{{ togqlname .Name  }}_lte: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}
		{{ togqlname .Name  }}_gt: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}
		{{ togqlname .Name  }}_gte: {{ sqltogqloptionaltype .Type .Col.IsPrimaryKey }}
		{{- end -}}
	{{- end -}}
{{- end }}
	}
{{- end }}

	input Insert{{ .Name }}Input {
{{- range .Fields -}}
	{{- if ne .Name $.PrimaryKey.Name }}
		{{ togqlname .Name }}: {{ sqltogqltype .Type .Col.IsPrimaryKey }}
	{{- end -}}
{{- end }}
	}

	input Update{{ .Name }}Input {
{{- range .Fields }}
		{{ togqlname .Name }}: {{ sqltogqltype (sqlniltype .Type) .Col.IsPrimaryKey }}
{{- end }}
	}

	input Delete{{ .Name }}Input {
{{- range .Fields -}}
	{{- if eq .Name $.PrimaryKey.Name }}
		{{ togqlname .Name }}: {{ sqltogqltype .Type .Col.IsPrimaryKey }}
	{{- end -}}
{{- end }}
	}
`

// {{ .Name }}Resolver defines the GraphQL resolver for '{{ .Name }}'.
type {{ .Name }}Resolver struct { node *{{ .Name }} }

func New{{ .Name }}Resolver(node *{{ .Name }}) *{{ .Name }}Resolver {
	return &{{ .Name }}Resolver{ node: node }
}

{{- $typeName := .Name }}

{{- range .Fields -}}
	{{- $field := . -}}
	{{- with (getforeignkey .Name $.ForeignKeys) }}
		func (r {{ $typeName }}Resolver) {{ fkname $field.Name }}(ctx context.Context) (*{{ .RefType.Name }}Resolver, error) {
			db, ok := ctx.Value(DBCtx).(XODB)
			if !ok {
				return nil, errors.New("db is not found in context")
			}
			{{- $ot := .RefField.Type -}}
			{{- $it := $field.Type -}}
			{{- $varname := ( togqlname $field.Name ) -}}
			{{- if (eq $it $ot) }}
				{{ $varname }} := r.node.{{$field.Name}}
			{{- else if (eq $it "int64") }}
				{{ $varname }} := ({{ $ot }})(r.node.{{$field.Name}})
			{{- else if (eq $it "sql.NullInt64") }}
				if !r.node.{{$field.Name}}.Valid {
					return nil, nil
				}
				{{ $varname }} := ({{ $ot }})(r.node.{{$field.Name}}.Int64)
			{{- else }}
				panic("TODO: implement in postgres.type.go.tpl {{ printf "input: %s, output %s" $it $ot }}")
			{{- end }}
			node, err := {{ .RefType.Name }}By{{ .RefField.Name }}(db, {{$varname}})
			if err != nil {
				return nil, errors.Wrap(err, "unable to retrieve {{ fkname $field.Name }}")
			}
			return &{{ .RefType.Name }}Resolver{node: node}, nil
		}
	{{- else }}
 		func (r {{ $typeName }}Resolver) {{ .Name }}() {{ sqltogotype .Type .Col.IsPrimaryKey }} { return {{ sqltogql .Type (print "r.node." .Name) .Col.IsPrimaryKey }} }
	{{- end }}
{{- end }}

{{- range .RefFKs }}
func (r {{ $typeName }}Resolver) {{ .FkReverseField }}(ctx context.Context, queryArgs *{{ .Type.Name }}QueryArguments) (*{{ .Type.Name }}ConnectionResolver, error){
	db, ok := ctx.Value(DBCtx).(XODB)
	if !ok {
		return nil, errors.New("db is not found in context")
	}

	if queryArgs != nil && (queryArgs.After != nil || queryArgs.First != nil || queryArgs.Before != nil || queryArgs.Last != nil) {
		return nil, errors.New("not implemented yet, use offset + limit for pagination")
	}

	queryArgs = Apply{{ .Type.Name }}QueryArgsDefaults(queryArgs)
{{ if (existsqlfilter $table $idxFields) }}
	filterArgs, err := get{{ .Type.Name }}Filter(queryArgs.Where)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{ .Type.Name }} filter")
	}
	queryArgs.filterArgs = filterArgs
{{ end }}
	{{ $varname := (togqlname .RefField.Name) -}}
	{{ $varname }} := r.node.{{.RefField.Name}}

	data, err := {{ plural .Type.Name }}By{{.Field.Name}}FK(db, {{ $varname }}, queryArgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{plural .Type.Name}}")
	}

	count, err := Count{{ plural .Type.Name }}By{{.Field.Name}}FK(db, {{ $varname }}, queryArgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{plural .Type.Name}} count")
	}

	return &{{.Type.Name}}ConnectionResolver{
		data: data,
		count: int32(count),
	}, nil
}
{{- end }}


// {{ .Name }}ConnectionResolver defines a GraphQL resolver for {{ .Name }}Connection
type {{ .Name }}ConnectionResolver struct {
	data  []*{{ .Name }}
	count int32
}

// PageInfo returns PageInfo
func (r {{ .Name }}ConnectionResolver) PageInfo() *PageInfoResolver {
	if len(r.data) == 0 {
		return nil
	}

	return &PageInfoResolver{
		startCursor:     encodeCursor("{{ .Name }}", int(r.data[0].{{ .PrimaryKey.Name }})),
		endCursor:       encodeCursor("{{ .Name }}", int(r.data[len(r.data)-1].{{ .PrimaryKey.Name }})),
		hasNextPage:     false, // TODO
		hasPreviousPage: false, // TODO
	}
}

// Edges returns standard GraphQL edges
func (r {{ .Name }}ConnectionResolver) Edges() *[]*{{ .Name }}EdgeResolver {
	edges := make([]*{{ .Name }}EdgeResolver, len(r.data))

	for i := range r.data {
		edges[i] = &{{ .Name }}EdgeResolver{node: r.data[i]}
	}
	return &edges
}

// TotalCount returns total count
func (r {{ .Name }}ConnectionResolver) TotalCount() *int32 {
	return &r.count
}

// {{ plural .Name }} returns the list of {{ .Name }}
func (r {{ .Name }}ConnectionResolver) {{ plural .Name }}() *[]*{{ .Name }}Resolver {
	data := make([]*{{ .Name }}Resolver, len(r.data))
	for i := range r.data {
		data[i] = &{{ .Name }}Resolver{r.data[i]}
	}
	return &data
}

// {{ .Name }}EdgeResolver defines the {{ .Name }} edge
type {{ .Name }}EdgeResolver struct {
	node *{{ .Name }}
}

// Node returns the {{ .Name }} node
func (r {{ .Name }}EdgeResolver) Node() *{{ .Name }}Resolver {
	return &{{ .Name }}Resolver{r.node}
}

// Cursor returns the cursor
func (r {{ .Name }}EdgeResolver) Cursor() graphql.ID {
	return encodeCursor("{{ .Name }}", int(r.node.{{ .PrimaryKey.Name }}))
}

// Insert{{ .Name }}Input defines the insert {{ .Name }} mutation input
type Insert{{ .Name }}Input struct {
{{- range .Fields -}}
	{{- if ne .Name $.PrimaryKey.Name }}
		{{ .Name }} {{ sqltogotype .Type .Col.IsPrimaryKey }}
	{{- end -}}
{{- end }}
}

// Update{{ .Name }}Input defines the update {{ .Name }} mutation input
type Update{{ .Name }}Input struct {
{{- range .Fields }}
		{{ .Name }} {{ sqltogotype (sqlniltype .Type) .Col.IsPrimaryKey }}
{{- end }}
}

// Delete{{ .Name }}Input defines the delete {{ .Name }} mutation input
type Delete{{ .Name }}Input struct {
{{- range .Fields -}}
	{{- if eq .Name $.PrimaryKey.Name }}
		{{ .Name }} {{ sqltogotype .Type .Col.IsPrimaryKey }}
	{{- end -}}
{{- end }}
}

// All{{ plural .Name }} is the GraphQL end point for GetAll{{ .Name }}
func All{{ plural .Name }}(ctx context.Context, queryArgs *{{ .Name }}QueryArguments) (*{{ .Name }}ConnectionResolver, error) { // nolint: gocyclo
	db, ok := ctx.Value(DBCtx).(XODB)
	if !ok {
		return nil, errors.New("db is not found in context")
	}

	if queryArgs != nil && (queryArgs.After != nil || queryArgs.First != nil || queryArgs.Before != nil || queryArgs.Last != nil) {
		return nil, errors.New("not implemented yet, use offset + limit for pagination")
	}

	queryArgs = Apply{{ .Name }}QueryArgsDefaults(queryArgs)
{{ if (existsqlfilter $table $idxFields) }}
	filterArgs, err := get{{ .Name }}Filter(queryArgs.Where)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{ .Name }} filter")
	}
	queryArgs.filterArgs = filterArgs
{{ end }}
	all{{ .Name }}, err := GetAll{{ .Name }}(db, queryArgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{ .Name }}")
	}

	count, err := CountAll{{ .Name }}(db, queryArgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get count")
	}

	return &{{ .Name }}ConnectionResolver{
		data:  all{{ .Name }},
		count: int32(count),
	}, nil
}

// Insert{{ .Name }}GraphQL is the GraphQL end point for Insert{{ .Name }}
func Insert{{ .Name }}GraphQL(ctx context.Context, items []Insert{{ .Name }}Input) ([]{{ .Name }}Resolver, error) {
	db, ok := ctx.Value(DBCtx).(XODB)
	if !ok {
		return nil, errors.New("db is not found in context")
	}
	results := make([]{{ .Name }}Resolver, len(items))
	for i := range items {
		input := items[i]
		{{ range $index, $field := .Fields -}}
			{{ $it := (sqltogotype .Type .Col.IsPrimaryKey) }}
			{{ if .Col.IsPrimaryKey }}
				{{/* primary key column skipped */}}
			{{ else if (eq $it "graphql.ID") -}}
				{{ print "f" $index }}, err := strconv.Atoi(string({{ gotosql .Type (print "input." .Name) }}))
				if err != nil {
					return nil, errors.New("{{ .Name }} must be an integer")
				}
			{{- else if (and (eq $it "*string") (eq .Type "sql.NullInt64")) -}}
				var {{ print "f" $index }} sql.NullInt64
				if {{ print "input." .Name }} != nil {
					n, err := strconv.ParseInt(*{{ print "input." .Name }}, 10, 0)
					if err != nil {
						return nil, errors.New("{{ .Name }} must be an integer")
					}
					{{ print "f" $index }} = sql.NullInt64{Int64: n, Valid: true}
				}
			{{- else if (and (eq $it "string") (eq .Type "int64")) -}}
				{{ print "f" $index }}, err := strconv.ParseInt({{ gotosql .Type (print "input." .Name) }}, 10, 0)
				if err != nil {
					return nil, errors.New("{{ .Name }} must be an integer")
				}
			{{- else -}}
				{{ print "f" $index }} := {{ gotosql .Type (print "input." .Name) }}
			{{- end -}}
		{{ end }}
		node:= &{{ .Name }}{
			{{- range $index, $field := .Fields -}}
				{{- if ne .Name $.PrimaryKey.Name -}}
					{{ .Name }}: {{ print "f" $index }},
				{{- end }}
			{{ end }}
		}
		if err := node.Insert(db); err != nil {
			return nil, errors.Wrap(err, "unable to insert {{ .Name }}")
		}
		results[i] = {{ .Name }}Resolver{ node: node }
	}
	return results, nil
}

// Update{{ .Name }}GraphQL is the GraphQL end point for Update{{ .Name }}
func Update{{ .Name }}GraphQL(ctx context.Context, items []Update{{ .Name }}Input) ([]{{ .Name }}Resolver, error) {
	db, ok := ctx.Value(DBCtx).(XODB)
	if !ok {
		return nil, errors.New("db is not found in context")
	}
	results := make([]{{ .Name }}Resolver, len(items))
	for i := range items {
		input := items[i]
		{{ gqlupdate . }}
		results[i] = {{ .Name }}Resolver{ node: node }
	}
	return results, nil
}

// Delete{{ .Name }}GraphQL is the GraphQL end point for Delete{{ .Name }}
func Delete{{ .Name }}GraphQL(ctx context.Context, items []Delete{{ .Name }}Input) ([]graphql.ID, error) {
	db, ok := ctx.Value(DBCtx).(XODB)
	if !ok {
		return nil, errors.New("db is not found in context")
	}
	// sql query
	const sqlstr = `DELETE FROM "{{ .Table.TableName }}" WHERE id = $1`

	results := make([]graphql.ID, len(items))

	for i := range items {
		input := items[i]
		{{ range $index, $field := .Fields -}}
			{{ $it := (sqltogotype .Type .Col.IsPrimaryKey) }}
			{{ if (not .Col.IsPrimaryKey) }}
			{{ else if (and (eq $it "graphql.ID") (eq .Type "int")) -}}
				id, err := strconv.Atoi(string({{ gotosql .Type (print "input." .Name) }}))
				if err != nil {
					return nil, errors.New("{{ .Name }} must be an integer")
				}
			{{- else if (and (eq $it "graphql.ID") (eq .Type "int64")) -}}
				id, err := strconv.ParseInt(string({{ gotosql .Type (print "input." .Name) }}), 10, 0)
				if err != nil {
					return nil, errors.New("{{ .Name }} must be an integer")
				}
			{{- else if (and (eq $it "*string") (eq .Type "sql.NullInt64")) -}}
				var id sql.NullInt64
				if {{ print "input." .Name }} != nil {
					n, err := strconv.ParseInt(*{{ print "input." .Name }}, 10, 0)
					if err != nil {
						return nil, errors.New("{{ .Name }} must be an integer")
					}
					id = sql.NullInt64{Int64: n, Valid: true}
				}
			{{- else if (and (eq $it "string") (eq .Type "int64")) -}}
				id, err := strconv.ParseInt({{ gotosql .Type (print "input." .Name) }}, 10, 0)
				if err != nil {
					return nil, errors.New("{{ .Name }} must be an integer")
				}
			{{- else -}}
				id := {{ gotosql .Type (print "input." .Name) }}
			{{- end -}}
		{{ end }}

		XOLog(sqlstr, id)
		_, err = db.Exec(sqlstr, id)
		if err != nil {
			return nil, err
		}
		results[i] = input.{{ .PrimaryKey.Name }}
	}
	return results, nil
}
