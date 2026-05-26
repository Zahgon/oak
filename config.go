package oak

import (
	"io"
)

// A Config defines the settings oak accepts on initialization. Some of these settings may be ignored depending
// on the target platform.
type Config struct {
	Driver Driver `json:"-"`
	// Assets defines where assets should be loaded from by default. Defaults to
	// 'assets/audio' and 'assets/images'.
	Assets           Assets           `json:"assets"`
	Debug            Debug            `json:"debug"`
	Screen           Screen           `json:"screen"`
	BatchLoadOptions BatchLoadOptions `json:"batchLoadOptions"`
	// FrameRate, representing the rate enter frame events are triggered, defaults to 60.
	FrameRate int `json:"frameRate"`
	// DrawFrameRate is ignored on JS. It defaults to 60.
	DrawFrameRate int `json:"drawFrameRate"`
	// IdleDrawFrameRate defaults to 60. When a window goes out of focus, this setting can be lowered to
	// reduce resource consumption by drawing.
	IdleDrawFrameRate int `json:"idleDrawFrameRate"`
	// Language defines the language oak logs are attempted to be translated to. Defaults to English.
	Language string `json:"language"`
	// Title defaults to 'Oak Window'.
	Title               string `json:"title"`
	BatchLoad           bool   `json:"batchLoad"`
	GestureSupport      bool   `json:"gestureSupport"`
	LoadBuiltinCommands bool   `json:"loadBuiltinCommands"`
	TrackInputChanges   bool   `json:"trackInputChanges"`
	// EnableDebugConsole is ignored on JS.
	EnableDebugConsole bool `json:"enableDebugConsole"`
	TopMost            bool `json:"topmost"`
	Borderless         bool `json:"borderless"`
	Fullscreen         bool `json:"fullscreen"`
	SkipRNGSeed        bool `json:"skip_rng_seed"`
	// UnlimitedDrawFrameRate is ignored on JS (it is effectively always true).
	UnlimitedDrawFrameRate bool `json:"unlimitedDrawFrameRate"`
}

// NewConfig creates a config from a set of transformation options.
func NewConfig(opts ...ConfigOption) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

func (c Config) setDefaults() Config { _ = "STUB: not implemented"; return *new(Config) }

// Assets is a json type storing paths to different asset folders
type Assets struct {
	AudioPath string `json:"audioPath"`
	ImagePath string `json:"imagePath"`
}

// Debug is a json type storing the starting debug filter and level
type Debug struct {
	Filter string `json:"filter"`
	Level  string `json:"level"`
}

// Screen is a json type storing the starting screen width and height
type Screen struct {
	X      int     `json:"X"`
	Y      int     `json:"Y"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Scale  float64 `json:"scale"`
}

// BatchLoadOptions is a json type storing customizations for batch loading.
// These settings do not take effect unless Config.BatchLoad is true.
type BatchLoadOptions struct {
	BlankOutAudio    bool  `json:"blankOutAudio"`
	MaxImageFileSize int64 `json:"maxImageFileSize"`
}

// FileConfig loads a config file, that could exist inside
// oak's binary data storage (see fileutil), to SetupConfig
func FileConfig(filePath string) ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// A ConfigOption transforms a Config object.
type ConfigOption func(Config) (Config, error)

// ReaderConfig reads a Config as json from the given reader.
func ReaderConfig(r io.Reader) ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

func (c Config) overwriteFrom(c2 Config) Config { _ = "STUB: not implemented"; return *new(Config) }

// Booleans can be directly overwritten-- all booleans in a Config
// default to false, if they were unset they will stay false.
