# wireshark tips

How to hide P_DATA_V2 packets:

```
!(_ws.col.info == "MessageType: P_DATA_V2")
```

# openvp
