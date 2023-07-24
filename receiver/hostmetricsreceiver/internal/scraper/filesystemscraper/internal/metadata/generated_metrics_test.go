// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemFilesystemInodesUsageDataPoint(ts, 1, "device-val", "mode-val", "mountpoint-val", "type-val", AttributeStateFree)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemFilesystemUsageDataPoint(ts, 1, "device-val", "mode-val", "mountpoint-val", "type-val", AttributeStateFree)

			allMetricsCount++
			mb.RecordSystemFilesystemUtilizationDataPoint(ts, 1, "device-val", "mode-val", "mountpoint-val", "type-val")

			metrics := mb.Emit()

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 0)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "system.filesystem.inodes.usage":
					assert.False(t, validatedMetrics["system.filesystem.inodes.usage"], "Found a duplicate in the metrics slice: system.filesystem.inodes.usage")
					validatedMetrics["system.filesystem.inodes.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "FileSystem inodes used.", ms.At(i).Description())
					assert.Equal(t, "{inodes}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mode")
					assert.True(t, ok)
					assert.EqualValues(t, "mode-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mountpoint")
					assert.True(t, ok)
					assert.EqualValues(t, "mountpoint-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "type-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.EqualValues(t, "free", attrVal.Str())
				case "system.filesystem.usage":
					assert.False(t, validatedMetrics["system.filesystem.usage"], "Found a duplicate in the metrics slice: system.filesystem.usage")
					validatedMetrics["system.filesystem.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Filesystem bytes used.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mode")
					assert.True(t, ok)
					assert.EqualValues(t, "mode-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mountpoint")
					assert.True(t, ok)
					assert.EqualValues(t, "mountpoint-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "type-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.EqualValues(t, "free", attrVal.Str())
				case "system.filesystem.utilization":
					assert.False(t, validatedMetrics["system.filesystem.utilization"], "Found a duplicate in the metrics slice: system.filesystem.utilization")
					validatedMetrics["system.filesystem.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Fraction of filesystem bytes used.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mode")
					assert.True(t, ok)
					assert.EqualValues(t, "mode-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mountpoint")
					assert.True(t, ok)
					assert.EqualValues(t, "mountpoint-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "type-val", attrVal.Str())
				}
			}
		})
	}
}