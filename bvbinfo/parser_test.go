package bvbinfo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/stretchr/testify/assert"
)

func Test_GetMatches_Handles3SetMatch(t *testing.T) {
	// Given
	input := `<br>Match 61: <b><a href="player.asp?ID=5214">Phil Dalhausser</a> / <a href="player.asp?ID=1931">Nick Lucena</a> (3)</b> def. <a href="player.asp?ID=13453">Trevor Crabb</a> / <a href="player.asp?ID=1163">Sean Rosenthal</a> (4) 23-25, 21-18, 15-10 (1:15)`

	// When
	_, matches := bvbinfo.GetMatches(strings.NewReader(input), 0)

	// Then
	match := matches[0]
	assert.Equal(t, 5214, match.PlayerA.Id)
	assert.Equal(t, "Phil Dalhausser", match.PlayerA.Name)
	assert.Equal(t, "http://bvbinfo.com/images/photos/5214.jpg", match.PlayerA.ImgUrl)
	assert.Equal(t, 1931, match.PlayerB.Id)
	assert.Equal(t, "Nick Lucena", match.PlayerB.Name)
	assert.Equal(t, "http://bvbinfo.com/images/photos/1931.jpg", match.PlayerB.ImgUrl)
	assert.Equal(t, 13453, match.PlayerC.Id)
	assert.Equal(t, "Trevor Crabb", match.PlayerC.Name)
	assert.Equal(t, "http://bvbinfo.com/images/photos/13453.jpg", match.PlayerC.ImgUrl)
	assert.Equal(t, 1163, match.PlayerD.Id)
	assert.Equal(t, "Sean Rosenthal", match.PlayerD.Name)
	assert.Equal(t, "http://bvbinfo.com/images/photos/1163.jpg", match.PlayerD.ImgUrl)
	assert.False(t, match.IsForfeit)
	assert.Equal(t, "23-25", match.Set1)
	assert.Equal(t, "21-18", match.Set2)
	assert.Equal(t, "15-10", match.Set3)
}

func Test_GetMatches_Handles2ndSetRetired(t *testing.T) {
	// Given
	input := `<br>Match 12: <b><a href="player.asp?ID=16546">Andrea Abbiati</a> / <a href="player.asp?ID=10736">Tiziano Andreatta</a> Italy (31, Q27)</b> def. <a href="player.asp?ID=7145">Lombardo Ontiveros</a> / <a href="player.asp?ID=8011">Juan Virgen</a> Mexico (Q6) 26-24 retired (0:29)`

	// When
	_, matches := bvbinfo.GetMatches(strings.NewReader(input), 0)

	// Then
	match := matches[0]
	assert.Equal(t, 16546, match.PlayerA.Id)
	assert.Equal(t, 10736, match.PlayerB.Id)
	assert.Equal(t, 7145, match.PlayerC.Id)
	assert.Equal(t, 8011, match.PlayerD.Id)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_Handles3rdSetRetired(t *testing.T) {
	// Given
	input := `<br>Match 30: <b><a href="player.asp?ID=7710">Leonardo Lunardi</a> / <a href="player.asp?ID=11131">Daniel Virkus</a> (Q18)</b> def. <a href="player.asp?ID=7960">Wayne Leever</a> / <a href="player.asp?ID=8777">Jared Tucker</a> (Q47) 21-16, 8-2 retired (0:32)`

	// When
	_, matches := bvbinfo.GetMatches(strings.NewReader(input), 0)

	// Then
	match := matches[0]
	assert.Equal(t, 7710, match.PlayerA.Id)
	assert.Equal(t, 11131, match.PlayerB.Id)
	assert.Equal(t, 7960, match.PlayerC.Id)
	assert.Equal(t, 8777, match.PlayerD.Id)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_HandlesForfeit(t *testing.T) {
	// Given
	input := `<br>Match 2: <b><a href="player.asp?ID=13513">Juan Beltran</a> / <a href="player.asp?ID=14187">Zack Kweder</a> (Q32)</b> def. <a href="player.asp?ID=10935">Alex Pepke</a> / <a href="player.asp?ID=15591">Drew Pitlik</a> (Q33) by Forfeit`

	// When
	_, matches := bvbinfo.GetMatches(strings.NewReader(input), 0)

	// Then
	match := matches[0]
	assert.Equal(t, 13513, match.PlayerA.Id)
	assert.Equal(t, 14187, match.PlayerB.Id)
	assert.Equal(t, 10935, match.PlayerC.Id)
	assert.Equal(t, 15591, match.PlayerD.Id)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_GetsDateAndYear(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2018-fivb-gstaad-major-mens-matches.html")

	// When
	tournament, _ := bvbinfo.GetMatches(file, 3465)

	// Then
	assert.Equal(t, 3465, tournament.Id)
	assert.Equal(t, "Men's FIVB US$300,000 Gstaad Major", tournament.Name)
	assert.Equal(t, "July 9-14, 2018", tournament.Dates)
	assert.Equal(t, 2018, tournament.Year)
	assert.Equal(t, "male", tournament.Gender)
}

func Test_GetMatches_GetsDateAndYearWithLocation(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2018-avp-new-york-mens-matches.html")

	// When
	tournament, _ := bvbinfo.GetMatches(file, 3487)

	// Then
	assert.Equal(t, 3487, tournament.Id)
	assert.Equal(t, "Men's $100,000 AVP Gold Series NYC", tournament.Name)
	assert.Equal(t, "June 7-10, 2018", tournament.Dates)
	assert.Equal(t, 2018, tournament.Year)
	assert.Equal(t, "male", tournament.Gender)
}

func Test_GetMatches_GetsFemaleGender(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2017-avp-manhattan-beach-womens-matches.html")

	// When
	tournament, _ := bvbinfo.GetMatches(file, 3333)

	// Then
	assert.Equal(t, 3333, tournament.Id)
	assert.Equal(t, "Women's AVP $112,500 Manhattan Beach Open", tournament.Name)
	assert.Equal(t, "August 17-20, 2017", tournament.Dates)
	assert.Equal(t, 2017, tournament.Year)
	assert.Equal(t, "female", tournament.Gender)
}

func Test_GetMatches_ReturnsCorrectMatchCounts(t *testing.T) {
	// Given
	var tournaments = []struct {
		expectedTotalMatches int
		filePath             string
	}{
		{76, "2014-avp-st-petersburg-mens-matches.html"},
		{107, "2015-avp-manhattan-beach-mens-matches.html"},
		{159, "2017-avp-manhattan-beach-mens-matches.html"},
		{79, "2018-fivb-gstaad-major-mens-matches.html"},
	}

	for _, tt := range tournaments {
		t.Run(tt.filePath, func(t *testing.T) {
			// Given
			file, _ := os.Open("./testdata/" + tt.filePath)

			// When
			_, matches := bvbinfo.GetMatches(file, 0)

			// Then
			assert.Equal(t, tt.expectedTotalMatches, len(matches))
		})
	}
}

func Test_GetTournaments(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2017-avp-tournaments.html")

	// When
	tournaments := bvbinfo.GetTournaments(file)

	// Then
	assert.Equal(t, 16, len(tournaments))
	assert.Equal(t, 3320, tournaments[0].Id)
	assert.Equal(t, "Huntington Beach", tournaments[0].Name)
}

func Test_GetSeasons(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/all-seasons.html")

	// When
	seasons := bvbinfo.GetSeasons(file)

	// Then
	assert.Equal(t, 269, len(seasons))
	assert.Equal(t, "3", seasons[0].AssocID)
	assert.Equal(t, "2019", seasons[0].Year)
}
