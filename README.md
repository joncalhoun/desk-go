# desk-go
Go client for the Desk.com API


## Example usage

```go
package main

import (
    "github.com/joncalhoun/desk"
    "github.com/joncalhoun/desk/kase"
    "fmt"
)

func main() {
    desk.Subdomain = "your-subdomain"
    desk.Username = "your@email.com"
    desk.Password = "your-password"

    // Look up a single case, its message, replies, and notes.
    id := 3
    c, err := kase.Get(id) // returns a pointer to a desk.Case
    if err != nil {
        panic(err)
    }
    kase.GetMessage(c)
    kase.GetReplies(c)
    kase.GetNotes(c)
    fmt.Printf("%+v", c)

    // Look up a list of cases, and get a []desk.Case. These will *not* have
    // their message, replies, or notes by default due to the design of Desk's API
    ps := desk.CaseListParams{
        // id, created_at, priority, received_at, status, updated_at. Defaults to created_at
        SortField: "created_at",

        // asc, desc. Defaults to asc
        SortDirection: "desc",

        // 1 - 500. Use Search() if you need higher pages. Defaults to 1
        Page: 1,

        // Limit not defined in API docs for Cases.
        PerPage: 10,
    }
    cs, err := kase.List(&ps)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v", cs)
}
```
