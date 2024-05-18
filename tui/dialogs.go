package tui

import (
	dialog "github.com/ajayd-san/teaDialog"
)

const (
	dialogRemoveContainer dialog.DialogType = iota
	dialogPruneContainers
	dialogRemoveImage
)

func getRemoveContainerDialog(storage map[string]string) dialog.Dialog {
	prompts := []dialog.Prompt{
		dialog.MakeTogglePrompt("remVols", "Remove volumes?"),
		dialog.MakeTogglePrompt("remLinks", "Remove links?"),
		dialog.MakeTogglePrompt("force", "Force?"),
	}

	return dialog.InitDialogue("Remove Container Options:", prompts, dialogRemoveContainer, storage)
}

func getPruneContainersDialog(storage map[string]string) dialog.Dialog {
	prompts := []dialog.Prompt{
		dialog.MakeOptionPrompt("confirm", "This will remove all stopped containers, are your sure?", []string{"Yes", "No"}),
	}

	return dialog.InitDialogue("Prune Containers: ", prompts, dialogPruneContainers, storage)
}

func getRemoveImageDialog(storage map[string]string) dialog.Dialog {
	prompts := []dialog.Prompt{
		dialog.MakeTogglePrompt("force", "Force"),
		dialog.MakeTogglePrompt("pruneChildren", "Prune Children"),
	}

	return dialog.InitDialogue("Remove Image Options:", prompts, dialogRemoveImage, storage)
}
