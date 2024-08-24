# Question
"Let's discuss a design problem to evaluate your understanding of low-level design (LLD). I’d like you to design a Snake and Ladder game. This will be a simplified version where two or more players can participate. The game will have the following features:

1. **Board Setup:** The board has 100 squares, numbered from 1 to 100. Some squares contain ladders that move a player forward, while others contain snakes that move a player backward.
  
2. **Players:** Multiple players take turns rolling a dice, and their tokens move according to the dice roll. The first player to reach square 100 wins the game.
  
3. **Game Play:** Players roll a 6-sided dice and move forward by the number of squares rolled. If a player lands at the bottom of a ladder, they move up to the ladder's top. If they land on the head of a snake, they slide down to the snake's tail.

4. **Turn Handling:** The game should be able to handle turns among multiple players until one player wins.

5. **Extensibility:** Consider how you might extend the game in the future, such as adding more complex rules or additional players.

Could you walk me through the low-level design of this system? Please describe the classes you would create, their responsibilities, and how they interact with each other. Also, consider how you would ensure the design is scalable and maintainable."

# Requirements
1. **How many people can play a game at a time?**
   - The game should support at least 2 players, but it should be designed in a way that can easily scale to allow more players if needed.

2. **Do we need to have rooms for multiple players to play the game?**
   - Yes, implementing rooms would allow multiple games to be played concurrently by different sets of players. Each room would represent a separate game instance.

3. **Where is this game played? In a browser? Just an API?**
   - For this exercise, let's assume the game will be played via an API. The front-end or user interface is out of scope, but you should design the API to support potential front-end integration, whether it's a web application or mobile app.

4. **Do we need to have an undo feature?**
   - Yes, an undo feature is required. It should allow players to revert to the previous state of the game, undoing the last move made.

5. **Do we need to have authentication? Should only authenticated users be able to play a particular game?**
   - Yes, authentication is required. Only authenticated users should be able to join or create a game room. You can assume there is an existing authentication service that you can integrate with.

6. **Do we need to store the player results in the database?**
   - Yes, storing player results in a database is required. This would allow for tracking of game outcomes, player performance, and other statistics.

7. **Do we need a scoreboard?**
   - Yes, implementing a scoreboard that tracks and displays the rankings of players based on their game results is required.

8. **If in a square there is both a snake's head and a ladder, which takes priority?**
   - In this design, let's assume that if a player lands on a square with both a snake's head and a ladder, the snake takes priority, meaning the player will slide down to the snake's tail.

---

These additional requirements refine the scope of the problem and provide a more comprehensive understanding of the system you need to design.

# Disucssion

- Main Entities
    - Game-Boad : details on game-board
    - Player : Action Dice Roal (can be exted to some generic action) 
    - GameMaster : details on all action by user and by board
    - Dice : Its a type of action 

## Dependency
game <- Player
game <- Game-Boad
player <- Dice

