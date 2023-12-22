// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package logs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsLogConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementLogConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsLogConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookLogConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementLogConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_LogConfig(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookLogConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_LogConfig(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_LogConfig(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_LogConfig(val, result))
}

func testDecodeRaw_LogConfig(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_LogConfig(vStringSlice, result))
}

func TestLogConfig_GetPFlagSet(t *testing.T) {
	val := LogConfig{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestLogConfig_SetFlags(t *testing.T) {
	actual := LogConfig{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_cloudwatch-enabled", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cloudwatch-enabled", testValue)
			if vBool, err := cmdFlags.GetBool("cloudwatch-enabled"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vBool), &actual.IsCloudwatchEnabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_cloudwatch-region", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cloudwatch-region", testValue)
			if vString, err := cmdFlags.GetString("cloudwatch-region"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.CloudwatchRegion)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_cloudwatch-log-group", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cloudwatch-log-group", testValue)
			if vString, err := cmdFlags.GetString("cloudwatch-log-group"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.CloudwatchLogGroup)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_cloudwatch-template-uri", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("cloudwatch-template-uri", testValue)
			if vString, err := cmdFlags.GetString("cloudwatch-template-uri"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.CloudwatchTemplateURI)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_kubernetes-enabled", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("kubernetes-enabled", testValue)
			if vBool, err := cmdFlags.GetBool("kubernetes-enabled"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vBool), &actual.IsKubernetesEnabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_kubernetes-url", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("kubernetes-url", testValue)
			if vString, err := cmdFlags.GetString("kubernetes-url"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.KubernetesURL)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_kubernetes-template-uri", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("kubernetes-template-uri", testValue)
			if vString, err := cmdFlags.GetString("kubernetes-template-uri"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.KubernetesTemplateURI)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_stackdriver-enabled", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("stackdriver-enabled", testValue)
			if vBool, err := cmdFlags.GetBool("stackdriver-enabled"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vBool), &actual.IsStackDriverEnabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_gcp-project", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("gcp-project", testValue)
			if vString, err := cmdFlags.GetString("gcp-project"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.GCPProjectName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_stackdriver-logresourcename", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("stackdriver-logresourcename", testValue)
			if vString, err := cmdFlags.GetString("stackdriver-logresourcename"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.StackdriverLogResourceName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_stackdriver-template-uri", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("stackdriver-template-uri", testValue)
			if vString, err := cmdFlags.GetString("stackdriver-template-uri"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.StackDriverTemplateURI)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_flyin-enabled", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("flyin-enabled", testValue)
			if vBool, err := cmdFlags.GetBool("flyin-enabled"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vBool), &actual.IsFlyinEnabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_flyin-template-uri", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("flyin-template-uri", testValue)
			if vString, err := cmdFlags.GetString("flyin-template-uri"); err == nil {
				testDecodeJson_LogConfig(t, fmt.Sprintf("%v", vString), &actual.FlyinTemplateURI)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
