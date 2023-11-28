package utils

import (
	"flag"

	rgbmatrix "github.com/zaggash/go-rpi-rgb-led-matrix"
)

var (
	gpio_mapping        = flag.String("led-gpio-mapping", "adafruit-hat-pwm", "Name of GPIO mapping used")
	gpio_slowdown       = flag.Int("led-slowdown-gpio", 0, "Slowdown GPIO. Needed for faster Pis and/or slower panels. range=0...4")
	rows                = flag.Int("led-rows", 32, "Panel rows. Typically 8, 16, 32 or 64.")
	cols                = flag.Int("led-cols", 32, "Panel columns. Typically 32 or 64.")
	chain               = flag.Int("led-chain", 1, "Number of daisy-chained panels.")
	parallel            = flag.Int("led-parallel", 1, "For A/B+ models or RPi2,3b: parallel chains. range=1..3")
	panel_type          = flag.String("led-panel-type", "", "Chipset of the panel. Typically null, FM6126A or FM6127")
	multiplexing        = flag.Int("led-multiplexing", 0, "Mux type: 0=direct; 1=Stripe; 2=Checkered... range=0...17")
	row_address_type    = flag.Int("led-row-addr-type", 0, "0 = default; 1 = AB-addressed panels; 2 = direct row select; 3 = ABC-addressed panels; 4 = ABC Shift + DE direct")
	pixel_mapper        = flag.String("led-pixel-mapper", "", "Semicolon-separated list of pixel-mappers to arrange pixels.")
	brightness          = flag.Int("led-brightness", 80, "Brightness in percent. range=0...100")
	pwm_bits            = flag.Int("led-pwm-bits", 11, "PWM bits range=1..11")
	show_refresh_rate   = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	limit_refresh       = flag.Int("led-limit-refresh", 0, "Limit refresh rate to this frequency in Hz. 0=no limit.")
	pwm_lsb_nanoseconds = flag.Int("led-pwm-lsb-nanoseconds", 130, "PWM Nanoseconds for LSB")
	pwm_dither_bits     = flag.Int("led-pwm-dither-bits", 0, "Time dithering of lower bits")
	no_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "Don't use hardware pin-pulse generation.")
	inverse_colors      = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	rgb_sequence        = flag.String("led-rgb-sequence", "RGB", "Switch if your matrix has led colors swapped.")
)

type Display struct {
	Hardware rgbmatrix.HardwareConfig
	Runtime  rgbmatrix.RuntimeConfig
	Matrix   rgbmatrix.Matrix
	Toolkit  *rgbmatrix.ToolKit
	Channel  chan bool
}

func NewDisplay() (*Display, error) {
	c := Display{}

	hwConfig := &rgbmatrix.DefaultConfig
	hwConfig.GPIOMapping = *gpio_mapping
	hwConfig.Rows = *rows
	hwConfig.Cols = *cols
	hwConfig.Parallel = *parallel
	hwConfig.PanelType = *panel_type
	hwConfig.Multiplexing = *multiplexing
	hwConfig.RowAddressType = *row_address_type
	hwConfig.ChainLength = *chain
	hwConfig.PixelMapperConfig = *pixel_mapper
	hwConfig.Brightness = *brightness
	hwConfig.PWMBits = *pwm_bits
	hwConfig.ShowRefreshRate = *show_refresh_rate
	hwConfig.LimitRefresh = *limit_refresh
	hwConfig.PWMLSBNanoseconds = *pwm_lsb_nanoseconds
	hwConfig.PWMDitherBits = *pwm_dither_bits
	hwConfig.DisableHardwarePulsing = *no_hardware_pulsing
	hwConfig.InverseColors = *inverse_colors
	hwConfig.RGBSequence = *rgb_sequence
	c.Hardware = *hwConfig

	rtConfig := &rgbmatrix.DefaultRtConfig
	rtConfig.GPIOSlowdown = *gpio_slowdown
	c.Runtime = *rtConfig

	// create a new Matrix instance with the DefaultConfig
	matrix, err := rgbmatrix.NewRGBLedMatrix(hwConfig, rtConfig)
	if err != nil {
		return &c, err
	}
	c.Matrix = matrix

	// create a ToolKit instance
	toolkit := rgbmatrix.NewToolKit(matrix)
	c.Toolkit = toolkit

	// clear canvas
	err = toolkit.Canvas.Clear()
	return &c, err
}
