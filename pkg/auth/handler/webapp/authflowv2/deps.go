package authflowv2

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	wire.Struct(new(InternalAuthflowV2SignupLoginHandler), "*"),
	wire.Struct(new(AuthflowV2LoginHandler), "*"),
	wire.Struct(new(AuthflowV2SignupHandler), "*"),
	wire.Struct(new(AuthflowV2ReauthHandler), "*"),
	wire.Struct(new(AuthflowV2EnterPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2EnterOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SetupOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowV2ViewRecoveryCodeHandler), "*"),
	wire.Struct(new(AuthflowV2ErrorHandler), "*"),
	wire.Struct(new(AuthflowV2NoAuthenticatorHandler), "*"),
	wire.Struct(new(AuthflowV2CreatePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2AccountStatusHandler), "*"),
	wire.Struct(new(AuthflowV2NotFoundHandler), "*"),
	wire.Struct(new(AuthflowV2SelectAccountHandler), "*"),
	wire.Struct(new(AuthflowV2VerifyBotProtectionHandler), "*"),
	wire.Struct(new(AuthflowV2EnterRecoveryCodeHandler), "*"),
	wire.Struct(new(AuthflowV2ChangePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ChangePasswordSuccessHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordLinkSentHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordOTPHandler), "*"),
	wire.Struct(new(AuthflowV2ResetPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ResetPasswordSuccessHandler), "*"),
	wire.Struct(new(AuthflowV2SetupTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2EnterTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2OOBOTPLinkHandler), "*"),
	wire.Struct(new(AuthflowV2VerifyLoginLinkOTPHandler), "*"),
	wire.Struct(new(AuthflowV2PromptCreatePasskeyHandler), "*"),
	wire.Struct(new(AuthflowV2UsePasskeyHandler), "*"),
	wire.Struct(new(AuthflowV2TerminateOtherSessionsHandler), "*"),
	wire.Struct(new(AuthflowV2PromoteHandler), "*"),
	wire.Struct(new(AuthflowV2FinishFlowHandler), "*"),
	wire.Struct(new(AuthflowV2WechatHandler), "*"),
	wire.Struct(new(AuthflowV2AccountLinkingHandler), "*"),
	wire.Struct(new(AuthflowV2LDAPLoginHandler), "*"),
)
