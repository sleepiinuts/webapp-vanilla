const picker = new easepick.create({
  element: "#datepicker",
  css: ["https://cdn.jsdelivr.net/npm/@easepick/bundle@1.2.1/dist/index.css"],
  zIndex: 12,
  firstDay: 0,
  plugins: ["RangePlugin"],
});
