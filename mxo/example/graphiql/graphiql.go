package graphiql

// GraphiQLPage the page used to query the graph API
var GraphiQLPage = []byte(`<!--
	*  Copyright (c) Facebook, Inc.
	*  All rights reserved.
	*
	*  This source code is licensed under the license found in the
	*  LICENSE file in the root directory of this source tree.
	-->
	<!DOCTYPE html>
	<html>
	
	<head>
		<meta charset="UTF-8">
		<style>
			body {
				height: 100%;
				margin: 0;
				width: 100%;
				overflow: hidden;
			}
	
			#graphiql {
				height: 100vh;
			}
		</style>
	
		<!--
		 This GraphiQL example depends on Promise and fetch, which are available in
		 modern browsers, but can be "polyfilled" for older browsers.
		 GraphiQL itself depends on React DOM.
		 If you do not want to rely on a CDN, you can host these files locally or
		 include them directly in your favored resource bunder.
	   -->
		<script src="//cdn.jsdelivr.net/es6-promise/4.0.5/es6-promise.auto.min.js"></script>
		<script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react-dom.min.js"></script>
	
		<!--
		 These two files can be found in the npm module, however you may wish to
		 copy them directly into your environment, or perhaps include them in your
		 favored resource bundler.
		-->
		<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.3/graphiql.min.css" />
		<script src="//cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.3/graphiql.min.js"></script>
	
	</head>
	
	<body>
		<div id="graphiql">Loading...</div>
		<script>
let parameters = {}
function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2)
  	return parts.pop().split(";").shift();
}
// When the query and variables string is edited, update the URL bar so
// that it can be easily shared
function onEditQuery(newQuery) {
  parameters.query = newQuery
  updateURL()
}
function onEditVariables(newVariables) {
  parameters.variables = newVariables
  updateURL()
}
function onEditOperationName(newOperationName) {
  parameters.operationName = newOperationName
  updateURL()
}
function updateURL() {}
// Defines a GraphQL fetcher using the fetch API. You're not required to
// use fetch, and could instead implement graphQLFetcher however you like,
// as long as it returns a Promise or Observable.
function graphQLFetcher(graphQLParams) {
  // This example expects a GraphQL server at the path /graphql.
  // Change this to point wherever you host your GraphQL server.

  let headers = {
    Accept: "application/json",
    "Content-Type": "application/json",
  }
  let token = window.localStorage.getItem("access_token")

  if (token) {
    headers["Authorization"] = "Bearer " + token
  }

  return fetch('/', {
    method: "post",
    headers,
    body: JSON.stringify(graphQLParams),
    credentials: "include",
  })
    .then(response => {
      return response.text()
    })
    .then(responseBody => {
      try {
        let d = JSON.parse(responseBody)
        if (
          d &&
          d.data &&
          d.data.logIn &&
          d.data.logIn.accessToken &&
          d.data.logIn.accessToken.token
        ) {
          window.localStorage.setItem(
            "access_token",
            d.data.logIn.accessToken.token,
          )
        }
        return d
      } catch (error) {
        return responseBody
      }
    })
}
GraphiQL.Logo = React.createClass({
  updateAccessToken(accessToken) {
    window.localStorage.setItem("access_token", accessToken)
    this.setState({ accessToken })
  },
  getInitialState() {
    const accessToken =
      getCookie("__Secure-id_token") ||
      getCookie("__Secure-access_token") ||
      window.localStorage.getItem("access_token") ||
      ""
    return { accessToken }
  },
  render() {
    return React.createElement(
      "div",
      {
        style: { display: "flex", flexDirection: "row", alignItems: "center" },
      },
      React.createElement(
        "div",
        { className: "title" },
        React.createElement(
          "span",
          null,
          "Graph",
          React.createElement("em", null, "i"),
          "QL",
        ),
      ),
      React.createElement("input", {
        type: "text",
        style: { marginLeft: "10px", width: "100px" },
        placeholder: "Access Token",
        value: this.state.accessToken,
        onChange: function(e) {
          this.updateAccessToken(e.target.value)
        }.bind(this),
      }),
    )
  },
})
// Render <GraphiQL /> into the body.
// See the README in the top level of this module to learn more about
// how you can customize GraphiQL by providing different values or
// additional child elements.
ReactDOM.render(
  React.createElement(GraphiQL, {
    fetcher: graphQLFetcher,
    query: parameters.query,
    variables: parameters.variables,
    operationName: parameters.operationName,
    onEditQuery,
    onEditVariables,
    onEditOperationName,
  }),
  document.getElementById("graphiql"),
)

		</script>
	</body>
	
	</html>`)
