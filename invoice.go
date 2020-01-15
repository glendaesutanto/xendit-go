package xendit

// Invoice contains data from Xendit's API response of invoice related request.
// For more details see https://xendit.github.io/apireference/?bash#invoices.
type Invoice struct {
	ID                        string                `json:"id"`
	ExternalID                string                `json:"external_id"`
	UserID                    string                `json:"user_id"`
	PayerEmail                string                `json:"payer_email"`
	Description               string                `json:"description"`
	Amount                    float64               `json:"amount"`
	MerchantName              string                `json:"merchant_name"`
	MerchantProfilePictureURL string                `json:"merchant_profile_picture_url"`
	InvoiceURL                string                `json:"invoice_url"`
	ExpiryDate                string                `json:"expiry_date"`
	AvailableBanks            []InvoiceBank         `json:"available_banks,omitempty"`
	AvailableEWallets         []InvoiceEWallet      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets    []InvoiceRetailOutlet `json:"available_retail_outlets,omitempty"`
	ShouldExcludeCreditCard   bool                  `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                  `json:"should_send_email"`
	Created                   string                `json:"created"`
	Updated                   string                `json:"updated"`
	BankCode                  string                `json:"bank_code,omitempty"`
	PaidAmount                string                `json:"paid_amount,omitempty"`
	AdjustedReceivedAmount    string                `json:"adjusted_received_amount,omitempty"`
	RecurringPaymentID        string                `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID        string                `json:"credit_card_charge_id,omitempty"`
	Currency                  string                `json:"currency,omitempty"`
	InitialCurrency           string                `json:"initial_currency,omitempty"`
	InitialAmount             string                `json:"initial_amount,omitempty"`
	PaidAt                    string                `json:"paid_at,omitempty"`
	MidLabel                  string                `json:"mid_label,omitempty"`
	PaymentChannel            string                `json:"payment_channel,omitempty"`
	PaymentMethod             string                `json:"payment_method,omitempty"`
	PaymentDestination        string                `json:"payment_destination,omitempty"`
	SuccessRedirectURL        string                `json:"success_redirect_url,omitempty"`
	FailureRedirectURL        string                `json:"failure_redirect_url,omitempty"`
	Items                     string                `json:"items,omitempty"`
	FixedVA                   string                `json:"fixed_va,omitempty"`
}

// InvoiceBank is data that contained in `Invoice` at AvailableBanks
type InvoiceBank struct {
	BankCode          string  `json:"bank_code"`
	CollectionType    string  `json:"collection_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string  `json:"bank_branch"`
	AccountHolderName string  `json:"account_holder_name"`
	IdentityAmount    int     `json:"identity_amount"`
}

// InvoiceEWallet is data that contained in `Invoice` at AvailableEWallets
type InvoiceEWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// InvoiceRetailOutlet is data that contained in `Invoice` at AvailableRetailOutlets
type InvoiceRetailOutlet struct {
	RetailOutletName string  `json:"retail_outlet_name"`
	PaymentCode      string  `json:"payment_code"`
	TransferAmount   float64 `json:"transfer_amount"`
	MerchantName     string  `json:"merchant_name,omitempty"`
}

// InvoiceItem is data that contained in `Invoice` at Items
type InvoiceItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
