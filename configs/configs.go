package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/onflow/flow-go-sdk"
)

type Config struct {
	// -- Feature flags --

	DisableRawTransactions   bool `env:"FLOW_WALLET_DISABLE_RAWTX"`
	DisableFungibleTokens    bool `env:"FLOW_WALLET_DISABLE_FT"`
	DisableNonFungibleTokens bool `env:"FLOW_WALLET_DISABLE_NFT"`
	DisableChainEvents       bool `env:"FLOW_WALLET_DISABLE_CHAIN_EVENTS"`

	// -- Admin account --

	AdminAddress    string `env:"FLOW_WALLET_ADMIN_ADDRESS,notEmpty"`
	AdminKeyIndex   int    `env:"FLOW_WALLET_ADMIN_KEY_INDEX" envDefault:"0"`
	AdminKeyType    string `env:"FLOW_WALLET_ADMIN_KEY_TYPE" envDefault:"local"`
	AdminPrivateKey string `env:"FLOW_WALLET_ADMIN_PRIVATE_KEY,notEmpty"`
	// This sets the number of proposal keys to be used on the admin account.
	// You can increase transaction throughput by using multiple proposal keys for
	// parallel transaction execution.
	AdminProposalKeyCount uint16 `env:"FLOW_WALLET_ADMIN_PROPOSAL_KEY_COUNT" envDefault:"1"`

	// -- Keys --

	// When "DefaultKeyType" is set to "local", private keys are generated by the API
	// and stored as encrypted text in the database.
	DefaultKeyType  string `env:"FLOW_WALLET_DEFAULT_KEY_TYPE" envDefault:"local"`
	DefaultKeyIndex int    `env:"FLOW_WALLET_DEFAULT_KEY_INDEX" envDefault:"0"`
	// If the default of "-1" is used for "DefaultKeyWeight"
	// the service will use flow.AccountKeyWeightThreshold from the Flow SDK.
	DefaultKeyWeight int    `env:"FLOW_WALLET_DEFAULT_KEY_WEIGHT" envDefault:"-1"`
	DefaultSignAlgo  string `env:"FLOW_WALLET_DEFAULT_SIGN_ALGO" envDefault:"ECDSA_P256"`
	DefaultHashAlgo  string `env:"FLOW_WALLET_DEFAULT_HASH_ALGO" envDefault:"SHA3_256"`
	// This symmetrical key is used to encrypt private keys
	// that are stored in the database.
	// It needs to be exactly 32 bytes long.
	EncryptionKey string `env:"FLOW_WALLET_ENCRYPTION_KEY,notEmpty"`

	// -- Database --

	DatabaseDSN  string `env:"FLOW_WALLET_DATABASE_DSN" envDefault:"wallet.db"`
	DatabaseType string `env:"FLOW_WALLET_DATABASE_TYPE" envDefault:"sqlite"`

	// -- Host and chain access --

	Host          string       `env:"FLOW_WALLET_HOST"`
	Port          int          `env:"FLOW_WALLET_PORT" envDefault:"3000"`
	AccessAPIHost string       `env:"FLOW_WALLET_ACCESS_API_HOST,notEmpty"`
	ChainID       flow.ChainID `env:"FLOW_WALLET_CHAIN_ID" envDefault:"flow-emulator"`

	// -- Templates --

	EnabledTokens []string `env:"FLOW_WALLET_ENABLED_TOKENS" envSeparator:","`

	// -- Workerpool --

	// Defines the maximum number of active jobs that can be queued before
	// new jobs are rejected.
	WorkerQueueCapacity uint `env:"FLOW_WALLET_WORKER_QUEUE_CAPACITY" envDefault:"1000"`
	// Number of concurrent workers handling incoming jobs.
	// You can increase the number of workers if you're sending
	// too many transactions and find that the queue is often backlogged.
	WorkerCount uint `env:"FLOW_WALLET_WORKER_COUNT" envDefault:"100"`

	// -- Google KMS --

	GoogleKMSProjectID  string `env:"FLOW_WALLET_GOOGLE_KMS_PROJECT_ID"`
	GoogleKMSLocationID string `env:"FLOW_WALLET_GOOGLE_KMS_LOCATION_ID"`
	GoogleKMSKeyRingID  string `env:"FLOW_WALLET_GOOGLE_KMS_KEYRING_ID"`
}

type Options struct {
	EnvFilePath string
}

// ParseConfig parses environment variables and flags to a valid Config.
func ParseConfig(opt *Options) (*Config, error) {
	if opt != nil && opt.EnvFilePath != "" {
		// Load variables from a file to the environment of the process
		if err := godotenv.Load(opt.EnvFilePath); err != nil {
			log.Printf("Could not load environment variables from file.\n%s\nIf running inside a docker container this can be ignored.\n\n", err)
		}
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
