# smtp-reader

A small attempt to receive data via smtp and process it automatically, like it Wunderlist does to add new item's
to a defined list.

```
curl smtp://127.0.0.1:2525 --mail-from myself@example.com --mail-rcpt receiver@example.com --upload-file email.txt --insecure
```