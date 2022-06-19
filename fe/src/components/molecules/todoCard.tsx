import { Card, CardActionArea, CardContent } from "@mui/material";
import { BaseButton } from "../atoms/baseButton";

type TodoCardProps = {
  todo: string;
  text: string;
  text2: string;
  onClick: any;
  onClick2: any;
  onClickCard: any;
};

export const TodoCard = (props: TodoCardProps) => {
  return (
    <Card sx={{ margin: "8px", display: "flex", alignItems: "center" }}>
      <CardActionArea onClick={props.onClickCard}>
        <CardContent>{props.todo}</CardContent>
      </CardActionArea>
    </Card>
  );
};
