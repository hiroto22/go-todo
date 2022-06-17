import { Card, CardContent } from "@mui/material";
import { BaseButton } from "../atoms/baseButton";

type TodoCardProps = {
  todo: string;
  text: string;
  onClick: any;
};

export const TodoCard = (props: TodoCardProps) => {
  return (
    <Card sx={{ margin: "8px", display: "flex", alignItems: "center" }}>
      <CardContent>{props.todo}</CardContent>
      <BaseButton text={props.text} onClick={props.onClick} />
    </Card>
  );
};
