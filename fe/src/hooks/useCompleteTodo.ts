import axios from "axios";
import { useGetStorageToken } from "./useGetStorageToken";
import { useGetTodos } from "./useGetTodos";

export const useCompleteTodo = () => {
  const CompleteTodo = (token: string, id: number) => {
    const URL = `http://127.0.0.1:8080/completetodo?id=${id}&isComplete=false`;
    axios
      .post(URL, null, { headers: { Authorization: token } })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  const ReturnTodo = (token: string, id: number) => {
    const URL = `http://127.0.0.1:8080/completetodo?id=${id}&isComplete=true`;
    axios
      .post(URL, null, { headers: { Authorization: token } })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  const DeleteTodo = (token: string, id: number) => {
    const URL = `http://127.0.0.1:8080/deletetodo?id=${id}`;
    axios
      .post(URL, null, { headers: { Authorization: token } })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  return { CompleteTodo, ReturnTodo, DeleteTodo };
};
