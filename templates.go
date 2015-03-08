package greeting

onst placesTemplateHTML = `
<html>
  <body>
    <table style="width:100%">
      {{range .}}
      <tr>
      	<td>{{.Name}}</td>
      	<td>{{.Vicinity}}</td>
	<td>{{.Types}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>
`
