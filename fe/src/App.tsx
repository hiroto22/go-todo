import React from "react";
import "./App.css";
import { Route, Routes } from "react-router-dom";
import { RecoilRoot } from "recoil";
import { TopPage } from "./pages/topPage";
import { AddTodoPage } from "./pages/addTodoPage";
import { EditTodoPage } from "./pages/editTodoPage";
import { LoginPage } from "./pages/loginPage";

const App = () => {
  return (
    <RecoilRoot>
      <Routes>
        <Route path="/" element={<TopPage />} />
        <Route path="/addtodo" element={<AddTodoPage />} />
        <Route path="edit-todo" element={<EditTodoPage />} />
        <Route path="/login" element={<LoginPage />} />
      </Routes>
    </RecoilRoot>
  );
};

export default App;
