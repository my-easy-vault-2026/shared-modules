package utils

import (
	"os"
	"shared-modules/common"
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v2"
)

var Config *Conf

type AppStoreConfig struct {
	Kid        string `yaml:"kid"`
	IssuerId   string `yaml:"issuerId"`
	PrivateKey string `yaml:"privateKey"`
	BundleID   string `yaml:"bundleID"`
	Sandbox    bool   `yaml:"sandbox"`
}

type RedisConfig struct {
	Addr       string `yaml:"addr"`
	Pwd        string `yaml:"pwd"`
	Db         int    `yaml:"db"`
	SSL        bool   `yaml:"ssl"`
	MaxConn    int    `yaml:"maxConn"`
	TLSVersion string `yaml:"tlsVersion"`
}

type MysqlConfig struct {
	DBs      []MysqlDBConfig `yaml:"dbs"`
	RowLimit int             `yaml:"rowLimit"`
}

type MysqlDBConfig struct {
	Name                  string        `yaml:"name"`
	Dsn                   string        `yaml:"dsn"`
	MaxOpenConn           int           `yaml:"maxOpenConn"`
	MaxIdleConn           int           `yaml:"maxIdleConn"`
	ConMaxLifeTimeSeconds time.Duration `yaml:"conMaxLifeTimeSeconds"`
}

type GinConfig struct {
	Listen         string   `yaml:"listen"`
	WsListen       string   `yaml:"wsListen"`
	Mode           string   `yaml:"mode"`
	TrustedProxies []string `yaml:"trustedProxies"`
}

type SystemConfig struct {
	BrandName                           string        `yaml:"brandName"`
	CdnDomain                           string        `yaml:"cdnDomain"`
	S3ImgDomin                          string        `yaml:"s3ImgDomain"`
	UploadPath                          string        `yaml:"uploadPath"`
	MailChannel                         string        `yaml:"mailChannel"`
	MailKey                             string        `yaml:"mailKey"`
	SendMail                            string        `yaml:"sendMail"`
	MailGunMailKey                      string        `yaml:"mailGunMailKey"`
	MailGunSendMail                     string        `yaml:"mailGunSendMail"`
	MailGunDomain                       string        `yaml:"mailGunDomain"`
	MailGunApiBase                      string        `yaml:"mailGunApiBase"`
	MailGunWebhookSigningKey            string        `yaml:"mailGunWebhookSigningKey"`
	MailGunApplyCardDomain              string        `yaml:"mailGunApplyCardDomain"`
	MerchantMail                        string        `yaml:"merchantMail"`
	TraderMail                          string        `yaml:"traderMail"`
	OrderMailNotify                     bool          `yaml:"orderMailNotify"`
	LockMicroseconds                    time.Duration `yaml:"lockMicroseconds"`
	LockPollingMicroseconds             time.Duration `yaml:"lockPollingMicroseconds"`
	LockWaitMicroseconds                time.Duration `yaml:"lockWaitMicroseconds"`
	RedisDaoLockMicroseconds            time.Duration `yaml:"redisDaoLockMicroseconds"`
	SaltLength                          int           `yaml:"saltLength"`
	RateUrl                             string        `yaml:"rateUrl"`
	RateUrl2                            string        `yaml:"rateUrl2"`
	VerifyServer                        string        `yaml:"verifyServer"`
	HashUrl                             string        `yaml:"hashUrl"`
	VerifyUrl                           string        `yaml:"verifyUrl"`
	L2CacheExpireSeconds                time.Duration `yaml:"l2CacheExpireSeconds"`
	RatePrecision                       int           `yaml:"ratePrecision"`
	RedisPopTimeoutMilliseconds         time.Duration `yaml:"redisPopTimeoutMilliseconds"`
	NodeSize                            int           `yaml:"nodeSize"`
	PromotionLink                       string        `yaml:"promotionLink"`
	PromotionLinkKey                    string        `yaml:"promotionLinkKey"`
	HTTPTimeout                         time.Duration `yaml:"httpTimeout"`
	WorkerPoolSize                      int           `yaml:"workerPoolSize"`
	SpinlockAttempts                    int           `yaml:"spinlockAttempts"`
	DefaultGoroutineTimeoutMilliseconds time.Duration `yaml:"defaultGoroutineTimeoutMilliseconds"`
	DBTimezoneOffset                    int           `yaml:"dbTimezoneOffset"`
	ExportLimit                         int64         `yaml:"exportLimit"`
}

type CoinsdoConfig struct {
	ApiKey             string                   `yaml:"apiKey"`
	SecretKey          string                   `yaml:"secretKey"`
	TargetDeviceUUID   string                   `yaml:"targetDeviceUuid"` // 這個是存幣的錢包UUID
	Url                string                   `yaml:"url"`
	GetAddressUrl      string                   `yaml:"getAddressUrl"`
	WithdrawUrl        string                   `yaml:"withdrawUrl"`
	DeviceInfoUrl      string                   `yaml:"deviceInfoUrl"`
	MaxLimit           int                      `yaml:"maxLimit"`
	SubmitAccount      string                   `yaml:"submitAccount"`
	WithdrawDeviceID   string                   `yaml:"withdrawDeviceId"` // 這個是提幣錢包的UUID
	CoinsendPrivateKey string                   `yaml:"coinsendPrivateKey"`
	ReviewUUID         string                   `yaml:"reviewUuid"`
	ReviewDevice       string                   `yaml:"reviewDevice"`
	WithdrawAutoLevel  common.WithdrawAutoLevel `yaml:"withdrawAutoLevel"`
	AssociatedMainnet  string                   `yaml:"associatedMainnet"`
}

type ChippayConfig struct {
	BaseURL           string `yaml:"baseUrl"`
	PublicKey         string `yaml:"publicKey"`
	PrivateKey        string `yaml:"privateKey"`
	CompanyID         uint64 `yaml:"companyId"`
	CallbackURL       string `yaml:"callbackUrl"`
	CallbackPublicKey string `yaml:"callbackPublicKey"`
}

type AuthConfig struct {
	ExpireTime                time.Duration `yaml:"expireTime"`
	LoginDataExpireHours      time.Duration `yaml:"loginDataExpireHours"`
	AdminExpireHours          time.Duration `yaml:"adminExpireHours"`
	AdminLoginDataExpireHours time.Duration `yaml:"adminLoginDataExpireHours"`
	PINUnlockMaxAttempts      int           `yaml:"pinUnlockMaxAttempts"`
}

type NotifyConfig struct {
	OTPExpireTime                time.Duration `yaml:"otpExpireTime"`
	CallbackIntervalMilliseconds time.Duration `yaml:"callbackIntervalMilliseconds"` // 系統同時發出多個callback，callback間需間隔多久時間
	CallbackTimeoutSeconds       time.Duration `yaml:"callbackTimeoutSeconds"`       // 單個callback超時秒數
	CallbackTaskTimeoutSeconds   time.Duration `yaml:"callbackTaskTimeoutSeconds"`   // 整個callback task(裡面包含許多callback) 超時秒數
}

type CardConfig struct {
	ExpireSeconds          time.Duration `yaml:"expireSeconds"`
	ConfirmLockSeconds     time.Duration `yaml:"confirmLockSeconds"`
	ReactivateAfterMinutes time.Duration `yaml:"reactivateAfterMinutes"`
	ReactivateWaitSeconds  time.Duration `yaml:"reactivateWaitSeconds"`
}

type QuoteConfig struct {
	ExpireSeconds time.Duration `yaml:"expireSeconds"`
}

type ServerConfig struct {
	Env string `yaml:"env"`
}

type ReapConfig struct {
	BaseURL                                 string        `yaml:"baseUrl"`
	APIKey                                  string        `yaml:"apiKey"`
	PrivateKey                              string        `yaml:"privateKey"`
	PublicKey                               string        `yaml:"publicKey"`
	AuthorizationLockSeconds                time.Duration `yaml:"authorizationLockSeconds"`
	AuthorizationAdviceLockWaitMicroseconds time.Duration `yaml:"authorizationAdviceLockWaitMicroseconds"`
	ThreedsNotificationCacheSeconds         time.Duration `yaml:"threedsNotificationCacheSeconds"`
	ThreedsNotificationTimeoutSeconds       time.Duration `yaml:"threedsNotificationTimeoutSeconds"`
	MerchantPayTimeoutMilliseconds          time.Duration `yaml:"merchantPayTimeoutMilliseconds"` // request的時候等待時間，通常為1.2秒
	MerchantPayWaitMilliseconds             time.Duration `yaml:"merchantPayWaitMilliseconds"`    // 授權通過去auth的callback等待時間，可以等很久
	MerchantPayTXNotifyAttempt              int           `yaml:"merchantPayTxNotifyAttempt"`     // callback次數
	MerchantPayTXNotifySeconds              time.Duration `yaml:"merchantPayTxNotifySeconds"`     // callback重試間隔
	MerchantCardStatusAttempt               int           `yaml:"merchantCardStatusAttempt"`
	MerchantCardStatusSeconds               time.Duration `yaml:"merchantCardStatusSeconds"`
	MerchantForward3dsAttempt               int           `yaml:"merchantForward3dsAttempt"`
	MerchantForward3dsSeconds               time.Duration `yaml:"merchantForward3dsSeconds"`
	MerchantWalletOTPAttempt                int           `yaml:"merchantWalletOTPAttempt"`
	MerchantWalletOTPSeconds                time.Duration `yaml:"merchantWalletOTPSeconds"`
	SignatureExpireSeconds                  time.Duration `yaml:"signatureExpireSeconds"`
	MerchantPrivateKey                      string        `yaml:"merchantPrivateKey"` // 與商戶溝通用的密鑰
	MerchantPublicKey                       string        `yaml:"merchantPublicKey"`  // 與商戶溝通用的密鑰
	AutoTopUpMinGap                         string        `yaml:"autoTopUpMinGap"`    // 自動充值時，錢包金額需要超過消費金額多少的最小差值(gap)
	AutoThreedsAnswerDelayMilliseconds      time.Duration `yaml:"autoThreedsAnswerDelayMilliseconds"`
	AutoThreedsAnswerAttempt                int           `yaml:"autoThreedsAnswerAttempt"`
	AutoThreedsTaskLockSeconds              time.Duration `yaml:"autoThreedsTaskLockMilliSeconds"`
}

type TopUpConfig struct {
	PreviewExpireSeconds time.Duration `yaml:"previewExpireSeconds"`
}
type TopDownConfig struct {
	PreviewExpireSeconds time.Duration `yaml:"previewExpireSeconds"`
}
type ExchangeConfig struct {
	PreviewExpireSeconds time.Duration `yaml:"previewExpireSeconds"`
}
type TransferConfig struct {
	PreviewExpireSeconds time.Duration `yaml:"previewExpireSeconds"`
}
type CardToCardConfig struct {
	PreviewExpireSeconds time.Duration `yaml:"previewExpireSeconds"`
}
type WithdrawConfig struct {
	CheckIntervalSeconds time.Duration `yaml:"checkIntervalSeconds"`
	WarnCriteriaSeconds  time.Duration `yaml:"warnCriteriaSeconds"`
}
type MerchantConfig struct {
	APIKeyLength           int           `yaml:"apiKeyLength"`
	RSAKeySize             int           `yaml:"rsaKeySize"`
	CardCategoryID         uint64        `yaml:"cardCategoryId"`
	BatchQueryBillSize     int           `yaml:"batchQueryBillSize"`
	BatchQueryUserSize     int           `yaml:"batchQueryUserSize"`
	BillCacheExpiryMinutes time.Duration `yaml:"billCacheExpiryMinutes"`
}

type ManualConfig struct {
	LockSeconds time.Duration `yaml:"lockSeconds"`
}

type PasskeyConfig struct {
	DisplayName string   `yaml:"displayName"`
	Domain      string   `yaml:"domain"`
	Sites       []string `yaml:"sites"`
	EncryptKey  string   `yaml:"encryptKey"`
}

type WebSocketConfig struct {
	BucketSize      int           `yaml:"bucketSize"`
	WriteWaitMs     time.Duration `yaml:"writeWaitMs"`
	PongWaitMs      time.Duration `yaml:"pongWaitMs"`
	PingPeriodMs    time.Duration `yaml:"pingPeriodMs"`
	MaxMsgSize      int64         `yaml:"maxMsgSize"`
	ReadBufferSize  int           `yaml:"readBufferSize"`
	WriteBufferSize int           `yaml:"writeBufferSize"`
	BroadcastSize   int64         `yaml:"broadcastSize"`
}

type PointConfig struct {
	DefaultExpiryMonths int `yaml:"defaultExpiryMonths"`
}

type FinancialConfig struct {
	SnapshotBatchSize int `yaml:"snapshotBatchSize"`
	UserBatchSize     int `yaml:"userBatchSize"`
}

type WhaleConfig struct {
	CoinAPPID                          string          `yaml:"coinAppId"`
	BaseURL                            string          `yaml:"baseUrl"`
	AESKey                             string          `yaml:"aesKey"`
	AESIV                              string          `yaml:"aesIv"`
	ActivationDeposit                  decimal.Decimal `yaml:"activationDeposit"`
	ActivationFee                      decimal.Decimal `yaml:"activationFee"`
	TransferLockMicroseconds           time.Duration   `yaml:"transferLockMicroseconds"`
	TransferLockWaitMicroseconds       time.Duration   `yaml:"transferLockWaitMicroseconds"`
	AdjustBalanceFee                   decimal.Decimal `yaml:"adjustBalanceFee"`
	CheckAdjustOrderFromMinutes        time.Duration   `yaml:"checkAdjustOrderFromMinutes"`
	CheckAdjustOrderToMinutes          time.Duration   `yaml:"checkAdjustOrderToMinutes"`
	PollingBalanceAttemps              int             `yaml:"pollingBalanceAttempts"`
	PollingBalanceIntervalMilliseconds time.Duration   `yaml:"pollingBalanceIntervalMilliseconds"`
}

type PaycryptoConfig struct {
	BaseURL           string          `yaml:"baseUrl"`
	APIKey            string          `yaml:"apiKey"`
	APISecret         string          `yaml:"apiSecret"`
	APIPassphrase     string          `yaml:"apiPassphrase"`
	ActivationDeposit decimal.Decimal `yaml:"activationDeposit"`
	PrivateKey        string          `yaml:"privateKey"`
	PublicKey         string          `yaml:"publicKey"`
}

type Conf struct {
	Server     ServerConfig     `yaml:"server"`
	Redis      RedisConfig      `yaml:"redis"`
	Coinface   CoinfaceConfig   `yaml:"coinface"`
	Mysql      MysqlConfig      `yaml:"mysql"`
	Gin        GinConfig        `yaml:"gin"`
	System     SystemConfig     `yaml:"system"`
	Auth       AuthConfig       `yaml:"auth"`
	Notify     NotifyConfig     `yaml:"notify"`
	Quote      QuoteConfig      `yaml:"quote"`
	Card       CardConfig       `yaml:"card"`
	Reap       ReapConfig       `yaml:"reap"`
	Coinsdo    CoinsdoConfig    `yaml:"coinsdo"`
	Chippay    ChippayConfig    `yaml:"chippay"`
	TopUp      TopUpConfig      `yaml:"topUp"`
	TopDown    TopDownConfig    `yaml:"topDown"`
	Exchange   ExchangeConfig   `yaml:"exchange"`
	Transfer   TransferConfig   `yaml:"transfer"`
	CardToCard CardToCardConfig `yaml:"cardToCard"`
	Passkey    PasskeyConfig    `yaml:"passkey"`
	Manual     ManualConfig     `yaml:"manual"`
	Merchant   MerchantConfig   `yaml:"merchant"`
	AppStore   AppStoreConfig   `yaml:"appStore"`
	Websocket  WebSocketConfig  `yaml:"websocket"`
	Firebase   FirebaseConfig   `yaml:"firebase"`
	S3         S3Config         `yaml:"s3"`
	Point      PointConfig      `yaml:"point"`
	Withdraw   WithdrawConfig   `yaml:"withdraw"`
	Changelly  ChangellyConfig  `yaml:"changelly"`
	Financial  FinancialConfig  `yaml:"financial"`
	Whale      WhaleConfig      `yaml:"whale"`
	Paycrypto  PaycryptoConfig  `yaml:"paycrypto"`
	Binance    BinanceConf      `yaml:"binance"`
	Telegram   TelegramConf     `yaml:"telegram"`
	Etherfi    EtherfiConf      `yaml:"etherfi"`
	Scroll     ScrollConf       `yaml:"scroll"`
	Report     ReportConfig     `yaml:"report"`
}

type FirebaseConfig struct {
	JsonFile string `yaml:"jsonFile"`
}

type CoinfaceConfig struct {
	Host                     string `yaml:"host"`
	SecretKey                string `yaml:"secretKey"`
	SyncUrl                  string `yaml:"syncUrl"`
	AsyncUrl                 string `yaml:"asyncUrl"`
	CompanyId                int64  `yaml:"companyId"`
	OpenUploadCertificates   bool   `yaml:"openUploadCertificates"`
	DuplicateRegisterCheck   bool   `yaml:"duplicateRegisterCheck"`
	KycOkUrl                 string `yaml:"kycOkUrl"`
	KycFailUrl               string `yaml:"kycFailUrl"`
	KycProgressUrl           string `yaml:"kycProgressUrl"`
	Kyc3SyncUrl              string `yaml:"kyc3SyncUrl"`
	Kyc3AsyncUrl             string `yaml:"kyc3AsyncUrl"`
	Kyc3Ocr                  bool   `yaml:"kyc3Ocr"`
	Kyc3UploadCertificate    bool   `yaml:"kyc3UploadCertificate"`
	Kyc3DuplicateCheck       bool   `yaml:"kyc3DuplicateCheck"`
	Kyc3BlacklistCheck       bool   `yaml:"kyc3BlacklistCheck"`
	Kyc3AmlCheck             bool   `yaml:"kyc3AmlCheck"`
	MerchantUserKyc3SyncUrl  string `yaml:"merchantUserKyc3SyncUrl"`
	MerchantUserKyc3AsyncUrl string `yaml:"merchantUserKyc3AsyncUrl"`
}

type S3Config struct {
	AccessKey     string `yaml:"accessKey"`
	SecretKey     string `yaml:"secretKey"`
	Region        string `yaml:"region"`
	BucketName    string `yaml:"bucketName"`
	BaseDirectory string `yaml:"baseDirectory"`
}

type ChangellyConfig struct {
	BaseURL    string `yaml:"baseUrl"`
	APIKey     string `yaml:"apiKey"`
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
}

type BinanceConf struct {
	Host            string `yaml:"host"`
	ApiKey          string `yaml:"apiKey"`
	SecretKey       string `yaml:"secretKey"`
	OrderExpireTime int    `yaml:"orderExpireTime"`
	Webhook         string `yaml:"webhook"`
}

type TelegramConf struct {
	BotToken  string `yaml:"botToken"`
	NotifyUrl string `yaml:"notifyUrl"`
}

type EtherfiConf struct {
	BaseURL       string `yaml:"baseUrl"`
	AccountId     string `yaml:"accountId"`
	SafeId        string `yaml:"safeId"`
	LegalMailFrom string `yaml:"legalMailFrom"`
	UsdtToken     string `yaml:"usdtToken"`
	UsdcToken     string `yaml:"usdcToken"`
}

type ScrollConf struct {
	Scheme  string             `yaml:"scheme"`
	BaseUrl string             `yaml:"baseUrl"`
	Path    string             `yaml:"path"`
	ApiKey  string             `yaml:"apikey"`
	Targets []ScrollTargetConf `yaml:"targets"`
}

type ScrollTargetConf struct {
	Currency string `yaml:"currency"`
	Address  string `yaml:"address"`
}

type ReportConfig struct {
	MaxDurationDays int `yaml:"maxDurationDays"`
	MaxMemoryMB     int `yaml:"maxMemoryMB"`
	MaxRows         int `yaml:"maxRows"`
	CSVPageSize     int `yaml:"csvPageSize"`
}

func GetConf(configPath string) error {

	c := &Conf{}

	yamlFile, err := os.ReadFile(configPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return err
	}

	Config = c

	return nil
}
