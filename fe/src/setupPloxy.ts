export const proxy = require("http-proxy-middleware");

module.exports = function (app: any) {
  const headers = {
    "Content-Type": "application/json",
  };
  app.use(
    proxy("/gettodo", {
      target: "http://127.0.0.1:8080",
      changeOrigin: true,
      secure: false,
      headers: headers,
    })
  );
};
