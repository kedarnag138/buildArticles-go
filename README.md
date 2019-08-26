1. To setup the application you'll need to have installed Go. You have install using GVM or just download the binary.
2. Set the GOPATH and GOROOT.
3. Clone the application into the src/ directory
4. GoDep has been setup to maintain all dependencies. Run `go get -v` to install all dependencies.
5. Please setup the respective databases(local and test), you can find in `api/v1/config/environment_setup.go`.
6. After this, run the migrations, using the following command `goose postgres "user=your_username dbname=build_articles_test sslmode=disable" up`
6. To run the application, run `go run main.go development`
7. To run the test cases, run `go test -v ./...` from the root directory of the project.

The following are the endpoints for the API.

1. To create an Article

Method - POST
URL - http://localhost:8080
Endpoint - /api/v1/articles
Headers - { "Content-Type": "application/json" }
Body - {"title": "Test Title", "author": "Test Author", "content": "Test Content"}

2. To list all articles

Method - GET
URL - http://localhost:8080
Endpoint - /api/v1/articles

3. To get a article by ID

Method - GET
URL - http://localhost:8080
Endpoint - /api/v1/articles/:id