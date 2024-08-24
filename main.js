
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
}