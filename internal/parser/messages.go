package parser

// Common header fields
type MessageHeader struct {
	MessageType    byte   // 1 byte
	StockLocate    uint16 // 2 bytes
	TrackingNumber uint16 // 2 bytes
	Timestamp      uint64 // 6 bytes (nanoseconds since midnight)
}

// 1.1 System Event Message ("S")
type SystemEvent struct {
	Header    MessageHeader
	EventCode byte // 1 byte
}

// 1.2.1 Stock Directory Message ("R")
type StockDirectory struct {
	Header                   MessageHeader
	Stock                    string // 8 bytes ASCII
	MarketCategory           byte   // 1 byte
	FinancialStatusIndicator byte   // 1 byte
	RoundLotSize             uint32 // 4 bytes
	RoundLotsOnly            byte   // 1 byte
	IssueClassification      byte   // 1 byte
	IssueSubType             string // 2 bytes
	Authenticity             byte   // 1 byte
	ShortSaleThreshold       byte   // 1 byte
	IPOFlag                  byte   // 1 byte
	LULDReferencePriceTier   byte   // 1 byte
	ETPFlag                  byte   // 1 byte
	ETPLeverageFactor        uint32 // 4 bytes
	InverseIndicator         byte   // 1 byte
}

// 1.2.2 Stock Trading Action ("H")
type StockTradingAction struct {
	Header       MessageHeader
	Stock        string // 8 bytes
	TradingState byte   // 1 byte
	Reserved     byte   // 1 byte
	Reason       string // 4 bytes
}

// 1.2.3 Reg SHO Short Sale Price Test RestrictedIndicator ("Y")
type RegSHOShortSalePriceTestRestrictedIndicator struct {
	Header       MessageHeader // stock locate is locate code here
	Stock        string        // 8 bytes for the stock symbol
	RegSHOAction byte          // 1 byte for the action code
}

// 1.2.4 Market Participant Position Message ("L")
type MarketParticipantPosition struct {
	Header                 MessageHeader
	MPID                   string // 4 bytes
	Stock                  string // 8 bytes
	PrimaryMarketMaker     byte   // 1 byte
	MarketMakerMode        byte   // 1 byte
	MarketParticipantState byte   // 1 byte
}

// 1.2.5 Market---Wide Circuit Breaker (MWCB) Messaging
// 1.2.5.1 MWCB Decline Level Message ("V")
type MWCBDeclineLevel struct {
	Header MessageHeader
	Level1 uint64 // 8 bytes
	Level2 uint64 // 8 bytes
	Level3 uint64 // 8 bytes
}

// 1.2.5.2 MWCB Status Message ("W")
type MWCBStatus struct {
	Header        MessageHeader
	BreachedLevel byte // 1 byte
}

// 1.2.6 Quoting Period Update ("K")
type QuotingPeriodUpdate struct {
	Header                       MessageHeader
	Stock                        string // 8 bytes
	IPOQuotationReleaseTime      uint32 // 4 bytes
	IPOQuotationReleaseQualifier byte   // 1 byte
	IPOPrice                     uint32 // 4 bytes
}

// 1.2.7 Limit Up–Limit Down (LULD) Auction Collar ("J")
type LULDAuctionCollar struct {
	Header                      MessageHeader
	Stock                       string // 8 bytes
	AuctionCollarReferencePrice uint32 // 4 bytes
	UpperAuctionCollarPrice     uint32 // 4 bytes
	LowerAuctionCollarPrice     uint32 // 4 bytes
	AuctionCollarExtension      uint32 // 4 bytes
}

// 1.2.8 Operational Halt ("h")
type OperationalHalt struct {
	Header                MessageHeader
	Stock                 string // 8 bytes
	MarketCode            byte   // 1 byte
	OperationalHaltAction byte   // 1 byte
}

// Add Order Message
// 1.3.1 Add Order - No MPID Attribution ("A")
type AddOrderNoMPIDAttribution struct {
	Header               MessageHeader
	OrderReferenceNumber uint64 // 8 bytes
	BuySellIndicator     byte   // 1 byte
	Shares               uint32 // 4 bytes
	Stock                string // 8 bytes
	Price                uint32 // 4 bytes
}

// 1.3.2 Add Order with MPID Attribution ("F")
type AddOrderWithMPIDAttribution struct {
	Header               MessageHeader
	OrderReferenceNumber uint64 // 8 bytes
	BuySellIndicator     byte   // 1 byte
	Shares               uint32 // 4 bytes
	Stock                string // 8 bytes
	Price                uint32 // 4 bytes
	Attribution          string // 4 bytes
}

// 1.4 Modify Order Messages
// 1.4.1 Order ExecutedMessage ("E")
