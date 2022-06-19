import { ChangeEventHandler } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useRecoilState } from "recoil";
import { useCompleteTodo } from "../../hooks/useCompleteTodo";
import { useEditTodo } from "../../hooks/useEditTodo";
import { editTodoState } from "../../state/editTodoState";
import { BaseButton } from "../atoms/baseButton";

export const EditTodo = () => {
  const location = useLocation();
  console.log(location.state);
  const [todo, setTodo] = useRecoilState(editTodoState);
  const editTodo = useEditTodo();
  const onChangeTodo: ChangeEventHandler<HTMLTextAreaElement> = (e: any) => {
    editTodo(todo.id, e.target.value);
  };
  const { CompleteTodo, ReturnTodo, DeleteTodo } = useCompleteTodo();
  const navigate = useNavigate();

  return (
    <div>
      <textarea onChange={onChangeTodo} defaultValue={todo.todo}></textarea>
      {todo.isComplete ? (
        <div>
          <BaseButton
            text="タスクに戻す"
            onClick={() => {
              ReturnTodo(todo.id);
              navigate("/");
            }}
          />
          <BaseButton
            text="削除"
            onClick={() => {
              DeleteTodo(todo.id);
              navigate("/");
            }}
          />
        </div>
      ) : (
        <div>
          <BaseButton
            text="タスク完了"
            onClick={() => {
              CompleteTodo(todo.id);
              navigate("/");
            }}
          />
          <BaseButton
            text="削除"
            onClick={() => {
              DeleteTodo(todo.id);
              navigate("/");
            }}
          />
        </div>
      )}
    </div>
  );
};
