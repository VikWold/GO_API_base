# Set up HTMX in a project

## Files
We will make a ui folder where we are going to put our html files and templates.

As we make the ui folder the first file we will make here is the index.html.
This is the file that have ouer base for ouer project and we will then use later to server the page.
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <script src="https://unpkg.com/htmx.org@2.0.0"></script>
    <title>Go App</title>
</head>
<body>
    <h1>Welcome</h1>
    <div id="display-area">
        <!-- Content will be swapped here -->
    </div>
    
    <button hx-get="/clicked" hx-target="#display-area">
        Click to Update
    </button>
</body>
</html> 
```

In the code we have a spot that will display the content we have fetched.  


## Code changes

### main.go
We are making a change to our application struct where we are adding the field `tmpl`
So we are left with this:

```go 
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	tmpl     *template.Template // newly added
}
```
Moving on from the struct, we also need to edit our main function to have this so we
can serve the HTMX to the frontend.

```go
	// added this to parse the templates
	tmpl, err := template.ParseGlob("./ui/html/*.html")
	if err != nil {
		logger.Error("unable to parse templates", "error", err)
		os.Exit(1)
	}

	app := &application{
		logger: instanceLogger,
		models: data.NewModels(db, &queryTimeout),
		tmpl:   tmpl, // added to initialize the tmpl field in the application struct
	}
```

### routes.go
We need to add a new route were we will make the backend send the htmx page to the client asking for it.
```go
	mux.HandleFunc("/htmx", app.htmx)
	mux.HandleFunc("GET /clicked", app.clickedHandler)
```

### utilitieshandlers.go
We need to add new handlers that server the page to the client asking for it.
This file is were we will place those functions that we can use to achieve this.

Here you see the function that serves the page to the client asking for it.

```go
func (app *application) htmxHandler(w http.ResponseWriter, r *http.Request) {
	err := app.tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		app.logger.Error("template error", "error", err)
	}
}
```