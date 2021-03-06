package mysensors

import (
	"time"

	ml "github.com/mycontroller-org/backend/v2/pkg/model"
	msgml "github.com/mycontroller-org/backend/v2/pkg/model/message"
	mtsml "github.com/mycontroller-org/backend/v2/plugin/metrics"
)

// Labels used in this provider
const (
	LabelImperialSystem    = "ms_imperial_system" // this is bool label, used to configure about the system metric or imperial
	LabelNodeID            = "ms_node_id"         // MySensors node id
	LabelSensorID          = "ms_sensor_id"       // MySensors sensor id
	LabelType              = "ms_type"            // MySensors type reference, can be used for sensor fields
	LabelTypeString        = "ms_type_string"     // MySensors type in string format
	LabelNodeType          = "ms_node_type"       // MySensors node type
	LabelLockedReason      = "ms_locked_reason"   // If the the node is locked, reason will be in this label
	LabelEraseEEPROM       = "ms_erase_eeprom"    // Supports only for MYSBootloader, on a reboot of node, the node eeprom will be erased
	LabelFirmwareTypeID    = "ms_type_id"         // MySensors firmware type id
	LabelFirmwareVersionID = "ms_version_id"      // MySensors firmware version id
)

// internal references
const (
	idBroadcastInt            = 255                        // broadcast id in MySensors
	idBroadcast               = "255"                      // broadcast id in MySensors
	payloadEmpty              = ""                         // Empty payload
	serialMessageSplitter     = '\n'                       // serial message splitter
	firmwarePurgeJobName      = "mysensors_firmware_store" // firmware purge job name, append with gateway id
	firmwarePurgeJobCron      = "0 */5 * * * *"            // purge loaded firmware, if not used for a while
	firmwarePurgeInactiveTime = 15 * time.Minute           // firmware inactive time, eligible for purging

	payloadON  = "1"
	payloadOFF = "0"

	// Command types and value
	cmdPresentation = "0"
	cmdSet          = "1"
	cmdRequest      = "2"
	cmdInternal     = "3"
	cmdStream       = "4"

	// Internal action type details
	actionTime                = "1"
	actionIDResponse          = "4"
	actionConfig              = "6"
	actionReboot              = "13"
	actionHeartBeatRequest    = "18"
	actionRequestPresentation = "19"
	actionDiscoverRequest     = "20"
	// Stream type details
	actionFirmwareConfigResponse = "1"
	actionFirmwareResponse       = "3"
)

// mysensors message data
// node-id/child-sensor-id/command/ack/type
type message struct {
	NodeID   string
	SensorID string
	Command  string
	Ack      string
	Type     string
	Payload  string
}

// Command mapping on received messages
var cmdMapForRx = map[string]string{
	"0": msgml.TypePresentation,
	"1": msgml.TypeSet,
	"2": msgml.TypeRequest,
	"3": msgml.TypeAction,
	"4": msgml.TypeAction,
}

// Presentation type mapping for received messages
var presentationTypeMapForRx = map[string]string{
	"0":  "S_DOOR",
	"1":  "S_MOTION",
	"2":  "S_SMOKE",
	"3":  "S_BINARY",
	"4":  "S_DIMMER",
	"5":  "S_COVER",
	"6":  "S_TEMP",
	"7":  "S_HUM",
	"8":  "S_BARO",
	"9":  "S_WIND",
	"10": "S_RAIN",
	"11": "S_UV",
	"12": "S_WEIGHT",
	"13": "S_POWER",
	"14": "S_HEATER",
	"15": "S_DISTANCE",
	"16": "S_LIGHT_LEVEL",
	"17": "S_ARDUINO_NODE",
	"18": "S_ARDUINO_REPEATER_NODE",
	"19": "S_LOCK",
	"20": "S_IR",
	"21": "S_WATER",
	"22": "S_AIR_QUALITY",
	"23": "S_CUSTOM",
	"24": "S_DUST",
	"25": "S_SCENE_CONTROLLER",
	"26": "S_RGB_LIGHT",
	"27": "S_RGBW_LIGHT",
	"28": "S_COLOR_SENSOR",
	"29": "S_HVAC",
	"30": "S_MULTIMETER",
	"31": "S_SPRINKLER",
	"32": "S_WATER_LEAK",
	"33": "S_SOUND",
	"34": "S_VIBRATION",
	"35": "S_MOISTURE",
	"36": "S_INFO",
	"37": "S_GAS",
	"38": "S_GPS",
	"39": "S_WATER_QUALITY",
}

// Set, Request type mapping for received messages
var setReqFieldMapForRx = map[string]string{
	"0":  "V_TEMP",
	"1":  "V_HUM",
	"2":  "V_STATUS",
	"3":  "V_PERCENTAGE",
	"4":  "V_PRESSURE",
	"5":  "V_FORECAST",
	"6":  "V_RAIN",
	"7":  "V_RAINRATE",
	"8":  "V_WIND",
	"9":  "V_GUST",
	"10": "V_DIRECTION",
	"11": "V_UV",
	"12": "V_WEIGHT",
	"13": "V_DISTANCE",
	"14": "V_IMPEDANCE",
	"15": "V_ARMED",
	"16": "V_TRIPPED",
	"17": "V_WATT",
	"18": "V_KWH",
	"19": "V_SCENE_ON",
	"20": "V_SCENE_OFF",
	"21": "V_HVAC_FLOW_STATE",
	"22": "V_HVAC_SPEED",
	"23": "V_LIGHT_LEVEL",
	"24": "V_VAR1",
	"25": "V_VAR2",
	"26": "V_VAR3",
	"27": "V_VAR4",
	"28": "V_VAR5",
	"29": "V_UP",
	"30": "V_DOWN",
	"31": "V_STOP",
	"32": "V_IR_SEND",
	"33": "V_IR_RECEIVE",
	"34": "V_FLOW",
	"35": "V_VOLUME",
	"36": "V_LOCK_STATUS",
	"37": "V_LEVEL",
	"38": "V_VOLTAGE",
	"39": "V_CURRENT",
	"40": "V_RGB",
	"41": "V_RGBW",
	"42": "V_ID",
	"43": "V_UNIT_PREFIX",
	"44": "V_HVAC_SETPOINT_COOL",
	"45": "V_HVAC_SETPOINT_HEAT",
	"46": "V_HVAC_FLOW_MODE",
	"47": "V_TEXT",
	"48": "V_CUSTOM",
	"49": "V_POSITION",
	"50": "V_IR_RECORD",
	"51": "V_PH",
	"52": "V_ORP",
	"53": "V_EC",
	"54": "V_VAR",
	"55": "V_VA",
	"56": "V_POWER_FACTOR",
}

// Stream message types mapping for received messages
var streamTypeMapForRx = map[string]string{
	"0": "ST_FIRMWARE_CONFIG_REQUEST",
	"1": "ST_FIRMWARE_CONFIG_RESPONSE",
	"2": "ST_FIRMWARE_REQUEST",
	"3": "ST_FIRMWARE_RESPONSE",
}

// Internal message types mapping for received messages
var internalTypeMapForRx = map[string]string{
	"0":  "I_BATTERY_LEVEL",
	"1":  "I_TIME",
	"2":  "I_VERSION",
	"3":  "I_ID_REQUEST",
	"4":  "I_ID_RESPONSE",
	"5":  "I_INCLUSION_MODE",
	"6":  "I_CONFIG",
	"7":  "I_FIND_PARENT",
	"8":  "I_FIND_PARENT_RESPONSE",
	"9":  "I_LOG_MESSAGE",
	"10": "I_CHILDREN",
	"11": "I_SKETCH_NAME",
	"12": "I_SKETCH_VERSION",
	"13": "I_REBOOT",
	"14": "I_GATEWAY_READY",
	"15": "I_SIGNING_PRESENTATION",
	"16": "I_NONCE_REQUEST",
	"17": "I_NONCE_RESPONSE",
	"18": "I_HEARTBEAT_REQUEST",
	"19": "I_PRESENTATION",
	"20": "I_DISCOVER_REQUEST",
	"21": "I_DISCOVER_RESPONSE",
	"22": "I_HEARTBEAT_RESPONSE",
	"23": "I_LOCKED",
	"24": "I_PING",
	"25": "I_PONG",
	"26": "I_REGISTRATION_REQUEST",
	"27": "I_REGISTRATION_RESPONSE",
	"28": "I_DEBUG",
	"29": "I_SIGNAL_REPORT_REQUEST",
	"30": "I_SIGNAL_REPORT_REVERSE",
	"31": "I_SIGNAL_REPORT_RESPONSE",
	"32": "I_PRE_SLEEP_NOTIFICATION",
	"33": "I_POST_SLEEP_NOTIFICATION",
}

// messages received from this internal type considered as a field data on the node
// Example battery level, isLocked?, RSSI, etc...
var internalValidFields = map[string]string{
	"I_BATTERY_LEVEL":          ml.FieldBatteryLevel,
	"I_DISCOVER_RESPONSE":      ml.FieldParentID,
	"I_HEARTBEAT_RESPONSE":     ml.FieldHeartbeat,
	"I_LOCKED":                 ml.FieldLocked,
	"I_SIGNAL_REPORT_RESPONSE": ml.FieldSignalStrength,
	"I_SKETCH_NAME":            ml.FieldName,
	"I_SKETCH_VERSION":         ml.LabelNodeVersion,
	"I_VERSION":                ml.LabelNodeLibraryVersion,
}

// MySensors should implement globally defined features for the request
// Some of the request could be very specific to MySensors
// Those features should be filtered here and should be implemented
// Other than this list all other requests will be ignored
var customValidActions = []string{
	// internal message type request
	"I_CONFIG",
	"I_ID_REQUEST",
	"I_TIME",

	// stream message type requests
	"ST_FIRMWARE_CONFIG_REQUEST",
	"ST_FIRMWARE_REQUEST",
}

// this struct used to construct payload metric type and unit
type payloadMetricTypeUnit struct{ Type, Unit string }

// map default metric types unit types for the fields
var metricTypeAndUnit = map[string]payloadMetricTypeUnit{
	"V_TEMP":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitCelsius},
	"V_HUM":                {mtsml.MetricTypeGaugeFloat, mtsml.UnitHumidity},
	"V_STATUS":             {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_PERCENTAGE":         {mtsml.MetricTypeGaugeFloat, mtsml.UnitPercent},
	"V_PRESSURE":           {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_FORECAST":           {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_RAIN":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_RAINRATE":           {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_WIND":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_GUST":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_DIRECTION":          {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_UV":                 {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_WEIGHT":             {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_DISTANCE":           {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_IMPEDANCE":          {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_ARMED":              {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_TRIPPED":            {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_WATT":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_KWH":                {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_SCENE_ON":           {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_SCENE_OFF":          {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_HVAC_FLOW_STATE":    {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_HVAC_SPEED":         {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_LIGHT_LEVEL":        {mtsml.MetricTypeGaugeFloat, mtsml.UnitPercent},
	"V_VAR1":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_VAR2":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_VAR3":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_VAR4":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_VAR5":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_UP":                 {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_DOWN":               {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_STOP":               {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_IR_SEND":            {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_IR_RECEIVE":         {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_FLOW":               {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_VOLUME":             {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_LOCK_STATUS":        {mtsml.MetricTypeBinary, mtsml.UnitNone},
	"V_LEVEL":              {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_VOLTAGE":            {mtsml.MetricTypeGaugeFloat, mtsml.UnitVoltage},
	"V_CURRENT":            {mtsml.MetricTypeGaugeFloat, mtsml.UnitAmpere},
	"V_RGB":                {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_RGBW":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_ID":                 {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_UNIT_PREFIX":        {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_HVAC_SETPOINT_COOL": {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_HVAC_SETPOINT_HEAT": {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_HVAC_FLOW_MODE":     {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_TEXT":               {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_CUSTOM":             {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_POSITION":           {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_IR_RECORD":          {mtsml.MetricTypeNone, mtsml.UnitNone},
	"V_PH":                 {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_ORP":                {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_EC":                 {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_VAR":                {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_VA":                 {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
	"V_POWER_FACTOR":       {mtsml.MetricTypeGaugeFloat, mtsml.UnitNone},
}
