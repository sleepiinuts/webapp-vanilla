function focus() {
  let el = document.getElementById("hover");
  //   el.onmouseover = function () {
  //     el.classList.add("hover");
  //   };

  //   el.onmouseleave = function () {
  //     el.classList.remove("hover");
  //   };
  el.onclick = function () {
    el.classList.toggle("clickedd");
    // alert(el.classList);
  };
}

focus();
