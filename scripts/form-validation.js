// scripts/form-validation.js

document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('form');
    form.addEventListener('submit', (event) => {
        let isValid = true;
        const swr = document.getElementById('swr');
        const gaindbi = document.getElementById('gaindbi');
        const transmitterPower = document.getElementById('transmitter_power');
        const feedlineLength = document.getElementById('feedline_length');
        const dutyCycle = document.getElementById('duty_cycle');
        const uncontrolledPercentage = document.getElementById('uncontrolled_percentage_30_minutes');

        // Helper function to check if a value is a float
        const isFloat = (value) => !isNaN(value) && parseFloat(value) == value;

        // Helper function to check if a value is an integer
        const isInt = (value) => !isNaN(value) && parseInt(value) == value;

        // Validate SWR (float64)
        if (!isFloat(swr.value)) {
            isValid = false;
            alert('SWR must be a float.');
        }

        // Validate Gain (dBi) (float64)
        if (!isFloat(gaindbi.value)) {
            isValid = false;
            alert('Gain (dBi) must be a float.');
        }

        // Validate Transmitter Power (int16)
        if (!isInt(transmitterPower.value)) {
            isValid = false;
            alert('Transmitter Power must be an integer.');
        }

        // Validate Feedline Length (int16)
        if (!isInt(feedlineLength.value)) {
            isValid = false;
            alert('Feedline Length must be an integer.');
        }

        // Validate Duty Cycle (float64)
        if (!isFloat(dutyCycle.value)) {
            isValid = false;
            alert('Duty Cycle must be a float.');
        }

        // Validate Uncontrolled Percentage (30 minutes) (float64)
        if (!isFloat(uncontrolledPercentage.value)) {
            isValid = false;
            alert('Uncontrolled Percentage (30 minutes) must be a float.');
        }

        if (!isValid) {
            event.preventDefault();
        }
    });
});