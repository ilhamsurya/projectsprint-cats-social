package config

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var cfg Config
var doOnce sync.Once

type Config struct {
	Application struct {
		Port  string `mapstructure:"PORT"`
		Group string `mapstructure:"GROUP"`
	} `mapstructure:"APPLICATION"`
	DB struct {
		Postgre struct {
			Host    string `mapstructure:"HOST"`
			Port    int    `mapstructure:"PORT"`
			Name    string `mapstructure:"NAME"`
			User    string `mapstructure:"USER"`
			Pass    string `mapstructure:"PASS"`
			Migrate bool   `mapstructure:"MIGRATE"`
		} `mapstructure:"POSTGRE"`
	} `mapstructure:"DB"`
	Auth struct {
		SecretKey               string `mapstructure:"SECRET_KEY"`
		AccessTokenExpiredTime  string `mapstructure:"ACCESS_TOKEN_EXPIRED_TIME"`
		RefreshTokenExpiredTime string `mapstructure:"REFRESH_TOKEN_EXPIRED_TIME"`
	} `mapstructure:"AUTH"`
	ObjectStorage struct {
		CloudName string `mapstructure:"NAME"`
		ApiKey    string `mapstructure:"API_KEY"`
		ApiSecret string `mapstructure:"API_SECRET"`
	} `mapstructure:"OBJECT_STORAGE"`
	Redis struct {
		Host             string `mapstructure:"HOST"`
		Password         string `mapstructure:"PASSWORD"`
		OTPExpiredTime   string `mapstructure:"OTP_EXPIRED_TIME"`
		TokenExpiredTime string `mapstructure:"TOKEN_EXPIRED_TIME"`
	} `mapstructure:"REDIS"`
	MailGun struct {
		Domain string `mapstructure:"DOMAIN"`
		ApiKey string `mapstructure:"API_KEY"`
	} `mapstructure:"MAILGUN"`
	Oauth2Google struct {
		ClientID     string `mapstructure:"CLIENT_ID"`
		ClientSecret string `mapstructure:"CLIENT_SECRET"`
		RedirectUrl  string `mapstructure:"REDIRECT_URL"`
	} `mapstructure:"OAUTH2_GOOGLE"`
	ShippingApi struct {
		URL    string `mapstructure:"URL"`
		APIKey string `mapstructure:"API_KEY"`
	} `mapstructure:"SHIPPING_API"`
	SeaLabsPayAPI struct {
		SeaMoneyUrl        string `mapstructure:"SEA_MONEY_URL"`
		PaymentUrl         string `mapstructure:"PAYMENT_URL"`
		RefundUrl          string `mapstructure:"REFUND_URL"`
		CancelUrl          string `mapstructure:"CANCEL_URL"`
		APIKey             string `mapstructure:"API_KEY"`
		MerchantCode       string `mapstructure:"MERCHANT_CODE"`
		VerifyCallbackUrl  string `mapstructure:"VERIFY_CALLBACK_URL"`
		TopUpCallbackUrl   string `mapstructure:"TOP_UP_CALLBACK_URL"`
		PaymentCallbackUrl string `mapstructure:"PAYMENT_CALLBACK_URL"`
		VerifyRedirectUrl  string `mapstructure:"VERIFY_REDIRECT_URL"`
		TopUpRedirectUrl   string `mapstructure:"TOP_UP_REDIRECT_URL"`
		PaymentRedirectUrl string `mapstructure:"PAYMENT_REDIRECT_URL"`
	} `mapstructure:"SEALABS_PAY_API"`
	ClientUrlRedirectOuath2 string `mapstructure:"OAUTH2_CLIENT_URL"`
	ResetPasswordURLFormat  string `mapstructure:"RESET_PASSWORD_URL_FORMAT"`
}

func Get() Config {
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.WithLevel(zerolog.FatalLevel).Msg(fmt.Sprintf("cannot read .yaml file: %v", err))
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.WithLevel(zerolog.FatalLevel).Msg("cannot unmarshaling config")
		}
	})

	return cfg
}
