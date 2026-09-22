# Snippetbox Documentation

This project is a basic web application built using Go's `net/http` standard library.

## 1. Routing with `http.ServeMux`

The application uses a **ServeMux** (HTTP request multiplexer) to match incoming request URLs against a list of registered patterns.

```go main.go
// ... existing code ...
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
    // ... rest of routes ...
}
```

**Note:** In Go, the pattern `"/"` acts as a catch-all. This is why the `home` handler checks `r.URL.Path != "/"` to manually return a 404 for paths that don't match exactly.

## 2. Handlers and Responses

Handlers are responsible for processing requests and writing responses via `http.ResponseWriter`.

### TDD Perspective: Testing Handlers

Before writing the `home` handler logic, a TDD approach suggests defining the expected behavior:

1. **Requirement:** A GET request to `/` should return status `200 OK`.
2. **Test:** Use the `net/http/httptest` package to record the response and assert the status code.

### Parsing Query Parameters

In `snippetView`, we extract data from the URL:

- `r.URL.Query().Get("id")`: Retrieves the value from the query string.
- `strconv.Atoi()`: Converts the string to an integer.

## 3. Restricting HTTP Methods

For the `snippetCreate` handler, we enforce the `POST` method.

```go main.go
// ... inside snippetCreate ...
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
// ...
```

**Best Practice:**

- **Headers first:** Always call `w.Header().Set()` before `w.WriteHeader()` or `w.Write()`.
- **Return early:** After sending an error with `http.Error()`, you must `return` to stop the handler execution.

## 4. Starting the Server

The server is started using `http.ListenAndServe`.

```go main.go
// ...
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
```

**Error Handling Note:** `log.Fatal(err)` is used here because `ListenAndServe` only returns an error if the server fails to start or crashes (e.g., the port is already in use). It always returns a non-nil error.
