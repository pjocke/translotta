package main

func MapSearch(s string) []string {
	words := make(map[string][]string)

	words["happy"] = []string{"smile", "slightly_smiling_face", "wink"}
	words["glad"] = []string{"smile", "slightly_smiling_face", "wink"}
	words["joyful"] = []string{"smile", "slightly_smiling_face", "wink"}
	words["sad"] = []string{"disappointed", "angry", "worried"}
	words["angry"] = []string{"disappointed", "angry", "worried"}
	words["pissed"] = []string{"disappointed", "angry", "worried"}
	words["scared"] = []string{"scream", "fearful", "cold_sweat"}
	words["afraid"] = []string{"scream", "fearful", "cold_sweat"}
	words["frightened"] = []string{"scream", "fearful", "cold_sweat"}

	return words[s]
}
