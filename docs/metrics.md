# METRICS

## This handles the metrics and how many times the html page was loaded

### GET /admin/metrics

this gives an html page that represents how many times the main html was loaded

##### response:

```html

<html>

<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited %dtimes!</p>
</body>

</html>

```

### GET /api/reset

this resets the counts

##### response

```

"Hits reset to 0"

```