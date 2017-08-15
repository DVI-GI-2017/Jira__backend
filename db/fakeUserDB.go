package db

import "github.com/DVI-GI-2017/Jira__backend/models"

var UsersListFromFakeDB = models.Users{
	models.User{
		Email: "mbazley1@a8.net", FirstName: "Jeremy", LastName: "Moore",
		Tasks: models.Tasks{}, Password: "??04*products*GRAIN*began*58??",
		Bio: `Spent childhood selling wooden tops in Pensacola, FL. In 2008 I
was testing the market for sheep in Miami, FL. Was quite successful at promoting
yard waste in Tampa, FL. Spent 2001-2006 implementing bullwhips in the government
sector. Had a brief career buying and selling bullwhips in Edison, NJ. A real dynamo
when it comes to selling action figures for farmers.`},

	models.User{
		Email: "rcattermull0@storify.com", FirstName: "Crawford", LastName: "Eustis",
		Tasks: models.Tasks{}, Password: "//56.belong.SURE.fresh.16//",
		Bio: `Once had a dream of creating marketing channels for jigsaw puzzles in
Gainesville, FL. Spent 2001-2008 building bathtub gin for the government. What gets
me going now is consulting about Yugos on Wall Street. Earned praise for marketing
jack-in-the-boxes in Mexico. At the moment I'm selling dogmas with no outside help.
Enthusiastic about getting my feet wet with tobacco in Jacksonville, FL.`},

	models.User{
		Email: "bputtan6@discovery.com", FirstName: "Kurtis", LastName: "Chambers",
		Tasks: models.Tasks{}, Password: "--06$last$REST$prepared$76--",
		Bio: `Spent childhood licensing banjos in Salisbury, MD. Spent 2001-2008
analyzing puppets in Ohio. Once had a dream of implementing mosquito repellent on
Wall Street. Managed a small team investing in hugs in New York, NY. Was quite
successful at supervising the production of glucose in Naples, FL. Have a strong
interest in getting my feet wet with psoriasis in Fort Lauderdale, FL.`},
}
