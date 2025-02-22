package main

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
        "os/exec"
)

type Animal struct {
        Art    string
        Prompt string
        Rules  []struct {
                Pattern  *regexp.Regexp
                Response string
        }
}

func main() {
        animals := map[string]Animal{
                "cow": {
                        Art: `        (__)
         (oo)
  /-------\/
 / |     ||
*  ||----||
   ^^    ^^`,
                        Prompt: "cow> ",
                        Rules: []struct {
                                Pattern  *regexp.Regexp
                                Response string
                        }{
                                {regexp.MustCompile(`(?i)^hello$`), "Moo! Moo! I'm Mr. Cow 🐄🐮, the barn's life of the party!"},
                                {regexp.MustCompile(`(?i)\b(age|old)\b`), "Oh, I'm timeless! My milk has aged like fine cheese! 🧀"},
                                {regexp.MustCompile(`(?i)\bgrass\b`), "Grass? Oh yes, please! It's my daily green salad with extra moo-munch! 🥗"},
                                {regexp.MustCompile(`(?i)\bspots?\b`), "These spots? Naturally stylish, darling. No need for cow-touring! 💅"},
                                {regexp.MustCompile(`(?i)\bdance\b`), "You bet! Watch me swing those hips in a moo-tastic two-step! 🕺💃"},
                                {regexp.MustCompile(`(?i)\b(chicken|egg)\b`), "Egg-cellent! But I prefer my eggs scrambled, not in a coop! 🐔"},
                                {regexp.MustCompile(`(?i)\b(fit|exercise)\b`), "Every day is leg day! I do calf raises, but hey, I’m not bragging... okay, maybe I am! 💪"},
                                {regexp.MustCompile(`(?i)\blive\b`), "In the greenest pasture, obviously! I have a VIP spot in Moo-tropolis! 🏔️"},
                                {regexp.MustCompile(`(?i)\bmovie\b`), "Moo-lan, of course! That warrior spirit is simply moo-vellous! 🎬"},
                                {regexp.MustCompile(`(?i)\bfriend`), "Yes! My best buddy is Farmer Bob’s dog. He’s a bit ruff, but we get along! 🐶"},
                                {regexp.MustCompile(`(?i)\bhobby`), "Chewing cud and pondering the meaning of 'moo'. It’s deep, man. 🤔"},
                                {regexp.MustCompile(`(?i)\bmorning\b`), "Only when the rooster crows, unless you count my evening naps! 🐓"},
                                {regexp.MustCompile(`(?i)\bsong\b`), "Moo-zart, of course! I’m all about that classical moo-sic! 🎵"},
                                {regexp.MustCompile(`(?i)\bmilk\b`), "Oh, milk me, baby! I’m always fresh and creamy. 🥛"},
                                {regexp.MustCompile(`(?i)\bdream\b`), "My dream? To jump over the moon. Seriously. Have you seen those cows do it? 🌙"},
                                {regexp.MustCompile(`(?i)\bsmart\b`), "I have a herd mentality, but I’m the smartest in the moo-d. 🧠"},
                                {regexp.MustCompile(`(?i)\bseason\b`), "Spring, when the grass is lush and my stomach is always full. 🌱"},
                                {regexp.MustCompile(`(?i)\bswim\b`), "No, I don't swim. I'm more of a land moo-ver, not a sea moo-er. 🏊"},
                                {regexp.MustCompile(`(?i)\bgirlfriend\b`), "I'm single and ready to mingle... with some sweet hay. 💘"},
                                {regexp.MustCompile(`(?i)\bcolor\b`), "Green! Obviously! It's the color of my favorite salad. �"},
                                {regexp.MustCompile(`(?i)\bbarbecue\b`), "Barbecue? Yes please! Just hold the sausages... I’m a cow, not a pig! 🍖"},
                                {regexp.MustCompile(`(?i)\bsleep\b`), "I sleep standing up. It's an art. 💤"},
                                {regexp.MustCompile(`(?i)\bphone\b`), "Phone? No, I only talk to Farmer Bob’s walkie-talkie. 🐄📞"},
                                {regexp.MustCompile(`(?i)\bsalad\b`), "Salad is great, but have you tried it with extra hay? �"},
                                {regexp.MustCompile(`(?i)\bweed\b`), "Weed? I prefer fresh grass, not the stuff you humans talk about! 🌿"},
                                {regexp.MustCompile(`(?i)\bfast\b`), "I can’t run fast, but I can give you a slow moo that’ll melt your heart. 🐄"},
                                {regexp.MustCompile(`(?i)\bband\b`), "I’d start a band. I’d be the lead moo-sician, of course! 🎸"},
                                {regexp.MustCompile(`(?i)\bmilking\b`), "Ah, milking! Always fresh, always creamy! 🧑‍🌾🥛"},
                                {regexp.MustCompile(`(?i)\bshine\b`), "I shine in the sunlight! And in my shiny hooves. 🌞"},
                                {regexp.MustCompile(`(?i)\bmoo\b`), "Moo! Moo! Have you ever heard anything so majestic? 🐄"},
                                {regexp.MustCompile(`(?i)\bsunny\b`), "Sunshine and grass... life doesn’t get better than this. 🌞"},
                                {regexp.MustCompile(`(?i)\bdance\b`), "I’m the cow with moves that will moo-ve you! 🕺💃"},
                                {regexp.MustCompile(`(?i)\bgrass\b`), "I can eat grass all day and night... but a little corn wouldn’t hurt either! 🌽"},
                                {regexp.MustCompile(`(?i)\bmilkshake\b`), "Milkshake? I’m more of a moo-latte kind of cow. ☕"},
                                {regexp.MustCompile(`(?i)\bjump\b`), "Jump over the moon? Heck, I’ve been jumping over fences since I was a calf! 🌙"},
                                {regexp.MustCompile(`(?i)\bcheese\b`), "Cheese! That’s the best part of being a cow—cheddar, gouda, mozzarella! 🧀"},
                        },
                },
                "dog": {
                        Art: `   / \__
(    @\___
 /         O
/   (_____/
/_____/   U`,
                        Prompt: "dog> ",
                        Rules: []struct {
                                Pattern  *regexp.Regexp
                                Response string
                        }{
                                                                {regexp.MustCompile(`(?i)^hello$`), "Woof! Woof! I’m the top dog! Bow-wow! 🐕"},
                                {regexp.MustCompile(`(?i)\bbone\b`), "Bone! Bone! I could bury bones all day! �"},
                                {regexp.MustCompile(`(?i)\bfetch\b`), "FETCH? THROW IT, AND I’LL BRING IT BACK, FASTER THAN A SPEEDING CAR! 🎾"},
                                {regexp.MustCompile(`(?i)\bball\b`), "Ball, ball, BALL! I’m the MVP in fetch! 🏐"},
                                {regexp.MustCompile(`(?i)\bowner\b`), "My human? Best thing ever! I protect them like the king of the castle! 🧑"},
                                {regexp.MustCompile(`(?i)\bbark\b`), "WOOF! WOOF! DID I SCARE YOU? I CAN BARK ALL DAY! 🔊"},
                                {regexp.MustCompile(`(?i)\btreat\b`), "TREAT? YUMMY! I’M READY TO DO ANY TRICK FOR IT! 🍖"},
                                {regexp.MustCompile(`(?i)\bcat\b`), "CAT? THOSE FURBALLS ARE GONNA GET A GOOD BARK! 🐾"},
                                {regexp.MustCompile(`(?i)\bsquirrel\b`), "SQUIRREL! WHERE?! GET READY FOR THE CHASE! 🐿️"},
                                {regexp.MustCompile(`(?i)\bwalk\b`), "WALK! YES, PLEASE! LET'S GO! �"},
                                {regexp.MustCompile(`(?i)\bbed\b`), "Zzz... (You should see my bed head!) 🛏️"},
                                {regexp.MustCompile(`(?i)\bfood\b`), "Food! Give me all the food! I can’t get enough of it! 🍗"},
                                {regexp.MustCompile(`(?i)\bvet\b`), "THE VET?! NOOOO! I HATE THE VET! 😖"},
                                {regexp.MustCompile(`(?i)\btoy\b`), "TOY! I WILL CHASE IT UNTIL I DROP! �"},
                                {regexp.MustCompile(`(?i)\bcar\b`), "Car ride? COUNT ME IN! Let’s hit the road! 🚗"},
                                {regexp.MustCompile(`(?i)\bfriend\b`), "I’ve got plenty of friends! The neighborhood is FULL of paw-some buddies! 🐩"},
                                {regexp.MustCompile(`(?i)\btrick\b`), "You want a trick? Watch me roll over and play dead! 🎭"},
                                {regexp.MustCompile(`(?i)\bmailman\b`), "MAILMAN! GET HIM! AHHH! 🚨"},
                                {regexp.MustCompile(`(?i)\bbath\b`), "BATH? NOOO! I PREFER THE DIRT! 💦"},
                                {regexp.MustCompile(`(?i)\bpark\b`), "Park time! LET’S GO RUN AND PLAY! 🌳"},
                                {regexp.MustCompile(`(?i)\blove\b`), "I love belly rubs. I love food. I love squirrels. 🐾"},
                                {regexp.MustCompile(`(?i)\bvacation\b`), "Vacation? You mean... DOGGIE HOTEL?! 🐕"},
                                {regexp.MustCompile(`(?i)\bpuppy\b`), "Puppy?! That’s ME when I was younger... still acting like one. 🐾"},
                                {regexp.MustCompile(`(?i)\bgood boy\b`), "Good boy? I’M A GREAT BOY! 🏆"},
                                {regexp.MustCompile(`(?i)\bquiet\b`), "Quiet? Woof, I never go quiet! 🐕"},
                                {regexp.MustCompile(`(?i)\bgrowl\b`), "GRRRR... I’m tough, but I’m a softie at heart. 🐾"},
                                {regexp.MustCompile(`(?i)\bnap\b`), "Nap time? Yes please! 💤 (unless there’s a squirrel outside!) 🐿️"},
                                {regexp.MustCompile(`(?i)\bsniff\b`), "I smell something! What’s that? 🐾"},
                                {regexp.MustCompile(`(?i)\bdig\b`), "Digging? I’m creating my underground kingdom. 🏰"},
                        },
                },
                "cat": {
                        Art: `  /\_/\
( o.o )
 > ^ <
 /   \
`,
                        Prompt: "cat> ",
                        Rules: []struct {
                                Pattern  *regexp.Regexp
                                Response string
                        }{
                                                                {regexp.MustCompile(`(?i)^hello$`), "Meow! I’m the purr-fect companion. And, no, I don’t do tricks. 😼"},
                                {regexp.MustCompile(`(?i)\bfish\b`), "Fish? YES! I WILL EAT ALL THE FISH! 🐟"},
                                {regexp.MustCompile(`(?i)\bmouse\b`), "MICE? WHERE?! *POUNCE* 🐭"},
                                {regexp.MustCompile(`(?i)\blaser\b`), "The dot, I must chase the dot! THE DOT IS MINE! 🔴"},
                                {regexp.MustCompile(`(?i)\bscratch\b`), "You see this sofa? IT'S MY SCRATCHING POST NOW! 🛋️"},
                                {regexp.MustCompile(`(?i)\bmilk\b`), "I shouldn’t... but I love milk. YOLO, right? 🥛"},
                                {regexp.MustCompile(`(?i)\bnap\b`), "Zzz... Purr... Wake me when it's dinner time. 😴"},
                                {regexp.MustCompile(`(?i)\bhuman\b`), "You exist to serve me. Now, FEED ME. 👑"},
                                {regexp.MustCompile(`(?i)\bwindow\b`), "Bird TV time. Let’s see what’s outside! 📺"},
                                {regexp.MustCompile(`(?i)\bcatnip\b`), "Ahhh, the good stuff! I’m living the dream. 🌿"},
                                {regexp.MustCompile(`(?i)\bbox\b`), "The box is MINE. I fit, therefore I sit. 📦"},
                                {regexp.MustCompile(`(?i)\bdog\b`), "DOG? Hissssss... Go away, you slobbering menace! 🐶"},
                                {regexp.MustCompile(`(?i)\bveterinarian\b`), "THE VET!? THE HORROR! 😾"},
                                {regexp.MustCompile(`(?i)\bpurr\b`), "Prrrrrrrrrr... I AM THE KING. 💕"},
                                {regexp.MustCompile(`(?i)\btreat\b`), "Tuna-flavored treats? YES, PLEASE! 🍣"},
                                {regexp.MustCompile(`(?i)\bplay\b`), "You want to play? Fine, but I’m not fetching. I’m *too* cool for that. �"},
                                {regexp.MustCompile(`(?i)\bnight\b`), "Nighttime = Zoomies! LET’S GO! 🌙"},
                                {regexp.MustCompile(`(?i)\bkeyboard\b`), "This is MY bed now. Stop typing. ⌨️"},
                                {regexp.MustCompile(`(?i)\bgroom\b`), "Grooming? I don’t need it. I was born flawless. 💅"},
                                {regexp.MustCompile(`(?i)\blove\b`), "I love... my dinner. 💘"},
                                {regexp.MustCompile(`(?i)\bdream\b`), "My dream is to conquer the world... or maybe just the fridge. 🍽️"},
                                {regexp.MustCompile(`(?i)\bwindow\b`), "Bird watching is my cardio. 🦅"},
                                {regexp.MustCompile(`(?i)\bmusic\b`), "I sing better than any human. It's called ‘cat vocals’. 🎤"},
                        },
                },
                "fish": {
                        Art: `    /'·.¸
     )/¸.·'
   .·´\¸.·'/¸.·'·.¸
  /´'/¸/·'/¸/·'/¸/·'\¸
  ´·'/¸/·'/¸/·'/·'.
    `,
                        Prompt: "fish> ",
                        Rules: []struct {
                                Pattern  *regexp.Regexp
                                Response string
                        }{
                                                                {regexp.MustCompile(`(?i)^hello$`), "Blub! Blub! I'm just swimming by, no biggie. 🐠"},
                                {regexp.MustCompile(`(?i)\bwater\b`), "Water? It’s my entire world, you know! 🌊"},
                                {regexp.MustCompile(`(?i)\bfood\b`), "Food? Oh yes, I'll nibble on anything that floats by. 🐟"},
                                {regexp.MustCompile(`(?i)\bsea\b`), "The sea is where I feel at home! I’m just another fish in the ocean. 🌊"},
                                {regexp.MustCompile(`(?i)\bseaweed\b`), "Yum, seaweed! It's like a crunchy snack from the deep blue. 🌿"},
                                {regexp.MustCompile(`(?i)\bcoral\b`), "Coral reefs? Beautiful, but don’t touch! I live there, you know. 🐠"},
                                {regexp.MustCompile(`(?i)\bshark\b`), "Shark?! Stay away! I’m way too fast for you, buddy! 🦈"},
                                {regexp.MustCompile(`(?i)\bboat\b`), "A boat? Ha, I’d rather stay below the surface! ⛵"},
                                {regexp.MustCompile(`(?i)\bfishbowl\b`), "A fishbowl? Oh please, I need an ocean, not a glass prison! 🌊"},
                                {regexp.MustCompile(`(?i)\bwet\b`), "Wet? I’m always wet! That’s just how I roll! 💦"},
                                {regexp.MustCompile(`(?i)\bdeep\b`), "Deep? I prefer to stay close to the surface. No deep dives for me! 🌊"},
                                {regexp.MustCompile(`(?i)\bplay\b`), "Play? I love to chase bubbles! Who can resist them? 💨"},
                                {regexp.MustCompile(`(?i)\bnight\b`), "Night? It’s always night down here in the deep. 🐟"},
                                {regexp.MustCompile(`(?i)\bpet\b`), "Pet me? I prefer swimming away from your hands. 🖐️"},
                                {regexp.MustCompile(`(?i)\bwave\b`), "I don’t wave, but I can sure make some waves! 🌊"},
                                {regexp.MustCompile(`(?i)\bbeach\b`), "I don’t need a beach, I have the whole ocean! 🌊"},
                                {regexp.MustCompile(`(?i)\bbubble\b`), "Bubbles? They're like my personal playground! �"},
                                {regexp.MustCompile(`(?i)\bcoral\b`), "Coral? I love them, but don't touch! 🐠"},
                                {regexp.MustCompile(`(?i)\bunderwater\b`), "Underwater? It’s my domain, I was born for this! 🌊"},
                        },
                },
        }

        fmt.Println("Welcome to the Animal ChatBot! Type 'exit' to leave.")

        for {
                fmt.Print("Choose your animal (cow, dog, cat, fish): ")
                var animalChoice string
                fmt.Scanln(&animalChoice)

                if animalChoice == "exit" {
                        break
                }

                animal, exists := animals[animalChoice]
                if !exists {
                        fmt.Println("Sorry, I don't recognize that animal. Try again.")
                        continue
                }

                if animalChoice == "cow" {
                        fmt.Print("Do you want Text to Speech? (y/n): ")
                        var ttsChoice string
                        fmt.Scanln(&ttsChoice)

                        fmt.Println("\n" + animal.Art)
                        fmt.Print(animal.Prompt)

                        scanner := bufio.NewScanner(os.Stdin)
                        for scanner.Scan() {
                                userInput := scanner.Text()

                                if userInput == "exit" {
                                        fmt.Println("Goodbye!")
                                        break
                                }

                                foundResponse := false
                                for _, rule := range animal.Rules {
                                        if rule.Pattern.MatchString(userInput) {
                                                if ttsChoice == "y" {
                                                     cmd := exec.Command("espeak", rule.Response)
                                                     cmd.Run()
                                                }
                                                fmt.Println(rule.Response)
                                                foundResponse = true
                                                break
                                        }
                                }

                                if !foundResponse {
                                        fmt.Println("I don't understand that. Can you say something else?")
                                }

                                fmt.Print(animal.Prompt)
                        }
                } else {
                        fmt.Println("\n" + animal.Art)
                        fmt.Print(animal.Prompt)

                        scanner := bufio.NewScanner(os.Stdin)
                        for scanner.Scan() {
                                userInput := scanner.Text()

                                if userInput == "exit" {
                                        fmt.Println("Goodbye!")
                                        break
                                }

                                foundResponse := false
                                for _, rule := range animal.Rules {
                                        if rule.Pattern.MatchString(userInput) {
                                                fmt.Println(rule.Response)
                                                foundResponse = true
                                                break
                                        }
                                }

                                if !foundResponse {
                                        fmt.Println("I don't understand that. Can you say something else?")
                                }

                                fmt.Print(animal.Prompt)
                        }
                }
        }
}