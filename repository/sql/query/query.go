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
	//go:embed scripts/subscriber/Create.sql
	CreateSubscriber string
	//go:embed scripts/newsletter/Read.sql
	ReadNewsletter string
	//go:embed scripts/subscriberNewsletter/SelectSubscribersByNewsletterId.sql
	ReadSubscribersByNewsletterId string
	//go:embed scripts/subscriberNewsletter/Subscribe.sql
	Subscribe string
	//go:embed scripts/subscriberNewsletter/GetVerificationString.sql
	GetVerificationString string
	//go:embed scripts/subscriberNewsletter/Unsubscribe.sql
	Unsubscribe string
)
