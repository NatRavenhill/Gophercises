module Gophercises/cli

replace cli/cmd => ./cmd

replace cli/database => ./database

go 1.16

require (
	cli/cmd v0.0.0-00010101000000-000000000000
	cli/database v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.2.1
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
)
