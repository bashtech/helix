# Chat Documentation

## Get Channel Chat Badges

This is an example of how to get channel chat badges

```go
client, err := helix.NewClient(&helix.Options{
    ClientID: "your-client-id",
})
if err != nil {
    // handle error
}

resp, err := client.GetChannelChatBadges(&helix.GetChatBadgeParams{
    BroadcasterID: "145328278",
})
if err != nil {
    // handle error
}

fmt.Printf("%+v\n", resp)
```

## Get Global Chat Badges

This is an example of how to get global chat badges

```go
client, err := helix.NewClient(&helix.Options{
    ClientID: "your-client-id",
})
if err != nil {
    // handle error
}

resp, err := client.GetGlobalChatBadges()
if err != nil {
    // handle error
}

fmt.Printf("%+v\n", resp)
```
