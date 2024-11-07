package main

import (
	"exposure-web/rfexposure"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func formHandler(c *gin.Context) {
	data := rfexposure.TestStub()
	c.HTML(http.StatusOK, "form.html", gin.H{
		"results": data,
	})
}

func main() {
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseFiles("form.html")))

	r.GET("/", formHandler)
	r.POST("/submit", submitHandler)
	r.Static("/static", "./static")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

type FormInput struct {
	Frequency                       string `form:"frequency" binding:"required"`
	SWR                             string `form:"swr" binding:"required"`
	GainDBI                         string `form:"gaindbi" binding:"required"`
	K1                              string `form:"k1" binding:"required"`
	K2                              string `form:"k2" binding:"required"`
	TransmitterPower                string `form:"transmitter_power" binding:"required"`
	FeedlineLength                  string `form:"feedline_length" binding:"required"`
	DutyCycle                       string `form:"duty_cycle" binding:"required"`
	UncontrolledPercentage30Minutes string `form:"uncontrolled_percentage_30_minutes" binding:"required"`
}

func submitHandler(c *gin.Context) {
	var input FormInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse form data
	frequency := input.Frequency
	swr := input.SWR
	gaindbi := input.GainDBI
	k1 := input.K1
	k2 := input.K2
	transmitter_power := input.TransmitterPower
	feedline_length := input.FeedlineLength
	duty_cycle := input.DutyCycle
	uncontrolled_percentage_30_minutes := input.UncontrolledPercentage30Minutes

	// Convert cable values to appropriate types
	k_1, err := strconv.ParseFloat(k1, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid k1 value"})
		return
	}

	k_2, err := strconv.ParseFloat(k2, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid k2 value"})
		return
	}

	var cableValues rfexposure.CableValues
	cableValues.K1 = k_1
	cableValues.K2 = k_2

	// Convert frequency values to appropriate types
	freq, err := strconv.ParseFloat(frequency, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid frequency value"})
		return
	}

	standing_wave_ratio, err := strconv.ParseFloat(swr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid swr value"})
		return
	}

	gain, err := strconv.ParseFloat(gaindbi, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gain dbi value"})
		return
	}

	var freqValues rfexposure.FrequencyValues
	freqValues.Freq = freq
	freqValues.SWR = standing_wave_ratio
	freqValues.GainDBI = gain

	// Convert other values to appropriate types
	transmitterPowerInt, err := strconv.ParseInt(transmitter_power, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transmitter power value"})
		return
	}
	transmitterPower := int16(transmitterPowerInt)

	feedlineLengthInt, err := strconv.ParseInt(feedline_length, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feedline length value"})
		return
	}
	feedlineLength := int16(feedlineLengthInt)

	dutyCycle, err := strconv.ParseFloat(duty_cycle, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duty cycle value"})
		return
	}

	uncontrolledPercentage30Minutes, err := strconv.ParseFloat(uncontrolled_percentage_30_minutes, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid uncontrolled percentage 30 minutes value"})
		return
	}

	// Execute CalculateUncontrolledSafeDistance
	distance := rfexposure.CalculateUncontrolledSafeDistance(freqValues, cableValues, transmitterPower, feedlineLength, dutyCycle, uncontrolledPercentage30Minutes)

	// Return the result
	c.JSON(http.StatusOK, gin.H{"distance": distance})
}
