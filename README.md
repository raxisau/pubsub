# PubSub 

Used to wait for events

```go
pubsub := NewPubSub( "PubSub for waiting for Msg" )
go func() {
    fmt.Println( "Waiting for a message" )
    msg := <-pubsub.Subscribe()
    fmt.Println( msg )
}
pubsub.Publish( "Event Happened" )
pubsub.Close()
```
