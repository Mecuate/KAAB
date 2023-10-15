package fixtures

type UserDataType = []struct {
	Name     string
	Nick     string
	Email    string
	Password string
	Picture  string
}

var UserData = UserDataType{
	{
		Name:     "Daniel Robles",
		Nick:     "danroblo",
		Email:    "dan@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdnb.artstation.com/p/assets/covers/images/015/183/937/smaller_square/sefki-ibrahim-creation-animationfinal.jpg",
	},
	{
		Name:     "Yoshi",
		Nick:     "yosh",
		Email:    "yosh@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdnb.artstation.com/p/assets/covers/images/036/647/067/smaller_square/j-hill-j-hill-focus-thumbnail.jpg",
	},
	{
		Name:     "Daniel Alberto",
		Nick:     "daniel",
		Email:    "daniel@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdna.artstation.com/p/assets/images/images/016/256/454/20190301160733/smaller_square/mahmoud-salah-persp-2.jpg",
	},
	{
		Name:     "Carolina",
		Nick:     "caro",
		Email:    "caro@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdnb.artstation.com/p/assets/images/images/054/674/837/smaller_square/vlx-minguillo-3dportrait01.jpg",
	},
	{
		Name:     "Cinthya Estrada",
		Nick:     "cin",
		Email:    "cin@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdna.artstation.com/p/assets/images/images/058/113/264/smaller_square/phung-nguy-n-quang-girl-paper.jpg",
	},
	{
		Name:     "Jose Lorenzo",
		Nick:     "pepe",
		Email:    "jose@mecuate.org",
		Password: "Password-7",
		Picture:  "https://cdna.artstation.com/p/assets/images/images/043/414/048/smaller_square/marco-di-lucca-img-8537-2.jpg",
	},
}
