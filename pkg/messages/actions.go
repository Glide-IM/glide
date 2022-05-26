package messages

type Action string

const (
	ActionChatMessage        Action = "message.chat"
	ActionChatMessageRecall         = "message.chat.recall" // TODO remove
	ActionChatMessageRetry          = "message.chat.retry"  // 消息重发, 服务器未ack
	ActionChatMessageResend         = "message.chat.resend" // 消息重发, 服务器已ack, 接收方未ack
	ActionGroupMessage              = "message.group"
	ActionGroupMessageRecall        = "message.group.recall" // TODO remove
	ActionMessageFailed             = "message.failed.send"
	ActionClientCustom              = "message.cli"

	ActionNotifyNeedAuth      = "notify.auth"
	ActionNotifyKickOut       = "notify.kickout"
	ActionNotifyNewContact    = "notify.contact"
	ActionNotifyGroup         = "notify.group"
	ActionNotifyAccountLogin  = "notify.login"
	ActionNotifyAccountLogout = "notify.logout"
	ActionNotifyError         = "notify.error"

	ActionAckRequest  = "ack.request"
	ActionAckGroupMsg = "ack.group.msg"
	ActionAckMessage  = "ack.message"
	ActionAckNotify   = "ack.notify"

	ActionApiAuth    = "api.auth"
	ActionHeartbeat  = "heartbeat"
	ActionApiFailed  = "api.failed"
	ActionApiSuccess = "api.success"
)
