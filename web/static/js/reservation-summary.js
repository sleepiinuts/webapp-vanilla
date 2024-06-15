function clicked(e) {
  console.log("clicked...");
  console.log(e.target);
  e.target.classList.toggle("clicked");
  console.log(e.target.classList);
}
