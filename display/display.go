package display

import (
	"flag"

	rgbmatrix "github.com/zaggash/go-rpi-rgb-led-matrix"
)

var (
	rows                     = flag.Int("led-rows", 32, "number of rows supported")
	cols                     = flag.Int("led-cols", 32, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 1, "number of daisy-chained panels")
	chainLength              = flag.Int("led-chain", 1, "number of displays daisy-chained")
	pwm_bits                 = flag.Int("pwm-bits", 11, "set PWM bits used for output")
	pwm_lsb_nanoseconds      = flag.Int("pwm-lsb-nanoseconds", 130, "set base time-unit for the on-time in the lowest significant bit")
	brightness               = flag.Int("brightness", 80, "brightness (0-100)")
	pixelMapperConfig        = flag.String("pixel-mapper-config", "", "semicolon-separated list of pixel-mappers to arrange pixels")
	hardware_mapping         = flag.String("led-gpio-mapping", "adafruit-hat-pwm", "name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "switch if your matrix has inverse colors on")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "don't use hardware pin-pulse generation")
)

type Config struct {
	Matrix   rgbmatrix.Matrix
	Toolkit  *rgbmatrix.ToolKit
	Hardware rgbmatrix.HardwareConfig
}

func New() (*Config, error) {
	c := Config{}

	hwConfig := &rgbmatrix.DefaultConfig
	hwConfig.Rows = *rows
	hwConfig.Cols = *cols
	hwConfig.Parallel = *parallel
	hwConfig.ChainLength = *chainLength
	hwConfig.PWMBits = *pwm_bits
	hwConfig.PWMLSBNanoseconds = *pwm_lsb_nanoseconds
	hwConfig.Brightness = *brightness
	hwConfig.PixelMapperConfig = *pixelMapperConfig
	hwConfig.HardwareMapping = *hardware_mapping
	hwConfig.ShowRefreshRate = *show_refresh
	hwConfig.InverseColors = *inverse_colors
	hwConfig.DisableHardwarePulsing = *disable_hardware_pulsing

	c.Hardware = *hwConfig

	// create a new Matrix instance with the DefaultConfig
	matrix, err := rgbmatrix.NewRGBLedMatrix(hwConfig)
	c.Matrix = matrix
	if err != nil {
		return &c, err
	}

	// create a ToolKit instance
	toolkit := rgbmatrix.NewToolKit(matrix)
	c.Toolkit = toolkit

	// clear canvas
	err = toolkit.Canvas.Clear()
	return &c, err
}
