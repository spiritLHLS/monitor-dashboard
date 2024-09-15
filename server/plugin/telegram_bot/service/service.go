package service

import (
	"errors"
	"fmt"
	"gopkg.in/telebot.v3"
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
		case "markdown":
			parseMode = telebot.ModeMarkdown
		default:
			parseMode = ""
		}
		msg, err := bot.Send(&telebot.Chat{ID: chatID}, content, &telebot.SendOptions{ParseMode: parseMode}, telebot.NoPreview)
		if err != nil {
			// 发送失败
			lastError = errors.New(fmt.Sprintf("bot send message failed for token%d: %v", index, err))
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
