module example.com/hello

go 1.24.0

// Optional: submodules
replace (
	example.com/hello => ./helloworld
	example.com/integers => ./integers
)

require github.com/google/uuid v1.6.0 // indirect
