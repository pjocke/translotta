package main

func StructSearch(s string) []string {
	type Words struct{
		Trigger		[]string
		Replacement	[]string
	}

	w := []Words{
		{
			Trigger: []string{
				"happy",
				"glad",
				"joyful",
			},
			Replacement: []string{
				"smile",
				"slightly_smiling_face",
				"wink",
			},
		},
		{
			Trigger: []string{
				"sad",
				"angry",
				"pissed",
			},
			Replacement: []string{
				"disappointed",
				"angry",
				"worried",
			},
		},
		{
			Trigger: []string{
				"scared",
				"afraid",
				"frightened",
			},
			Replacement: []string{
				"scream",
				"fearful",
				"cold_sweat",
			},
		},
	}

	for v := range w {
		for t := range w[v].Trigger {
			if w[v].Trigger[t] == s {
				return w[v].Replacement
			}
		}
	}

	return nil
}
