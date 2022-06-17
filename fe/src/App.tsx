import React from "react";
import "./App.css";
import { Route, Routes } from "react-router-dom";
import { RecoilRoot } from "recoil";
import { TopPage } from "./pages/topPage";

const App = () => {
  return (
    <RecoilRoot>
      <Routes>
        <Route path="/" element={<TopPage />} />
      </Routes>
    </RecoilRoot>
  );
};

export default App;
