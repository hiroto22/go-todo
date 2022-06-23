import { useState } from "react";
import { useCreateTodo } from "../../hooks/useCreateTodo";
import { BaseButton } from "../atoms/baseButton";
import { TextArea } from "../atoms/textArea";

export const AddTodo = () => {
  const [text, setText] = useState("");
  const CreateTodo = useCreateTodo();
  const token = "Bearer " + sessionStorage.getItem("token");

  const onClickCreate = () => {
    CreateTodo(token, text);
  };

  return (
    <div>
      <TextArea
        onChange={(e: any) => setText(e.target.value)}
        defaultValue=""
      />

      <BaseButton text="タスクを追加" onClick={onClickCreate} />
    </div>
  );
};
