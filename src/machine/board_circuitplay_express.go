// +build sam,atsamd21,circuitplay_express

package machine

// GPIO/Analog Pins
const (
	A0      = 2 // PWM available, also ADC/AIN[0]
	A1      = 5 // ADC/AIN[5]
	A2      = 6 // PWM available, also ADC/AIN[6]
	A3      = 7 // PWM available, also ADC/AIN[7]
	A4      = 3 // PORTB so does not work yet
	A5      = 2 // PORTB so does not work yet
	A6      = 9 // PORTB so does not work yet
	A7      = 8 // PORTB so does not work yet
	BUTTONA = 28
	BUTTONB = 14
	SLIDE   = 15 // built-in slide switch
	D13     = 17 // PWM available

	D10 = 0xfc // not connected
	D11 = 0xfe // not connected
	D12 = 0xff // not connected

	BUTTON  = BUTTONA
	BUTTON1 = BUTTONB
)

const (
	LED = D13
)

// Analog Pins
const (
	A8 = 11 // ADC/AIN[19]
	A9 = 9  // ADC/AIN[17]

	LIGHTSENSOR = A8
	TEMPSENSOR  = A9
)

// UART0 pins
const (
	UART_TX_PIN = 0xfa // 8 // PORTB so does not work yet
	UART_RX_PIN = 0xfb // 9 // PORTB so does not work yet
)

// I2C pins
const (
	SDA_PIN = 0 // SDA: SERCOM3/PAD[0]
	SCL_PIN = 1 // SCL: SERCOM3/PAD[1]
)
