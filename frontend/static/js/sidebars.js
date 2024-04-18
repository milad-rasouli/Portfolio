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
  document.querySelector("#nightMode").addEventListener("click",()=>{
    if(currentMode?.indexOf("night-mode")==0){
      moveToDayMode();
    }else{
      moveToNightMode();
    }
  })
  if(currentMode?.indexOf("night-mode")==0)
  {
    document.getElementById('nightMode').checked = true;
    moveToNightMode();
  }
  else{
    document.getElementById('nightMode').checked = false;
    moveToDayMode();
  }
}

function Swap(target,oldOne,newOne){
  const content =  document.querySelector(target);
  if (content?.classList.contains(oldOne)){
    content.classList.replace(oldOne,newOne);
  }else if(content?.classList.contains(newOne)==false){
    content.classList.add(newOne);
  }
}


function SwapAll(target,oldOne,newOne){
  const contents =  document.querySelectorAll(target);
  contents.forEach(content=>{
    if (content?.classList.contains(oldOne)){
      content.classList.replace(oldOne,newOne);
    }else if(content?.classList.contains(newOne)==false){
      content.classList.add(newOne);
    }
  })

}
function moveToNightMode(){
  Swap(".b-content-divider","b-content-divider-day","b-content-divider-night");
  SwapAll(".btn-dark","btn-dark","btn-light");
  Swap(".btn-outline-dark","btn-outline-dark","btn-outline-light");
  Swap(".summary","summary-day","summary-night");
  SwapAll(".bg-light","bg-light","bg-dark");
  Swap(".active","text-light","text-dark");
  Swap(".side-nav","bg-light","bg-dark");
  // SwapAll(".nav-link","link-dark","link-light");
  SwapAll(".link-dark","link-dark","link-light");
  SwapAll(".text-dark","text-dark","text-light");
  SwapAll(".border-secondary","border-secondary","border-light");
  SwapAll(".dropdown-menu-light","dropdown-menu-light","dropdown-menu-dark");
  SwapAll(".form-control-day","form-control-day","form-control-night")
  SwapAll(".form-control","bg-light","bg-secondary");
  // Swap(".navbar-toggler","bg-dark","bg-light");
  console.log("move to night mode");

  localStorage.setItem('theme', 'night');
  currentMode = "night-mode";
}

function moveToDayMode(){
  Swap(".b-content-divider","b-content-divider-night","b-content-divider-day");
  SwapAll(".btn-light","btn-light","btn-dark");
  Swap(".btn-outline-light","btn-outline-light","btn-outline-dark");
  Swap(".summary","summary-night","summary-day");  
  Swap(".bg-dark","bg-dark","bg-light");
  Swap(".active","text-dark","text-light");
  Swap(".side-nav","bg-dark","bg-light");
  // SwapAll(".nav-link","link-light","link-dark");
  SwapAll(".link-light","link-light","link-dark");
  SwapAll(".text-light","text-light","text-dark");
  SwapAll(".border-light","border-light","border-secondary");
  SwapAll(".dropdown-menu-dark","dropdown-menu-dark","dropdown-menu-light");
  SwapAll(".form-control-night","form-control-night","form-control-day");
  SwapAll(".form-control","bg-secondary","bg-light");
  // Swap(".navbar-toggler","bg-light","bg-dark");
  console.log("move to day mode");

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
