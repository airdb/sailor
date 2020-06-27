package sailor

// Delimiter const value.
const (
	DelimiterNewline   = "\n"
	DelimiterComma     = ","
	DelimiterDot       = "."
	DelimiterAt        = "@"
	DelimiterStar      = "*"
	DelimiterDash      = "-"
	DelimiterSlash     = "/"
	DelimiterUnderline = "_"
	DelimiterStarHex   = "\\052"
)

// Operation const value.
const (
	OperationCreate string = "CREATE"
	OperationUpdate string = "UPDATE"
	OperationDelete string = "DELETE"
	OperationUpsert string = "UPSERT"
	OperationPatch  string = "PATCH"
	OperationPut    string = "PUT"
)
