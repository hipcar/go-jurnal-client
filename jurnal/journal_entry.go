package jurnal

import (
	"time"
	"fmt"
)

type JournalEntry struct {
	client *Client
}

type CreateJournalEntryRequest struct {
	JournalEntry JournalEntryRequest `json:"journal_entry,omitempty"`
}

type JournalEntryRequest struct {
	TransactionDate string `json:"transaction_date"`
	TransactionNo string `json:"transaction_no"`
	Memo string `json:"memo,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
	TransactionAccountLinesAttributes []TransactionAccountLinesAttributeRequest `json:"transaction_account_lines_attributes"`
}

type TransactionAccountLinesAttributeRequest struct {
	Id int `json:"id,omitempty"`
	AccountName string `json:"account_name"`
	Description string `json:"description,omitempty"`
	Debit float64 `json:"debit,omitempty"`
	Credit float64 `json:"credit,omitempty"`
}

type JournalEntriesResponse struct {
	JournalEntries []JournalEntryResponse `json:"journal_entries,omitempty"`
}

type JournalEntryByIdResponse struct {
	JournalEntry JournalEntryResponse `json:"journal_entry,omitempty"`
}

type JournalEntryResponse struct {
	Id int `json:"id,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
	TransactionNo string `json:"transaction_no,omitempty"`
	Token string `json:"token,omitempty"`
	Memo string `json:"memo,omitempty"`
	Source string `json:"source,omitempty"`
	Deletable bool `json:"deletable,omitempty"`
	Editable bool `json:"editable,omitempty"`
	AuditedBy string `json:"audited_by,omitempty"`
	TransactionDate string `json:"transaction_date,omitempty"`
	TransactionStatus TransactionStatusResponse `json:"transaction_status,omitempty"`
	TransactionAccountLines []TransactionAccountLineResponse `json:"transaction_status,omitempty"`
	TotalDebit float64 `json:"total_debit,omitempty"`
	TotalCredit float64 `json:"total_credit,omitempty"`
	TotalDebitCurrencyFormat string `json:"total_debit_currency_format,omitempty"`
	TotalCreditCurrencyFormat string `json:"total_credit_currency_format,omitempty"`
	Locked bool `json:"locked,omitempty"`
	IsReconciled bool `json:"is_reconciled,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	CurrencyRate float64 `json:"currency_rate,omitempty"`
}

type TransactionStatusResponse struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	NameBahasa string `json:"name_bahasa,omitempty"`
}

type TransactionAccountLineResponse struct {
	Id int `json:"id, omitempty"`
	Description string `json:"description,omitempty"`
	Debit float64 `json:"debit,omitempty"`
	Credit float64 `json:"credit,omitempty"`
	CreditCurrencyFormat string `json:"credit_currency_format,omitempty"`
	DebitCurrencyFormat string `json:"debit_currency_format,omitempty"`
	Account AccountResponse `json:"account,omitempty"`
}

type AccountResponse struct {
	Id int `json:"id, omitempty"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	Category CategoryResponse `json:"category,omitempty"`
}

type CategoryResponse struct {
	Id int `json:"id, omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *JournalEntry) GetJournalEntries () (JournalEntriesResponse, error) {
	res := new(JournalEntriesResponse)

	endpoint := "journal_entries"

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *JournalEntry) GetJournalEntryById (id string) (JournalEntryByIdResponse, error) {
	res := new(JournalEntryByIdResponse)

	endpoint := fmt.Sprintf("journal_entries/%s", id)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *JournalEntry) CreateJournalEntry (data CreateJournalEntryRequest) (JournalEntryByIdResponse, error) {
	res := new(JournalEntryByIdResponse)

	endpoint := "journal_entries"

	err := c.client.Request("POST", endpoint, data, res)
	return *res, err
}

func (c *JournalEntry) UpdateJournalEntry (id string, data CreateJournalEntryRequest) (JournalEntryByIdResponse, error) {
	res := new(JournalEntryByIdResponse)

	endpoint := fmt.Sprintf("journal_entries/%s", id)

	err := c.client.Request("PATCH", endpoint, data, res)
	return *res, err
}

func (c *JournalEntry) DeleteJournalEntry (id string) (JournalEntryByIdResponse, error) {
	res := new(JournalEntryByIdResponse)

	endpoint := fmt.Sprintf("journal_entries/%s", id)

	err := c.client.Request("DELETE", endpoint, nil, res)
	return *res, err
}
