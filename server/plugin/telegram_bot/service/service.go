package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"server/global"
	"server/model/ecsusers"
	"server/model/privmsg"
	"server/plugin/telegram_bot/model"
	"strconv"
	"strings"
	"time"
)

type TelegramBotService struct{}

// SendTgMessage 发送TG消息
// 回传结果 res 为消息ID或报错文本内容 err 为具体报错内容
func (e *TelegramBotService) SendTgMessage(tokens, chatId, content, messageType string) (res string, err error) {
	// 英文逗号分隔token
	tokenList := strings.Split(tokens, ",")
	var lastError error
	// 多个token的时候轮询直至发送成功
	for index, token := range tokenList {
		bot, err := telebot.NewBot(telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 3 * time.Second},
		})
		if err != nil {
			// 初始化失败
			lastError = errors.New(fmt.Sprintf("bot initialise failed for token%d: %v", index, err))
			continue
		}
		chatID, err := strconv.ParseInt(chatId, 10, 64)
		if err != nil {
			// chat_id 转换失败
			lastError = errors.New(fmt.Sprintf("chatID cover failed for token%d: %v with %v", index, err, chatId))
			continue
		}
		var parseMode telebot.ParseMode
		switch messageType {
		case "html":
			parseMode = telebot.ModeHTML
			content = strings.ReplaceAll(content, "<br>", "")
			content = strings.ReplaceAll(content, "<\\br>", "\n")
		case "markdown":
			parseMode = telebot.ModeMarkdown
		default:
			parseMode = ""
		}
		msg, merr := bot.Send(&telebot.Chat{ID: chatID}, content, &telebot.SendOptions{ParseMode: parseMode}, telebot.NoPreview)
		if merr != nil {
			// 发送失败
			if strings.Contains(merr.Error(), "chat not found") && strings.Contains(strings.ToLower(content), "code") {
				lastError = errors.New("请先打开 https://t.me/ecs_register_bot 私聊后再进行注册，否则收不到验证码")
			} else {
				lastError = errors.New(fmt.Sprintf("bot send message failed from token%d: %s", index, merr.Error()))
			}
			continue
		}
		// 发送成功
		return strconv.Itoa(msg.ID), nil
	}
	// 全部token发送都失败了
	return "bot send message failed", lastError
}

// EditTgMessage 修改TG消息
func (e *TelegramBotService) EditTgMessage(tokens, chatId, messageId, content, messageType string) (res string, err error) {
	// 英文逗号分隔token
	tokenList := strings.Split(tokens, ",")
	var lastError error
	// 多个token的时候轮询直至修改成功
	for index, token := range tokenList {
		bot, err := telebot.NewBot(telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 3 * time.Second},
		})
		if err != nil {
			// 初始化失败
			lastError = errors.New(fmt.Sprintf("bot initialise failed for token%d: %v", index, err))
			continue
		}
		chatID, err1 := strconv.ParseInt(chatId, 10, 64)
		if err1 != nil {
			// chat_id 转换失败
			lastError = errors.New(fmt.Sprintf("chatID cover failed for token%d: %v with %v", index, err1, chatId))
			continue
		}
		messageID, err2 := strconv.Atoi(messageId)
		if err2 != nil {
			// chat_id 转换失败
			lastError = errors.New(fmt.Sprintf("messageId cover failed for token%d: %v with %v", index, err2, chatId))
			continue
		}
		var parseMode telebot.ParseMode
		switch messageType {
		case "html":
			parseMode = telebot.ModeHTML
			content = strings.ReplaceAll(content, "<br>", "")
			content = strings.ReplaceAll(content, "<\\br>", "\n")
		case "markdown":
			parseMode = telebot.ModeMarkdown
		default:
			parseMode = ""
		}
		msg, err := bot.Edit(&telebot.Message{ID: messageID, Chat: &telebot.Chat{ID: chatID}}, content,
			&telebot.SendOptions{ParseMode: parseMode}, telebot.NoPreview)
		if err != nil {
			// 修改失败
			lastError = errors.New(fmt.Sprintf("bot edit message failed for token%d: %v", index, err))
			continue
		}
		// 修改成功
		return strconv.Itoa(msg.ID), nil
	}
	return "", lastError
}

// IsTgMember 判断是否为TG频道用户
func (e *TelegramBotService) IsTgMember(tokens, userID, channelID string) (res string, err error) {
	// 英文逗号分隔token
	tokenList := strings.Split(tokens, ",")
	var lastError error
	// 多个token的时候轮询直至发送成功
	for index, token := range tokenList {
		bot, err := telebot.NewBot(telebot.Settings{
			Token: token,
		})
		if err != nil {
			// 初始化失败
			lastError = errors.New(fmt.Sprintf("bot initialise failed for token%d: %v", index, err))
			continue
		}
		// 获取频道中的成员
		userIDInt, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			// user_id 转换失败
			lastError = errors.New(fmt.Sprintf("chatID cover failed for token%d: %v with %v", index, err, userID))
			continue
		}
		channelIDInt, err := strconv.ParseInt(channelID, 10, 64)
		if err != nil {
			// channel_id 转换失败
			lastError = errors.New(fmt.Sprintf("channelID cover failed for token%d: %v", index, err))
			continue
		}
		member, err := bot.ChatMemberOf(telebot.ChatID(channelIDInt), telebot.ChatID(userIDInt))
		if err != nil {
			// 查询不到用户
			if strings.Contains(err.Error(), "user not found") {
				return "false", nil
			}
			// 查询失败，但原因不是查不到用户
			lastError = errors.New(fmt.Sprintf("bot get channel member failed for token%d: %v", index, err))
			continue
		} else {
			// 查询到用户
			if member != nil {
				return "true", nil
			}
		}
	}
	return "bot find member failed", lastError
}

func (e *TelegramBotService) CheckTgBotWithUUID(clientIP string, uuid uuid.UUID) model.TgBotCheckResult {
	// 检查Redis客户端是否初始化
	if global.GVA_REDIS == nil {
		global.GVA_LOG.Error("Redis客户端未初始化")
		return model.TgBotCheckResult{Success: false, Message: "服务配置错误"}
	}
	// 创建跟踪IP Telegram Bot测试次数的Redis键
	rateLimitKey := fmt.Sprintf("tgbot_test_rate_limit:%s", clientIP)
	// 检查当前尝试次数
	ctx := context.Background()
	currentAttempts, err := global.GVA_REDIS.Get(ctx, rateLimitKey).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		global.GVA_LOG.Error("查询Redis出错!", zap.Error(err))
		return model.TgBotCheckResult{Success: false, Message: "服务错误"}
	}
	// 如果尝试次数超过3次，返回限流错误
	if currentAttempts >= 3 {
		ttl, _ := global.GVA_REDIS.TTL(ctx, rateLimitKey).Result()
		return model.TgBotCheckResult{
			Success: false,
			Message: fmt.Sprintf("超过测试限制，请在 %d 分钟后重试", int(ttl.Minutes())+1),
		}
	}
	// 获取TG推送配置
	var pushers []privmsg.PusherConfig
	err = global.GVA_DB.Where("push_type = ?", "telegram_bot").Find(&pushers).Error
	if err != nil {
		global.GVA_LOG.Error("查询配置时出错!", zap.Error(err))
		return model.TgBotCheckResult{Success: false, Message: "push_type查询配置时出错"}
	}
	// 获取用户信息
	var user ecsusers.EcsUsers
	result := global.GVA_DB.Model(&ecsusers.EcsUsers{}).Where("uuid = ?", uuid.String()).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("未找到对应的用户", zap.String("uuid", uuid.String()))
			return model.TgBotCheckResult{Success: false, Message: "未找到对应的用户"}
		}
		global.GVA_LOG.Error("uuid查询配置时出错!", zap.Error(result.Error))
		return model.TgBotCheckResult{Success: false, Message: fmt.Sprintf("查询配置时出错 %v", result.Error)}
	}
	if user.TGID == "" {
		return model.TgBotCheckResult{Success: false, Message: "用户未配置TGID，无法发信"}
	}
	var tgBotSent bool
	for _, pusher := range pushers {
		_, err := e.SendTgMessage(pusher.ConfigValue, user.TGID, "订阅成功，恭喜你订阅成功，本消息无需回复。", "html")
		if err != nil {
			global.GVA_LOG.Error("发送Telegram消息失败!", zap.Error(err), zap.String("token", pusher.ConfigValue))
			continue
		}
		tgBotSent = true
		break
	}
	// 如果没有成功发送消息
	if !tgBotSent {
		return model.TgBotCheckResult{Success: false, Message: "Telegram消息发送失败，请检查配置"}
	}
	// 在Redis中增加尝试次数
	err = global.GVA_REDIS.Incr(ctx, rateLimitKey).Err()
	if err != nil {
		global.GVA_LOG.Error("Redis计数器增加失败!", zap.Error(err))
	}
	// 如果是第一次尝试，设置15分钟过期
	if currentAttempts == 0 {
		err = global.GVA_REDIS.Expire(ctx, rateLimitKey, 15*time.Minute).Err()
		if err != nil {
			global.GVA_LOG.Error("设置Redis过期时间失败!", zap.Error(err))
		}
	}
	return model.TgBotCheckResult{Success: true, Message: "发送成功"}
}
