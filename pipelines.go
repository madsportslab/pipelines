package main


const (
	APP_NAME										= "pipelines"
	APP_VERSION									= "1.0"
	APP_SYNC_PERIOD             = 12
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
	STR_EMPTY										= ""
	STR_PERIOD                  = "."
	STR_TAB											= "\t"
)

const (
	DATE_YEAR_ONLY							= "2006"
	NBA_DATE_FORMAT							= "20060102"
)

const (
	JOB_BOXSCORE               	= "boxscore"
	JOB_PLAYBYPLAY							= "playbyplay"
)

const (
	GAME_TYPE_PRESEASON					= 0
	GAME_TYPE_REGULAR           = 1
	GAME_TYPE_INSEASON          = 2
	GAME_TYPE_ALLSTAR           = 3
	GAME_TYPE_PLAYIN            = 4
	GAME_TYPE_PLAYOFF           = 5
	GAME_TYPE_RISING_STARS      = 6
	GAME_TYPE_UNKNOWN           = 7
)

const (
	GAME_LABEL_ALLSTAR              = "All-Star Game"
	GAME_LABEL_INSEASON           	= "In-Season Tournament"
	GAME_LABEL_INTL									= "NBA"
	GAME_LABEL_PLAYIN               = "Play-In"
	GAME_LABEL_PRESEASON						= "Preseason"
	GAME_LABEL_RISING_STARS					= "Rising Stars"
	GAME_STATUS_FINAL               = "Final"
	WEEK_NAME_ALLSTAR               = "All-Star"
	WEEK_NAME_PLAYIN                = "Play-In"
	
)

const (
	PLAYERS_PREFIX      		= "players"
	GAMES_PREFIX      			= "games"
	LEADERS_PREFIX      		= "leaders"
	STANDINGS_PREFIX    		= "standings"
  PLAYOFF_PREFIX          = "playoffs"
  REGULAR_PREFIX          = "regular"
	PRESEASON_PREFIX        = "preseason"
)

const (
	RESOURCE_BOXSCORE				= 0
	RESOURCE_PLAYBYPLAY     = 1
)

var currentSeason	string
