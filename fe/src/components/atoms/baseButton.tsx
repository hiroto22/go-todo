import { Button } from "@mui/material";
import styled from "styled-components";

type onClickBaseButton = {
  text: string;
  onClick: any;
};

export const BaseButton = (props: onClickBaseButton) => {
  return (
    <Button variant="outlined" sx={{ height: "80%" }} onClick={props.onClick}>
      {props.text}
    </Button>
  );
};

// const SButton = styled.button`
//   background-color: blue;
// `;
