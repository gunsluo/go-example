package storage

const (
	SchemaString = `
	schema {
		query: Query
	}

    type Query {
	  # Get a thing by it's ID.
	  user(id: ID!): User
	  users(limit: Int!, offset: Int!): [User]!
	}

	# Thing is pretty simple.
	type User {
	  id: ID
	  name: String
	  fullname: String
	  friends(): [Friend]!
	}

	type Friend {
	  friendId: String
	  user(): User
	}
`
)
