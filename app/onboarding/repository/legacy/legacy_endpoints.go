package legacy

const (
	ONBOARDING = "/onboarding"

	SIMOBI_VERIFY_USER  = ONBOARDING + "/simobiplus/verifyuser"
	SIMOBI_VERIFY_PHONE = ONBOARDING + "/simobiplus/verifyphone"
	SIMOBI_OTP          = ONBOARDING + "/simobiplus/otp"
	SIMOBI_REGISTER     = ONBOARDING + "/simobiplus/register"

	ATM_CARD     = ONBOARDING + "/atm/card"
	ATM_OTP      = ONBOARDING + "/atm/otp"
	ATM_PIN      = ONBOARDING + "/atm/pin"
	ATM_REGISTER = ONBOARDING + "/atm/register"

	NO_ATM_VERIFY_USER = ONBOARDING + "/noatm/verifyuser"
	NO_ATM_OTP         = ONBOARDING + "/noatm/otp"
	NO_ATM_EMAIL       = ONBOARDING + "/noatm/email"
	NO_ATM_REGISTER    = ONBOARDING + "/noatm/register"

	FRESH_VERIFY_PHONE = ONBOARDING + "/fresh/verifyphone"
	FRESH_OTP          = ONBOARDING + "/fresh/otp"
	FRESH_PASSWORD     = ONBOARDING + "/fresh/password"

	RESET_VERIFY_PHONE = ONBOARDING + "/reset/verifyphone"
	RESET_OTP          = ONBOARDING + "/reset/otp"
	RESET_PASSWORD     = ONBOARDING + "/reset/password"

	LOGIN_VERIFY_PASSWORD = ONBOARDING + "/login/verifypassword"
)
