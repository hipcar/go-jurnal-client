# go-jurnal-client
Jurnal Client written in Go. Un-official Jurnal API Wrapper. 
- [Have trouble ?](https://github.com/hipcar/go-jurnal-client/issues)
- [Submit changes/features ?](https://github.com/hipcar/go-jurnal-client/pulls)

Documentation
=============

## Getting Started
```
go get github.com/hipcar/go-jurnal-client
```

## Init
```go
package main

import (
	"github.com/hipcar/go-jurnal-client"
)

func main() {
  jurnalClient := jurnal.NewClient(nil)
  jurnalClient.APIKey = "YOUR_JURNAL_API_KEY"
}
```

## Journal Entry

### Get Journal Entries
```go
res, err := jurnalClient.JournalEntry.GetJournalEntries()
```

### Get Journal Entry By Id / Transaction No
```go
res, err := jurnalClient.JournalEntry.GetJournalEntryById("1")
```

### Create Journal Entry
```go
journalEntry := jurnal.JournalEntryRequest {
    TransactionDate: "17/09/2018",
    TransactionNo: "TEST-CREATE-1",
    Memo: "test creating journal entry",
    TransactionAccountLinesAttributes: []jurnal.TransactionAccountLinesAttributeRequest{
        {
            AccountName: "Account Name 1",
            Description: "Desc",
            Debit: 100000,
        },
        {
            AccountName: "Account Name 2",
            Description: "Desc",
            Credit: 100000,
        },
    },
}

data := jurnal.CreateJournalEntryRequest{
    JournalEntry: journalEntry,
}

result, err := jurnalApiClient.JournalEntry.CreateJournalEntry(data)
```

### Update Journal Entry
```go
journalEntry := jurnal.JournalEntryRequest {
    TransactionDate: "17/09/2018",
    TransactionNo: "TEST-CREATE-1",
    Memo: "test creating journal entry",
    TransactionAccountLinesAttributes: []jurnal.TransactionAccountLinesAttributeRequest{
        {
            AccountName: "Account Name 1",
            Description: "Desc",
            Debit: 100000,
        },
        {
            AccountName: "Account Name 2",
            Description: "Desc",
            Credit: 100000,
        },
    },
}

data := jurnal.CreateJournalEntryRequest{
    JournalEntry: journalEntry,
}

result, err := jurnalApiClient.JournalEntry.UpdateJournalEntry("TEST-CREATE-1", data)
```

### Delete Journal Entry By Id / Transaction No
```go
_, err := jurnalApiClient.JournalEntry.DeleteJournalEntry("1")
```

