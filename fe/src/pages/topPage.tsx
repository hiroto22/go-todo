import { Header } from "../components/organisms/header";
import styled from "styled-components";
import { Todos } from "../components/organisms/todos";
import { AddIcon } from "../components/atoms/addIcon";
import { useNavigate } from "react-router-dom";

export const TopPage = () => {
  const navigate = useNavigate();
  return (
    <div>
      <Header />
      <Todos />
      <AddIcon onClick={() => navigate("/addtodo")} />
    </div>
  );
};
