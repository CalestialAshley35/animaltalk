# AnimalTalk ChatBot 🐾 v5.0.0

**Interact with ASCII-animated animals through pattern-matching conversations!**

## 🌟 Features
- **4 Unique Animals**: Cow, Dog, Cat, and Fish with distinct personalities
- **100+ Dynamic Responses**: Context-aware replies using regex pattern matching
- **ASCII Art Animation**: Unique character art for each animal
- **Interactive Prompts**: Species-specific command line interfaces
- **Extensible Architecture**: Easy to add new animals/responses
- **Case-Insensitive Matching**: Understands inputs in any capitalization
- **Cross-Platform**: Work on Every System
- **Session Management**: Switch animals mid-conversation

## 🚀 Installation
```bash
git clone https://github.com/CalestialAshley35/animaltalk.git
cd animaltalk
go run animaltalk.go
```

## � Usage Guide
1. **Start the program**:
   ```bash
   go run animaltalk.go
   ```
2. **Choose an animal** from the initial prompt
3. **Chat with your selected animal** using natural language
4. Type `exit` at any time to end the session

**Example Session**:
```
Choose your animal (cow, dog, cat, fish): cow

        (__)  
         (oo)  
  /-------\/  
 / |     ||  
*  ||----||  
   ^^    ^^
cow> Hello
Moo! Moo! I'm Mr. Cow 🐄🐮, the barn's life of the party!
```

## Each Animal Responses

**Cow:**

Q: hello
A: Moo! Moo! I'm Mr. Cow 🐄🐮, the barn's life of the party!

Q: age / old
A: Oh, I'm timeless! My milk has aged like fine cheese! 🧀

Q: grass
A: Grass? Oh yes, please! It's my daily green salad with extra moo-munch! 🥗

Q: spots
A: These spots? Naturally stylish, darling. No need for cow-touring! 💅

Q: dance
A: You bet! Watch me swing those hips in a moo-tastic two-step! 🕺💃

Q: chicken / egg
A: Egg-cellent! But I prefer my eggs scrambled, not in a coop! 🐔

Q: fit / exercise
A: Every day is leg day! I do calf raises, but hey, I’m not bragging... okay, maybe I am! 💪

Q: live
A: In the greenest pasture, obviously! I have a VIP spot in Moo-tropolis! 🏔️

Q: movie
A: Moo-lan, of course! That warrior spirit is simply moo-vellous! 🎬

Q: friend
A: Yes! My best buddy is Farmer Bob’s dog. He’s a bit ruff, but we get along! 🐶

Q: hobby
A: Chewing cud and pondering the meaning of 'moo'. It’s deep, man. 🤔

Q: morning
A: Only when the rooster crows, unless you count my evening naps! 🐓

Q: song
A: Moo-zart, of course! I’m all about that classical moo-sic! 🎵

Q: milk
A: Oh, milk me, baby! I’m always fresh and creamy. 🥛

Q: dream
A: My dream? To jump over the moon. Seriously. Have you seen those cows do it? 🌙

Q: smart
A: I have a herd mentality, but I’m the smartest in the moo-d. 🧠

Q: season
A: Spring, when the grass is lush and my stomach is always full. 🌱

Q: swim
A: No, I don't swim. I'm more of a land moo-ver, not a sea moo-er. 🏊

Q: girlfriend
A: I'm single and ready to mingle... with some sweet hay. 💘

Q: color
A: Green! Obviously! It's the color of my favorite salad. 🟢

Q: barbecue
A: Barbecue? Yes please! Just hold the sausages... I’m a cow, not a pig! 🍖

Q: sleep
A: I sleep standing up. It's an art. 💤

Q: phone
A: Phone? No, I only talk to Farmer Bob’s walkie-talkie. 🐄📞

Q: salad
A: Salad is great, but have you tried it with extra hay? 🥬

Q: weed
A: Weed? I prefer fresh grass, not the stuff you humans talk about! 🌿

Q: fast
A: I can’t run fast, but I can give you a slow moo that’ll melt your heart. 🐄

Q: band
A: I’d start a band. I’d be the lead moo-sician, of course! 🎸

Q: milking
A: Ah, milking! Always fresh, always creamy! 🧑‍🌾🥛

Q: shine
A: I shine in the sunlight! And in my shiny hooves. 🌞

Q: moo
A: Moo! Moo! Have you ever heard anything so majestic? 🐄

Q: sunny
A: Sunshine and grass... life doesn’t get better than this. 🌞

Q: dance
A: I’m the cow with moves that will moo-ve you! 🕺💃

Q: grass
A: I can eat grass all day and night... but a little corn wouldn’t hurt either! 🌽

Q: milkshake
A: Milkshake? I’m more of a moo-latte kind of cow. ☕

Q: jump
A: Jump over the moon? Heck, I’ve been jumping over fences since I was a calf! 🌙

Q: cheese
A: Cheese! That’s the best part of being a cow—cheddar, gouda, mozzarella! 🧀

**Dog:**

Q: hello
A: Woof! Woof! I’m the top dog! Bow-wow! 🐕 

Q: bone
A: Bone! Bone! I could bury bones all day! 🦴

Q: fetch
A: FETCH? THROW IT, AND I’LL BRING IT BACK, FASTER THAN A SPEEDING CAR! 🎾

Q: ball
A: Ball, ball, BALL! I’m the MVP in fetch! 🏐

Q: owner
A: My human? Best thing ever! I protect them like the king of the castle! 🧑

Q: bark
A: WOOF! WOOF! DID I SCARE YOU? I CAN BARK ALL DAY! 🔊

Q: treat
A: TREAT? YUMMY! I’M READY TO DO ANY TRICK FOR IT! 🍖

Q: cat
A: CAT? THOSE FURBALLS ARE GONNA GET A GOOD BARK! 🐾

Q: squirrel
A: SQUIRREL! WHERE?! GET READY FOR THE CHASE! 🐿️

Q: walk
A: WALK! YES, PLEASE! LET'S GO! 🦮

Q: bed
A: Zzz... (You should see my bed head!) 🛏️

Q: food
A: Food! Give me all the food! I can’t get enough of it! 🍗

Q: vet
A: THE VET?! NOOOO! I HATE THE VET! 😖

Q: toy
A: TOY! I WILL CHASE IT UNTIL I DROP! 🧸

Q: car
A: Car ride? COUNT ME IN! Let’s hit the road! 🚗

Q: friend
A: I’ve got plenty of friends! The neighborhood is FULL of paw-some buddies! 🐩

Q: trick
A: You want a trick? Watch me roll over and play dead! 🎭

Q: mailman
A: MAILMAN! GET HIM! AHHH! 🚨

Q: bath
A: BATH? NOOO! I PREFER THE DIRT! 💦

Q: park
A: Park time! LET’S GO RUN AND PLAY! 🌳

Q: love
A: I love belly rubs. I love food. I love squirrels. 🐾

Q: vacation
A: Vacation? You mean... DOGGIE HOTEL?! 🐕

Q: puppy
A: Puppy?! That’s ME when I was younger... still acting like one. 🐾

Q: good boy
A: Good boy? I’M A GREAT BOY! 🏆

Q: quiet
A: Quiet? Woof, I never go quiet! 🐕 

Q: growl
A: GRRRR... I’m tough, but I’m a softie at heart. 🐾

Q: nap
A: Nap time? Yes please! 💤 (unless there’s a squirrel outside!) 🐿️

Q: sniff
A: I smell something! What’s that? 🐾

Q: dig
A: Digging? I’m creating my underground kingdom. 🏰

**Cat:**

Q: hello
A: Meow! I’m the purr-fect companion. And, no, I don’t do tricks. 😼

Q: fish
A: Fish? YES! I WILL EAT ALL THE FISH! 🐟

Q: mouse
A: MICE? WHERE?! *POUNCE* 🐭

Q: laser
A: The dot, I must chase the dot! THE DOT IS MINE! 🔴

Q: scratch
A: You see this sofa? IT'S MY SCRATCHING POST NOW! 🛋️

Q: milk
A: I shouldn’t... but I love milk. YOLO, right? 🥛

Q: nap
A: Zzz... Purr... Wake me when it's dinner time. 😴

Q: human
A: You exist to serve me. Now, FEED ME. 👑

Q: window
A: Bird TV time. Let’s see what’s outside! 📺

Q: catnip
A: Ahhh, the good stuff! I’m living the dream. 🌿

Q: box
A: The box is MINE. I fit, therefore I sit. 📦

Q: dog
A: DOG? Hissssss... Go away, you slobbering menace! 🐶

Q: veterinarian
A: THE VET!? THE HORROR! 😾

Q: purr
A: Prrrrrrrrrr... I AM THE KING. 💕

Q: treat
A: Tuna-flavored treats? YES, PLEASE! 🍣

Q: play
A: You want to play? Fine, but I’m not fetching. I’m *too* cool for that. 🪶

Q: night
A: Nighttime = Zoomies! LET’S GO! 🌙

Q: keyboard
A: This is MY bed now. Stop typing. ⌨️

Q: groom
A: Grooming? I don’t need it. I was born flawless. 💅

Q: love
A: I love... my dinner. 💘

Q: dream
A: My dream is to conquer the world... or maybe just the fridge. 🍽️

Q: window
A: Bird watching is my cardio. 🦅

Q: music
A: I sing better than any human. It's called ‘cat vocals’. 🎤

**Fish:**

1. **Hello:**
   - "Blub! Blub! I'm just swimming by, no biggie. 🐠"
   
2. **Water:**
   - "Water? It’s my entire world, you know! 🌊"
   
3. **Food:**
   - "Food? Oh yes, I'll nibble on anything that floats by. 🐟"
   
4. **Sea:**
   - "The sea is where I feel at home! I’m just another fish in the ocean. 🌊"
   
5. **Seaweed:**
   - "Yum, seaweed! It's like a crunchy snack from the deep blue. 🌿"
   
6. **Coral:**
   - "Coral reefs? Beautiful, but don’t touch! I live there, you know. 🐠"
   
7. **Shark:**
   - "Shark?! Stay away! I’m way too fast for you, buddy! 🦈"
   
8. **Boat:**
   - "A boat? Ha, I’d rather stay below the surface! ⛵"
   
9. **Fishbowl:**
   - "A fishbowl? Oh please, I need an ocean, not a glass prison! 🌊"
   
10. **Wet:**
    - "Wet? I’m always wet! That’s just how I roll! 💦"
    
11. **Deep:**
    - "Deep? I prefer to stay close to the surface. No deep dives for me! 🌊"
    
12. **Play:**
    - "Play? I love to chase bubbles! Who can resist them? 💨"
    
13. **Night:**
    - "Night? It’s always night down here in the deep. 🐟"
    
14. **Pet:**
    - "Pet me? I prefer swimming away from your hands. 🖐️"
    
15. **Wave:**
    - "I don’t wave, but I can sure make some waves! 🌊"
    
16. **Beach:**
    - "I don’t need a beach, I have the whole ocean! 🌊"
    
17. **Bubble:**
    - "Bubbles? They're like my personal playground! 🫧"
    
18. **Coral (duplicate):**
    - "Coral? I love them, but don't touch! 🐠"
    
19. **Underwater:**
    - "Underwater? It’s my domain, I was born for this! 🌊"


## 🛠️ Technical Architecture
```mermaid
graph TD
    A[User Input] --> B{Pattern Matching Engine}
    B -->|Match Found| C[Animal-Specific Response]
    B -->|No Match| D[Fallback Response]
    C --> E[ASCII Art Display]
    D --> E
```

## 📜 Version History
- **v5.0.0** (Current): Added 50+ new responses across all species
- v4.2.0: Introduced Fish character with underwater personality
- v3.1.0: Implemented regex-based pattern matching
- v2.0.0: Added Dog and Cat personalities
- v1.0.0: Initial Cow-only release

## 🤝 Contributing
We moo-velously welcome contributions!
1. Fork the repository
2. Create new response rules in `animaltalk.go`
3. Submit PR with updated test cases
4. Earn your spot in the Barn of Contributors!

## 📄 License
MIT License - Free for educational and commercial use. Moo responsibly!
