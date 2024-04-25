
(() => {
  'use strict'
  const tooltipTriggerList = Array.from(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
  tooltipTriggerList.forEach(tooltipTriggerEl => {
    new bootstrap.Tooltip(tooltipTriggerEl)
  })
})()

var currentMode;
function ReadNightMode(){
  var storedTheme = localStorage.getItem('theme');
  if (storedTheme) {
    currentMode = storedTheme + '-mode'
  }
}

function SetNightBtn(status){
  const nightButtons = document.querySelectorAll('.nightModeButton');
  nightButtons.forEach(nightBtn=>{
    nightBtn.checked = status;
  })
}

function ActiveNightMode()
{
  const nightButtons = document.querySelectorAll(".nightModeButton");
  nightButtons.forEach(nightBtn=>{
    nightBtn.addEventListener("click",()=>{
      if(currentMode?.indexOf("night-mode")==0){
        moveToDayMode();
      }else{
        moveToNightMode();
      }
    })
  })

  if(currentMode?.indexOf("night-mode")==0)
  {
    moveToNightMode();
  }
  else{
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
function RemoveAll(target,theOne){
  const search = document.querySelectorAll(target);
  search.forEach(item=>{
    if(item?.classList.contains(theOne))
      item.classList.remove("bg-dark");
  });
}
function moveToNightMode(){
  SetNightBtn(true);

  Swap(".b-content-divider","b-content-divider-day","b-content-divider-night");
  SwapAll(".btn-dark","btn-dark","btn-light");
  Swap(".btn-outline-dark","btn-outline-dark","btn-outline-light");
  Swap(".summary","summary-day","summary-night");
  SwapAll(".bg-light","bg-light","bg-dark");
  SwapAll(".link-dark","link-dark","link-light");
  SwapAll(".text-dark","text-dark","text-light");
  SwapAll(".border-secondary","border-secondary","border-light");
  SwapAll(".dropdown-menu-light","dropdown-menu-light","dropdown-menu-dark");
  SwapAll(".form-control-day","form-control-day","form-control-night")
  RemoveAll(".form-control","bg-dark")
  SwapAll(".form-control","bg-light","bg-secondary");
  Swap(".navbar-toggler","bg-dark","bg-dark")
  RemoveAll(".active",".link-light")
  SwapAll(".active","link-light","link-dark");
  SwapAll(".active","bg-dark","bg-light");
  console.log("move to night mode");

  localStorage.setItem('theme', 'night');
  currentMode = "night-mode";
}

function moveToDayMode(){
  SetNightBtn(false);

  Swap(".b-content-divider","b-content-divider-night","b-content-divider-day");
  SwapAll(".btn-light","btn-light","btn-dark");
  Swap(".btn-outline-light","btn-outline-light","btn-outline-dark");
  Swap(".summary","summary-night","summary-day");  
  SwapAll(".bg-dark","bg-dark","bg-light");
  SwapAll(".link-light","link-light","link-dark");
  SwapAll(".text-light","text-light","text-dark");
  SwapAll(".border-light","border-light","border-secondary");
  SwapAll(".dropdown-menu-dark","dropdown-menu-dark","dropdown-menu-light");
  SwapAll(".form-control-night","form-control-night","form-control-day");
  RemoveAll(".form-control","bg-light")
  SwapAll(".form-control","bg-secondary","bg-light");
  Swap(".navbar-toggler","bg-light","bg-dark");
  RemoveAll(".active","link-dark");
  SwapAll(".active","link-dark","link-light");
  SwapAll(".active","bg-light","bg-dark");
  console.log("move to day mode");

  localStorage.setItem('theme', 'day');
  currentMode = "day-mode";
}
window.onload = ()=>{
  ReadNightMode();
  ActiveNightMode();

  const textElement = document.querySelector('.interval-text');
  if(textElement==undefined){
    return;
  }
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

// TODO: add this to base.templ
function sendRefreshTokenRequest() {
  fetch('/user/refresh-token', {
      method: 'POST',
      credentials: 'include',
  })
  .then(response => {
      if (response === null) {
          console.log('No response from server');
          return;
      }
      return true;
  })
  .catch(error => {
      console.log('Error refreshing token:', error);
  });
}
sendRefreshTokenRequest();
setInterval(sendRefreshTokenRequest, 5000);