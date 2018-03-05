{{- $short := (shortname .Type.Name "err" "sqlstr" "db" "q" "res" "XOLog" .Fields) -}}
{{- $table := (schema .Type.Table.TableName) }}

// {{ .FuncName }} retrieves a row from '{{ $table }}' as a {{ .Type.Name }}.
//
// Generated from index '{{ .Index.IndexName }}'.
func {{ .FuncName }}(db XODB{{ goparamlist .Fields true true }}) ({{ if not .Index.IsUnique }}[]{{ end }}*{{ .Type.Name }}, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`{{ colnames .Type.Fields }} ` +
		`FROM "{{ $table }}" ` +
		`WHERE {{ colnamesquery .Fields " AND " }}`

	// run query
	XOLog(sqlstr{{ goparamlist .Fields true false }})
{{- if .Index.IsUnique }}
	{{ $short }} := {{ .Type.Name }}{
	{{- if .Type.PrimaryKey }}
		_exists: true,
	{{ end -}}
	}

	err = db.QueryRow(sqlstr{{ goparamlist .Fields true false }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
	if err != nil {
		return nil, err
	}

	return &{{ $short }}, nil
{{- else }}
	q, err := db.Query(sqlstr{{ goparamlist .Fields true false }})
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*{{ .Type.Name }}{}
	for q.Next() {
		{{ $short }} := {{ .Type.Name }}{
		{{- if .Type.PrimaryKey }}
			_exists: true,
		{{ end -}}
		}

		// scan
		err = q.Scan({{ fieldnames .Type.Fields (print "&" $short) }})
		if err != nil {
			return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
{{- end }}
}

// {{ .FuncName }}GraphQL retrieves a row from '{{ $table }}' as a {{ .Type.Name }}.
//
// Generated from index '{{ .Index.IndexName }}'.
func {{ .FuncName }}GraphQL(ctx context.Context, args struct{ 
{{- range .Fields }}
	{{ .Name }} {{ sqltogotype .Type .Col.IsPrimaryKey }}
{{- end -}}
	}) ({{ if not .Index.IsUnique }}*[]{{ else }}*{{ end }}{{ .Type.Name }}Resolver, error) {
	db, ok := ctx.Value(DBCtx).(XODB)
	if !(ok) {
		return nil, errors.New("db is not found in context")
	}

{{ range $index, $field := .Fields }}
{{ if eq .Type "int" -}}
	arg{{ $index }}, err := strconv.Atoi(string(args.{{ .Name }}))
	if err != nil {
		return nil, errors.Wrap(err, "{{ .Name }} should be integer")
	}
{{ else if eq .Type "sql.NullString" -}}
	arg{{ $index }} := StringPointer(args.{{ .Name }})
{{ else if eq .Type "int64" -}}
	arg{{ $index }}, err := strconv.ParseInt(string(args.{{ .Name }}), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "{{ .Name }} should be int64")
	}
{{ else if eq .Type "string" -}}
	arg{{ $index }} := args.{{ .Name }}
{{ else if eq .Type "sql.NullInt64" -}}
	if args.{{.Name}} == nil {
		return nil, nil
	}
	n, err := strconv.ParseInt(*args.{{.Name}}, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "{{ .Name }} should be int64")
	}
	arg{{ $index }} := sql.NullInt64{Int64: n, Valid: true}
{{ else }}
	panic("fix me in xo template postgres.index.go.tpl for {{.Type}}")
{{- end }}
{{ end }}

	data, err := {{ .FuncName }}(db, 
	{{- range $index, $field := .Fields -}}
		arg{{ $index }},
	{{- end -}})
	if err != nil {
		return nil, errors.Wrap(err, "unable to get {{ $table }}")
	}

{{ if .Index.IsUnique }}
	return &{{ .Type.Name }}Resolver{ node: data }, nil
{{ else }}
	ret := make([]{{ .Type.Name }}Resolver, len(data))
	for i, row := range data {
		ret[i] = {{ .Type.Name }}Resolver{ node: row }
	}
	return &ret, nil
{{ end }}
}
