import { Header } from "../components/organisms/header";
import styled from "styled-components";
import { Todos } from "../components/organisms/todos";
import { AddIcon } from "../components/atoms/addIcon";

export const TopPage = () => {
  return (
    <div>
      <Header />
      <Todos />
      <AddIcon />
    </div>
  );
};
