const picker = new easepick.create({
  element: "#datepicker",
  css: ["https://cdn.jsdelivr.net/npm/@easepick/bundle@1.2.1/dist/index.css"],
  // this setup doesnt work with go template either??
  setup: (picker) => {
    picker.on("select", (e) => {
      const { start, end } = e.detail;

      console.log("e", e);
      console.log("start: ", start);
      console.log("end: ", end);

      // fetch(`/check-room-avail?datepicker=${start}-${end}`).then((resp) => {
      //   console.log(`Response: ${resp.text}`);
      // });
    });
  },
  zIndex: 12,
  firstDay: 0,
  plugins: ["RangePlugin"],
});

// why console.log doesnt work with go template
console.log("please print me: inside datepicker");
