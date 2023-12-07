fooApp/ // The root directory of the project
  circle.yml // A configuration file for CircleCI
  Dockerfile // A file to build a Docker image for the project
  cmd/ // A directory to store the main application entry point files
    foosrv/ // A directory for the foosrv binary
      main.go // The main file for the foosrv binary
    foocli/ // A directory for the foocli binary
      main.go // The main file for the foocli binary
  pkg/ // A directory to store code that is OK for other services to consume
    fs/ // A package for file system operations
      fs.go // The main file for the fs package
      fs_test.go // The test file for the fs package
      mock.go // A file to provide mock implementations for the fs package
      mock_test.go // A test file for the mock implementations
    merge/ // A package for merging data
      merge.go // The main file for the merge package
      merge_test.go // The test file for the merge package
    api/ // A package for API clients
      api.go // The main file for the api package
      api_test.go // The test file for the api package
  internal/ // A directory to store code that is specific to the function of the service and not shared with other services
    auth/ // A package for authentication
      auth.go // The main file for the auth package
      auth_test.go // The test file for the auth package
  serverlib/ // A directory to store code that is used by the server binary
    lib.go // The main file for the serverlib package
  go.mod // A file to define the module path and dependency requirements
  go.sum // A file to store the checksums of the dependencies
  README.md // A file to provide documentation for the project
  LICENSE // A file to specify the license of the project
