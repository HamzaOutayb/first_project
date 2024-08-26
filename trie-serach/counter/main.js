username = window.prompt(`inter ur username my bro`)
username = username.trim();
firstchar = username.charAt(0);
firstchar = firstchar.toUpperCase();
lastword = username.slice(1);
lastword = lastword.toLowerCase();
username = firstchar + lastword;
console.log(username)

let count =  document.getElementById("count");
let decreasebtn = document.getElementById("decreasebtn");
let resetbtn = document.getElementById("resetbtn");
let increasebtnbtn = document.getElementById("increasebtn");

let counter = 0;
increasebtnbtn.onclick = function(){
    count.textContent = counter++
    }
    
    decreasebtn.onclick = function(){
        count.textContent = counter--
        }
        
        resetbtn.onclick = function() {
            count.textContent = 0
            counter = 0
}

