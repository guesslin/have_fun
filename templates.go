package greeting

const AddresesForm = `
<html>
  <body>
    <form action="/inplaces" method="post">
      <div><textarea name="address1" rows="3" cols="60"></textarea></div>
      <div><textarea name="address2" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Find FUNs"></div>
    </form>
  </body>
</html>
`

const placesTemplateHTML = `
<html>
  <body>
    <p>First Address is</p>
    <pre>{{.}}</pre>
  </body>
</html>
`
