package core

// credentials
var (
	SmtpHost = "smtp.mail.ru"
	SmtpPort = 465
	Username = "aalex110@mail.ru"
	Password = ""
)

// message params
var (
	MessageHeaderFrom       = "From"
	MessageHeaderTo         = "To"
	MessageAddressHeaderCc  = "Cc"
	MessageHeaderSubject    = "Subject"
	MessageSubjectValue     = "password confirmation"
	MessageBodyType         = "text/html"
	ConfirmCodeBodyTemplate = "Hello, <b>%s</b>, that your confirmation code: <b>%d</b> "
)
