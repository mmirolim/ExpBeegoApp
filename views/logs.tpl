<ul>
{{ range $log := .logs }}
   <li>{{ $log }}</li>
{{ end }}
</ul>