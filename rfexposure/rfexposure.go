package rfexposure

import (
	"errors"
	"math"
)

type Result struct {
	Frequency float64
	Distance  float64
}

func TestStub() []Result {

	var xmtr_power int16 = 1000
	var feedline_length int16 = 73
	var duty_cycle float64 = 0.5
	var per_30 float64 = 0.5

	c1 := CableValues{
		K1: 0.122290,
		K2: 0.000260,
	}

	all_frequency_values := []FrequencyValues{
		{
			Freq:    7.3,
			SWR:     2.25,
			GainDBI: 1.5,
		},
		{
			Freq:    14.35,
			SWR:     1.35,
			GainDBI: 1.5,
		},
		{
			Freq:    18.1,
			SWR:     3.7,
			GainDBI: 1.5,
		},
		{
			Freq:    21.45,
			SWR:     4.45,
			GainDBI: 1.5,
		},
		{
			Freq:    24.99,
			SWR:     4.1,
			GainDBI: 1.5,
		},
		{
			Freq:    29.7,
			SWR:     2.18,
			GainDBI: 4.5,
		},
	}

	var uncontrolled_safe_distances []Result

	for _, f := range all_frequency_values {
		distance := CalculateUncontrolledSafeDistance(f, c1, xmtr_power, feedline_length, duty_cycle, per_30)

		uncontrolled_safe_distances = append(uncontrolled_safe_distances, Result{Frequency: f.Freq, Distance: distance})
	}

	return uncontrolled_safe_distances
}

type CableValues struct {
	K1 float64
	K2 float64
}

type FrequencyValues struct {
	Freq    float64
	SWR     float64
	GainDBI float64
}

func CalculateUncontrolledSafeDistance(freq_values FrequencyValues, cable_values CableValues, transmitter_power int16,
	feedline_length int16, duty_cycle float64, uncontrolled_percentage_30_minutes float64) float64 {

	gamma := CalculateReflectionCoefficient(freq_values)

	feedline_loss_per_100ft_at_frequency := CalculateFeedlineLossPer100ftAtFrequency(freq_values, cable_values)

	feedline_loss_for_matched_load_at_frequency := CalculateFeedlineLossForMatchedLoadAtFrequency(feedline_length, feedline_loss_per_100ft_at_frequency)

	feedline_loss_for_matched_load_at_frequency_percentage := CalculateFeedlineLossForMatchedLoadAtFrequencyPercentage(feedline_loss_for_matched_load_at_frequency)

	gamma_squared := math.Pow(math.Abs(gamma), 2)

	feedline_loss_for_swr := CalculateFeedlineLossForSWR(feedline_loss_for_matched_load_at_frequency_percentage, gamma_squared)

	feedline_loss_for_swr_percentage := CalculateFeedlineLossForSWRPercentage(feedline_loss_for_swr)

	power_loss_at_swr := feedline_loss_for_swr_percentage * float64(transmitter_power)

	peak_envelope_power_at_antenna := float64(transmitter_power) - power_loss_at_swr

	uncontrolled_average_pep := peak_envelope_power_at_antenna * duty_cycle * uncontrolled_percentage_30_minutes

	mpe_s := 180 / (math.Pow(freq_values.Freq, 2))

	gain_decimal := math.Pow(10, freq_values.GainDBI/10)

	return math.Sqrt((0.219 * uncontrolled_average_pep * gain_decimal) / mpe_s)
}

func CalculateReflectionCoefficient(freq_values FrequencyValues) float64 {
	return math.Abs(float64((freq_values.SWR - 1) / (freq_values.SWR + 1)))
}

func CalculateFeedlineLossForMatchedLoadAtFrequency(feedline_length int16, feedline_loss_per_100ft_at_frequency float64) float64 {
	return ((float64(feedline_length) / 100.0) * feedline_loss_per_100ft_at_frequency)
}

func CalculateFeedlineLossForMatchedLoadAtFrequencyPercentage(feedline_loss_for_matched_load float64) float64 {
	return math.Pow(10, (-(feedline_loss_for_matched_load) / 10.0))
}

func CalculateFeedlineLossPer100ftAtFrequency(freq_values FrequencyValues, cable_values CableValues) float64 {
	return cable_values.K1 * (math.Sqrt(freq_values.Freq + cable_values.K2*freq_values.Freq))
}

func CalculateFeedlineLossForSWR(feedline_loss_for_matched_load_percentage float64, gamma_squared float64) float64 {
	return -10 * math.Log10(feedline_loss_for_matched_load_percentage*
		((1-gamma_squared)/(1-math.Pow(feedline_loss_for_matched_load_percentage, 2)*gamma_squared)))
}

func CalculateFeedlineLossForSWRPercentage(feedline_loss_for_swr float64) float64 {
	return (100 - 100/(math.Pow(10, feedline_loss_for_swr/10))) / 100
}

func ValidateParameters(freq_values FrequencyValues, cable_values CableValues, transmitter_power int16,
	feedline_length int16, duty_cycle float64, uncontrolled_percentage_30_minutes float64) error {

	if freq_values.Freq <= 0 {
		return errors.New("frequency must be greater than 0")
	}
	if cable_values.K1 < 0 || cable_values.K2 < 0 {
		return errors.New("cable values must be non-negative")
	}
	if transmitter_power <= 0 {
		return errors.New("transmitter power must be greater than 0")
	}
	if feedline_length <= 0 {
		return errors.New("feedline length must be greater than 0")
	}
	if duty_cycle < 0 || duty_cycle > 1 {
		return errors.New("duty cycle must be between 0 and 1")
	}
	if uncontrolled_percentage_30_minutes < 0 || uncontrolled_percentage_30_minutes > 1 {
		return errors.New("uncontrolled percentage must be between 0 and 1")
	}
	return nil
}
