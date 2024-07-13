function clickk(el) {
  el.classList.toggle("clickedd");
}

function tabClick(el) {
  let els = document.getElementsByClassName("tab");

  Array.from(els).forEach((element) => {
    element.classList.remove("active");
  });

  el.classList.add("active");

  // redirect to corersponding room calendar
  // ?rid={roomId}&y={year}&m={month}

  // get query params
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);

  // console.log("id: ", el);

  const y = urlParams.get("y");
  const m = urlParams.get("m");

  window.location.href = `http://localhost:8080/admin/dashboard?rid=${el.id}&y=${y}&m=${m}`;
}

function calMonth(added, d, rid) {
  // add/subtract 1 month
  d.setMonth(d.getMonth() + added);
  const ny = d.getYear() + 1900;
  let nm = d.getMonth() + 1;
  nm = nm.toString().padStart(2, "0");
  window.location.href = `http://localhost:8080/admin/dashboard?rid=${rid}&y=${ny}&m=${nm}`;
}

function nextMonth() {
  // get query params
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);

  const rid = urlParams.get("rid");
  const y = parseInt(urlParams.get("y"));
  const m = parseInt(urlParams.get("m"));

  const d = new Date(`${y}-${m}-01`);
  calMonth(1, d, rid);
}
function prevMonth() {
  // get query params
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);

  const rid = urlParams.get("rid");
  const y = parseInt(urlParams.get("y"));
  const m = parseInt(urlParams.get("m"));

  const d = new Date(`${y}-${m}-01`);
  calMonth(-1, d, rid);
}

function currMonth() {
  // get query params
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);

  const rid = urlParams.get("rid");
  const d = new Date();
  calMonth(0, d, rid);
}
