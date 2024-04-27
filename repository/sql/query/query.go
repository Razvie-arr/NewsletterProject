package query

import (
	_ "embed"
)

var (
	//go:embed scripts/editor/Read.sql
	ReadEditor string
	//go:embed scripts/editor/ReadByEmail.sql
	ReadEditorByEmail string
	//go:embed scripts/subscriber/ReadByEmail.sql
	ReadSubscriberByEmail string
	//go:embed scripts/newsletter/Read.sql
	ReadNewsletter string
	//go:embed scripts/subscriberNewsletter/SelectSubscribersByNewsletterId.sql
	ReadSubscribersByNewsletterId string
	//go:embed scripts/subscriberNewsletter/GetVerificationString.sql
	GetVerificationString string
	//go:embed scripts/newsletter/ExistsNewsletterWithEditor.sql
	ExistsNewsletterWithEditor string
	//go:embed scripts/newsletter/ReadNewsletterInfoWithLimit.sql
	ReadNewsletterInfoWithLimit string
)
