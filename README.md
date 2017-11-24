# Data Representation and Querying 
# Project Eliza-ChatBot

####Author: Conor McGrath
This repository was created as part of a project for the module [Data Representation and Querying](https://data-representation.github.io/)

The project outcomes were to create a chat-bot web application using Golang.
The project description can be found [here](https://data-representation.github.io/problems/project.html) 

#### How to run the program
You must first install the Go package on your machine to run this program.

Click on the following link to download and install Go. [INSTALL GO](https://golang.org/dl/)

To clone the repository to your local machine, in command prompt enter 
```
git clone https://github.com/g00291461/Eliza.git
```

To run this program these are the following steps you must complete: 
1. **Build and Run** Navigate to the Eliza-Chatbot folder and enter the following to compile the code 
```
go build Eliza.go
```
This will create a .exe file in your current directory. 
Then type the name of that file without any file extension.
```
Eliza
```
2. **Run** Type go run plus the file name and its extension.
```
go run chatbot.go
```  
The webpage will now be served to the local host. 
To view the homepage , you must open you browser and enter to the local address
```
127.0.0.1:8080
```
Or enter
```
localhost:8080
```
### The Project Process
#### Research
As with any project, my first task was to break it down into manageable sections or pieces. I went online and found a few different resources on Regex, JQuery, and Go. I also looked on various repositories and websites as to how other people had built their own version of Eliza using other languages such as Java and Python. Once I had finished my information gathering I looked back on previous examples I had completed during the module itself and tried to take from them what I could and implement it into my own program. 

#### Design
For the design process I basically looked at resources the lecturer had put up and also looked at other examples I had done in the module. Although the code isn’t as object oriented as I would want it to be, I felt I completed it as best I could at this moment in time with all the knowledge of Go, Regex, and JQuery I have. I wrote out in pseudo code how I was going to implement the algorithm to take in user input, have that data manipulated and then returned using a response from the Eliza chat bot using AJAX.

#### Implementation
I first wanted to have the program serve a HTML file to the localhost. Once this was working I moved onto using AJAX and JQUERY to pass the user input from the HTML, to the server, and then back to the HTML page again. I then built the main Eliza.go program but, ultimately, I would have preferred to of had a separate JSON file that held the structure for user inputs and Eliza’s responses. Maybe a data file holding all the patterns for the regular expressions and the keys for words to look out for and the responses to match those keywords.

### Features
After I had the general workings of my program in order minus a few regular expressions I decided to design the look of the web application. I used bootstrap, CSS and JQuery to make the web app look easier on the eye. I added in an ‘nth-child’ parameter in CSS to capture certain list elements when they created “on the fly” in JQuery, and I styled them accordingly. I didn’t want to have too much going on with the background, so I left out any colour or imaging and solely focused on how the messages would appear on the screen. Between the use of CSS and JQuery I was able to manipulate the colour, position, the timing and various other aspects of how each message appears on the screen. With better understanding and more practice, I would have liked to added features like, “Eliza is Responding...”, the use of emojis or the capturing foreign languages also.

### References
I used many different sites in order to learn about Eliza and learn new algorithims that were needed to develop the chatbot. Any adapted code has been referenced within  the code.
Here is a list of a few resourcses I visited most frequently.
+ (https://stackoverflow.com/)
+ (https://www.w3schools.com/jquery/)
+ (http://api.jquery.com/jquery.ajax/)
+ (https://data-representation.github.io/)
+ (https://github.com/codeanticode/eliza)


![Eliza](/static/img/eliza.png?raw=true "Screen Capture of Eliza")
