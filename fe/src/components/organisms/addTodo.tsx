import axios from "axios";
import { useState } from "react";
import { useCreateTodo } from "../../hooks/useCreateTodo";
import { useGetTodos } from "../../hooks/useGetTodos";
import { BaseButton } from "../atoms/baseButton";

export const AddTodo = () => {
  const [text, setText] = useState("");
  const createTodo = useCreateTodo();

  const onClickCreate = () => {
    createTodo(text);
  };

  return (
    <div>
      <textarea onChange={(e) => setText(e.target.value)}></textarea>
      <BaseButton text="タスクを追加" onClick={onClickCreate} />
    </div>
  );
};
