# desk-go
Go client for the Desk.com API


## Example usage

```go
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
cs, err := kase.List(1, 1)
if err != nil {
  panic(err)
}
fmt.Printf("%+v", cs)
```
