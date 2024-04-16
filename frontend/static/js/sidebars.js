/* global bootstrap: false */
(() => {
  'use strict'
  const tooltipTriggerList = Array.from(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
  tooltipTriggerList.forEach(tooltipTriggerEl => {
    new bootstrap.Tooltip(tooltipTriggerEl)
  })
})()

var currentMode;
function ReadNightMode(){
  // TODO: Set the night mode switch
  var storedTheme = localStorage.getItem('theme');
  if (storedTheme) {
    currentMode = storedTheme + '-mode'
  }
}

function ActiveNightMode()
{
  console.log("ActiveNightMode: ", currentMode)
  document.querySelector("#nightMode").addEventListener("click",()=>{
    console.log("click event: ", currentMode)
    if(currentMode.includes("night-mode")==true){
      moveToDayMode();
    }else{
      moveToNightMode();
    }
  })
  if(currentMode.includes("night-mode")==true)
  {
    moveToNightMode();
  }else{
    moveToDayMode();
  }
}

function Swap(target,oldOne,newOne){
  const content =  document.querySelector(target);
  if (content?.classList?.contains(oldOne)){
    content.classList.replace(oldOne,newOne);
  }else{
    classList.add(newOne);
  }
}
function moveToNightMode(){
  Swap(".b-content-divider","b-content-divider-day","b-content-divider-night");
  // Swap("#btn-github","btn-dark","btn-light");
  // Swap(".btn-here","btn-outline-dark","btn-outline-light");

  localStorage.setItem('theme', 'night');
  currentMode = "night-mode";
}

function moveToDayMode(){
  Swap(".b-content-divider","b-content-divider-night","b-content-divider-day");
  // Swap("#btn-github","btn-light","btn-dark");
  // Swap(".btn-here","btn-outline-light","btn-outline-dark");
  localStorage.setItem('theme', 'day');
  currentMode = "day-mode";
}
window.onload = ()=>{
  ReadNightMode();
  ActiveNightMode();

  const textElement = document.querySelector('.interval-text');
  let text = textElement.textContent;
  textElement.textContent = '';
  let i = 0;

  function typeWriter() {
    if (i < text.length) {
      textElement.innerHTML += text.charAt(i);
      i++;
      setTimeout(typeWriter, 60);
    }
  }
  typeWriter();
}
