package mutation

import (
	_ "embed"
)

var (
	//go:embed scripts/editor/Create.sql
	CreateEditor string
	//go:embed scripts/subscriber/Create.sql
	CreateSubscriber string
	//go:embed scripts/subscriberNewsletter/Unsubscribe.sql
	Unsubscribe string
	//go:embed scripts/subscriberNewsletter/Subscribe.sql
	Subscribe string
	//go:embed scripts/post/Create.sql
	CreatePost string
	//go:embed scripts/newsletter/Create.sql
	CreateNewsletter string
	//go:embed scripts/newsletter/Delete.sql
	DeleteNewsletter string
	//go:embed scripts/newsletter/Update.sql
	UpdateNewsletter string
)
