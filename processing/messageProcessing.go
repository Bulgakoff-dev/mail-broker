package processing

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"mail-broker/config"
	"mail-broker/logger"
	"mail-broker/mail"
	"strings"
	"time"
)

func ProcessMail(message map[string]string, rdb *redis.Client) error {
	cfg := config.GetConfig()
	logg := logger.GetLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var withName bool = true

	// Fields check
	email, ok := message["email"]
	if !ok {
		logg.Error().Msg("Field 'email' not found in config")
		return errors.New("required field 'email' not found")
	}

	confirmationLink, ok := message["confirmationLink"]
	if !ok {
		logg.Error().Msg("Field 'confirmationLink' not found in config")
		return errors.New("required field 'confirmationLink' not found")
	}

	locale, ok := message["locale"]
	if !ok {
		logg.Warn().Msg("Required field 'locale' not found, locale would be set up to EN")
		locale = "EN"
	}
	if locale == "" {
		locale = "EN"
	}

	if confirmationLink == "" || email == "" {
		logg.Error().Msg("One from two required fields is empty")
		return errors.New("empty required field 'email' or 'confirmationLink'")
	}

	firstName, ok := message["firstName"]
	if !ok {
		logg.Warn().Msg("Field 'firstName' not found in config")
		withName = false
	}

	lastName, ok := message["lastName"]
	if !ok {
		logg.Warn().Msg("Field 'lastName' not found in config")
		withName = false
	}

	if firstName == "" || lastName == "" {
		logg.Warn().Msg("One from two name fields is empty")
		withName = false
	}

	// Get Mail Template and Subject from Redis
	if withName {
		template, err := rdb.HGetAll(ctx, "MAIL:sign_up_link:"+locale).Result()
		if err != nil {
			logg.Err(err).Msg("Error getting template from Redis")
			return err
		}
		mailBody := template["body"]

		mailBody = strings.ReplaceAll(mailBody, "{{first_name}}", firstName)
		mailBody = strings.ReplaceAll(mailBody, "{{last_name}}", lastName)
		mailBody = strings.ReplaceAll(mailBody, "{{confirm_link}}", confirmationLink)

		to := []string{email}
		err = mail.SendMail(to, template["subject"], mailBody, cfg.SMTPMail, cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPAddr, cfg.SMTPPort)

		if err != nil {
			logg.Log().Err(err).Msg("Error sending mail")
			return err
		}
	} else {
		template, err := rdb.HGetAll(ctx, "MAIL:sign_up_link_unnamed:"+locale).Result()
		if err != nil {
			logg.Log().Err(err).Msg("Error getting template from Redis")
			return err
		}
		mailBody := template["body"]
		mailBody = strings.ReplaceAll(mailBody, "{{confirm_link}}", confirmationLink)

		to := []string{email}
		err = mail.SendMail(to, template["subject"], mailBody, cfg.SMTPMail, cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPAddr, cfg.SMTPPort)

		if err != nil {
			logg.Log().Err(err).Msg("Error sending mail")
			return err
		}
	}
	return nil
}
