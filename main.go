package main

import (
	"exposure-web/rfexposure"
	"fmt"
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

func submitHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	// Parse form data
	frequency := c.PostForm("frequency")
	swr := c.PostForm("swr")
	gaindbi := c.PostForm("gaindbi")
	k1 := c.PostForm("k1")
	fmt.Println(k1)
	k2 := c.PostForm("k2")
	transmitter_power := c.PostForm("transmitter_power")
	feedline_length := c.PostForm("feedline_length")
	duty_cycle := c.PostForm("duty_cycle")
	uncontrolled_percentage_30_minutes := c.PostForm("uncontrolled_percentage_30_minutes")

	// Convert cable values to appropriate types
	k_1, err := strconv.ParseFloat(k1, 64)
	fmt.Println(err)
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

func main() {
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseFiles("form.html")))

	r.GET("/", formHandler)
	r.POST("/submit", submitHandler)
	r.Static("/static", "./static")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
