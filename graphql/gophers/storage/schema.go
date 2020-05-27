package storage

const (
	SchemaString = `
	schema {
		query: Query
	}

    type Query {
	  # Get a thing by it's ID.
	  user(id: ID!): User
	  users(): [User]!
	}

	# Thing is pretty simple.
	type User {
	  id: ID
	  name: String
	  fullname: String
	  firends(): [Firend]!
	}

	type Firend {
	  firendId: String
	  user(): User
	}
`
)

//firends(): String
//
//type Firend {
//}
