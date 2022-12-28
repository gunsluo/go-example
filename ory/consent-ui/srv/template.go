package srv

var loginHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1 id="login-title">Please log in</h1>
  <form action="login" method="POST">
  		<input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
		<input type="hidden" name="challenge" value="{{ .Challenge }}">
    <table>
      <tr>
        <td><input type="email" id="email" name="email" value="email@foobar.com" placeholder="email@foobar.com"></td>
        <td>(it's "foo@bar.com")</td>
      </tr>
      <tr>
        <td><input type="password" id="password" name="password"></td>
        <td>(it's "foobar")</td>
      </tr>
    </table><input type="checkbox" id="remember" name="remember" value="1"><label for="remember">Remember
      me</label><br>
	  <input type="submit" id="accept" name="submit" value="Log in">
	  <input type="submit" id="reject" name="submit" value="Deny access">
  </form>
</body>

</html>
`

var consentHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1>An application requests access to your data!</h1>
  <form action="consent" method="POST">
    <input type="hidden" name="challenge" value="{{ .Challenge }}">
    <input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
    <p>Hi foo@bar.com, application <strong>auth-code-client</strong> wants access resources on your behalf and to:
    </p>
    <input class="grant_scope" type="checkbox" id="openid" value="openid" name="grant_scope">
    <label for="openid">openid</label><br>
    <input class="grant_scope" type="checkbox" id="offline" value="offline" name="grant_scope">
    <label for="offline">offline</label><br>
    <p>Do you want to be asked next time when this application wants to access your data? The application will
      not be able to ask for more permissions without your consent.</p>
    <ul></ul>
    <p><input type="checkbox" id="remember" name="remember" value="1"><label for="remember">Do not ask me again</label>
    </p>
    <p><input type="submit" id="accept" name="submit" value="Allow access"><input type="submit" id="reject"
        name="submit" value="Deny access"></p>
  </form>
</body>

</html>
`

var logoutHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1>Do you wish to log out?</h1>
  <form action="logout" method="POST">
    <input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
    <input type="hidden" name="challenge" value="{{ .Challenge }}">
    <input type="submit" id="accept" value="Yes"><input type="submit" id="reject" value="No">
  </form>
</body>

</html>`
