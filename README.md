# go-shippable #

go-shippable is a Go client library for accessing the [Shippable API][].

[![Build Status](https://api.shippable.com/projects/556734d0edd7f2c052ff35b4/badge?branchName=master)](https://app.shippable.com/projects/556734d0edd7f2c052ff35b4/builds/latest)
[![Coverage Status](http://gocover.io/_badge/github.com/PeoplePerHour/go-shippable?upd=1)](http://gocover.io/github.com/PeoplePerHour/go-shippable)
[![GoDoc](https://godoc.org/github.com/PeoplePerHour/go-shippable?status.svg)](https://godoc.org/github.com/PeoplePerHour/go-shippable)

go-shippable requires Go version 1.1 or greater.

## Usage ##

```go
import "github.com/PeoplePerHour/go-shippable"
```

Create a new Shippable API client by passing an authentication token (you can create
a token from your Shippable account settings if you are on a subscription plan)

```go
token := "mytoken"
client := shippable.NewClient(token)
projects, _, err := client.Projects.GetProjects()
```

## Documentation ##

For complete usage of go-shippable, see the full [package docs][].

[Shippable API]: http://docs.shippable.com/api/
[package docs]: https://godoc.org/github.com/PeoplePerHour/go-shippable


## Roadmap ##

This library is being initially developed for a DevOps CI/CD internal application at
[PeoplePerHour][] and [SuperTasker][], so features 
are implemented as needed by those projects.

Contributions are more than welcomed. PRs bundled with tests will get merged ASAP.

[PeoplePerHour]: https://www.peopleperhour.com
[SuperTasker]: https://www.supertasker.com


## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
