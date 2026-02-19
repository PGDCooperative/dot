package client

type UIState map[string]bool

func GetUIState() UIState {
	uistate := UIState{
		"GamePlay":     false,
		"MainMenu":     true,
		"SettingsMenu": false,
	}
	return uistate
}
