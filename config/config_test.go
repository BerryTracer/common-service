package config_test

import (
	"errors"
	"testing"

	"github.com/BerryTracer/common-service/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEnvLoader struct {
	mock.Mock
}

func (m *MockEnvLoader) Getenv(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func (m *MockEnvLoader) Load() error {
	args := m.Called()
	return args.Error(0)
}

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name          string
		envName       string
		mockEnvValue  string
		mockLoadErr   error
		expectedValue string
		expectedError string
	}{
		{
			name:          "Environment variable exists",
			envName:       "EXISTING_VAR",
			mockEnvValue:  "value1",
			mockLoadErr:   nil,
			expectedValue: "value1",
			expectedError: "",
		},
		{
			name:          "Environment variable does not exist, .env load fails",
			envName:       "MISSING_VAR",
			mockEnvValue:  "",
			mockLoadErr:   errors.New("failed to load .env"),
			expectedValue: "",
			expectedError: "failed to load .env",
		},
		{
			name:          "Environment variable does not exist, .env load succeeds",
			envName:       "NEW_VAR",
			mockEnvValue:  "value2",
			mockLoadErr:   nil,
			expectedValue: "value2",
			expectedError: "",
		},
		{
			name:          "Environment variable and .env both missing",
			envName:       "NON_EXISTENT_VAR",
			mockEnvValue:  "",
			mockLoadErr:   nil,
			expectedValue: "",
			expectedError: "NON_EXISTENT_VAR environment variable not set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Here you'd set up your mocks based on tt.mockEnvValue and tt.mockLoadErr
			mockEnvLoader := new(MockEnvLoader)
			mockEnvLoader.On("Getenv", tt.envName).Return(tt.mockEnvValue)
			mockEnvLoader.On("Load").Return(tt.mockLoadErr)
			value, err := config.LoadEnv(mockEnvLoader, tt.envName)

			if tt.expectedError != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.expectedValue, value)
		})
	}
}

func TestLoadEnvWithDefault(t *testing.T) {
	tests := []struct {
		name          string
		envName       string
		mockEnvValue  string
		mockLoadErr   error
		defaultValue  string
		expectedValue string
		expectedError string
	}{
		{
			name:          "Environment variable exists",
			envName:       "EXISTING_VAR",
			mockEnvValue:  "value1",
			mockLoadErr:   nil,
			defaultValue:  "",
			expectedValue: "value1",
			expectedError: "",
		},
		{
			name:          "Environment variable does not exist, .env load fails, default value is empty",
			envName:       "MISSING_VAR",
			mockEnvValue:  "",
			mockLoadErr:   errors.New("failed to load .env"),
			defaultValue:  "",
			expectedValue: "",
			expectedError: "failed to load .env",
		},
		{
			name:          "Environment variable does not exist, .env load fails, default value is not empty",
			envName:       "MISSING_VAR",
			mockEnvValue:  "",
			mockLoadErr:   errors.New("failed to load .env"),
			defaultValue:  "default",
			expectedValue: "default",
			expectedError: "",
		},
		{
			name:          "Environment variable does not exist, .env load succeeds, default value is empty",
			envName:       "NEW_VAR",
			mockEnvValue:  "value2",
			mockLoadErr:   nil,
			defaultValue:  "",
			expectedValue: "value2",
			expectedError: "",
		},
		{
			name:          "Environment variable does not exist, .env load succeeds, default value is not empty",
			envName:       "NEW_VAR",
			mockEnvValue:  "value2",
			mockLoadErr:   nil,
			defaultValue:  "default",
			expectedValue: "value2",
			expectedError: "",
		},
		{
			name:          "Environment variable and .env both missing, default value is empty",
			envName:       "NON_EXISTENT_VAR",
			mockEnvValue:  "",
			mockLoadErr:   nil,
			defaultValue:  "",
			expectedValue: "",
			expectedError: "NON_EXISTENT_VAR environment variable not set",
		},
		{
			name:          "Environment variable and .env both missing, default value is not empty",
			envName:       "NON_EXISTENT_VAR",
			mockEnvValue:  "",
			mockLoadErr:   nil,
			defaultValue:  "default",
			expectedValue: "default",
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Here you'd set up your mocks based on tt.mockEnvValue and tt.mockLoadErr
			mockEnvLoader := new(MockEnvLoader)
			mockEnvLoader.On("Getenv", tt.envName).Return(tt.mockEnvValue)
			mockEnvLoader.On("Load").Return(tt.mockLoadErr)
			value, err := config.LoadEnvWithDefault(mockEnvLoader, tt.envName, tt.defaultValue)

			if tt.expectedError != "" {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.expectedValue, value)
		})
	}

}
