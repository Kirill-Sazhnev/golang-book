package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type player struct {
	name   string
	reader *bufio.Reader
}

type action func(*player, []string) []string

func drawScreen(s string) {
	cls()
	fmt.Println(s)
	fmt.Println()
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func slowPrint(str string) {
	for _, sym := range str {
		fmt.Print(string(sym))
		time.Sleep(5e7)
	}
	fmt.Println()
}

func (player *player) input() string {
	input, err := player.reader.ReadString('\n')
	if err != nil {
		log.Println("error input: ", err)
	}
	return input[:len(input)-2]
}

func (player *player) branch(points []string, actions []action) {
	pointMap := make(map[string]int, len(points))
	for i, val := range points {
		pointMap[val] = i
	}
	for len(points) > 0 {
		for i, val := range points {
			fmt.Printf("[%d] %s\n", i+1, val)
			time.Sleep(1e9)
		}
		fmt.Print("Enter the number of your choice: ")
		point := player.input()

		slct, err := strconv.Atoi(point)
		if err != nil || slct < 1 || slct > len(points) {
			slowPrint("Stop conveing edge tests. Just pick the number:")
			continue
		}

		if ix, ok := pointMap[points[slct-1]]; ok {
			points = actions[ix](player, points)
		} else {
			slowPrint("It's not that hard. Just pick the number:")
		}
	}
}

func remove(sl []string, item string) []string {
	ix := 0
	for i, val := range sl {
		if val == item {
			ix = i + 1
			break
		}
	}
	switch ix {
	case 1:
		return sl[1:]
	case len(sl):
		return sl[:ix-1]
	default:
		return append(sl[:ix-1], sl[ix:]...)
	}
}

func main() {
	drawScreen(dragon)
	fmt.Println()
	slowPrint("Welcome to the DOSABLO 4.5! How should I call you?")
	fmt.Print("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
inpName:
	name, err := reader.ReadString('\n')
	if err != nil || len(name) < 3 {
		slowPrint("Stop conveing edge tests. Just enter a name:")
		goto inpName
	}
	name = name[:len(name)-2]
	player := &player{
		name:   name,
		reader: reader,
	}

	slowPrint(fmt.Sprintf("The time has come %s. The evil is here to fight", name))
	time.Sleep(2e9)
	drawScreen(wizard)

	slowPrint(fmt.Sprintf("Pick your hero's role %s. We have 3 roles for you to choose from.", name))
	slowPrint("They fit your playstyle the best:")
	time.Sleep(1e8)
	player.branch([]string{"Healer", "Healer", "Healer"}, []action{plug, plug, plug})

	slowPrint(fmt.Sprintf("Great choice! So be it - %s the Healer", name))
	time.Sleep(2e9)

	drawScreen(maps)
	slowPrint("Before you face the Dreadful Dosablo, you need to feel your prougress baaaaar and earn some exp.\n")
	slowPrint("Here is the world map with following points of interest to investigate:")
	time.Sleep(1e9)

	points := []string{"Mountains", "Forest", "Ship"}
	actions := []action{toMountains, toForest, toShip}
	player.branch(points, actions)

	slowPrint("Finaly! You've earned enough exp and gear score to face main antagonist: The DOSABLO!")
	time.Sleep(3e9)
	drawScreen(demon)
	slowPrint(fmt.Sprintf("Dreadful demon confronts you %s! What you gonna do?", player.name))

	points = []string{"Compare him with his predesessor DOSABLO 3", "Aim for the head to land a headshot", "Use your ultimate"}
	actions = []action{compare, headshot, ultimate}
	player.branch(points, actions)
	drawScreen(hbj)
	time.Sleep(8e9)
}

func plug(player *player, points []string) []string { return []string{} }

func toMountains(player *player, points []string) []string {
	drawScreen(mountain)
	slowPrint("What a view! Mountains, fresh air and sound of a free wind.")
	slowPrint(fmt.Sprintf("The spirit of exploration grants a small amount of exp to %s the Healer.", player.name))
	player.branch([]string{"Shot some selfies", "Come back to the map"}, []action{selfie, plug})

	drawScreen(maps)
	if len(points) > 1 {
		slowPrint("Remaining points of interest to investigate:")
	}
	return remove(points, "Mountains")
}

func selfie(player *player, points []string) []string {
	drawScreen(mntnFlash)
	time.Sleep(1e7)
	drawScreen(mountain)
	time.Sleep(1e9)
	drawScreen(mntnFlash)
	time.Sleep(1e7)
	drawScreen(mountain)
	time.Sleep(1e8)
	drawScreen(mntnFlash)
	time.Sleep(1e7)
	drawScreen(mountain)
	slowPrint("The stories have been posted on IG, let's go back now.")
	time.Sleep(1e8)
	return remove(points, "Shot some selfies")
}

func toForest(player *player, points []string) []string {
	drawScreen(forest)
	slowPrint(fmt.Sprintf("%s slowly approaches dark woods, admiring the dynamically animated branches of surrounding bushes and trees,\n", player.name))
	time.Sleep(1e9)
	slowPrint("while suddenly...")
	time.Sleep(3e9)

	drawScreen(arachna)
	slowPrint(fmt.Sprintf("A fierceful ARACHNA attacks you %s! What you gonna do?", player.name))
	time.Sleep(1e8)
	newPoints := []string{"Put a curse: GAVNO OBOSSANNOE", "Charge forward and scream out loud: \"Gde BLYAT' vse?\"", "Cast a hex: LOOZOR"}
	actions := []action{curse, charge, hex}
	player.branch(newPoints, actions)
	time.Sleep(1e9)

	drawScreen(maps)
	if len(points) > 1 {
		slowPrint("Remaining points of interest to investigate:")
	}
	return remove(points, "Forest")
}

func curse(player *player, points []string) []string {
	drawScreen(archnStink)
	slowPrint("ARACHNA suffers severe shit damage with a pisscrit!\n")
	if len(points) < 3 {
		slowPrint("The monster defeated! You earn a decent amount of exp.\n")
		time.Sleep(1e9)
		return []string{}
	}
	slowPrint("It bloody stinks though, it survives the attack. \n")
	slowPrint(fmt.Sprintf("What's your next move %s?\n", player.name))
	return remove(points, "Put a curse: GAVNO OBOSSANNOE")
}

func charge(player *player, points []string) []string {
	slowPrint("\nYou charge forward boldly but the thing is you are healer and melee hit chance is ridiculously low.\n")
	slowPrint("MISS!\n")
	slowPrint(fmt.Sprintf("What's your next move %s?\n", player.name))
	return remove(points, "Charge forward and scream out loud: \"Gde BLYAT' vse?\"")
}

func hex(player *player, points []string) []string {
	slowPrint("\nARACHNA can't live the fact it is not that special anymore.\n")
	slowPrint("The hex drains ARACHNA's life energy.\n")

	if len(points) < 3 {
		slowPrint("The monster defeated! You earn a decent amount of exp.\n")
		time.Sleep(1e9)
		return []string{}
	}
	slowPrint("Making it a number 2 doesn't kill it\n")
	slowPrint(fmt.Sprintf("What's your next move %s?\n", player.name))
	return remove(points, "Cast a hex: LOOZOR")
}

func toShip(player *player, points []string) []string {
	drawScreen(ship)
	slowPrint("Ship is drifting along the shore and looks abandoded.\n")
	slowPrint("Seems like crew gone missing in the forest.\n")

	player.branch([]string{"Search ship hold for a loot", "Sail away from here to a new horizons", "Come back to main map"}, []action{loot, sail, plug})

	drawScreen(maps)
	if len(points) > 1 {
		slowPrint("Remaining points of interest to investigate:")
	}
	return remove(points, "Ship")
}

func loot(player *player, points []string) []string {
	drawScreen(equipment)
	slowPrint(fmt.Sprintf("%s searched every crate, chest and barrel as his friend Jafree taught him.\n", player.name))
	slowPrint("You earn some decent gear and potions.\n")
	time.Sleep(1e9)
	slowPrint("What would be the next?")

	return remove(points, "Search ship hold for a loot")
}

func sail(player *player, points []string) []string {
	slowPrint("\nWhat?! You can't just leave this land for good.\n")
	slowPrint(fmt.Sprintf("%s you were destined to save the life of all the people!\n", player.name))
	slowPrint("Let's move further:")
	return remove(points, "Sail away from here to a new horizons")
}

func compare(player *player, points []string) []string {
	for i := 0; i < 4; i++ {
		drawScreen(demonLaugh)
		time.Sleep(1e8)
		drawScreen(demon)
		time.Sleep(1e8)
	}
	slowPrint("\nYou've listed all the differences in game mechanic, design as well as new features.\n")
	slowPrint("This cuts DOSABLO's enrage timer by half! You have to hurry!\n")

	return remove(points, "Compare him with his predesessor DOSABLO 3")
}

func headshot(player *player, points []string) []string {
	for i := 0; i < 4; i++ {
		drawScreen(demonLaugh)
		time.Sleep(1e8)
		drawScreen(demon)
		time.Sleep(1e8)
	}
	slowPrint(fmt.Sprintf("%s is aiming with so much effort and tension. But it is sooo hard to aim in this gaaame.\n", player.name))
	slowPrint("After a few attemts to land a headshot you were able to melt down a half of his HP bar.\n")
	points = remove(points, "Aim for the head to land a headshot")
	if len(points) == 0 || len(points) == 1 && points[0] == "Compare him with his predesessor DOSABLO 3" {
		time.Sleep(2e9)
		drawScreen(angels)
		slowPrint("The monster is defeated and inprisoned by the High Angels!\n")
		slowPrint("You are the saver of this land!.\n")
		time.Sleep(4e9)
		return []string{}
	}
	return points
}

func ultimate(player *player, points []string) []string {
	slowPrint(fmt.Sprintf("\n%s is eager to use his ultimate but a rapid Dosablo cast of ancient disease \"ЗАЖОПИТЬ\" kicks in.\n", player.name))
	slowPrint("Luckily you picked a healer role and easily dispel the disease.\n")
	slowPrint("A devastating ultimate strike eats half of Dosablo HP bar!\n")

	points = remove(points, "Use your ultimate")
	if len(points) == 0 || len(points) == 1 && points[0] == "Compare him with his predesessor DOSABLO 3" {
		time.Sleep(2e9)
		drawScreen(angels)
		slowPrint("The monster is defeated and inprisoned by the High Angels!\n")
		slowPrint("You are the saver of this land!.\n")
		time.Sleep(4e9)
		return []string{}
	}
	slowPrint(fmt.Sprintf("What's your next move %s?\n", player.name))
	return points
}
