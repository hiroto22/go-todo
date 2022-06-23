import { ChangeEventHandler } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useRecoilState } from "recoil";
import { useCompleteTodo } from "../../hooks/useCompleteTodo";
import { useEditTodo } from "../../hooks/useEditTodo";
import { editTodoState } from "../../state/editTodoState";
import { BaseButton } from "../atoms/baseButton";
import { TextArea } from "../atoms/textArea";

export const EditTodo = () => {
  const token = "Bearer " + sessionStorage.getItem("token");

  const location = useLocation();
  console.log(location.state);
  const [todo, setTodo] = useRecoilState(editTodoState);
  const editTodo = useEditTodo();
  const onChangeTodo: ChangeEventHandler<HTMLTextAreaElement> = (e: any) => {
    editTodo(token, todo.id, e.target.value);
  };
  const { CompleteTodo, ReturnTodo, DeleteTodo } = useCompleteTodo();
  const navigate = useNavigate();
  console.log(todo);

  return (
    <div>
      <TextArea onChange={onChangeTodo} defaultValue={todo.todo}></TextArea>
      {todo.isComplete ? (
        <div>
          <BaseButton
            text="タスクに戻す"
            onClick={() => {
              ReturnTodo(token, todo.id);
              navigate("/");
            }}
          />
          <BaseButton
            text="削除"
            onClick={() => {
              DeleteTodo(token, todo.id);
              navigate("/");
            }}
          />
        </div>
      ) : (
        <div>
          <BaseButton
            text="タスク完了"
            onClick={() => {
              CompleteTodo(token, todo.id);
              navigate("/");
            }}
          />
          <BaseButton
            text="削除"
            onClick={() => {
              DeleteTodo(token, todo.id);
              navigate("/");
            }}
          />
        </div>
      )}
    </div>
  );
};
