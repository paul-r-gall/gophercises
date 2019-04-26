Goal: Create a blackjack game that is hosted on a server, and allows multiplayer action. Initially, build a command line blackjack, with just single player and dealer. 

Need: 
 - Player struct?
 - Dealer struct? 
 - Use the deck package already developed. 

**Single Player**

On Start:
 - spin up 2 Player structs (one for human, one for dealer)
 - within Game.run(), need a for not win condition loop. 

Order of Play:
 1. each player bets
 2. deal 1 card to each player, then deal a second card to each player
 3. Each player makes their decisions.
 4. Dealer completes hand.
 5. Resolve bets.

Need "Leave Game" option.