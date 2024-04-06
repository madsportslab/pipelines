package main


const (
	APP_NAME										= "pipelines"
	APP_VERSION									= "1.0"
	APP_SYNC_PERIOD             = 6
)

const (
	PIPELINE_DEFAULT_HOST				= "127.0.0.1"
	PIPELINE_DEFAULT_PORT       = "8686"
)

const (
	FORM_PARAM_IDS							= "ids"
	FORM_PARAM_LEAGUE						= "league"
	FORM_PARAM_SEASON						= "season"
	FORM_PARAM_YEAR							= "year"
)

const (
	DELIMITER_COMMA							= ","
)

const (
	SCHEDULE_JSON               = "schedule.json"
	EXT_JSON										= ".json"
	EXT_PBP                     = ".pbp"
	EXT_PARQUET									= ".parquet"
)

const (
	GAME_FINAL									= "Final"
)

const (
	STR_TAB											= "\t"
	STR_EMPTY										= ""
)

const (
	DATE_YEAR_ONLY							= "2006"
	NBA_DATE_FORMAT							= "20060102"
)


var currentSeason	string
