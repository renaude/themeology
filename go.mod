module themeology

go 1.23.4

require (
	themeology.local/config v0.0.0
	themeology.local/discord v0.0.0
	themeology.local/waybar v0.0.0
)

replace themeology.local/config => ./config

replace themeology.local/discord => ./discord

replace themeology.local/waybar => ./waybar
