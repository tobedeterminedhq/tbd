SELECT *
FROM {{ .Model }}
WHERE {{ .Column }} < {{ .Value }}
