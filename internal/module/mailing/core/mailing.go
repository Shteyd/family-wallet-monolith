package core

type (
	EmailUseCase interface {
		SendMail()
		SendConfirmationCode()
		ValidateConfirmationCode()
	}

	EmailRepository interface {
		SendMail()
		GetConfirmationCode()
		SaveConfirmationCode()
		DeleteConfirmationCode()
	}
)
