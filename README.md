# go-env

Environment variable module. 

Variable expansion allows overriding the get and set methods and includes
setting defaults with `${VALUE:-default}`and returning errors with `${VALUE?required}`

The Getx, Setx, Deletex are for working with windows environment variables
where you wish to set user or machine level variables.

There are functions for dealing with the PATH variable such as PrependPath
and AppendPath. 

MIT License