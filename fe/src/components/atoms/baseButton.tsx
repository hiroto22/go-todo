import { Button } from "@mui/material";

type onClickBaseButton = {
  text: string;
  onClick: React.MouseEventHandler<HTMLButtonElement>;
};

export const BaseButton = (props: onClickBaseButton) => {
  return (
    <Button
      variant="outlined"
      sx={{ height: "80%", margin: "3px" }}
      onClick={props.onClick}
    >
      {props.text}
    </Button>
  );
};

// const SButton = styled.button`
//   background-color: blue;
// `;
