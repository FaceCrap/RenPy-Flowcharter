package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleTags(t *testing.T) {
	t.Parallel()
	detect := initializeDetectors()

	testCases := []struct {
		id         int
		line       string
		updatedTag Tag
	}{
		// Empty
		{0, "jump truc:", Tag{}},
		// Basic : Ignore Label
		{1, "label truc: # Flowcharter: IGNORE", Tag{ignore: true}},
		// Lot of comments
		{2, "label truc(variable='arg') : # bla Flowcharter: this is a title", Tag{title: true}},
		// Comments but nothing special
		{3, " narr \"This is a test\" # nothing special in comments", Tag{}},
		// triggers Flowcharter but no more
		{4, " narr \"This is a test\" # Flowcharter: nothing special", Tag{}},
		// case sensitivity
		{5, " narr \"This is a test\" # renpY-grapHvIz: TITLE", Tag{title: true}},
		// flow break
		{6, "  # renpY-grapHvIz: BREAK", Tag{breakFlow: true}},
		// GAME OVER FLOW
		{7, " label truc: # renpY-grapHvIz: GAMEOVER", Tag{gameOver: true}},
		// SKIPLINK
		{8, " label truc: # renpY-grapHvIz: SKIPLINK", Tag{skipLink: true}},
		// INGAME_LABEL
		{9, "# Flowcharter: INGAME_LABEL(0,fake_label)", Tag{inGameLabel: true, inGameIndent: 0}},
		// INGAME_JUMP
		{10, " # Flowcharter: INGAME_JUMP(4,to_label)", Tag{inGameJump: true, inGameIndent: 4}},
		{10, " # Flowcharter: INGAME_JUMP (8,  to_label ) ", Tag{inGameJump: true, inGameIndent: 8}},
		// FAKE_LABEL
		{11, " # Flowcharter: FAKE_LABEL(label_name)", Tag{fakeLabel: true}},
		{11, " # Flowcharter: FAKE_LABEL(label_name)", Tag{fakeLabel: true}},
		// FAKE_JUMP
		{12, " # Flowcharter: FAKE_JUMP(fake_label,to_label)", Tag{fakeJump: true}},
		{12, " # Flowcharter: FAKE_JUMP ( fake_label , to_label ) ", Tag{fakeJump: true}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			context := Context{}
			err := context.handleTags(tc.line, detect)
			require.NoError(t, err, "Error in tags test %v", tc.id)
			require.Equal(t, tc.updatedTag, context.tags, "Error in tags test %v", tc.id)
		})
	}
}
