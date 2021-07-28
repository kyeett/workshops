
## Creating the app

Setup package manager for the app (similar to npm, or crate).
```
go mod init
```


For this workshop, we will use github.com/go-chi/chi. There are many other frameworks, in Go. See [List of Go web Frameworks](https://github.com/mingrammer/go-web-framework-stars)
```
go get -u github.com/go-chi/chi/v5
```

```
cat go.mod
```

### Create main.go
Open `main.go` in your favorite editor, and copy the example code from [github.com/go-chi/chi](https://github.com/go-chi/chi)
```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}
```

Start and build the app
```
go run .
```


# (Advanced) Deploy

```
YOUR_APP=magnus-app
PROJECT=$(gcloud config get-value project)
TAG=eu.gcr.io/$PROJECT/workshop/$YOUR_APP
gcloud config set run/platform managed
gcloud config set run/region europe-west1
gcloud builds submit --tag $TAG
```

```
gcloud run deploy $YOUR_APP --image $TAG --platform=managed --region=europe-west1 --allow-unauthenticated
```
